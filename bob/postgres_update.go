package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/update/qm"
)

func PostgresSimpleUpdate() (string, []any) {
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
