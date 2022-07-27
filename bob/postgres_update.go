package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
)

func PostgresSimpleUpdate() (string, []any) {
	qm := psql.UpdateQM{}

	sql, args, err := psql.Update(
		qm.Table("items"),
		qm.SetArg("name", "test"),
		qm.SetArg("address", "111 Test Addr"),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
