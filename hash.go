package main

import (
	"crypto/sha256"
)

func sha256Hash()  []byte {
	sha256Hash := sha256.New()
	sha256Hash.Write(gongyao())
	return sha256Hash.Sum(nil)
}

func sha() []byte {
	sha256Hash := sha256.New()
	sha256Hash.Write(gongyao())
	sha256Hash.Sum(nil)
	sha256Hash.Reset() //重置
	sha256Hash.Write(banBen())
	hash1 := sha256Hash.Sum(nil)
	sha256Hash.Reset()
	sha256Hash.Write(hash1)
	return sha256Hash.Sum(nil)
}



