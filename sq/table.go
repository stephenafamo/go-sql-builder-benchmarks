package sq

import "github.com/bokwoon95/sq"

type users struct {
	sq.TableStruct
	ID        sq.NumberField
	Name      sq.StringField
	FirstName sq.StringField
	LastName  sq.StringField
}

type items struct {
	sq.TableStruct
	Name    sq.StringField
	Address sq.StringField
}
