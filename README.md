# learn-go

## Useful stuff

```sh
# Install errcheck linting
go get -u github.com/kisielk/errcheck

# Will find missing test cases in your test
errcheck .
```

[Don't just check errors, handle them](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

[Mutex vs Channel](https://github.com/golang/go/wiki/MutexOrChannel)

```sh
# Run test and check race conditions
go test -race
```
