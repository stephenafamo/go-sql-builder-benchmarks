package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/select/qm"
)

func PostgresSimpleSelect() (string, []any) {
	sql, args, err := psql.Select(
		qm.Columns("id", "name"),
		qm.From("users"),
		qm.Where(psql.X("id").In(psql.Arg(100, 200, 300))),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
