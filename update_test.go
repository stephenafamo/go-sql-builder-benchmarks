package main

import (
	"testing"

	"github.com/stephenafamo/sqlbuilderbenchmarks/bob"
	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/sq"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
)

func BenchmarkPostgresSimpleUpdate(bUp *testing.B) {
	for _, x := range []benchset{
		{name: "bob", f: bob.PostgresSimpleUpdate},
		{name: "goqu", f: goqu.PostgresSimpleUpdate},
		{name: "sq", f: sq.PostgresSimpleUpdate},
		{name: "squirrel", f: squirrel.PostgresSimpleUpdate},
	} {
		bUp.Run(x.name, func(b *testing.B) {
			bench(b, x.f)
		})
	}
}
