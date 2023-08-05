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

## Full results: 2023-08-05

```sh
goos: darwin
goarch: arm64
pkg: github.com/stephenafamo/sqlbuilderbenchmarks
BenchmarkPostgresBulkInsert
BenchmarkPostgresBulkInsert/bob
BenchmarkPostgresBulkInsert/bob   	  605054	      1955 ns/op	    2288 B/op	      67 allocs/op
BenchmarkPostgresBulkInsert/bob-2 	  676304	      1744 ns/op	    2288 B/op	      67 allocs/op
BenchmarkPostgresBulkInsert/bob-4 	  685435	      1801 ns/op	    2288 B/op	      67 allocs/op
BenchmarkPostgresBulkInsert/goqu
BenchmarkPostgresBulkInsert/goqu           	  425469	      2701 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-2         	  441964	      3115 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/goqu-4         	  511267	      2344 ns/op	    2912 B/op	      82 allocs/op
BenchmarkPostgresBulkInsert/sq
BenchmarkPostgresBulkInsert/sq             	  543387	      2232 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-2           	  607789	      1928 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/sq-4           	  595689	      1944 ns/op	    3264 B/op	      41 allocs/op
BenchmarkPostgresBulkInsert/squirrel
BenchmarkPostgresBulkInsert/squirrel       	  216866	      5688 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-2     	  242967	      4818 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/squirrel-4     	  243682	      4795 ns/op	    4376 B/op	     106 allocs/op
BenchmarkPostgresBulkInsert/huandu
BenchmarkPostgresBulkInsert/huandu         	  702351	      1905 ns/op	    1576 B/op	      39 allocs/op
BenchmarkPostgresBulkInsert/huandu-2       	  757645	      1584 ns/op	    1576 B/op	      39 allocs/op
BenchmarkPostgresBulkInsert/huandu-4       	  706246	      1563 ns/op	    1576 B/op	      39 allocs/op
BenchmarkPostgresSimpleSelect
BenchmarkPostgresSimpleSelect/bob
BenchmarkPostgresSimpleSelect/bob          	  635319	      1939 ns/op	    2176 B/op	      60 allocs/op
BenchmarkPostgresSimpleSelect/bob-2        	  680466	      1747 ns/op	    2176 B/op	      60 allocs/op
BenchmarkPostgresSimpleSelect/bob-4        	  681163	      1727 ns/op	    2176 B/op	      60 allocs/op
BenchmarkPostgresSimpleSelect/goqu
BenchmarkPostgresSimpleSelect/goqu         	  375877	      3134 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-2       	  407896	      2821 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/goqu-4       	  410548	      2799 ns/op	    3488 B/op	      87 allocs/op
BenchmarkPostgresSimpleSelect/sq
BenchmarkPostgresSimpleSelect/sq           	  534285	      2275 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-2         	  587409	      1989 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/sq-4         	  593128	      1987 ns/op	    2984 B/op	      43 allocs/op
BenchmarkPostgresSimpleSelect/squirrel
BenchmarkPostgresSimpleSelect/squirrel     	  343189	      3546 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-2   	  375693	      3111 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/squirrel-4   	  376072	      3090 ns/op	    2776 B/op	      61 allocs/op
BenchmarkPostgresSimpleSelect/huandu
BenchmarkPostgresSimpleSelect/huandu       	 1201245	       998.2 ns/op	     816 B/op	      25 allocs/op
BenchmarkPostgresSimpleSelect/huandu-2     	 1429120	       838.8 ns/op	     816 B/op	      25 allocs/op
BenchmarkPostgresSimpleSelect/huandu-4     	 1420030	       840.3 ns/op	     816 B/op	      25 allocs/op
BenchmarkPostgresSimpleUpdate
BenchmarkPostgresSimpleUpdate/bob
BenchmarkPostgresSimpleUpdate/bob          	  864763	      1478 ns/op	    1456 B/op	      42 allocs/op
BenchmarkPostgresSimpleUpdate/bob-2        	  914834	      1283 ns/op	    1456 B/op	      42 allocs/op
BenchmarkPostgresSimpleUpdate/bob-4        	  914828	      1283 ns/op	    1456 B/op	      42 allocs/op
BenchmarkPostgresSimpleUpdate/goqu
BenchmarkPostgresSimpleUpdate/goqu         	  575066	      2159 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-2       	  611200	      1928 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/goqu-4       	  620728	      1906 ns/op	    2192 B/op	      57 allocs/op
BenchmarkPostgresSimpleUpdate/sq
BenchmarkPostgresSimpleUpdate/sq           	  875047	      1425 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-2         	  943785	      1242 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/sq-4         	  944354	      1235 ns/op	    1920 B/op	      29 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel
BenchmarkPostgresSimpleUpdate/squirrel     	  392042	      3260 ns/op	    2648 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-2   	  418280	      2784 ns/op	    2648 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/squirrel-4   	  416548	      2766 ns/op	    2648 B/op	      62 allocs/op
BenchmarkPostgresSimpleUpdate/huandu
BenchmarkPostgresSimpleUpdate/huandu       	 6102924	       196.3 ns/op	     200 B/op	       8 allocs/op
BenchmarkPostgresSimpleUpdate/huandu-2     	 6643478	       183.3 ns/op	     200 B/op	       8 allocs/op
BenchmarkPostgresSimpleUpdate/huandu-4     	 6638896	       180.4 ns/op	     200 B/op	       8 allocs/op
PASS
ok  	github.com/stephenafamo/sqlbuilderbenchmarks	60.960s
```

## Queries

### Simple Select

```sql
---Bob---
Query: SELECT
id, name
FROM users
WHERE ("id" IN ($1, $2, $3))
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
