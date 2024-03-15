module gen-go/tutorial

go 1.18

replace github.com/apache/thrift/tutorial/go/gen-go/shared => ../shared

replace github.com/apache/thrift/tutorial/go/gen-go/tutorial => ../tutorial

require (
	github.com/apache/thrift v0.18.1
	github.com/apache/thrift/tutorial/go/gen-go/shared v0.0.0-00010101000000-000000000000
)
