# Hashbag

Your all-in-one tool for generating random hashes and strings.

## Setup

```bash
go get -u github.com/inventorandy/hashbag
```

## Examples

### Generating a Random String

Generate a random string of 32 chars using lowercase, uppercase and numeric characters:

```go
package main

import (
	"fmt"

	"github.com/inventorandy/hashbag"
	"github.com/inventorandy/hashbag/charset"
)

func main() {
	str := hashbag.RandomString(charset.LowercaseAlpha, charset.UppercaseAlpha, charset.Numeric)
	fmt.Println("Random String:", str)
}
```

### Hashing a Password with a Random Salt

```go
package main

import (
	"fmt"

	"github.com/inventorandy/hashbag"
	"github.com/inventorandy/hashbag/charset"
)

func main() {
	salt := hashbag.RandomString(charset.LowercaseAlpha, charset.UppercaseAlpha, charset.Numeric, charset.Special)
	password := "myPassword123"
	hashedPassword := hashbag.SHA256HashString(password, salt)

	fmt.Printf("Password: %s Salt: %s Hashed Password: %s", password, salt, hashedPassword)
}
```