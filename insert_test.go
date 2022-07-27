package main

import (
	"testing"

	"github.com/stephenafamo/sqlbuilderbenchmarks/bob"
	"github.com/stephenafamo/sqlbuilderbenchmarks/goqu"
	"github.com/stephenafamo/sqlbuilderbenchmarks/sq"
	"github.com/stephenafamo/sqlbuilderbenchmarks/squirrel"
)

func BenchmarkPostgresBulkInsert(bUp *testing.B) {
	for _, x := range []benchset{
		{name: "bob", f: bob.PostgresBulkInsert},
		{name: "goqu", f: goqu.PostgresBulkInsert},
		{name: "sq", f: sq.PostgresBulkInsert},
		{name: "squirrel", f: squirrel.PostgresBulkInsert},
	} {
		bUp.Run(x.name, func(b *testing.B) {
			bench(b, x.f)
		})
	}
}
