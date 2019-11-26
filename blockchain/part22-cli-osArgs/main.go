package main

import (
	"fmt"
	"os"
)

/**
 * @Author: singHwang
 * @File:  main
 * @Version: 1.0.0
 * @Date: 11/26/19 7:49 PM
 */
func main() {
	args := os.Args

	fmt.Printf("%v/n", args)

	fmt.Printf("%v\n", args[1])

}
