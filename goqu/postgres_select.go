package goqu

import (
	"github.com/doug-martin/goqu/v9"
	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func PostgresSimpleSelect() (string, []any) {
	dialect := goqu.Dialect("postgres")
	sql, args, err := dialect.Select("id", "name").
		From("users").
		Where(goqu.Ex{"id": []int{100, 200, 300}}).
		Prepared(true).
		ToSQL()
	if err != nil {
		panic(err)
	}

	return sql, args
}
