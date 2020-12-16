package Text

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func gongyao() []byte {
	curve := elliptic.P256()
	//ecdsa.GenerateKey(curve,rand.Reader)
	//x和y可以组成公钥
	_, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	//将x和y组成公钥，转换为[]byte类型
	//公钥：x坐标+y坐标
	pubkey := append(x.Bytes(),y.Bytes()...)
	return pubkey
}

