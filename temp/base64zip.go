package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	// 假设有需要压缩的原始数据
	originalData := "Hello, this is some data that needs to be compressed!"

	// 将原始数据编码为Base64
	encodedData := base64.StdEncoding.EncodeToString([]byte(originalData))

	// 对Base64编码后的数据进行压缩
	var compressedData bytes.Buffer
	writer := zlib.NewWriter(&compressedData)
	writer.Write([]byte(encodedData))
	writer.Close()

	// 此时可以将compressedData保存在文件中或进行传输
	// ...

	// 解压缩过程
	// 从文件或传输接收到compressedData后
	// ...

	// 解压缩数据
	reader, err := zlib.NewReader(&compressedData)
	if err != nil {
		fmt.Println("解压缩出错：", err)
		return
	}
	defer reader.Close()

	decompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("解压缩读取出错：", err)
		return
	}

	// 对解压缩后的数据进行Base64解码
	decodedData, err := base64.StdEncoding.DecodeString(string(decompressedData))
	if err != nil {
		fmt.Println("Base64解码出错：", err)
		return
	}

	fmt.Println("原始数据：", originalData)
	fmt.Println("解压缩后的数据：", string(decodedData))
}
