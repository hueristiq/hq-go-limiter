# hq-go-limiter

![made with go](https://img.shields.io/badge/made%20with-Go-0000FF.svg) [![license](https://img.shields.io/badge/license-MIT-gray.svg?color=0000FF)](https://github.com/hueristiq/hq-go-limiter/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0000FF.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hq-go-limiter.svg?style=flat&color=0000FF)](https://github.com/hueristiq/hq-go-limiter/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hq-go-limiter.svg?style=flat&color=0000FF)](https://github.com/hueristiq/hq-go-limiter/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-0000FF.svg)](https://github.com/hueristiq/hq-go-limiter/blob/master/CONTRIBUTING.md)

A [Go(Golang)](https://golang.org/) package for handling rate limiting.

## Resource

* [Features](#features)
* [Usage](#usage)
* [Contributing](#contributing)
* [Licensing](#licensing)

## Features

## Usage

```bash
go get -v -u go.source.hueristiq.com/limiter
```

Here's a simple example demonstrating how to use `hq-go-limiter`:

```go
package main

import (
	"fmt"
	"go.source.hueristiq.com/limiter"
)

func main() {
	options := &limiter.Options{
		RequestsPerMinute: 40,
		MinimumDelayInSeconds: 2,
	}

	ltr := limiter.New(options)

	// Make 10 requests and ensure that they are rate limited.
	for i := 1; i <= 10; i++ {
		ltr.Wait()
		fmt.Printf("Request %d made at %v\n", i, time.Now())
	}
}
```

## Contributing

Feel free to submit [Pull Requests](https://github.com/hueristiq/hq-go-limiter/pulls) or report [Issues](https://github.com/hueristiq/hq-go-limiter/issues). For more details, check out the [contribution guidelines](https://github.com/hueristiq/hq-go-limiter/blob/master/CONTRIBUTING.md).

Huge thanks to the [contributors](https://github.com/hueristiq/hq-go-limiter/graphs/contributors) thus far!

![contributors](https://contrib.rocks/image?repo=hueristiq/hq-go-limiter&max=500)

## Licensing

This package is licensed under the [MIT license](https://opensource.org/license/mit). You are free to use, modify, and distribute it, as long as you follow the terms of the license. You can find the full license text in the repository - [Full MIT license text](https://github.com/hueristiq/hq-go-limiter/blob/master/LICENSE).