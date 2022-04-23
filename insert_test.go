package main

import (
	"testing"

	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
	"github.com/stephenafamo/sqlbuilderbenchmarks/typesql"
)

func BenchmarkPostgresBulkInsert(bUp *testing.B) {
	for name, Func := range map[string]benchFunc{
		"typesql":  typesql.PostgresBulkInsert,
		"squirrel": squirrel.PostgresBulkInsert,
		"goqu":     goqu.PostgresBulkInsert,
	} {
		bUp.Run(name, func(b *testing.B) {
			bench(b, Func)
		})
	}
}
