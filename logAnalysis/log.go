package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	fileName = "logAnalysisResult.txt"
	blk0     = "0000"
	blk2     = "0002"
	blk4     = "0004"
	blk6     = "0006"
	blk8     = "0008"
	blk10    = "0010"
	blk12    = "0012"
	blk14    = "0014"
)

var (
	badBlkMap  map[string]int
	logFileDir string
)

func main() {
	flag.StringVar(&logFileDir, "logpath", `.\logpath`, "log文件的目录")
	flag.Parse()
	fmt.Println("start analysis the log in ", `"`+logFileDir+`"`)

	badBlkMap = make(map[string]int)

	err := analyzeLog(logFileDir)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	marshal, _ := json.Marshal(badBlkMap)
	fmt.Println(string(marshal))
	fmt.Printf("\r\n良品中各个坏块的占比 :\r\n")
	fmt.Printf("blk0: %f%%, blk2: %f%%, "+
		"blk4: %f%%, blk6: %f%%, "+
		"blk8: %f%%,  blk10: %f%%, "+
		"blk12: %f%%, blk14: %f%% \r\n",
		float64(float64(badBlkMap["blk0"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk2"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk4"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk6"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk8"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk10"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk12"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk14"]*100)/float64(badBlkMap["count"])))
	fmt.Printf("\r\n良品记录坏块的占比(记录单次) :\r\n")
	fmt.Printf("blk: %f%% \r\n",
		float64(float64(badBlkMap["match_bad_blk"]*100)/float64(badBlkMap["count"])))

	f, err := os.OpenFile(fileName,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, fmt.Sprintf("start analysis log at the %s \r\n", `"`+logFileDir+`"`))
	fmt.Fprintln(w, fmt.Sprintf("记录总数 : %d，良品共有%d片芯片出现坏块，占比 %f%% \r\n",
		badBlkMap["count"], badBlkMap["match_bad_blk"],
		float64(float64(badBlkMap["match_bad_blk"]*100)/float64(badBlkMap["count"]))))
	fmt.Fprintln(w, fmt.Sprintf("良品中各个坏块的占比 :"))
	fmt.Fprintln(w, fmt.Sprintf("blk0: %f%%, blk2: %f%%, "+
		"blk4: %f%%, blk6: %f%%, "+
		"blk8: %f%%,  blk10: %f%%, "+
		"blk12: %f%%, blk14: %f%% \r\n",
		float64(float64(badBlkMap["blk0"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk2"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk4"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk6"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk8"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk10"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk12"]*100)/float64(badBlkMap["count"])),
		float64(float64(badBlkMap["blk14"]*100)/float64(badBlkMap["count"]))))

	//fmt.Fprintln(w, fmt.Sprintf("良品记录坏块的占比(记录单次) :"))
	//fmt.Fprintln(w, fmt.Sprintf("blk: %f%% \r\n",
	//	float64(float64(badBlkMap["match_bad_blk"]*100)/float64(badBlkMap["count"]))))
	w.Flush()
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

		badBlkMap["count"] = badBlkMap["count"] + 1

		if logSplit[2] == "BIN1" {
			if logSplit[5] != "0" {
				needAnalyStr = strings.TrimRight(logSplit[7], " ;")
			}
			if logSplit[6] != "0" {
				needAnalyStr += " " + strings.TrimRight(logSplit[8], " ;")
			}
			if logSplit[12] != "0" {
				splitArr := strings.Split(strings.TrimRight(logSplit[13],
					" ;"), " ")
				for _, matchNo := range splitArr {
					if !strings.Contains(needAnalyStr, matchNo) {
						needAnalyStr += " " + matchNo
					}
				}
			}

			matchBlkNo(badBlkMap, needAnalyStr)
		}
	}
}

func matchBlkNo(badBlkMaps map[string]int, needMatchs string) {
	var isMatch bool
	if strings.Contains(needMatchs, blk0) {
		badBlkMaps["blk0"] = badBlkMaps["blk0"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk2) {
		badBlkMaps["blk2"] = badBlkMaps["blk2"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk4) {
		badBlkMaps["blk4"] = badBlkMaps["blk4"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk6) {
		badBlkMaps["blk6"] = badBlkMaps["blk6"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk8) {
		badBlkMaps["blk8"] = badBlkMaps["blk8"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk10) {
		badBlkMaps["blk10"] = badBlkMaps["blk10"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk12) {
		badBlkMaps["blk12"] = badBlkMaps["blk12"] + 1
		isMatch = true
	}
	if strings.Contains(needMatchs, blk14) {
		badBlkMaps["blk14"] = badBlkMaps["blk14"] + 1
		isMatch = true
	}

	if isMatch {
		badBlkMaps["match_bad_blk"] = badBlkMaps["match_bad_blk"] + 1
	}
}
