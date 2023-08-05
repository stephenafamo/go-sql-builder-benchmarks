package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

func PostgresSimpleUpdate() (string, []any) {
	sql, args, err := psql.Update(
		um.Table("items"),
		um.Set("name").ToArg("test"),
		um.Set("address").ToArg("111 Test Addr"),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
