package sq

import (
	"bytes"
	"context"

	"github.com/bokwoon95/sq"
)

func PostgresBulkInsert() (string, []any) {
	t := sq.New[users]("")
	q := sq.InsertInto(t).
		Columns(t.FirstName, t.LastName).
		Values("Greg", "Farley").
		Values("Jimmy", "Stewart").
		Values("Jeff", "Jeffers")

	buf := bytes.NewBuffer(nil)
	var args []any

	err := q.WriteSQL(context.Background(), sq.DialectPostgres, buf, &args, nil)
	if err != nil {
		panic(err)
	}

	return buf.String(), args
}
