package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// bufio 包实现了带缓存的 I/O 操作

/**
 * 首先看reader和writer基本的结构
 * // Reader implements buffering for an io.Reader object.
 * type Reader struct {
 *	buf          []byte
 *	rd           io.Reader // reader provided by the client
 *	r, w         int       // buf read and write positions
 *	err          error
 *	lastByte     int
 *	lastRuneSize int
 * }
 *
 *
 * // Writer implements buffering for an io.Writer object.
 * // If an error occurs writing to a Writer, no more data will be
 * // accepted and all subsequent writes will return the error.
 * // After all data has been written, the client should call the
 * // Flush method to guarantee all data has been forwarded to
 * // the underlying io.Writer.
 * type Writer struct {
 *	err error
 *	buf []byte
 *	n   int
 *	wr  io.Writer
 * }
 *
 *
 *
 * // ReadWriter 集成了 bufio.Reader 和 bufio.Writer, 实现了 io.ReadWriter 接口
 * type ReadWriter struct {
 *	*Reader
 *	*Writer
 * }
 */

func main() {

	buf := bytes.NewBuffer(make([]byte, 1000))

	// 1: 使用bufio.NewReader构造一个reader
	reader := bufio.NewReader(buf)

	// 2: 使用bufio.NewWriter构造一个writer
	writer := bufio.NewWriter(buf)

	writer.Write([]byte{1, 23, 4, 4, 5, 5, 6})

	// 3: 函数Peek函数: 返回缓存的一个Slice(引用,不是拷贝)，引用缓存中前n字节数据
	// > 如果引用的数据长度小于 n，则返回一个错误信息
	// > 如果 n 大于缓存的总大小，则返回 ErrBufferFull
	// 通过Peek的返回值，可以修改缓存中的数据, 但不能修改底层io.Reader中的数据
	b, err := reader.Peek(5)
	if err != nil {
		fmt.Printf("Read data error")
		return
	}
	// 修改第一个字符
	b[0] = 'A'
	// 重新读取
	b, _ = reader.Peek(5)
	writer.Write(b)
	writer.Flush()

	// 4: Read函数, 每次读取一定量的数据, 这个由buf大小觉得, 所以我们可以循环读取数据, 直到Read返回0说明读取数据结束
	for {
		b1 := make([]byte, 3)
		n1, _ := reader.Read(b1)
		if n1 <= 0 {
			break
		}
		fmt.Println(n1, string(b1))
	}
}
