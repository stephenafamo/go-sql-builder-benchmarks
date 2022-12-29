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
BenchmarkPostgresBulkInsert/bob   	  578324	      2049 ns/op	    2296 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-2 	  636212	      1859 ns/op	    2296 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-4 	  632023	      1873 ns/op	    2296 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/goqu
BenchmarkPostgresBulkInsert/goqu           	  481170	      2596 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-2         	  508986	      2351 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-4         	  514045	      2313 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/sq
BenchmarkPostgresBulkInsert/sq             	  537190	      2263 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-2           	  582798	      1988 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-4           	  592009	      1995 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/squirrel
BenchmarkPostgresBulkInsert/squirrel       	  220479	      5752 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-2     	  240870	      4971 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-4     	  237212	      4976 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresSimpleSelect
BenchmarkPostgresSimpleSelect/bob
BenchmarkPostgresSimpleSelect/bob          	  613389	      1988 ns/op	    2664 B/op	      59 allocs/op
BenchmarkPostgresSimpleSelect/bob-2        	  659967	      1759 ns/op	    2664 B/op	      59 allocs/op
BenchmarkPostgresSimpleSelect/bob-4        	  666402	      1762 ns/op	    2664 B/op	      59 allocs/op
BenchmarkPostgresSimpleSelect/goqu
BenchmarkPostgresSimpleSelect/goqu         	  389308	      3080 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-2       	  419546	      2771 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-4       	  423420	      2782 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/sq
BenchmarkPostgresSimpleSelect/sq           	  514083	      2380 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-2         	  564770	      2070 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-4         	  567272	      2077 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/squirrel
BenchmarkPostgresSimpleSelect/squirrel     	  347695	      3627 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-2   	  364162	      3205 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-4   	  365632	      3231 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleUpdate
BenchmarkPostgresSimpleUpdate/bob
BenchmarkPostgresSimpleUpdate/bob          	  824893	      1565 ns/op	    1504 B/op	      50 allocs/op
BenchmarkPostgresSimpleUpdate/bob-2        	  864768	      1364 ns/op	    1504 B/op	      50 allocs/op
BenchmarkPostgresSimpleUpdate/bob-4        	  867516	      1365 ns/op	    1504 B/op	      50 allocs/op
BenchmarkPostgresSimpleUpdate/goqu
BenchmarkPostgresSimpleUpdate/goqu         	  583744	      2169 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-2       	  611674	      1895 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-4       	  628237	      1916 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/sq
BenchmarkPostgresSimpleUpdate/sq           	  841258	      1480 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-2         	  902778	      1307 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-4         	  903562	      1309 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel
BenchmarkPostgresSimpleUpdate/squirrel     	  377302	      3372 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-2   	  413792	      2901 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-4   	  412737	      2884 ns/op	    2616 B/op	      62 allocs/op
PASS
ok  	github.com/stephenafamo/sqlbuilderbenchmarks	44.999s
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
