package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 遍历所有以 .md 结尾的文件
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// 如果是 .md 文件，则处理
		if strings.HasSuffix(path, ".md") {
			// 打开文件
			file, err := os.OpenFile(path, os.O_RDWR, 0644)
			if err != nil {
				fmt.Println(err)
				return err
			}
			defer file.Close()
			// 读取文件内容并删除首行
			scanner := bufio.NewScanner(file)
			lines := []string{}
			tempNum := 0
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "将以下英文翻译为中文：") && tempNum == 0 {
					tempNum++
				} else {
					lines = append(lines, line)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
				return err
			}
			// 清空文件内容并写入新内容
			if err := file.Truncate(0); err != nil {
				fmt.Println(err)
				return err
			}
			if _, err := file.Seek(0, 0); err != nil {
				fmt.Println(err)
				return err
			}
			writer := bufio.NewWriter(file)
			for _, line := range lines {
				if _, err := writer.WriteString(line + "\n"); err != nil {
					fmt.Println(err)
					return err
				}
			}
			if err := writer.Flush(); err != nil {
				fmt.Println(err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
