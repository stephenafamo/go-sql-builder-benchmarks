package main

import (
	"fmt"
	"strings"

	"github.com/stephenafamo/sqlbuilderbenchmarks/bob"
	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/sq"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
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
	fmt.Println("---Bob---")
	Print(bob.PostgresSimpleSelect())

	fmt.Println("---Goqu---")
	Print(goqu.PostgresSimpleSelect())

	fmt.Println("---Sq---")
	Print(sq.PostgresSimpleSelect())

	fmt.Println("---Squirrel---")
	Print(squirrel.PostgresSimpleSelect())
}

func PostgresBulkInsert() {
	fmt.Println("---Bob---")
	Print(bob.PostgresBulkInsert())

	fmt.Println("---Goqu---")
	Print(goqu.PostgresBulkInsert())

	fmt.Println("---Sq---")
	Print(sq.PostgresBulkInsert())

	fmt.Println("---Squirrel---")
	Print(squirrel.PostgresBulkInsert())
}

func PostgresSimpleUpdate() {
	fmt.Println("---Bob---")
	Print(bob.PostgresSimpleUpdate())

	fmt.Println("---Goqu---")
	Print(goqu.PostgresSimpleUpdate())

	fmt.Println("---Sq---")
	Print(sq.PostgresSimpleUpdate())

	fmt.Println("---Squirrel---")
	Print(squirrel.PostgresSimpleUpdate())
}
