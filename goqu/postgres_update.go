package goqu

import (
	"github.com/doug-martin/goqu/v9"
	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func PostgresSimpleUpdate() (string, []any) {
	dialect := goqu.Dialect("postgres")
	sql, args, err := dialect.Update("items").
		Set(goqu.Record{
			"name":    "Test",
			"address": "111 Test Addr",
		}).
		Prepared(true).
		ToSQL()

	if err != nil {
		panic(err)
	}

	return sql, args
}
