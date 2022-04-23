package typesql

import (
	"github.com/stephenafamo/typesql/dialect/psql"
	"github.com/stephenafamo/typesql/expr"
	"github.com/stephenafamo/typesql/query"
)

func PostgresBulkInsert() (string, []any) {
	qm := psql.InsertQM{}
	sql, args, err := query.Build(psql.Insert(
		qm.Into("user", "first_name", "last_name"),
		qm.Values(expr.Arg("Greg", "Farley")),
		qm.Values(expr.Arg("Jimmy", "Stewart")),
		qm.Values(expr.Arg("Jeff", "Jeffers")),
	))
	if err != nil {
		panic(err)
	}

	return sql, args
}
