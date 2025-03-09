# Go Syntax

A Go file consists of the following pars:

- Package declaration

- Import packages

- Functions

- Statements and expressions

Look at the following code, to understand it better:

```go
package main

import ("fmt")

func main() {
	fmt.Println("Hello world")
}
```

**Line 1:** In Go, every program is part of a package. We define this using the `package` keyword. In this example, the program belongs to the `main` package.

**Line 2:** import ("fmt") let us import files included in the `fmt` package

**Lone 3:** A blank line. Go ignores white space. Having white spaces in code make it more readable.

**Line 4:** `func main() {}` is a function. Any code inside its curly brakes `{}`
will be executed.

**Line 5:** `fmt.Println()` is a function made available from the `fmt` package. It is used to output/pint text. In our example it will outout "Hello world!".

> **Note:** In Go, any executable code belongs to the `main` function.

# Go Statements

`fmt.Printl("Hello World!") is a statement.

In Go, Statements are separated by ending a line (hitting the Enter key) or by semicolon ";".

Hitting the Enter key ads "`;`" to the end of the line implicitly (does not show uo in the source code).

The lef curly bracket `{` cannot come at the start of a line. 

Run the following code and see what happens:

### example 

```go
package main
import ("fmt")

func main()
{
  fmt.Println("Hello World!")
}
```

# Go Compact Code
You can write more compact code, like shown bellow(we can write code down bellow, but is is difficult to read)

```go
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}
```