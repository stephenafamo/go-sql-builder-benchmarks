package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func PostgresSimpleSelect() (string, []any) {
	sql, args, err := psql.Select(
		sm.Columns("id", "name"),
		sm.From("users"),
		sm.Where(psql.Quote("id").In(psql.Arg(100, 200, 300))),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
