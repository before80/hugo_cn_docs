package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 获取当前程序运行目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 判断文件是否以 .md 结尾
		if strings.HasSuffix(path, ".md") {
			// 打开文件
			file, err := os.OpenFile(path, os.O_RDWR, 0644)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			defer file.Close()
			// 读取文件内容
			scanner := bufio.NewScanner(file)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			// 在文件开头插入指定的字符串
			lines = append([]string{"将以下英文翻译为中文："}, lines...)
			// 重写文件
			file.Truncate(0)
			file.Seek(0, 0)
			for _, line := range lines {
				fmt.Fprintln(file, line)
			}
		}
		return nil
	})
}
