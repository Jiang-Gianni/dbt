package db

import (
	"database/sql"
	"fmt"

	"github.com/Jiang-Gianni/dbt/parse"
)

type DbExecutor struct {
	DB   *sql.DB
	Scan *parse.Scanner
}

type QueryResults struct {
	Name    string
	Query   string
	Columns []string
	Rows    [][]string
}

func (exe *DbExecutor) Run(typeToRun string) ([]*QueryResults, error) {
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
		if len(args) > 0 {
			out[i], err = exe.ExecQuery(name, query, args...)
		} else {
			out[i], err = exe.ExecQuery(name, query)
		}
		if err != nil {
			return out, err
		}
	}
	return out, err
}

func (exe *DbExecutor) ExecQuery(name string, query string, args ...any) (*QueryResults, error) {
	qr := &QueryResults{
		Query: query,
		Name:  name,
	}
	var err error
	var rows *sql.Rows
	fmt.Printf("\n---------\n")
	fmt.Println(name, query)
	fmt.Println(args...)
	fmt.Printf("\n---------\n")
	if len(args) > 0 {
		rows, err = exe.DB.Query(query, args...)
	} else {
		rows, err = exe.DB.Query(query)
	}
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
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
			return nil, err
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
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return qr, nil
}
