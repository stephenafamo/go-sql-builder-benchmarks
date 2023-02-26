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
* [Huandu](https://github.com/huandu/go-sqlbuilder)

## Help improve this

Kindly send in PRs with new query builders or queries to benchmark.

## Full results: 2022-07-28

```sh
goos: darwin
goarch: arm64
pkg: github.com/stephenafamo/sqlbuilderbenchmarks
BenchmarkPostgresBulkInsert
BenchmarkPostgresBulkInsert/bob
BenchmarkPostgresBulkInsert/bob   	  574155	      2050 ns/op	    2320 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-2 	  630021	      1830 ns/op	    2320 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/bob-4 	  642820	      1835 ns/op	    2320 B/op	      77 allocs/op
BenchmarkPostgresBulkInsert/goqu
BenchmarkPostgresBulkInsert/goqu           	  477339	      2607 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-2         	  501019	      2317 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-4         	  508168	      2317 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/sq
BenchmarkPostgresBulkInsert/sq             	  514275	      2384 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-2           	  565558	      2056 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-4           	  574106	      2060 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/squirrel
BenchmarkPostgresBulkInsert/squirrel       	  213913	      5763 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-2     	  232371	      5000 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-4     	  234081	      4976 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/huandu
BenchmarkPostgresBulkInsert/huandu         	  705835	      1882 ns/op	    1688 B/op	      37 allocs/op
BenchmarkPostgresBulkInsert/huandu-2       	  747028	      1574 ns/op	    1688 B/op	      37 allocs/op
BenchmarkPostgresBulkInsert/huandu-4       	  748208	      1580 ns/op	    1688 B/op	      37 allocs/op
BenchmarkPostgresSimpleSelect
BenchmarkPostgresSimpleSelect/bob
BenchmarkPostgresSimpleSelect/bob          	  684865	      1811 ns/op	    1968 B/op	      58 allocs/op
BenchmarkPostgresSimpleSelect/bob-2        	  720649	      1618 ns/op	    1968 B/op	      58 allocs/op
BenchmarkPostgresSimpleSelect/bob-4        	  728940	      1619 ns/op	    1968 B/op	      58 allocs/op
BenchmarkPostgresSimpleSelect/goqu
BenchmarkPostgresSimpleSelect/goqu         	  371544	      3148 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-2       	  407264	      2794 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-4       	  409846	      2793 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/sq
BenchmarkPostgresSimpleSelect/sq           	  501062	      2462 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-2         	  536715	      2171 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-4         	  543694	      2184 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/squirrel
BenchmarkPostgresSimpleSelect/squirrel     	  332554	      3666 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-2   	  368530	      3232 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-4   	  357468	      3233 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/huandu
BenchmarkPostgresSimpleSelect/huandu       	 1000000	      1002 ns/op	     840 B/op	      23 allocs/op
BenchmarkPostgresSimpleSelect/huandu-2     	 1394989	       861.5 ns/op	     840 B/op	      23 allocs/op
BenchmarkPostgresSimpleSelect/huandu-4     	 1391994	       859.4 ns/op	     840 B/op	      23 allocs/op
BenchmarkPostgresSimpleUpdate
BenchmarkPostgresSimpleUpdate/bob
BenchmarkPostgresSimpleUpdate/bob          	  809067	      1580 ns/op	    1520 B/op	      50 allocs/op
BenchmarkPostgresSimpleUpdate/bob-2        	  851959	      1386 ns/op	    1520 B/op	      50 allocs/op
BenchmarkPostgresSimpleUpdate/bob-4        	  843014	      1384 ns/op	    1520 B/op	      50 allocs/op
BenchmarkPostgresSimpleUpdate/goqu
BenchmarkPostgresSimpleUpdate/goqu         	  571532	      2170 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-2       	  609662	      1918 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-4       	  621321	      1915 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/sq
BenchmarkPostgresSimpleUpdate/sq           	  805285	      1554 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-2         	  859828	      1376 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-4         	  886830	      1366 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel
BenchmarkPostgresSimpleUpdate/squirrel     	  369390	      3370 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-2   	  397760	      2907 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-4   	  400417	      2908 ns/op	    2616 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/huandu
BenchmarkPostgresSimpleUpdate/huandu       	 5548072	       215.8 ns/op	     360 B/op	       8 allocs/op
BenchmarkPostgresSimpleUpdate/huandu-2     	 5880306	       202.5 ns/op	     360 B/op	       8 allocs/op
BenchmarkPostgresSimpleUpdate/huandu-4     	 5894546	       202.3 ns/op	     360 B/op	       8 allocs/op
PASS
ok  	github.com/stephenafamo/sqlbuilderbenchmarks	57.843s
```

## Queries

### Simple Select

```sql
---Bob---
Query: SELECT
id, name
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

---Huandu---
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

---Huandu---
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

---Huandu---
Query: INSERT INTO users (first_name,last_name) VALUES ($1,$2),($3,$4),($5,$6)
Args: []interface {}{"Greg", "Farley", "Jimmy", "Stewart", "Jeff", "Jeffers"}
```
