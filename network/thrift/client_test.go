package thrift

import (
	"github.com/apache/thrift/lib/go/thrift"
	"testing"
)

func TestRunClient(t *testing.T) {

	protocolFactory := thrift.NewTJSONProtocolFactory()
	transportFactory := thrift.NewTTransportFactory()
	if err := runClient(transportFactory, protocolFactory, "localhost:9090", false, nil); err != nil {
		t.Log("error running client:", err)
	}
}
