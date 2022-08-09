# sharding
**sharding** is a universal library that rewrite SQL into a sharded form.

It provide simply interface that different database client can use.
```go
interface {
    Rewrite(sql string) string
    ...
}

// e.g.
// input: select t.c1, t.c2, t.c3 from t where t.c1=? and t.c2>?;
// output: select t_0001.c1, t_0001.c2, t_0001.c3 from t_0001 where t_0001.c1=? and t_0001.c2>?;
```