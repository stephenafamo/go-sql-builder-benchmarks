package squirrel

import sq "github.com/Masterminds/squirrel"

func PostgresBulkInsert() (string, []any) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Insert("users").
		Columns("first_name", "last_name").
		Values("Greg", "Farley").
		Values("Jimmy", "Stewart").
		Values("Jeff", "Jeffers").
		ToSql()
	if err != nil {
		panic(err)
	}

	return sql, args
}
