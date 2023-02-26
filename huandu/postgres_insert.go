package huandu

import (
	"github.com/huandu/go-sqlbuilder"
)

func PostgresBulkInsert() (string, []any) {
	sql, args := sqlbuilder.PostgreSQL.NewInsertBuilder().
		InsertInto("users").
		Cols("first_name", "last_name").
		Values("Greg", "Farley").
		Values("Jimmy", "Stewart").
		Values("Jeff", "Jeffers").
		Build()

	return sql, args
}
