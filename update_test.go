package main

import (
	"testing"

	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
	"github.com/stephenafamo/sqlbuilderbenchmarks/typesql"
)

func BenchmarkPostgresSimpleUpdate(bUp *testing.B) {
	for name, Func := range map[string]benchFunc{
		"typesql":  typesql.PostgresSimpleUpdate,
		"squirrel": squirrel.PostgresSimpleUpdate,
		"goqu":     goqu.PostgresSimpleUpdate,
	} {
		bUp.Run(name, func(b *testing.B) {
			bench(b, Func)
		})
	}
}
