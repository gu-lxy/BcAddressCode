package main

import (
	"BcAddressCode/Base58"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	fmt.Println("hello word")
	//第一步：生成私钥公钥
	curve := elliptic.P256()
	//ecdsa.GenerateKey(curve,rand.Reader)
	//x和y可以组成公钥
	_, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//将x和y组成公钥，转换为[]byte类型
	//公钥：x坐标+y坐标
	pubkey := append(x.Bytes(),y.Bytes()...)

	//2、hash计算
	sha256Hash := sha256.New()
	sha256Hash.Write(pubkey)
	pubHash256 := sha256Hash.Sum(nil)
	//ripemd160: gethub下载
	ripmd := ripemd160.New()
	ripmd.Write(pubHash256)
	pubRipmd160 := ripmd.Sum(nil)
	//3、添加版本号作为前缀
	//append([]byte{}, pubRipmd160...)
	versionPubRipemd160 := append([]byte{0x00}, pubRipmd160...)

	//4、计算校验位
	sha256Hash.Reset() //重置
	sha256Hash.Write(versionPubRipemd160)
	hash1 := sha256Hash.Sum(nil)
	sha256Hash.Reset()
	sha256Hash.Write(hash1)
	hash2 := sha256Hash.Sum(nil)
	//c、取前4个字节
	//如何截取[]byte的前4个字节
	//hash[开始:结尾]：前闭后开
	check := hash2[0:4]

	//5、拼接校验位,得到地址
	addBytes := append(versionPubRipemd160,check...)
	fmt.Println("地址", addBytes)

	//第六步，对地址进行base58编码
	//github：go base58
	address := Base58.Encode(addBytes)
	fmt.Println("生成的新的比特币地址：", address)
}
