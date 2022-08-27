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
	fmt.Printf("Query: %s\nArgs: %#v\n", strings.TrimSpace(sql), args)
}

func main() {
	PostgresSimpleSelect()
	fmt.Printf("\n\n")
	PostgresSimpleUpdate()
	fmt.Printf("\n\n")
	PostgresBulkInsert()
	fmt.Printf("\n\n")
}

func PostgresSimpleSelect() {
	fmt.Printf("### Simple Select\n\n```sql")
	fmt.Println("\n---Bob---")
	Print(bob.PostgresSimpleSelect())

	fmt.Println("\n---Goqu---")
	Print(goqu.PostgresSimpleSelect())

	fmt.Println("\n---Sq---")
	Print(sq.PostgresSimpleSelect())

	fmt.Println("\n---Squirrel---")
	Print(squirrel.PostgresSimpleSelect())
	fmt.Printf("```")
}

func PostgresBulkInsert() {
	fmt.Printf("### Bulk Insert\n\n```sql")
	fmt.Println("\n---Bob---")
	Print(bob.PostgresBulkInsert())

	fmt.Println("\n---Goqu---")
	Print(goqu.PostgresBulkInsert())

	fmt.Println("\n---Sq---")
	Print(sq.PostgresBulkInsert())

	fmt.Println("\n---Squirrel---")
	Print(squirrel.PostgresBulkInsert())
	fmt.Printf("```")
}

func PostgresSimpleUpdate() {
	fmt.Printf("### Simple Update\n\n```sql")
	fmt.Println("\n---Bob---")
	Print(bob.PostgresSimpleUpdate())

	fmt.Println("\n---Goqu---")
	Print(goqu.PostgresSimpleUpdate())

	fmt.Println("\n---Sq---")
	Print(sq.PostgresSimpleUpdate())

	fmt.Println("\n---Squirrel---")
	Print(squirrel.PostgresSimpleUpdate())
	fmt.Printf("```")
}
