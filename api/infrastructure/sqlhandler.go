package infrastructure

import (
	".../interfaces/database"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlRows struct {
	Rows *sql.Rows
}

type SqlResult struct {
	Result sql.Result
}

/*
	DBを接続するため、抽象化した構造体を返す
	return SqlHandler
*/
func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "user:password@)(localhost:3306)/task_share")
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

/*
	エラーを返す
	return error
*/
func (handler *SqlHandler) ErrNoRows() error {
	return sql.ErrNoRows
}

/*
	検索結果を取得
*/
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return rows, nil
}

/*
	pointerを次の行へ進める
	return error
*/
func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

/*
	次の行にデータがあるのかをチェック
	return bool
*/
func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

/*
	DB接続を閉じる
*/
func (r SqlRows) Close() error {
	return r.Rows.Close()
}

/*
	検索以外のクエリを実行
*/
func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	stmt, err := handler.Conn.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return res, err
	}
	exe, err := stmt.Exec(args...)
	if err != nil {
		return res, err
	}
	res.Result = exe
	return res, nil
}

/*
 insert後、登録されたIDを返す
*/
func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

/*
	insertでインクリメントされたIDを返す
*/
func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}