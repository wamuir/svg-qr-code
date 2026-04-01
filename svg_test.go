package qrsvg

import (
	"image/color"
	"os"
	"testing"
)

func mustReadFile(f string) []byte {
	b, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return b
}

func TestNew1(t *testing.T) {
	qr, err := New("")
	if err == nil {
		t.Errorf("err = nil; want non-nil")
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
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-highest.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestQRStringLevelLow(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code", WithRecoveryLevel(Low))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-low.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestQRStringLevelMedium(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code", WithRecoveryLevel(Medium))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-medium.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestQRStringLevelHigh(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code", WithRecoveryLevel(High))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-high.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestQRStringLevelHighest(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code", WithRecoveryLevel(Highest))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-highest.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestQRStringWithBlocksize(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code", WithBlocksize(8))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-blocksize8.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestQRStringWithBorderwidth(t *testing.T) {
	qr, err := New("https://github.com/wamuir/svg-qr-code", WithBorderwidth(2))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-borderwidth2.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}

func TestNewInvalidOptions(t *testing.T) {
	tests := []struct {
		name string
		opts []func(*opts)
	}{
		{"blocksize zero", []func(*opts){WithBlocksize(0)}},
		{"blocksize negative", []func(*opts){WithBlocksize(-1)}},
		{"borderwidth negative", []func(*opts){WithBorderwidth(-1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr, err := New("https://github.com/wamuir/svg-qr-code", tt.opts...)
			if err == nil {
				t.Error("err is nil, want non-nil")
			}
			if qr != nil {
				t.Errorf("qr = %v, want nil", qr)
			}
		})
	}
}

func TestQRStringWithStyle(t *testing.T) {
	style := `.dark{fill:#370240}`

	qr, err := New("https://github.com/wamuir/svg-qr-code", WithStyle(style))
	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}
	if qr == nil {
		t.Error("qr is nil, want non-nil")
	}

	expected := mustReadFile("testdata/example-styled.svg")

	svg := qr.String()
	if svg != string(expected) {
		t.Errorf("svg = %v, want %v", svg, string(expected))
	}
}
