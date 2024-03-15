package thrift

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	protocolFactory := thrift.NewTSimpleJSONProtocolFactoryConf(nil)
	transportFactory := thrift.NewTTransportFactory()
	if err := runServer(transportFactory, protocolFactory, "localhost:9090", false); err != nil {
		fmt.Println("error running server:", err)
	}
}
