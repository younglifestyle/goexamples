package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"log"
)

/*
 * 1、良率 = 良品数量 / 测试数量 * 100%——该良率统计时良品数量须剔除重复UID统计，
则每一个UID以最后1次测试（测试时间最晚）的测试结果为准进行统计计算；
每一行表示一次测试，BIN1/BIN2表示良品，
其他为不良品；——仅有BIN1为良品，BIN2、BIN3、BIN4为不良品；
 * 2、统计BIN1, BIN2, BIN4的数量及百分比，以及非BIN1类别的errcode占比。
——BIN别统计希望涵盖工具定义的所有BIN别，BIN1、BIN2、BIN3、BIN4全部都进行对应统计,
以防有出现BIN3没有统计到的情况。
 * 3、分别统计块坏数量、坏块的block number出现的所占比例（或者统计每个blk是坏块的比例），
最后形成一个坏块的blk number 占比分布图呈现坏块的大致分布情况。
 * 4．HY 2Gb 需分别按照ECC卡0bit、ECC卡2bit这 2种情况进行以上数据统计，
形成数据结果对比。
*/
// 0 UID    2 bin  3  【error code】
// 5       6        7     8
// 原厂数   新增坏块   原厂   新增坏块

const (
	fileName = "logAnalysisResult.csv"
)

var (
	count  int64
	blkMap map[string]*analysisTestInfo
)

type analysisTestInfo struct {
	binStr   string
	errStr   string
	testTime int64
	FacN     int
	lowN     int
	FacNo    *[]string
	LowNo    *[]string
}

func oldCode() {
	//flag.StringVar(&logFileDir, "logpath", `.\logpath`, "log文件的目录")
	//flag.Parse()
	//fmt.Println("start analysis the log in ", `"`+logFileDir+`"`)
	//
	//badBlkMap = make(map[string]int)
	//
	//err := analyzeLog(logFileDir)
	//if err != nil {
	//	fmt.Println("error :", err)
	//	return
	//}
	//
	//marshal, _ := json.Marshal(badBlkMap)
	//fmt.Println(string(marshal))
	//fmt.Printf("\r\n良品中各个坏块的占比 :\r\n")
	//fmt.Printf("blk0: %f%%, blk2: %f%%, "+
	//	"blk4: %f%%, blk6: %f%%, "+
	//	"blk8: %f%%,  blk10: %f%%, "+
	//	"blk12: %f%%, blk14: %f%% \r\n",
	//	float64(float64(badBlkMap["blk0"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk2"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk4"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk6"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk8"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk10"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk12"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk14"]*100)/float64(badBlkMap["count"])))
	//fmt.Printf("\r\n良品记录坏块的占比(记录单次) :\r\n")
	//fmt.Printf("blk: %f%% \r\n",
	//	float64(float64(badBlkMap["match_bad_blk"]*100)/float64(badBlkMap["count"])))
	//
	//f, err := os.OpenFile(fileName,
	//	os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	//if err != nil {
	//	return
	//}
	//w := bufio.NewWriter(f)
	//fmt.Fprintln(w, fmt.Sprintf("start analysis log at the %s \r\n", `"`+logFileDir+`"`))
	//fmt.Fprintln(w, fmt.Sprintf("记录总数 : %d，良品共有%d片芯片出现坏块，占比 %f%% \r\n",
	//	badBlkMap["count"], badBlkMap["match_bad_blk"],
	//	float64(float64(badBlkMap["match_bad_blk"]*100)/float64(badBlkMap["count"]))))
	//fmt.Fprintln(w, fmt.Sprintf("良品中各个坏块的占比 :"))
	//fmt.Fprintln(w, fmt.Sprintf("blk0: %f%%, blk2: %f%%, "+
	//	"blk4: %f%%, blk6: %f%%, "+
	//	"blk8: %f%%,  blk10: %f%%, "+
	//	"blk12: %f%%, blk14: %f%% \r\n",
	//	float64(float64(badBlkMap["blk0"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk2"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk4"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk6"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk8"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk10"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk12"]*100)/float64(badBlkMap["count"])),
	//	float64(float64(badBlkMap["blk14"]*100)/float64(badBlkMap["count"]))))
	//
	////fmt.Fprintln(w, fmt.Sprintf("良品记录坏块的占比(记录单次) :"))
	////fmt.Fprintln(w, fmt.Sprintf("blk: %f%% \r\n",
	////	float64(float64(badBlkMap["match_bad_blk"]*100)/float64(badBlkMap["count"]))))
	//w.Flush()
}

func main() {

	fmt.Println("开始进行log分析")

	fileTarget, err := listAll(`.\logpath`)
	if err != nil {
		log.Fatal(`error,`, err)
	}
	if len(fileTarget) == 0 {
		log.Fatal(`当前logpath目录下没有文件夹`)
	}

	// Create a csv file
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal("创建文件夹失败！, " + fileName + " 该文件或被占用")
		return
	}
	defer f.Close()

	// 写入UTF-8 BOM，防止中文乱码
	f.WriteString("\xEF\xBB\xBF")
	wsv := csv.NewWriter(f)

	for _, filePaths := range fileTarget {
		blkMap = make(map[string]*analysisTestInfo)
		analyzeLog(filePaths)
		err := spritfWriteFile(filePaths, wsv)
		if err != nil {
			log.Printf("分析文件夹：%s，出现错误\r\n", filePaths)
		}
		count = 0
	}

	wsv.Flush()

	fmt.Println("完成log分析")
}

func spritfWriteFile(logFileDir string, wsvRes *csv.Writer) error {

	var (
		bin1Count int
		bin2Count int
		bin3Count int
		bin4Count int
		errCount  int
	)

	//binMap = make(map[string]map[string]int)
	errMap := make(map[string]int)
	allBinErrMap := make(map[string]int)
	facNMap := make(map[string]int)
	lowNMap := make(map[string]int)

	wsvRes.Write([]string{""})
	wsvRes.Write([]string{fmt.Sprintf("start analysis log at the %s", logFileDir)})

	for _, info := range blkMap {
		switch info.binStr {
		case "BIN1":
			bin1Count = bin1Count + 1
			allBinErrMap[info.errStr] = allBinErrMap[info.errStr] + 1
		case "BIN2":
			bin2Count = bin2Count + 1
			errMap[info.errStr] = errMap[info.errStr] + 1
			allBinErrMap[info.errStr] = allBinErrMap[info.errStr] + 1
			errCount = errCount + 1
		case "BIN3":
			bin3Count = bin3Count + 1
			errMap[info.errStr] = errMap[info.errStr] + 1
			allBinErrMap[info.errStr] = allBinErrMap[info.errStr] + 1
			errCount = errCount + 1
		case "BIN4":
			bin4Count = bin4Count + 1
			errMap[info.errStr] = errMap[info.errStr] + 1
			allBinErrMap[info.errStr] = allBinErrMap[info.errStr] + 1
			errCount = errCount + 1
		}

		if info.FacN != 0 {
			for _, facNo := range *info.FacNo {
				facNMap[facNo] = facNMap[facNo] + 1
			}
		}
		if info.lowN != 0 {
			for _, lowNo := range *info.LowNo {
				lowNMap[lowNo] = lowNMap[lowNo] + 1
			}
		}
	}

	wsvRes.WriteAll([][]string{{"总测试数量", strconv.FormatInt(count, 10)}, {}})

	wsvRes.WriteAll([][]string{{"BIN类型占比数量"},
		{"BIN别", "数量", "占比"},
		{"bin1", strconv.FormatInt(int64(bin1Count), 10),
			fmt.Sprintf("%f%%", float64(float64(bin1Count*100)/float64(count)))},
		{"bin2", strconv.FormatInt(int64(bin2Count), 10),
			fmt.Sprintf("%f%%", float64(float64(bin2Count*100)/float64(count)))},
		{"bin3", strconv.FormatInt(int64(bin3Count), 10),
			fmt.Sprintf("%f%%", float64(float64(bin3Count*100)/float64(count)))},
		{"bin4", strconv.FormatInt(int64(bin4Count), 10),
			fmt.Sprintf("%f%%", float64(float64(bin4Count*100)/float64(count)))},
		{},
		{"非bin1类别errcode总数", strconv.FormatInt(int64(errCount), 10)},
		{"errorCode", "数量", "占比"},
	})

	// error的分布
	for errStr, errCnt := range errMap {
		wsvRes.Write([]string{errStr, strconv.FormatInt(int64(errCnt), 10),
			fmt.Sprintf("%f%%", float64(float64(errCnt*100)/float64(errCount)))})
	}

	wsvRes.WriteAll([][]string{{}, {"包含bin1类别errcode总数"},
		{"errorCode", "数量", "占比"},
	})
	for errStr, errCnt := range allBinErrMap {
		wsvRes.Write([]string{errStr, strconv.FormatInt(int64(errCnt), 10),
			fmt.Sprintf("%f%%", float64(float64(errCnt*100)/float64(count)))})
	}

	// 导入excel
	split := strings.Split(logFileDir, `\`)
	facNFileName := split[len(split)-1] + "_facN.csv"
	// Create a csv file
	f, err := os.OpenFile(facNFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	wsv := csv.NewWriter(f)
	// write csv
	for facNo, count := range facNMap {
		var record []string
		record = append(record, facNo, strconv.FormatInt(int64(count), 10))
		wsv.Write(record)
	}
	wsv.Flush()

	lowNFileName := split[len(split)-1] + "_lowN.csv"
	// Create a csv file
	f, err = os.OpenFile(lowNFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	wsv = csv.NewWriter(f)
	// write csv
	for lowNo, count := range lowNMap {
		var record []string
		record = append(record, lowNo, strconv.FormatInt(int64(count), 10))
		wsv.Write(record)
	}
	wsv.Flush()

	wsvRes.WriteAll([][]string{{},
		{fmt.Sprintf("将原厂坏块信息导入CSV文件: %s", facNFileName)},
		{fmt.Sprintf("将新增坏块信息导入CSV文件: %s", lowNFileName)},
	})

	return nil
}

func analyzeLog(logFileDir string) (err error) {

	filepath.Walk(logFileDir, func(pathStr string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(pathStr)
		if err != nil {
			return err
		}
		defer file.Close()

		err = analyzeLine(file)
		//_, err = io.Copy(writer, file)
		return err
	})

	return
}

func analyzeLine(file *os.File) (err error) {

	var (
		trim         string
		tm2          time.Time
		timeStamp    int64
		firstLine    bool
		needAnalyStr string
	)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				return nil
			}
			return err
		}
		logSplit := strings.Split(line, "\t")
		if len(logSplit) < 13 || len(logSplit) > 15 {
			continue
		}

		// 第一行不需要
		if !firstLine {
			firstLine = true
			continue
		}

		trim = strings.Trim(fmt.Sprint(logSplit[4]), "[]")
		tm2, _ = time.ParseInLocation("2006-01-02 15:04:05.0000",
			trim[:len(logSplit[4])-7]+"."+trim[len(logSplit[4])-6:], time.Local)
		timeStamp = tm2.UnixNano() / 1000000

		info, ok := blkMap[logSplit[0]]
		if ok {
			if info.testTime > timeStamp {
				continue
			}
			info.binStr = logSplit[2]
			info.errStr = logSplit[3]
			FacN, _ := strconv.Atoi(logSplit[5])
			lowN, _ := strconv.Atoi(logSplit[6])
			info.FacN = FacN
			info.lowN = lowN
			if info.FacN != 0 {
				needAnalyStr = strings.TrimRight(logSplit[7], " ;")
				facNoStr := strings.Split(needAnalyStr, " ")
				info.FacNo = &facNoStr
			}
			if info.lowN != 0 {
				needAnalyStr = strings.TrimRight(logSplit[8], " ;")
				lowNoStr := strings.Split(needAnalyStr, " ")
				info.LowNo = &lowNoStr
			}
			blkMap[logSplit[0]] = info
		} else {
			info = &analysisTestInfo{}
			info.testTime = timeStamp
			info.binStr = logSplit[2]
			info.errStr = logSplit[3]
			FacN, _ := strconv.Atoi(logSplit[5])
			lowN, _ := strconv.Atoi(logSplit[6])
			info.FacN = FacN
			info.lowN = lowN
			if info.FacN != 0 {
				needAnalyStr = strings.TrimRight(logSplit[7], " ;")
				facNoStr := strings.Split(needAnalyStr, " ")
				info.FacNo = &facNoStr
			}
			if info.lowN != 0 {
				needAnalyStr = strings.TrimRight(logSplit[8], " ;")
				lowNoStr := strings.Split(needAnalyStr, " ")
				info.LowNo = &lowNoStr
			}
			blkMap[logSplit[0]] = info
			count = count + 1
		}
	}

	return nil
}

func listAll(path string) (fileTarget []string, err error) {

	fileTarget = make([]string, 0)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fi := range files {

		if fi.IsDir() {
			fileTarget = append(fileTarget, filepath.Join(path, fi.Name()))
		}
	}
	return fileTarget, err
}
