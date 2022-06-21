# GoLang

### Commands
Execute a go file:
```sh
go run main.go
```

Execute a go program:
```sh
go run *.go
```

Documentation:
```sh
go doc fmt
```

Check program Health:
```sh
go vet main.go
```

Build a Program:
```sh
go build main.go
```

Install a package:
```sh
go install -u github.com/go-sql-driver/mysql
```

Tests Coverage:
```sh
go test -cover
```

Tests Coverage per function:
```sh
go tool cover -func=result.out
```

Tests Coverage Result:
```sh
go test -coverprofile=result.out
go tool cover -html=result.out
```