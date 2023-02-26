package main

import (
	"testing"

	"github.com/stephenafamo/sqlbuilderbenchmarks/bob"
	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/huandu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/sq"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
)

func BenchmarkPostgresSimpleSelect(bUp *testing.B) {
	for _, x := range []benchset{
		{name: "bob", f: bob.PostgresSimpleSelect},
		{name: "goqu", f: goqu.PostgresSimpleSelect},
		{name: "sq", f: sq.PostgresSimpleSelect},
		{name: "squirrel", f: squirrel.PostgresSimpleSelect},
		{name: "huandu", f: huandu.PostgresSimpleSelect},
	} {
		bUp.Run(x.name, func(b *testing.B) {
			bench(b, x.f)
		})
	}
}
