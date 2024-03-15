package thrift

import (
	"github.com/apache/thrift/lib/go/thrift"
	"testing"
)

func TestRunServer(t *testing.T) {

	protocolFactory := thrift.NewTJSONProtocolFactory()
	transportFactory := thrift.NewTTransportFactory()
	if err := runServer(transportFactory, protocolFactory, "localhost:9090", false); err != nil {
		t.Log("error running server:", err)
	}
}
