package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/insert/qm"
)

func PostgresBulkInsert() (string, []any) {
	sql, args, err := psql.Insert(
		qm.Into("users", "first_name", "last_name"),
		qm.Values(psql.Arg("Greg", "Farley")),
		qm.Values(psql.Arg("Jimmy", "Stewart")),
		qm.Values(psql.Arg("Jeff", "Jeffers")),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
