# Benchmark of Various SQL Query Builders

Run the bechmarks:

```sh
go test -v -run=XXX -cpu 1,2,4 -benchmem -bench=.
```

View generated queries:

```sh
go run .
```

## Packages

* [Bob](https://github.com/stephenafamo/bob)
* [Goqu](https://github.com/doug-martin/goqu)
* [Sq](https://github.com/bokwoon95/sq)
* [Squirrel](https://github.com/Masterminds/squirrel)

## Help improve this

Kindly send in PRs with new query builders or queries to benchmark.

## Full results: 2022-07-28

```sh
goos: darwin
goarch: arm64
pkg: github.com/stephenafamo/sqlbuilderbenchmarks
BenchmarkPostgresBulkInsert
BenchmarkPostgresBulkInsert/bob
BenchmarkPostgresBulkInsert/bob   	  289826	      4567 ns/op	    2232 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-2 	  267148	      4459 ns/op	    2232 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-4 	  264900	      4243 ns/op	    2232 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/goqu
BenchmarkPostgresBulkInsert/goqu           	  204159	      5781 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-2         	  214806	      5418 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-4         	  217070	      5143 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/sq
BenchmarkPostgresBulkInsert/sq             	  239257	      5264 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-2           	  249280	      4393 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-4           	  242902	      4604 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/squirrel
BenchmarkPostgresBulkInsert/squirrel       	  104138	     12253 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-2     	  102178	     11084 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-4     	  107812	     10792 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresSimpleSelect
BenchmarkPostgresSimpleSelect/bob
BenchmarkPostgresSimpleSelect/bob          	  273201	      4346 ns/op	    2688 B/op	      59 allocs/op
BenchmarkPostgresSimpleSelect/bob-2        	  279733	      3839 ns/op	    2688 B/op	      59 allocs/op
BenchmarkPostgresSimpleSelect/bob-4        	  279439	      4107 ns/op	    2688 B/op	      59 allocs/op
BenchmarkPostgresSimpleSelect/goqu
BenchmarkPostgresSimpleSelect/goqu         	  156697	      6761 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-2       	  189513	      6078 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-4       	  191656	      6051 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/sq
BenchmarkPostgresSimpleSelect/sq           	  229897	      5217 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-2         	  256777	      4558 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-4         	  244849	      4492 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/squirrel
BenchmarkPostgresSimpleSelect/squirrel     	  148010	      7784 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-2   	  166308	      7135 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-4   	  167236	      6966 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleUpdate
BenchmarkPostgresSimpleUpdate/bob
BenchmarkPostgresSimpleUpdate/bob          	  308307	      3568 ns/op	    1520 B/op	      54 allocs/op
BenchmarkPostgresSimpleUpdate/bob-2        	  361322	      3177 ns/op	    1520 B/op	      54 allocs/op
BenchmarkPostgresSimpleUpdate/bob-4        	  387152	      3214 ns/op	    1520 B/op	      54 allocs/op
BenchmarkPostgresSimpleUpdate/goqu
BenchmarkPostgresSimpleUpdate/goqu         	  239624	      4657 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-2       	  278362	      4072 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-4       	  279052	      4120 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/sq
BenchmarkPostgresSimpleUpdate/sq           	  410455	      3161 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-2         	  404882	      2912 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-4         	  381766	      2877 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel
BenchmarkPostgresSimpleUpdate/squirrel     	  158436	      7192 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-2   	  178832	      6204 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-4   	  166838	      6281 ns/op	    2616 B/op	      62 allocs/op
PASS
ok  	github.com/stephenafamo/sqlbuilderbenchmarks	44.058s
```

## Queries

### Simple Select

```sql
---Bob---
Query: SELECT id, name
FROM users
WHERE (id IN ($1, $2, $3))
Args: []interface {}{100, 200, 300}

---Goqu---
Query: SELECT "id", "name" FROM "users" WHERE ("id" IN ($1, $2, $3))
Args: []interface {}{100, 200, 300}

---Sq---
Query: SELECT users.id, users.name FROM users WHERE users.id IN ($1, $2, $3)
Args: []interface {}{100, 200, 300}

---Squirrel---
Query: SELECT id, name FROM users WHERE id IN ($1,$2,$3)
Args: []interface {}{100, 200, 300}
```

### Simple Update

```sql
---Bob---
Query: UPDATE items SET
"name" = $1,
"address" = $2
Args: []interface {}{"test", "111 Test Addr"}

---Goqu---
Query: UPDATE "items" SET "address"=$1,"name"=$2
Args: []interface {}{"111 Test Addr", "Test"}

---Sq---
Query: UPDATE items SET name = $1, address = $2
Args: []interface {}{"test", "111 Test Addr"}

---Squirrel---
Query: UPDATE items SET name = $1, address = $2
Args: []interface {}{"test", "111 Test Addr"}
```

### Bulk Insert

```sql
---Bob---
Query: INSERT INTO users ("first_name", "last_name")
VALUES ($1, $2), ($3, $4), ($5, $6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}

---Goqu---
Query: INSERT INTO "users" ("first_name", "last_name") VALUES ($1, $2), ($3, $4), ($5, $6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}

---Sq---
Query: INSERT INTO users (firstname, lastname) VALUES ($1, $2), ($3, $4), ($5, $6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}

---Squirrel---
Query: INSERT INTO users (first_name,last_name) VALUES ($1,$2),($3,$4),($5,$6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}
```
