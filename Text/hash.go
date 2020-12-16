package Text

import (
	"crypto/sha256"
)

func sha256Hash()  []byte {
	sha256Hash := sha256.New()
	sha256Hash.Write(gongyao())
	return sha256Hash.Sum(nil)
}



