package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//"github.com/gpmgo/gopm/modules/cae/zip"
)

func main() {
	//file, err := os.OpenFile("test.zip", os.O_RDWR|os.O_CREATE, 0755)
	//fmt.Println(err)

	//zipArc := zip.New(file)
	//
	//err = zipArc.AddFile("test",
	//	"/home/zanghong/zipTest/test")
	//err = zipArc.AddFile("test1",
	//	"/home/zanghong/zipTest/test")
	//fmt.Println(err)
	//
	//err = zipArc.Close()
	//fmt.Println(err)
	log.Println(filepath.Dir("/home/one"))
}

func CompressZip() {
	const dir = "./testDir/"
	//获取源文件列表
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create("img-50.zip")
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(dir + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		n, err := fw.Write(filecontent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}

func DeCompressZip() {
	const File = "img-50.zip"
	const dir = "img/"
	os.Mkdir(dir, 0777) //创建一个目录

	cf, err := zip.OpenReader(File) //读取zip文件
	if err != nil {
		fmt.Println(err)
	}
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}

		f, err := os.Create(dir + file.Name)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		n, err := io.Copy(f, rc)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}

}
