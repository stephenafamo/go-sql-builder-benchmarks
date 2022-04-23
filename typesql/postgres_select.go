package typesql

import (
	"github.com/stephenafamo/typesql/dialect/psql"
	"github.com/stephenafamo/typesql/expr"
	"github.com/stephenafamo/typesql/query"
)

func PostgresSimpleSelect() (string, []any) {
	qm := psql.SelectQM{}
	sql, args, err := query.Build(psql.Select(
		qm.Select("id", "name"),
		qm.From("users"),
		qm.Where(expr.IN("id", expr.Arg(100, 200, 300))),
	))
	if err != nil {
		panic(err)
	}

	return sql, args
}
