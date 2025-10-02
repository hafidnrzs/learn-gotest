Check test coverage

```bash
go test -cover ./
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test-Driven Development Tips

1. Write tests first
2. Keep tests focused
3. Use clear test names
4. Make tests independent
5. Avoid test dependencies
6. Update tests with code changes

References:
- https://www.ceos3c.com/golang/go-unit-testing-a-practical-guide-for-writing-reliable-tests/
- https://dev.to/zpeters/testing-in-go-with-table-drive-tests-and-testify-kd4
- https://medium.com/@ullauri.byron/using-httptest-server-in-go-to-mock-and-test-external-api-calls-68ce444cf934

## Project: To-Do List API

How to run the tests:
```bash
go test ./todolist -v
```

Run the actual server:
```bash
go run ./todolist
```