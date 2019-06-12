package codec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/link"
)

type Msg struct {
	Field1 string
	Field2 int
}

func ByteTestProtocol() *ByteProtocol {
	protocol := Byte()
	return protocol
}

func ByteTest(t *testing.T, protocol link.Protocol) {
	var stream bytes.Buffer

	codec, _ := protocol.NewCodec(&stream)

	sendMsg1 := Msg{
		Field1: "abc",
		Field2: 123,
	}

	bytes1, _ := json.Marshal(sendMsg1)

	err := codec.Send(bytes1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(bytes1)

	recvMsg1, err := codec.Receive()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(recvMsg1, err)

	b, _ := recvMsg1.([]byte)

	if string(bytes1) != string(b) {
		t.Fatalf("message not match: %v, %v", sendMsg1, recvMsg1)
	}
}

func Test_Byte(t *testing.T) {
	protocol := ByteTestProtocol()
	ByteTest(t, protocol)
}
