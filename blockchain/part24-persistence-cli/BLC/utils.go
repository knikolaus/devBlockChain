package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

/**
 * @Author: singHwang
 * @File:  CLI
 * @Version: 1.0.0
 * @Date: 11/26/19 8:43 PM
 */

// 将int64转换为字节数组
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
