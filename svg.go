package qrsvg

import (
	"encoding/xml"
	"fmt"
	"image"
	"image/color"

	"github.com/skip2/go-qrcode"
)

// defaults
var (
	blocksize   = 16
	borderwidth = 4
)

// QR holds a QR Code (from github.com/skip2/go-qrcode) and SVG settings.
type QR struct {
	Code        *qrcode.QRCode // underlying QR Code
	Blocksize   int            // size of each block in pixels, default = 16 pixels
	Borderwidth int            // size of the border in blocks, default = 4 blocks
	Borderfill  color.Color    // fill color for the border
}

// SVG is the vector representation of a QR code, as a Go struct.
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	NS      string   `xml:"xmlns,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Style   string   `xml:"style,attr"`
	Blocks  []Block  `xml:"rect"`
}

// Block is a color block in the rendered QR code.
type Block struct {
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Fill   string `xml:"fill,attr"`
}

// SVG returns the vector representation of a QR code, as a Go struct.
// This could, for instance, be marshaled with encoding/xml.
func (q *QR) SVG() *SVG {

	q.Code.DisableBorder = true
	var i image.Image = q.Code.Image(0)
	var w int = i.Bounds().Max.X

	var svg SVG
	svg.NS = "http://www.w3.org/2000/svg"
	svg.Width = (w + 2*q.Borderwidth) * q.Blocksize
	svg.Height = svg.Width

	svg.Blocks = make([]Block, 1+w*w)
	svg.Blocks[0] = Block{0, 0, svg.Width, svg.Height, hex(q.Borderfill)}

	for x := 0; x < w; x++ {
		for y := 0; y < w; y++ {
			svg.Blocks[1+x*w+y] = Block{
				X:      (x + q.Borderwidth) * q.Blocksize,
				Y:      (y + q.Borderwidth) * q.Blocksize,
				Width:  q.Blocksize,
				Height: q.Blocksize,
				Fill:   hex(i.At(x, y)),
			}
		}
	}

	return &svg
}

// String() returns the SVG as a string and satisfies the fmt.Stringer interface.
func (s *SVG) String() string {
	x, _ := xml.Marshal(s)
	return string(x)
}

// String() returns the QR as an SVG string and satisfies the fmt.Stringer interface.
func (q *QR) String() string {
	return q.SVG().String()
}

// New returns the QR for the provided string, with default settings for blocksize,
// borderwidth and background color.  Call .String() to obtain the SVG string.
func New(s string) (*QR, error) {

	code, err := qrcode.New(s, qrcode.Highest)
	if err != nil {
		return nil, err
	}

	q := QR{code, blocksize, borderwidth, code.BackgroundColor}
	return &q, nil
}

func hex(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}
