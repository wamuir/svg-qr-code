package qrsvg

import (
	"io/ioutil"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = mustReadFile("example.svg")

func mustReadFile(f string) []byte {

	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return b
}

func TestNew(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code")
	assert.NotNil(t, qr)
	assert.NoError(t, err)
}

func TestHex(t *testing.T) {
	h := hex(color.Black)
	assert.Equal(t, h, "#000000")
}

func TestQRString(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code")
	assert.NotNil(t, qr)
	assert.NoError(t, err)

	svg := qr.String()
	assert.Equal(t, svg, string(example))
}
