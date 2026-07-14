# Slide:
    - https://www.dropbox.com/s/ylgspq86se4fm9s/AG%20-%20A%20De%20Sarker%20-%20Common%20Patterns%20for%20Bounds%20Check%20Elimination.pdf?dl=0


## Check boundaries
```
    $go build -gcflags=”-d=ssa/check_bce/debug=1” foo_test.go
    OR
    $go tool compile -d=ssa/check_bce/debug=1 foo_test.go
```

## Run Benchmarks
    - go test -bench=BenchmarkSum
   