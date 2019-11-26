package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

/**
 * @Author: singHwang
 * @File:  main
 * @Version: 1.0.0
 * @Date: 11/26/19 7:49 PM
 */

func printUsage() {
	fmt.Println("Usage:")

	fmt.Println("\taddblock -data DATA --　交易数据.")

	fmt.Println("\tprintchain --　输出区块信息")

}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()

		os.Exit(1)

	}
}
func main() {

	isValidArgs()

	addBlodCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	flagAddBlockData := addBlodCmd.String("data", "http://liyunchun.org", "交易数据.....")

	switch os.Args[1] {
	case "addBlock":
		err := addBlodCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":

		//数组第二个数字以后的数
		err := addBlodCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1)
	}

	if addBlodCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}

		fmt.Println(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("输出所有区块数据....")

	}

}
