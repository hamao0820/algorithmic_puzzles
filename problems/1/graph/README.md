```sh
go run main.go > graph.dot
```

```sh
go install github.com/goccy/go-graphviz/cmd/dot@latest
```

```sh
dot -T png -o graph.png graph.dot
```
