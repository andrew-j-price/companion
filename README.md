# companion

This is my `companion` repo to my [journey](https://github.com/andrew-j-price/journey) repo to learn more about Golang packaging.
* Some practices applied may not be ideal

commmands
```bash
# run
go run .


# setup
go mod init companion
go get -d github.com/andrew-j-price/journey@v0.2.0

# package and reference upkeep
go mod tidy  # will remove `// indirect` references in `go.mod` if package is incoporated


# download reference
ls -la ~/go/pkg/mod/github.com/andrew-j-price/journey@v0.2.0/

```
