# hqgolimit

[![license](https://img.shields.io/badge/license-MIT-gray.svg?color=0040FF)](https://github.com/hueristiq/hqgolimit/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hqgolimit.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqgolimit/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hqgolimit.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqgolimit/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-0040ff.svg)](https://github.com/hueristiq/hqgolimit/blob/master/CONTRIBUTING.md)

A [Go(Golang)](https://golang.org/) package for handling rate limiting.

## Installation

```
go get -v -u github.com/hueristiq/hqgolimit
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/hueristiq/hqgolimit"
)

func main() {
	options := &hqgolimit.Options{
		RequestsPerMinute: 40,
		MinimumDelayInSeconds: 2,
	}

	limiter := hqgolimit.New(options)

	// Make 10 requests and ensure that they are rate limited.
	for i := 1; i <= 10; i++ {
		limiter.Wait()
		fmt.Printf("Request %d made at %v\n", i, time.Now())
	}
}
```

## Contributing

[Issues](https://github.com/hueristiq/hqgolimit/issues) and [Pull Requests](https://github.com/hueristiq/hqgolimit/pulls) are welcome! **Check out the [contribution guidelines](./CONTRIBUTING.md)**.

## Licensing

This package is distributed under the [MIT license](https://github.com/hueristiq/hqgolimit/blob/master/LICENSE).