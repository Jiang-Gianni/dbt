/*
Copyright © 2024 Gianni Jiang

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"database/sql"
	"fmt"

	"github.com/Jiang-Gianni/dbt/db"
	"github.com/Jiang-Gianni/dbt/parse"
	_ "github.com/lib/pq"
)

func main() {
	s, err := parse.New("./dbt")
	if err != nil {
		panic(err)
	}
	_ = s

	sqlDB, err := sql.Open("postgres", "postgresql://root:my-secret-pw@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	exe := db.DbExecutor{
		DB:   sqlDB,
		Scan: s,
	}
	rr, err := exe.Run("test")
	if err != nil {
		panic(err)
	}
	for _, r := range rr {
		fmt.Printf("\n\n")
		if r != nil {
			fmt.Print(*r)
		}
	}

	// cmd.Execute()

	// input := `MyFunc('Hello', {1,2,3}, 123, '{1,2,3}')`
	// name, args, err := parse.ParseFunctionCall(input)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("NAME", name)
	// fmt.Println("ARGS", args)
	// fmt.Println("LEN ARGS", len(args))
}
