package helpers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetQueryResults(db *sql.DB, statement string, args ...any) (*sql.Rows, error) {
	execStatement, err := db.Prepare(statement)
	if err != nil {
		return nil, err
	}

	rows, err := execStatement.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, err
}

func ExecuteStatement(db *sql.DB, statement string, args ...any) (sql.Result, error) {
	execStatement, err := db.Prepare(statement)
	if err != nil {
		return nil, err
	}

	res, err := execStatement.Exec(args...)
	if err != nil {
		return nil, err
	}

	return res, err
}