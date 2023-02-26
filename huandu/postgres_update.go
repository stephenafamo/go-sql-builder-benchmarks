package huandu

import (
	"github.com/huandu/go-sqlbuilder"
)

func PostgresSimpleUpdate() (string, []any) {
	sql, args := sqlbuilder.PostgreSQL.NewUpdateBuilder().
		Set("name", "test").
		Set("address", "111 Test Addr").
		Build()

	return sql, args
}
