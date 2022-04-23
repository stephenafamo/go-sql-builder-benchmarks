package main

import (
	"testing"

	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
	"github.com/stephenafamo/sqlbuilderbenchmarks/typesql"
)

func BenchmarkPostgresSimpleSelect(bUp *testing.B) {
	for name, Func := range map[string]benchFunc{
		"typesql":  typesql.PostgresSimpleSelect,
		"squirrel": squirrel.PostgresSimpleSelect,
		"goqu":     goqu.PostgresSimpleSelect,
	} {
		bUp.Run(name, func(b *testing.B) {
			bench(b, Func)
		})
	}
}
