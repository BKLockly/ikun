package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const (
	width         = 470
	height        = 120
	pageLimit     = 1479
	numCharacters = height * width
)

func main() {
	var cache = make(map[string][]byte)

	// 枚举所有文件，并将内容缓存到 `cache`
	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, "ASCII-cxk_") {
			b, err := ioutil.ReadFile(name)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
			cache[name] = b
		}
	}

	// 遍历所有页面
	for i := 1; i <= pageLimit; i++ {
		name := fmt.Sprintf("ASCII-cxk_%06d.txt", i)

		// 从缓存中读取文本
		strBytes := cache[name][:numCharacters]
		str := string(strBytes)
		time.Sleep(22 * time.Millisecond)
		fmt.Print(str)

		// 清屏
		fmt.Print("\033[H\033[2J")
	}

}
