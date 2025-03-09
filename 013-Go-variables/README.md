# Go variable Types

In Go, there are different **types** of variables, for example:

- `int` - stores integers (whole numbers), such as 123 or -123

- `float32` - stores floating point number, with decimal, such as 19.99 or -19.99

- `string` - stores text, such as "Hello World". String values are surrounded by double quotes `""`(not single quotes `''`) or backticks ` `` `

- `bool` - stories values with two states: true or false

# Declaring (Creating) Variables

In Go, there are two wats to declare a variable:

## 1. With the `var` keyword:

Use the `var` keyword, followed by variable name and type:

```go
var variablename type = value
```

**Note:** You always have to specify either `type` or `value` (or both)

## 2. With the `:=` sign:

Use the `:=` sign, followed by the variable value:

```go
variablename := value
```

**Note:** In this case, the type of the variable is **inferred** form the value (means that the compiler decides the type of the variable, based on the value)

**Note:** It is not possible to declare a variable using `:=`, without assigning a value to it.

# Variable Declaration With Initial Value

In the value of variable is from the start, you can declare the variable and assign a value to it on one line:

```go
package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" // type is string
	var student2 = "jane"        // type is inferred
	x := 2                       // type is inferred
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(`davit`)
}
```

**Note:** Thea variable types of `student2` and `x` is **inferred** from their values.

# Variable Declaration Without Initial Value

In Go, all variable are initialized, So, if you declare a variable without an initial value, its value will be set to the default value of its type:

```go
package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
```

In the example there are 3 variable:

- `a`

- `b`

- `c`

These variables are declared but they have not been assigned initial values.

By running the code, we can see that they already have the default values of their respective types:

- `a` is `""`

- `b` is `0`

- `c` is `false`

# Value Assignment After Declaration

It is possible to assign a value to a variable after it is declared. This is helpful cases the value is not initially known.

```go
package main

import (
	"fmt"
)

func main() {
	var student3 string
	student3 = "John"
	fmt.Println(student3)
}
```

**result:**

```go
John
```

> **Note:** It is not possible to declare a variable using "`:=`" without assigning a value to it

# Difference Between var and :=

There are some small difference between the `var` var `:=`:

| var                                                                  | :=                                                                                                      |
| -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| Can be used **inside** and **outside** of functions                  | Can only be used **inside** functions                                                                   |
| Variable declaration and value assignment **can be done separately** | Variable declaration and value assignment **cannot be done separately** (must be done in the same line) |

```go
package main

import (
	"fmt"
)

var z int
var v int = 2
var n = 3

func main() {

	z = 1
	fmt.Println(z)
	fmt.Println(v)
	fmt.Println(n)
}
```

**result:**

```go
1
2
3
```

Since `:=` is used outside of a function, running the program result in ar error.

```go
package main

import (
	"fmt"
)

j := 7

func main() {
	fmt.Println(j)
}
```

**result:**

```bash
.\hello.go:7:1: syntax error: non-declaration statement outside function body
```
