/*
读取管道数据，提供了标准、安全url的base64的加密，与解密
*/
package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

//stdBase64 标准base64加密
func stdBase64() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	var words []byte
	for scanner.Scan() {
		words = append(words, scanner.Bytes()...)
	}
	encod := base64.StdEncoding.EncodeToString(words)
	fmt.Print(encod)
}

//urlBase64 安全url base64加密
func urlBase64() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	var words []byte
	for scanner.Scan() {
		words = append(words, scanner.Bytes()...)
	}
	encod := base64.URLEncoding.EncodeToString(words)
	fmt.Print(encod)
}

//deBase64 解密base64字符串
func deBase64() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var str []string
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	for i := range str {
		stdStr := strings.ReplaceAll(str[i], " ", "")
		stdStr = strings.ReplaceAll(stdStr, "-", "/")
		stdStr = strings.ReplaceAll(stdStr, "_", "+")
		for len(stdStr)%4 != 0 {
			stdStr += "="
		}
		decod, err := base64.StdEncoding.DecodeString(stdStr)
		if err != nil {
			fmt.Printf("can not decode strings: %s\n", str[i])
		} else {
			fmt.Print(string(decod) + "\n")
		}
	}
}

//showHelp 打印帮助命令
func showHelp() {
	fmt.Print("Usage: base64 [OPTION]\n")
	fmt.Print("Base64 encode standard input, to standard output.Or decode base64 strings\n")
	fmt.Print("-h,    show help\n")
	fmt.Print("-u,    url encode\n")
	fmt.Print("-d,    data decode\n")
}

func main() {
	if len(os.Args) == 1 {
		fmt.Print("warnning: please use cmd ,not powershell\n")
		stdBase64()
	} else if os.Args[1] == "-h" {
		showHelp()
	} else if os.Args[1] == "-u" {
		fmt.Print("warnning: please use cmd ,not powershell\n")
		urlBase64()
	} else if os.Args[1] == "-d" {
		deBase64()
	} else {
		fmt.Printf("can not reconize \"%s\"\n", os.Args[1])
		showHelp()
	}
}
