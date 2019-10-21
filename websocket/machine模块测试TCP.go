package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
	"unsafe"

	"gitlab.com/Utils/ErrCode"

	"gitlab.com/Utils/structdefine"

	"gitlab.com/link"
	"gitlab.com/link/codec"
)

const (
	BIBSNLen      = 32
	TestPortTotal = 80
)

type tyServerCmdHead struct {
	dwTag     uint32 //固定为0x55AAAA55
	bOP       uint8  //命令码
	bOPb      uint8  //命令反码
	wDataLen  uint16 //数据长度
	wDataLenb uint16 //数据长度反码
}

type tTestPortInfo struct {
	bTestCtrlUnitBootVer   [4]uint8 //测试控制单元Bootload软件版本
	bTestCtrlUnitVer       [4]uint8 //测试控制单元控制软件版本
	dwTestCtrlUnitBinCheck uint32   //测试控制单元控制SPI的bin校验值

	bDeviceScanBinVer     [4]uint8 //被测器件自扫描软件版本
	bDeviceLogicBinVer    [4]uint8 //被测器件逻辑控制软件版本
	dwDeviceLogicBinCheck uint32   //被测器件逻辑控制SPI的bin校验值
}

type tyBIBInfo struct {
	tCmd      tyServerCmdHead //命令码：0x81
	bBIBState uint8           //是否有BIB板插入：0：没有BIB，BIB已被拨出  1:有BIB，BIB已插入
	bBIBSN    [BIBSNLen]uint8 //BIB唯一序列号编码

	dwPCBinCheck uint32 //PC服务器bin的校验值

	bCommCtrlUnitBootVer   [4]uint8 //通讯控制单元Bootload软件版本
	bCommCtrlUnitVer       [4]uint8 //通讯控制单元控制软件版本
	dwCommCtrlUnitBinCheck uint32   //通讯控制控制SPI的bin校验值

	tPortInfo [TestPortTotal]tTestPortInfo

	bRecordCtrlUnitBootVer   [4]uint8 //测试记录单元Bootload软件版本
	bRecordCtrlUnitVer       [4]uint8 //测试记录单元控制软件版本
	dwRecordCtrlUnitBinCheck uint32   //测试记录控制单元控制SPI的bin校验值
	wCheck                   uint16   //数据的校验值
}

// 获取结构体真实数据的大小
var sizeOfMyStruct = int(unsafe.Sizeof(tyBIBInfo{}))

// 填充[]byte的数据结构
// 结构体的数据指针也就是一个4字节的int类型（c基础知识！）
func MyStructToBytes(s *tyBIBInfo) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

// unsafe.Pointer(&b)：取[]byte首地址
// (*reflect.SliceHeader)(unsafe.Pointer(&b)) ： 强制转换其为reflect.SliceHeader指针
// (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data ： 将slice的数据指针取出来
// unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data) : 将uint指针转成任意指针
// (*MyStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))：成功转换
func BytesToMyStruct(b []byte) *tyBIBInfo {
	return (*tyBIBInfo)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
}

func main() {
	json := codec.Json()
	json.RegisterName("structdefine/Ctrl",
		structdefine.Ctrl{})
	//server, err := link.Listen("tcp", "0.0.0.0:0", json, 0 /* sync send */, link.HandlerFunc(serverSessionLoop))
	//checkErr(err)
	//addr := server.Listener().Addr().String()
	//fmt.Println(addr)
	//go server.Serve()

	client, err := link.Dial("tcp",
		"172.16.9.229:9996", json, 0)
	checkErr(err)
	err = client.Send(&structdefine.Ctrl{
		Cmd:     "init", // update
		Mid:     13,
		Agentid: "AGENT_00FF4AAA2255",
	})
	//err = client.Send(
	//	map[string]interface{}{
	//		"Head": "gitlab.com/Utils/structdefine/Ctrl",
	//		"Body": map[string]string{"cmd": "hello"},
	//	})

	fmt.Println(err)

	for {
		rsp, err := client.Receive()
		checkErr(err)

		fmt.Println("receive :", rsp)

		ctrl := rsp.(*structdefine.Ctrl)

		fmt.Printf("%+v \n", ctrl)

		err = client.Send(&structdefine.Ctrl{
			Tag:     ctrl.Tag,
			ErrCode: ErrCode.ErrCode{0, "ok"},
		})
	}

	//clientSessionLoop(client)
}

func clientSessionLoop(session *link.Session) {
	for i := 0; i < 10; i++ {

		bibInfo := &tyBIBInfo{
			tCmd: tyServerCmdHead{
				dwTag:    0x55AAAA55,
				bOP:      0x81,
				wDataLen: 123,
			},
			bBIBState: uint8(i),
		}
		bytes := MyStructToBytes(bibInfo)

		err := session.Send(bytes)
		checkErr(err)
		log.Printf("Send: %d ", i)

		time.Sleep(time.Millisecond * 1)
		//rsp, err := session.Receive()
		//checkErr(err)
		//bytes, ok := rsp.([]byte)
		//fmt.Println(ok)
		//if ok {
		//	myStruct := BytesToMyStruct(bytes)
		//	fmt.Println(myStruct)
		//} else {
		//	fmt.Println("test")
		//}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
