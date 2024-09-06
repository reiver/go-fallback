# go-fallback

Package **fallback** provides tool todefault to a contingency-value, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fallback

[![GoDoc](https://godoc.org/github.com/reiver/go-fallback?status.svg)](https://godoc.org/github.com/reiver/go-fallback)

## Examples

```golang
import "github.com/reiver/go-fallback"

// ...

// food returns false if there is no food to be found.
func food() (string, bool) {
	return "", false
}

// ...

var value string = fallback.DefaultTo[string]{"apple pie"}.WrapBool(food())

```

```golang
import "github.com/reiver/go-fallback"

// ...

// food returns an error if there is no food to be found.
func food() (string, error) {
	return "", errors.New("oh no! there is no food")
}

// ...

var value string = fallback.DefaultTo[string]{"apple pie"}.WrapError(food())

```

## Import

To import package **fallback** use `import` code like the following:
```
import "github.com/reiver/go-fallback"
```

## Installation

To install package **fallback** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-fallback
```

## Author

Package **fallback** was written by [Charles Iliya Krempeaux](http://reiver.link)
