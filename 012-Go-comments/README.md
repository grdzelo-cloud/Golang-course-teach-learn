# Go Comments

A comment is the text that is ignored upon execution.

Comments can be used to explain the code, and to make it more readable.

Comments cna also be used to prevent code execution when testing an alternative code.

Go supports single-line or multi-line comments.

# Go Single-line Comments

Single line comments start with two forward slash (`//`).

Ant text between `//` and the end of the line is ignored by the compiler (will not be executed).

```go
// This is a comment
package main

import "fmt"

func main() {
	// this is a comment

	fmt.Println("Hello world!")
}
```

# Go Multi-line Comments

Multi-line comments start with `/*` and ands with `*/`.

Any text between `/*` and `*/` will be ignored by the compiler:

```go
// This is a comment
package main

import "fmt"

func main() {
	// this is a comment

     /* The code below will print Hello World
         to the screen, and it is amazing */
	fmt.Println("Hello world!")
}
```
> **Tip:** It is uot to you wish to use. Normally, we use `//` for short comments, and `/**/` for longer comments.

# Comment to Prevent Code Execution

You can also use comments to prevent the code from being executed.

The commented code can be saved for later reference and  troubleshooting.

```go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
  // fmt.Println("This line does not execute")
}
```