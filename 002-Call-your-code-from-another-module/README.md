## Go ენის მოდულების შექმნა და დაკავშირება
```bash
go mod edit -replace example.com/greetings=../greetings
```

```go
module example.com/hello
replace example.com/greetings => ../greetings
```

```bash
go mod tidy
```

**go:** found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```go


module example.com/hello

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000

require example.com/greetings v1.1.0
```
### ***კოდის გაშვება***
```bash
go run .
```
output
```
Hi, Gladys. Welcome!
```
