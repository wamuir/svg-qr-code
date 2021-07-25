svg-qr-code
=====

[![Go Reference](https://pkg.go.dev/badge/github.com/wamuir/svg-qr-code.svg)](https://pkg.go.dev/github.com/wamuir/svg-qr-code)
[![Build Status](https://github.com/wamuir/svg-qr-code/actions/workflows/go.yml/badge.svg?branch=main&event=push)](https://github.com/wamuir/svg-qr-code/actions/workflows/go.yml?query=event%3Apush+branch%3Amain)
[![codecov](https://codecov.io/gh/wamuir/svg-qr-code/branch/main/graph/badge.svg)](https://codecov.io/gh/wamuir/svg-qr-code)
[![Go Report Card](https://goreportcard.com/badge/github.com/wamuir/svg-qr-code)](https://goreportcard.com/report/github.com/wamuir/svg-qr-code)

# Description

`svg-qr-code` is a Go module that encodes QR Codes in SVG format

# Installation

This module can be installed with the `go get` command:

    go get -u github.com/wamuir/svg-qr-code


# Example Usage

```go

  qr, err := qrsvg.New("https://github.com/wamuir/svg-qr-code")
  if err != nil {
     panic(err)
  }

  // qr satisfies fmt.Stringer interface (or call qr.String() for a string)
  fmt.Println(qr)
```

# Example Result

![Quick Response (QR) Code](https://github.com/wamuir/svg-qr-code/raw/main/example.svg)
