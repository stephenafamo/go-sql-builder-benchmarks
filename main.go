package main

import (
	"fmt"
	"strings"

	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
	"github.com/stephenafamo/sqlbuilderbenchmarks/typesql"
)

func Print(sql string, args []any) {
	fmt.Printf("Query: %s\nArgs: %#v\n\n", strings.TrimSpace(sql), args)
}

func main() {
	PostgresSimpleSelect()
	fmt.Printf("\n\n")
	PostgresSimpleUpdate()
	fmt.Printf("\n\n")
	PostgresBulkInsert()
}

func PostgresSimpleSelect() {
	fmt.Println("---TypeSql---")
	Print(typesql.PostgresSimpleSelect())

	fmt.Println("---Goqu---")
	Print(goqu.PostgresSimpleSelect())

	fmt.Println("---Squirrel---")
	Print(squirrel.PostgresSimpleSelect())
}

func PostgresBulkInsert() {
	fmt.Println("---TypeSql---")
	Print(typesql.PostgresBulkInsert())

	fmt.Println("---Goqu---")
	Print(goqu.PostgresBulkInsert())

	fmt.Println("---Squirrel---")
	Print(squirrel.PostgresBulkInsert())
}

func PostgresSimpleUpdate() {
	fmt.Println("---TypeSql---")
	Print(typesql.PostgresSimpleUpdate())

	fmt.Println("---Goqu---")
	Print(goqu.PostgresSimpleUpdate())

	fmt.Println("---Squirrel---")
	Print(squirrel.PostgresSimpleUpdate())
}
