# Golang Testability

Golang testability demo

## Golang code coverage

### list all packages

```
go list ./...
go test -v ./... 
```

### terminal coveragte report

```
go test -v -cover ./... 
```

### html coverage report
```
go test -coverprofile=z_coverage.out ./... 
go tool cover -html=z_coverage.out -o z_coverage.html
```


### inside editor with VSCode

Open a .go file to vscode editor, then

```
Shift + Command + P
Go: Toggle Test Coverage in Current Package
```

### project layout

https://github.com/golang-standards/project-layout





