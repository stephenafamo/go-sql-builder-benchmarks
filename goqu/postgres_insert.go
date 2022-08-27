package goqu

import (
	"github.com/doug-martin/goqu/v9"
	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func PostgresBulkInsert() (string, []any) {
	dialect := goqu.Dialect("postgres")
	sql, args, err := dialect.Insert("users").
		Cols("first_name", "last_name").
		Vals(
			goqu.Vals{"Greg", "Farley"},
			goqu.Vals{"Jimmy", "Stewart"},
			goqu.Vals{"Jeff", "Jeffers"},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		panic(err)
	}

	return sql, args
}
