package squirrel

import sq "github.com/Masterminds/squirrel"

func PostgresSimpleSelect() (string, []any) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Select("id", "name").
		From("users").
		Where("id IN (?,?,?)", 100, 200, 300).
		ToSql()

	if err != nil {
		panic(err)
	}

	return sql, args
}
