# Benchmark of Various SQL Query Builders

Run the bechmarks:

```sh
go test -v -run=XXX -benchmem -bench=.
```

View generated queries:

```sh
go run .
```

## Packages

* [TypeSQL](https://github.com/stephenafamo/typesql)
* [Squirrel](https://github.com/Masterminds/squirrel)
* [Goqu](https://github.com/doug-martin/goqu)

## Help improve this

Kindly send in PRs with new query builders or queries to benchmark.

## Full results

```sh
goos: darwin
goarch: arm64
pkg: github.com/stephenafamo/sqlbuilderbenchmarks
BenchmarkPostgresBulkInsert
BenchmarkPostgresBulkInsert/typesql
BenchmarkPostgresBulkInsert/typesql-8         	  363344	      3106 ns/op	    2160 B/op	      73 allocs/op
BenchmarkPostgresBulkInsert/squirrel
BenchmarkPostgresBulkInsert/squirrel-8        	  133378	      8623 ns/op	    4249 B/op	     105 allocs/op
BenchmarkPostgresBulkInsert/goqu
BenchmarkPostgresBulkInsert/goqu-8            	  306794	      4080 ns/op	    2928 B/op	      82 allocs/op
BenchmarkPostgresSimpleSelect
BenchmarkPostgresSimpleSelect/typesql
BenchmarkPostgresSimpleSelect/typesql-8       	  475926	      3025 ns/op	    1984 B/op	      53 allocs/op
BenchmarkPostgresSimpleSelect/squirrel
BenchmarkPostgresSimpleSelect/squirrel-8      	  174225	      5921 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/goqu
BenchmarkPostgresSimpleSelect/goqu-8          	  236059	      4874 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleUpdate
BenchmarkPostgresSimpleUpdate/typesql
BenchmarkPostgresSimpleUpdate/typesql-8       	  635308	      1797 ns/op	    1032 B/op	      38 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel
BenchmarkPostgresSimpleUpdate/squirrel-8      	  258829	      4779 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/goqu
BenchmarkPostgresSimpleUpdate/goqu-8          	  371764	      3268 ns/op	    2192 B/op	      57 allocs/op
```

## Queries

### Simple Select

```sql
---TypeSql---
Query: SELECT id, name
FROM users
WHERE id IN ($1, $2, $3)
Args: []interface {}{100, 200, 300}

---Goqu---
Query: SELECT "id", "name" FROM "users" WHERE ("id" IN ($1, $2, $3))
Args: []interface {}{100, 200, 300}

---Squirrel---
Query: SELECT id, name FROM users WHERE id IN ($1,$2,$3)
Args: []interface {}{100, 200, 300}
```

### Simple Update

```sql
---TypeSql---
Query: UPDATE items SET
name = $1,
address = $2
Args: []interface {}{"test", "111 Test Addr"}

---Goqu---
Query: UPDATE "items" SET "address"=$1,"name"=$2
Args: []interface {}{"111 Test Addr", "Test"}

---Squirrel---
Query: UPDATE items SET name = $1, address = $2
Args: []interface {}{"test", "111 Test Addr"}
```

### Bulk Insert

```sql
---TypeSql---
Query: INSERT INTO user("first_name", "last_name")
VALUES ($1, $2), ($3, $4), ($5, $6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}

---Goqu---
Query: INSERT INTO "user" ("first_name", "last_name") VALUES ($1, $2), ($3, $4), ($5, $6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}

---Squirrel---
Query: INSERT INTO user (first_name,last_name) VALUES ($1,$2),($3,$4),($5,$6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}
```
