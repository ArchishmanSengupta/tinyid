##  tinyid
A super tiny, high performant fast Go unique string ID generator:

- *Safe*: Utilizes cryptographically strong random APIs to ensure security.
- *Compact*: Employs a larger alphabet than UUID ```(A-Za-z0-9_-)```, potentially reducing ID size from 36 to 31.
- *Efficient*: Benchmarks at approximately ```489.3 ns/op``` for generating IDs.
- *100% TC*: Achieves full test coverage.
- *Small*: <80 LoC.
## Benchmark:
```
goos: darwin
goarch: arm64
pkg: github.com/ArchishmanSengupta/tinyid
BenchmarkNewTinyID-10         	 2419586	       489.3 ns/op
BenchmarkGenerateTinyID-10    	 2457799	       489.7 ns/op
PASS
ok  	github.com/ArchishmanSengupta/tinyid	3.535s
```

## Test Coverage:
```
go test -v
=== RUN   TestGenerateTinyID
--- PASS: TestGenerateTinyID (0.00s)
=== RUN   TestNewTinyID
--- PASS: TestNewTinyID (0.00s)
PASS
ok  	github.com/ArchishmanSengupta/tinyid	0.420s
```

## Use:
```go
id, err := tinyid.NewTinyID()
	if err != nil {
		fmt.Print("some problem")
}
fmt.Println("Generated ID:", id)
```
