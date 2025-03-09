# Go Multiple Variable Declaration

In Go, it is possible to declare multiple variables on the same line:

```go
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 1, 2, 3, 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
```

**result:**

```
1
2
3
4
```

> **Note:** If you use `type` keyword, it is only possible to declare **one type** of variable per line.

If the `type` keyword is not specified, you can declare different types of variables on the same line:

```go
package main

import (
	"fmt"
)

func main() {
	var a, b = 5, "Hello"
	c, d := 7, "world"

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(d)
}
```

**result:**

```bash
5
7
Hello
world
```

# Go Variable Declaration in a Block

```go
package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
```

**return:**

```bash
0
1
hello
```
