package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

func PostgresSimpleUpdate() (string, []any) {
	sql, args, err := psql.Update(
		um.Table("items"),
		um.SetArg("name", "test"),
		um.SetArg("address", "111 Test Addr"),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
