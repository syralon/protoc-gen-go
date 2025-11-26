package proto

//go:generate protoc -I . --go_out=paths=source_relative:. ./syralon/errors/*.proto
//go:generate protoc -I . --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. ./syralon/http/*.proto
