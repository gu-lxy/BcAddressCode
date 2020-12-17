package main

func banBen() []byte {
	return append([]byte{0x00}, sha256Hash()...)
}
