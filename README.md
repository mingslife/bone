# Bone

[![Go](https://github.com/mingslife/bone/actions/workflows/go.yml/badge.svg)](https://github.com/mingslife/bone/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mingslife/bone)](https://goreportcard.com/report/github.com/mingslife/bone)
[![MIT License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/mingslife/bone/branch/master/graph/badge.svg?token=ZTQOIOB07x)](https://codecov.io/gh/mingslife/bone)
[![GoDoc](https://pkg.go.dev/badge/github.com/mingslife/bone?status.svg)](https://pkg.go.dev/github.com/mingslife/bone?tab=doc)

🦴 Non-intrusive Go micro-service framework to combine modules organically, base on go-kit.

## Installation

```sh
$ go get -u github.com/mingslife/bone
```

## Documentation

[Go Packages](https://pkg.go.dev/github.com/mingslife/bone?tab=doc)

## Quick start

```go
package main

import "github.com/mingslife/bone"

func main() {
	options := bone.DefaultApplicationOptions()
	application := bone.NewApplication(options)
	// application.Use(...components)
	application.Run() // Listen on 127.0.0.1:8080
}
```

## Example

[bone-example](https://github.com/mingslife/bone-example)

## License

&copy; 2022 Ming

Released under the [MIT License](LICENSE)
