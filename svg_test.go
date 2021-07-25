package qrsvg

import (
	"image/color"
	"io/ioutil"
	"testing"
)

var example = mustReadFile("example.svg")

func mustReadFile(f string) []byte {

	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return b
}

func TestNew1(t *testing.T) {
	qr, err := New("")
	if err == nil {
		t.Errorf("err = %v; want non-nil", err)
	}
	if qr != nil {
		t.Errorf("qr = %v, want %v", qr, nil)
	}
}

func TestNew2(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code")
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Errorf("qr = %v, want non-nil", qr)
	}
}

func TestHex(t *testing.T) {
	h := hex(color.Black)
	if h != "#000000" {
		t.Errorf("h = %v, want %v", h, "#000000")
	}
}

func TestQRString(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code")
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Errorf("qr = %v, want non-nil", qr)
	}

	svg := qr.String()
	if svg != string(example) {
		t.Errorf("svg = %v, want %v", svg, string(example))
	}
}
