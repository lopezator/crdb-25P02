# crdb-25P02

How to reproduce the error:

1. Initialize database:

$> make db-init

2. Reproduce the error, execute the tests a few times to reproduce the error:

$> go test -v -count=1 ./...