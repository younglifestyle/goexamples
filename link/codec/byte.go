package codec

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"gitlab.com/link"
)

const (
	cmdSizeT     = 12
	bibInfoSizeT = 2000
	bibTestSizeT = 256
)

type ByteProtocol struct {
	Data chan interface{}
	//data []byte
}

func (b *ByteProtocol) NewCodec(rw io.ReadWriter) (link.Codec, error) {
	codec := &byteCodec{
		p: b,
		//r: rw,
		//w: rw,
		r: bufio.NewReader(rw),
		w: bufio.NewWriter(rw),
	}

	codec.closer, _ = rw.(io.Closer)
	return codec, nil
}

func Byte() *ByteProtocol {
	return &ByteProtocol{
		//data: make([]byte, 0),
		Data:make(chan interface{}, 20),
	}
}

type byteCodec struct {
	//r      io.Reader
	//w      io.Writer
	r      *bufio.Reader
	w      *bufio.Writer
	p      *ByteProtocol
	closer io.Closer
}

type tyServerCmdHead struct {
	dwTag     uint32 //固定为0x55AAAA55
	bOP       uint8  //命令码
	bOPb      uint8  //命令反码
	wDataLen  uint16 //数据长度
	wDataLenb uint16 //数据长度反码
}

func bytesToTyServerCmdHead(b []byte) *tyServerCmdHead {
	return (*tyServerCmdHead)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
}

//
func (c *byteCodec) Receive() (interface{}, error) {

	var (
		data []byte
		cnt  int
		err  error
	)

	recvData := make([]byte, 0)

	for {
		//读取指令头 返回输入流的前12个字节，不会移动读取位置
		data, err = c.r.Peek(12)
		if len(data) == 0 || err != nil {
			return nil, err
		}

		fmt.Println("come in...")

		// 读取小于12字节会报错
		//if len(data) < 12 {
		//	continue
		//}

		cmdHead := bytesToTyServerCmdHead(data)

		byteSize := 0
		// 验证固定头   验证数据大小
		if cmdHead.bOP == 0x81 {
			byteSize = bibInfoSizeT
		} else if cmdHead.bOP == 0x82 {
			byteSize = bibTestSizeT
		}

		data = make([]byte, byteSize)
		cnt, err = c.r.Read(data)

		fmt.Println(cnt, len(recvData))

		if cnt == 0 {
			return recvData, nil
		}
		if cnt != byteSize {
			return nil, errors.New("data length error")
		}
		if err != nil {
			return nil, err
		}

		c.p.Data <- recvData
	}

	// 将数据存到一起
	recvData = append(recvData, data...)

	//if len(recvData) > 100 {
	//	return recvData, nil
	//}

	return recvData, err

	//recvData := make([]byte, 4092)
	//
	//cnt, err := c.r.Read(recvData)
	//
	//return recvData[:cnt], err
}

func (c *byteCodec) Send(msg interface{}) error {

	b, ok := msg.([]byte)
	if !ok {
		return errors.New("Send Byte Format Error")
	}

	_, err := c.w.Write(b)

	c.w.Flush()
	return err
}

func (c *byteCodec) Close() error {
	if c.closer != nil {
		return c.closer.Close()
	}
	return nil
}
