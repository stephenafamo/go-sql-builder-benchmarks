package squirrel

import sq "github.com/Masterminds/squirrel"

func PostgresSimpleUpdate() (string, []any) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Update("items").
		Set("name", "test").
		Set("address", "111 Test Addr").
		ToSql()

	if err != nil {
		panic(err)
	}

	return sql, args
}
