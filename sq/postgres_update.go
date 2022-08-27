package sq

import (
	"bytes"
	"context"

	"github.com/bokwoon95/sq"
)

func PostgresSimpleUpdate() (string, []any) {
	t := sq.New[items]("")
	q := sq.Update(t).
		Set(t.Name.SetString("test")).
		Set(t.Address.SetString("111 Test Addr"))

	buf := bytes.NewBuffer(nil)
	var args []any

	err := q.WriteSQL(context.Background(), sq.DialectPostgres, buf, &args, nil)
	if err != nil {
		panic(err)
	}

	return buf.String(), args
}
