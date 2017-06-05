package postgresql

import (
	"bytes"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Query interface is a simple wrapper around a SQL query
// for internal use
type Query interface {
	SQL() string
	Parameters() []interface{}
	Result() QueryResult
}

// QueryResult wraps misc. data (sqlx.Row, sqlx.Rows, ) and
// implements sql.Result
type QueryResult struct {
	lastInsertedID int64
	rowsAffected   int64
	Data           interface{}
}

// LastInsertId implements sql.Result
func (qr *QueryResult) LastInsertId() (int64, error) {
	return qr.lastInsertedID, nil
}

// RowsAffected implements sql.Result
func (qr *QueryResult) RowsAffected() (int64, error) {
	return qr.rowsAffected, nil
}

// SimpleQuery is a wrapper around a SQL query and result
type SimpleQuery struct {
	SimpleSQL        string
	SimpleParameters []interface{}
	SimpleResult     QueryResult
}

// SQL implements Query interface
func (q *SimpleQuery) SQL() string {
	return q.SimpleSQL
}

// Parameters implements Query interface
func (q *SimpleQuery) Parameters() []interface{} {
	return q.SimpleParameters
}

// Result implements Query interface
func (q *SimpleQuery) Result() QueryResult {
	return q.SimpleResult
}

func NewCacheableQuery(sql string, parameters []interface{}) *CacheableQuery {
	return &CacheableQuery{
		SimpleSQL: sql,
		SimpleParameters: parameters,
	}
}

// A CacheableQuery implements the cache.Cacheable interface
// allowing the query and result to be cached
type CacheableQuery struct {
	SimpleQuery
}

// MarshalKey converts the SQL query complete with parameters
// and converts it to a []byte key.
// Depending on size of keys it may be worth investigating
// using a hashed version of the query as the key
func (q *CacheableQuery) MarshalKey() ([]byte, error) {
	// Simple plain-text encoding
	var b bytes.Buffer
	fmt.Fprintln(&b, q.SQL(), q.Parameters())
	return b.Bytes(), nil
}

func (q *CacheableQuery) MarshalValue() ([]byte, error) {
	// Simple plain-text encoding
	var b bytes.Buffer
	fmt.Fprintln(&b, q.Result().Data)
	return b.Bytes(), nil
}

// UnmarshalKey extracts the []byte key from the cache and rebuilds
// the SQL query and parameters
func (q *CacheableQuery) UnmarshalKey(data []byte) error {
	// Simple plain-text encoding
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &q.SimpleSQL, &q.SimpleParameters)
	return err
}

func (q *CacheableQuery) UnmarshalValue(data []byte) error {
	// Simple plain-text encoding
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &q.SimpleResult.Data)
	return err
}

type CacheableRows struct {
	CacheableQuery
}

func (cr *CacheableRows) Rows() (sqlx.Rows, error) {
	return cr.Result().Data.(sqlx.Rows), nil
}

type CacheableRow struct {
	CacheableQuery
}

func (cr *CacheableRow) Row() (sqlx.Row, error) {
	return cr.Result().Data.(sqlx.Row), nil
}
