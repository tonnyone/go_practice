module thrift

go 1.18

require (
	github.com/apache/thrift v0.18.1
	github.com/apache/thrift/tutorial/go/gen-go/shared v0.0.0-00010101000000-000000000000
	github.com/apache/thrift/tutorial/go/gen-go/tutorial v0.0.0-00010101000000-000000000000
)

replace github.com/apache/thrift/tutorial/go/gen-go/shared => ./gen-go/shared

replace github.com/apache/thrift/tutorial/go/gen-go/tutorial => ./gen-go/tutorial
