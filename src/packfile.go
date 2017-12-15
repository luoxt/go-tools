// @brief 解决根据路径列表打包到对应的文件夹里,方便ftp上传代码发布
// 
// @author luoxt
// @date 20171215
//
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func main() {

	fmt.Println("请输入根路径:")
	firstInput := bufio.NewScanner(os.Stdin)

	firstPath := ""
	if firstInput.Scan() {
		firstPath = firstInput.Text()
	}

	fmt.Println("请输入文件列表和'end'并按回车结束输入：")
	var pathstr []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" { break }

		textstr := input.Text()
		if len(textstr) >5 {
			pathstr = append(pathstr, textstr)
		}
	}

	fmt.Println("打包文件列表为：")
	for inputId, pathval := range pathstr {
		inputId++
		fmt.Println("【序号】", inputId)
		making(pathval, firstPath)
	}
}

func making(file string, firstPath string)  {
	srcname := file
	fmt.Println("源文件：", srcname)

	//目标文件
	dstname := "/zh_"+time.Now().Format("200601021504")+"/"+strings.Trim(srcname, ".")

	counter,_ := CopyFile(dstname, firstPath+srcname)
	dir, _ := os.Getwd()
	isexist, err := PathExists(dir+dstname)
	if isexist {
		fmt.Println("成功：√", "大小：", counter)
	} else {
		fmt.Println("失败：X", "原因：", err)
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("【打开源文件失败】", err)
	}
	defer src.Close()

	//创建目录
	dstPath := path.Dir(dstName)
	dstPathNew, err := makeDir(dstPath)
	if err != nil {
		fmt.Println("【创建目录失败】", err)
	}

	//创建文件
	filePath := dstPathNew+"/"+path.Base(dstName)
	dst, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("【创建新文件失败】", err)
	}
	defer dst.Close()

	fmt.Println("新文件：", filePath)

	//拷贝文件
	return io.Copy(dst, src) //
}

func makeDir(path string)(dirNew string, err error)  {

	dir, _ := os.Getwd()
	dirNew = dir+path
	err = os.MkdirAll(dirNew, os.ModePerm)  //生成多级目录
	return dirNew, err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
