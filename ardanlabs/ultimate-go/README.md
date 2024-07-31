# Ultimate Go foundations

- [link Course](https://courses.ardanlabs.com/courses/take/ultimate-go-advanc-concepts/lessons/7419523-5-2-1-decoupling-part-1)

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
