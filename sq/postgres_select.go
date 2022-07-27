package sq

import (
	"bytes"
	"context"

	"github.com/bokwoon95/sq"
)

func PostgresSimpleSelect() (string, []any) {
	t := sq.New[users]("")
	q := sq.
		From(t).
		Select(t.ID, t.Name).
		Where(t.ID.In([]int{100, 200, 300}))

	var buf = bytes.NewBuffer(nil)
	var args []any

	err := q.WriteSQL(context.Background(), sq.DialectPostgres, buf, &args, nil)
	if err != nil {
		panic(err)
	}

	return buf.String(), args
}
