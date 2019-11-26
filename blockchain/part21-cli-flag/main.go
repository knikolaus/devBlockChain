package main

import (
	"flag"
	"fmt"
)

/**
 * @Author: singHwang
 * @File:  main
 * @Version: 1.0.0
 * @Date: 11/26/19 7:49 PM
 */
func main() {
	flagString := flag.String("printchain", "", "输出所有的区块信息.....")

	flagInt := flag.Int("number", 6, "输出一个整数")

	flagBool := flag.Bool("open", false, "判断真假")

	flag.Parse()

	fmt.Printf("%s\n", *flagString)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)

}
