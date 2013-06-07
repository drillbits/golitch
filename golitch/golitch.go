package golitch

import (
	"io"
	"io/ioutil"
	"math/rand"
)

func Glitch(r io.Reader, num int) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return b, err
	}

	for i := 0; i < num; i++ {
		r := (int(rand.Float64()*float64(len(b)-128)) | 0) + 128
		b[r] = ^b[r] & 0xff
	}

	return b, nil
}
