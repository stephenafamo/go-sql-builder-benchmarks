package typesql

import (
	"github.com/stephenafamo/typesql/dialect/psql"
	"github.com/stephenafamo/typesql/query"
)

func PostgresSimpleUpdate() (string, []any) {
	qm := psql.UpdateQM{}
	sql, args, err := query.Build(psql.Update(
		qm.Table("items"),
		qm.SetArg("name", "test"),
		qm.SetArg("address", "111 Test Addr"),
	))

	if err != nil {
		panic(err)
	}

	return sql, args
}
