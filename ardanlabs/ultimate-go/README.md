# Ultimate Go foundations

- [link Course](hhttps://courses.ardanlabs.com/courses/take/ultimate-go-advanc-concepts/lessons/7655374-intro-goroutines)

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
