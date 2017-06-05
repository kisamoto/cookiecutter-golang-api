package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"{{cookiecutter.repo}}/pkg/cache"
)

// NewTable creates a helper to interact with a certain table in a provided *sqlx.DB
func NewTable(db *sqlx.DB, tableName string) *Table {
	return &Table{
		db:        db,
		cache:     nil,
		tableName: tableName,
	}
}

// NewCacheableTable creates a helper to interact with a
// certain table in a provided *sqlx.DB. By providing a cache.Cache
// certain retrieval queries can have their results cached.
// It makes no sense to cache insertions/updates/deletes although
// the caller should trigger a relevant cache bust if relevant.
func NewCacheableTable(db *sqlx.DB, cache cache.Cache, tableName string) *Table {
	return &Table{
		db:        db,
		cache:     cache,
		tableName: tableName,
	}
}

// Table wraps *sqlx.DB with some added functionality
type Table struct {
	db          *sqlx.DB
	cache       cache.Cache
	tableName   string
	transaction *sqlx.Tx
}

func (t *Table) newTransactionIfNeeded(tx *sqlx.Tx) (*sqlx.Tx, bool, error) {
	var err error
	wrapInSingleTransaction := false

	if tx != nil {
		return tx, wrapInSingleTransaction, nil
	}

	tx, err = t.db.Beginx()
	if err == nil {
		wrapInSingleTransaction = true
	}

	if err != nil {
		return nil, wrapInSingleTransaction, err
	}

	return tx, wrapInSingleTransaction, nil
}

func (t *Table) whereSQL(where string) string {
	sql := fmt.Sprintf("SELECT * FROM %s", t.tableName)
	if where != "" {
		sql = fmt.Sprintf("%s WHERE %s ", sql, where)
	}
	return sql
}

// SelectWhere creates or uses an existing transaction selecting * into 
// the destination interface using a provided set of
// WHERE conditions and arguments.
// e.g. err := SelectWhere(nil, &dest{}, "enabled=$1", true)
func (t *Table) SelectWhere(tx *sqlx.Tx, dest interface{}, where string, args ...interface{}) error {

	if t.tableName == "" {
		return errors.New("Table must not be empty.")
	}

	sql := t.whereSQL(where)

	tx, wrapInSingleTransaction, err := t.newTransactionIfNeeded(tx)
	if tx == nil {
		return errors.New("Transaction struct must not be empty.")
	}
	if err != nil {
		return err
	}

	destType := reflect.TypeOf(dest)
	// Check if slice and use Select
	if destType.Kind() == reflect.Slice {
		tx.Select(dest, sql, args...)
		// If not try QueryRow as only expecting single
	} else {
		tx.QueryRow(sql, args...).Scan(dest)
	}
	logutil.DebugSQLQuery(env.Log, sql)
	if wrapInSingleTransaction {
		err = tx.Commit()
	}
	return err
}

// SelectByID wraps SelectWhere providing the correct WHERE statement to use the given ID
func (t *Table) SelectByID(tx *sqlx.Tx, dest interface{}, id int) error {
	return t.SelectWhere(tx, dest, "id=$1", id)
}

// CachedSelectWhere performs the same as SelectWhere however will check
// the cache for a result and set one if necessary.
// To avoid complications, Cached* methods do not accept transactions
// e.g. err := CachedSelectWhere(&dest{}, "enabled=$1", true)
func (t *Table) CachedSelectWhere(dest interface{}, where string, args ...interface{}) error {
	if t.cache == nil {
		// log that no cache available and fall back to uncached SELECT
		// TODO: add logging
		return t.SelectWhere(nil, dest, where, args)
	}

	// TODO: Caching Logic

	return nil
}

// CachedSelectByID wraps CachedSelectWhere providing the correct 
// WHERE statement to use the given ID
func (t *Table) CachedSelectByID(dest interface{}, id int) error {
	return t.CachedSelectWhere(tx, dest, "id=$1", id)
}

// Insert inserts a map of column headers to values into a table
func (t *Table) Insert(tx *sqlx.Tx, colNamesValues map[string]interface{}) (sql.Result, error) {

	if t.tableName == "" {
		return nil, errors.New("Table must not be empty.")
	}

	tx, wrapInSingleTransaction, err := t.newTransactionIfNeeded(tx)
	if tx == nil {
		return nil, errors.New("Transaction struct must not be empty.")
	}
	if err != nil {
		return nil, err
	}

	var (
		columnNames  []string
		values       []interface{}
		dollarValues []string
		loopCounter  = 1
		insertResult = &sqlResult{}
	)

	for column, value := range colNamesValues {
		columnNames = append(columnNames, column)
		values = append(values, value)
		// Todo: there has to be a better way to do this than create an array each time?
		dollarValues = append(dollarValues, fmt.Sprintf("$%d", loopCounter))
		loopCounter++
	}

	// As we control the tables assume all tables have an "id" column
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		t.tableName,
		strings.Join(columnNames, ", "),
		strings.Join(dollarValues, ", "),
	)
	err = tx.QueryRow(query, values).Scan(&insertResult.lastInsertID)
	if err != nil {
		return nil, err
	}
	insertResult.rowsAffected = 1
	logutil.DebugSQLQuery(env.Log, query)
	if wrapInSingleTransaction {
		err = tx.Commit()
	}
	return insertResult, err
}

// Exec runs arbritary SQL in a transaction
func (t *Table) Exec(tx *sqlx.Tx, query string, args ...interface{}) (sql.Result, error) {

	if t.tableName == "" {
		return nil, errors.New("Table must not be empty.")
	}

	tx, wrapInSingleTransaction, err := t.newTransactionIfNeeded(tx)
	if tx == nil {
		return nil, errors.New("Transaction struct must not be empty.")
	}
	if err != nil {
		return nil, err
	}
	res, err := tx.Exec(query, args...)
	if logutil.LogError(env.Log, err, "table.Exec()") != nil {
		return nil, err
	}
	logutil.DebugSQLQuery(env.Log, query)
	if wrapInSingleTransaction {
		err = tx.Commit()
	}
	return res, err
}

// UpdateWhere uses Exec to UPDATE values based on WHERE conditions
func (t *Table) UpdateWhere(tx *sqlx.Tx, values map[string]interface{}, where string, args ...interface{}) (sql.Result, error) {

	if t.tableName == "" {
		return nil, errors.New("Table must not be empty.")
	}

	sql := fmt.Sprintf("UPDATE %s SET ", t.tableName)
	var (
		valueStrings []string
	)
	// Use the positional nature of arguments array. Start with len of args
	// and append to args array for simplicity.
	loopcounter := len(args)
	for k, v := range values {
		loopcounter++
		valueStrings = append(valueStrings, fmt.Sprintf("%s = $%d", k, loopcounter))
		args = append(args, v)
	}
	sql = fmt.Sprintf("%s %s", sql, strings.Join(valueStrings, ", "))
	if where != "" {
		sql = fmt.Sprintf("%s WHERE %s ", sql, where)
	}
	// Put the update arguments at the front and any passed args at the end
	return t.Exec(tx, sql, args...)
}

// UpdateByID is shorthand to UpdateWhere only specifying the required ID
func (t *Table) UpdateByID(tx *sqlx.Tx, values map[string]interface{}, id int) (sql.Result, error) {
	return t.UpdateWhere(tx, values, "id=$1", id)
}
