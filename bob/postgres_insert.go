package bob

import (
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/im"
)

func PostgresBulkInsert() (string, []any) {
	sql, args, err := psql.Insert(
		im.Into("users", "first_name", "last_name"),
		im.Values(psql.Arg("Greg", "Farley")),
		im.Values(psql.Arg("Jimmy", "Stewart")),
		im.Values(psql.Arg("Jeff", "Jeffers")),
	).Build()
	if err != nil {
		panic(err)
	}

	return sql, args
}
