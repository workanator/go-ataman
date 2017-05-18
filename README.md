# go-ataman
**Another Text Attribute Manipulator**

[![GoDoc](https://godoc.org/gopkg.in/workanator/go-ataman.v0?status.svg)](https://godoc.org/gopkg.in/workanator/go-ataman.v0)

The goal of the project is to help render colored text in terminal applications
with text attribute manipulation contained in text template.
The idea is similar to HTML, e.g. `<span style="color:green">Hello!</span>`

## Installation

To install the package use `go get github.com/workanator/go-ataman...`

## Quick Example

The next example shows basic usage of the package. The renderer here uses
basic decoration scheme.

```go
rndr := ataman.NewRenderer(basic.Style())
tpl := "<light:green>%s<green>, <b:light:yellow,black,bold> <<%s>> <-><red>!"
fmt.Println(rndr.MustRenderf(tpl, "Hello", "Terminal World"))

```

This example produces colored text like this.

[![Example Output](https://s4.postimg.org/60xukhl71/2017-05-19_0.58.15.png)](https://postimg.org/image/s0497p215/)
