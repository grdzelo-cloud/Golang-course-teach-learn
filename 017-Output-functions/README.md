# Go Output Functions

Go has three functions to output text:

- `Print()`

- `Println()`

- `Printf()`

# The Print() Function

The `Print()` function print it arguments with their default format.

```go
package main

import (
    "fmt"
)

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}
```

**result**

```bash
HelloWorld
```

If we wat to print the arguments in new lines, we need to use \n.

```go
package main

import (
    "fmt"
)

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}
```

**result:**

```
Hello
World
```

> **Tip:** `n` create new lines.

It is also possible to only use one `Print()` for printing multiple variables.

```go
package main

import (
    "fmt"
)

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n",j)
}
```

**result**

```
Hello
World
```

# The Println() Function

The `Println()` function is similar to `Print()` with the difference that a whitespace is added between the arguments, and a newline is added at the end:

```go
package main

import (
    "fmt"
)

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
```

**result**

```
Hello World
```

# The Printf() Function

The `Printf()` function first formats its argument based on the given formatting verb and then prints them.

Here we will use to formatting verbs:

- `%v` is used to print the **value** of the arguments

- `%T` is used to print the **type** of the arguments

```go
package main

import (
	"fmt"
)

func main() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)
}
```

**result**

```
i has value: Hello and type: string
j has value: 15 and type: int
```

> **Tip:** Look at all the formatting verbs in the [Go Formatting Verbs](https://www.w3schools.com/go/go_formatting_verbs.php) chapter.