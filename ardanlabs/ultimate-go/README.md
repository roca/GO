# Ultimate Go foundations

- [link Course](https://courses.ardanlabs.com/courses/take/ultimate-go-advanc-concepts/lessons/8721834-14-5-execution-tracing)

- [link2 course repo](https://github.com/ardanlabs/gotraining)

- [link3: Go source](https://github.com/golang/go)

## See escape analysis and inlining decisions

```sh
go build -gcflags -m=2
```

## Bench marks

```sh
go test -run none -bench . -benchtime 3s
```