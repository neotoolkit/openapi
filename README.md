# OpenAPI

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

OpenAPI specification object model

## Features
- Easy to integrate.

## Installation
```shell
go get github.com/neotoolkit/openapi
```

## Usage
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/neotoolkit/openapi"
)

func main() {
	file, err := ioutil.ReadFile("openapi.yml")
	if err != nil {
		log.Fatalln(err)
	}

	oapi, err := openapi.Parse(file)
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(oapi.OpenAPI)
}

```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/neotoolkit/openapi/workflows/build/badge.svg
[build-url]: https://github.com/neotoolkit/openapi/actions
[pkg-img]: https://pkg.go.dev/badge/neotoolkit/openapi
[pkg-url]: https://pkg.go.dev/github.com/neotoolkit/openapi
[reportcard-img]: https://goreportcard.com/badge/neotoolkit/openapi
[reportcard-url]: https://goreportcard.com/report/neotoolkit/openapi
[coverage-img]: https://codecov.io/gh/neotoolkit/openapi/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/neotoolkit/openapi
[version-img]: https://img.shields.io/github/v/release/neotoolkit/openapi
[version-url]: https://github.com/neotoolkit/openapi/releases

## Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=dotenv-linter">
    <img src="https://raw.githubusercontent.com/neotoolkit/.github/main/assets/sponsored_by_evrone.svg"
      alt="Sponsored by Evrone">
  </a>
</p>
