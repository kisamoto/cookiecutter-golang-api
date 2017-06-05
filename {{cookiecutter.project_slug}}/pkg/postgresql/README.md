# PostgreSQL Service

## ORM

No ORM here. Raw SQL manipulation with Struct scanning powered by [`jmoiron/sqlx`](https://github.com/jmoiron/sqlx). 

## Layout

General layout has been adapted after getting some ideas from a [Medium Post](https://medium.com/bumpers/our-go-is-fine-but-our-sql-is-great-b4857950a243) on managing SQL in Go. 

- SQL queries are kept in a `\{\{ model_name \}\}_queries.go` file as string constants
- Services prepare statements - `*sqlx.NamedStmt` or `*sqlx.Stmt` - as struct attributes upon service initialisation
- Transactions - `*sqlx.Tx` - use these statements to query the database in the service method implementations

Following this pattern keeps the SQL and Go code suitably separated while preparing SQL statements verifies any sytanx errors upon service initialisation rather than during execution time.
