# go-ataman
**Another Text Attribute Manipulator**

[![GoDoc](https://godoc.org/gopkg.in/workanator/go-ataman.v1?status.svg)](https://godoc.org/gopkg.in/workanator/go-ataman.v1)

The goal of the project is to help render colored text in terminal applications
with text attribute manipulation contained in text template.
The idea is similar to HTML, e.g. `<span style="color:green">Hello!</span>`

## Installation

To install the package use `go get github.com/workanator/go-ataman...`

## Quick Example

The next example shows basic usage of the package. The renderer here uses
basic decoration style.

```go
rndr := ataman.NewRenderer(ataman.BasicStyle())
tpl := "<light+green>%s<->, <b+light+yellow,black,bold> <<%s>> <-><red>!"
fmt.Println(rndr.MustRenderf(tpl, "Hello", "Terminal World"))

```

This example produces colored text like this.

[![Example Output](https://s24.postimg.org/cpl13bvp1/2017-05-19_15.56.34.png)](https://postimg.org/image/6onc6992p/)

Renderer is able to pre-render templates so in further output operations they
can be reused without parsing and rendering. In this example the renderer uses
curly brackets decoration style.

```go
rndr := ataman.NewRenderer(ataman.CurlyStyle())
prep := rndr.MustPrepare("{light_green}%s{-}, {bg_light_yellow+black+bold} <%s> {-}{red}!")
fmt.Println(rndr.MustRenderf(tpl, "Hi", "Everybody"))
```

The pre-rendered template implements `fmt.Stringer` interface so it can be used
in output operations without additional code, e.g.

```go
prep := rndr.MustPrepare("{red}Red {green}Green {blue}Blue{-}")
fmt.Println(prep)
```
