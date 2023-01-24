# go-uuid

A simple, stdlib only, go module for generating UUIDs (Universally Unique IDentifiers).

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
package main

import (
	"fmt"

	"github.com/cmackenzie1/uuid"
)

func main() {
	v4, _ := uuid.NewV4()
	fmt.Printf("UUIDv4: %s", v4)

	v7, _ := uuid.NewV7()
	fmt.Printf("UUIDv7: %s", v7)
}

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](./LICENSE.md)