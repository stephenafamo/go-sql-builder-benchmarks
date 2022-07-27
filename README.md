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
BenchmarkPostgresBulkInsert/bob   	  479548	      2679 ns/op	    2248 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-2 	  463983	      2615 ns/op	    2248 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-4 	  437846	      2759 ns/op	    2248 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/goqu
BenchmarkPostgresBulkInsert/goqu           	  367628	      3415 ns/op	    2928 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-2         	  318274	      3398 ns/op	    2928 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-4         	  344468	      3609 ns/op	    2928 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/sq
BenchmarkPostgresBulkInsert/sq             	  378704	      2864 ns/op	    3280 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-2           	  435078	      2772 ns/op	    3280 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-4           	  420333	      3072 ns/op	    3280 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/squirrel
BenchmarkPostgresBulkInsert/squirrel       	  162462	      7014 ns/op	    4424 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-2     	  156394	      7008 ns/op	    4424 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-4     	  161126	      7013 ns/op	    4424 B/op	     106 allocs/op
BenchmarkPostgresSimpleSelect
BenchmarkPostgresSimpleSelect/bob
BenchmarkPostgresSimpleSelect/bob          	  472821	      2425 ns/op	    2072 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/bob-2        	  498748	      2610 ns/op	    2072 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/bob-4        	  415248	      2446 ns/op	    2072 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/goqu
BenchmarkPostgresSimpleSelect/goqu         	  297669	      4243 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-2       	  272984	      4064 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-4       	  297141	      4051 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/sq
BenchmarkPostgresSimpleSelect/sq           	  410154	      3049 ns/op	    3000 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-2         	  351970	      3026 ns/op	    3000 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-4         	  399742	      2881 ns/op	    3000 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/squirrel
BenchmarkPostgresSimpleSelect/squirrel     	  231603	      4931 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-2   	  258865	      4592 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-4   	  249937	      4517 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleUpdate
BenchmarkPostgresSimpleUpdate/bob
BenchmarkPostgresSimpleUpdate/bob          	  672798	      1680 ns/op	    1088 B/op	      41 allocs/op
BenchmarkPostgresSimpleUpdate/bob-2        	  734822	      1602 ns/op	    1088 B/op	      41 allocs/op
BenchmarkPostgresSimpleUpdate/bob-4        	  726020	      1651 ns/op	    1088 B/op	      41 allocs/op
BenchmarkPostgresSimpleUpdate/goqu
BenchmarkPostgresSimpleUpdate/goqu         	  455451	      2686 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-2       	  452808	      2698 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-4       	  373696	      2823 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/sq
BenchmarkPostgresSimpleUpdate/sq           	  514360	      2350 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-2         	  559113	      2347 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-4         	  467106	      2534 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel
BenchmarkPostgresSimpleUpdate/squirrel     	  199029	      5550 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-2   	  215467	      5522 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-4   	  214370	      5773 ns/op	    2616 B/op	      62 allocs/op
PASS
ok  	github.com/stephenafamo/sqlbuilderbenchmarks	44.910s
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
name = $1,
address = $2
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
Query: INSERT INTO users("first_name", "last_name")
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
