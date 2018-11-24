package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/toolkits/file"
)

//  RemoveAll 无论文件夹下存在

func Control(workdir string, arg []string) (string, error) {
	cmd := exec.Command("./mkscript", arg...)
	cmd.Dir = workdir
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("cd %s; ./mkscript %s fail %v. output: %s", workdir, arg, err, string(bs))
	}
	return string(bs), err
}

func main() {

	outStr, err := Control(path.Join(file.SelfDir()),
		[]string{"test", "1", "1", "123"})
	if err != nil && !strings.Contains(outStr, "success") {
		return
	}

	splitArr := strings.Split(outStr, " ")
	fmt.Println(splitArr, splitArr[0], splitArr[1])
	fmt.Println("测试：", splitArr[0], splitArr[1])
	fmt.Println(splitArr[0])

	dir, files := filepath.Split(splitArr[0])
	fmt.Println(dir, files)

	err = os.Remove(dir)
	fmt.Println("error1 :", err)

	err = os.RemoveAll(dir)
	fmt.Println("error2 :", err)
}
