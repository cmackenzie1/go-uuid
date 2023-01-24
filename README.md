# go-uuid

[![Go Reference](https://pkg.go.dev/badge/github.com/cmackenzie1/go-uuid.svg)](https://pkg.go.dev/github.com/cmackenzie1/go-uuid)
![go workflow](https://github.com/cmackenzie1/go-uuid/actions/workflows/go.yml/badge.svg)

A simple, stdlib only, go module for generating UUIDs (**U**niversally **U**nique **ID**entifiers).

## Installation

```bash
go get github.com/cmackenzie1/go-uuid
```

## Supported versions

| Version     | Variant | Details                                                                                                                                                                                                      |
|-------------|---------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Version 4` | `10`    | Pure random as defined in [RFC4122](https://www.rfc-editor.org/rfc/rfc4122).                                                                                                                                 |
| `Version 7` | `10`    | Time-sortable as defined in a [working draft]( https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-01.html#name-uuid-version-7) meant to update [RFC4122](https://www.rfc-editor.org/rfc/rfc4122). |

## Usage

```go
// example/example.go
package main

import (
	"fmt"

	"github.com/cmackenzie1/go-uuid"
)

func main() {
	v4, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Printf("UUIDv4: %s\n", v4) // c07526de-40e5-418f-93d1-73ba20d2ac2c

	v7, _ := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	fmt.Printf("UUIDv7: %s\n", v7) // 0185e1af-a3c1-704f-80f5-6fd2f8387f09
}

```

## FAQ

### What are the benefits of this library over X?

- A single library with no external dependencies for multiple types of UUIDs.
- `UUID` type is defined as a fixed-size, `[16]byte`, array which can be used as a map (instead of the 36 byte
  string representation). Over 2x space savings for memory!
- Limited API. As per RFC4122, UUIDs (while containing embedded information), should be treated as opaque
  values. There is no temptation to build dependencies on the embedded information if you can't easily access it. 😉

### When should I use UUIDv7 over UUIDv4?

> Non-time-ordered UUID versions such as UUIDv4 have poor database index locality. This means that new
> values created in succession are not close to each other in the index and thus require inserts to be performed at
> random locations. The negative performance effects of which on common structures used for this (B-tree and its
> variants) can be dramatic. [1]

**tl;dr**: if you intend to use the UUID as a database key, use UUIDv7. If you require
purely random IDs, use UUIDv4.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](./LICENSE.md)

[1]: https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-01.html#section-2.1