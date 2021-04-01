# Miro API in Go

![tests](https://github.com/kekelinks/go-golinks/workflows/tests/badge.svg)
[![codecov](https://codecov.io/gh/kekelinks/go-golinks/branch/master/graph/badge.svg)](https://codecov.io/gh/kekelinks/go-golinks)
[![Documentation](https://godoc.org/github.com/yangwenmai/how-to-add-badge-in-github-readme?status.svg)](https://pkg.go.dev/mod/github.com/kekelinks/go-golinks)

Go written [Golinks](https://docs.golinks.io/) API client.

## Installation

Include this is your code as below:

```go
import "github.com/kekelinks/go-golinks/golinks"
```

Using `go get`:

```console
$ go get github.com/kekelinks/go-golinks
```

## Usage

Using the client:

```go
client := golinks.NewClient("token")
```

API's are very simple and easy to understand.

```go
client.Links.Retrieve(10)
```

## Copyright and License

Please see the LICENSE file for the included license information.
Copyright 2021 by Keisuke Yamashita.
