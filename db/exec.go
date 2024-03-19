package db

import (
	"database/sql"
	"strconv"

	"github.com/Jiang-Gianni/dbt/parse"
)

type QueryExecutor struct {
	DB   *sql.DB
	Scan *parse.Scanner
}

type QueryResults struct {
	// in .sql files: -- test: YOUR_TEST_NAME
	Name string

	// sql query
	Query string

	// column of the query output
	Columns []string

	// rows of the query output
	Rows [][]string

	// file name containing the query
	FileName string

	// line number of the file
	Line string

	// error message (if present)
	Error string

	// arguments used in the query
	Args []any
}

func (exe *QueryExecutor) Run(typeToRun string) ([]*QueryResults, error) {
	var err error
	var query string
	var args []any
	nameList := exe.Scan.MapList[typeToRun]
	out := make([]*QueryResults, len(nameList))
	for i, name := range nameList {
		query, args, err = exe.Scan.Extract(name)
		if err != nil {
			return nil, err
		}
		out[i], err = exe.ExecQuery(name, query, args...)
		if err != nil {
			return out, err
		}
	}
	return out, err
}

func (exe *QueryExecutor) ExecQuery(name string, query string, args ...any) (qr *QueryResults, err error) {
	qr = &QueryResults{
		Query:    query,
		Name:     name,
		Line:     strconv.Itoa(exe.Scan.Queries[name].Line),
		FileName: exe.Scan.Queries[name].FileName,
		Args:     args,
	}

	// If there is an error we don't bubble it up but rather insert it in the result struct
	defer func() {
		if err != nil {
			qr.Error = err.Error()
			err = nil
		}
	}()

	var rows *sql.Rows
	rows, err = exe.DB.Query(query, args...)
	if err != nil {
		return qr, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return qr, err
	}
	qr.Columns = cols

	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		result := make([]string, len(cols))
		if err := rows.Scan(dest...); err != nil {
			return qr, err
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		qr.Rows = append(qr.Rows, result)
	}

	if err := rows.Close(); err != nil {
		return qr, err
	}
	if err := rows.Err(); err != nil {
		return qr, err
	}

	return qr, nil
}
