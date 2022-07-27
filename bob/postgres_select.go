package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
)

func PostgresSimpleSelect() (string, []any) {
	qm := psql.SelectQM{}

	sql, args, err := psql.Select(
		qm.Select("id", "name"),
		qm.From("users"),
		qm.Where(psql.X("id").In(psql.Arg(100, 200, 300))),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
