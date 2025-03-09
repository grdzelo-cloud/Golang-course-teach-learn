# Go Constants

If a variable should have fixed value that cannot be changed, you can use `const` keyword.

The `const` keyword declares the variable as "constant", which means that it is **unchangeable and read-only.**

```go
const CONSTANTNAME type = value
```

> **Note:** The value of a constant must be assigned when you declare it.

# Declaring a Constant

Here is example of declaring a constant in Go:

```go
package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	fmt.Println(PI)
}
```

**result:**

```bash
3.14
```

# Constant Rules

- Constant names follows the same naming rules as variables

- Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)

# Const Types

There are two types of constants:

- Typed constants

- Untyped constants

# Typed Constants

Typed constants are declared with a defined type:

```go
package main

import (
	"fmt"
)

const A int = 1

func main() {
	fmt.Println(A)
}
```

**result:**

```bash
1
```

# Untyped Constants

Untyped constants are declared without a type:

```go
package main

import (
	"fmt"
)

const A = 1

func main() {
	fmt.Println(A)
}
```

**result:**

```bash
1
```

> **Note:** In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value)

# Constants: Unchangeable and Read-only

When a constant is declared, it is not possible to change the value later:

package main

```go
package main

import (
	"fmt"
)

func main() {
    const A int = 1
	A = 2
    fmt.Println(A)
}
```

**result**

```bash
./hello.go:9:7: cannot assign to A
```

# Multiple Constants Declaration

Multiple constants can be grouped together into a block for readability:

```go
package main

import (
	"fmt"
)

const (
	A int = 1
	B     = 3.14
	c     = "Hi!"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(c)
}
```

**result:**

```
1
3.14
Hi!
```
