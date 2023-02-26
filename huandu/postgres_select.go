package huandu

import "github.com/huandu/go-sqlbuilder"

func PostgresSimpleSelect() (string, []any) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sql, args := sb.
		Select("id", "name").
		From("users").
		Where(sb.In("id", 100, 200, 300)).
		Build()

	return sql, args
}
