# Basic Data Types In Go

- bool: represents a boolean value and is either true or false
- Numeric: represents integer types, float point values, and complex types
- string: represents a string values

## Boolean Types

A boolean data type is declared with the bool keyword and can only take the values true or false.

The default value of a boolean data type is false.

### Example:

```go
package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}
```

## Integer Types
| Type  | Size                           | Range                                                                 |
|-------|--------------------------------|-----------------------------------------------------------------------|
| int   | Depends on platform: 32 bits in 32-bit systems and 64 bits in 64-bit systems | -2,147,483,648 to 2,147,483,647 in 32-bit systems and -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 in 64-bit systems |
| int8  | 8 bits / 1 byte                | -128 to 127                                                           |
| int16 | 16 bits / 2 bytes              | -32,768 to 32,767                                                     |
| int32 | 32 bits / 4 bytes              | -2,147,483,648 to 2,147,483,647                                       |
| int64 | 64 bits / 8 bytes              | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807               |

### Example:


```go
package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)

    var i int = 42
    var f float64 = 3.142
    var c complex128 = complex(1, 2)
    fmt.Printf("int: %d, float64: %f, complex128: %v\n", i, f, c)
}
```

## Float Types

|   Type    |   Size    |          Range        |
|-----------|-----------|-----------------------|
|   float32 |   32 bits | -3.4e+38 to 3.4e+38.  |
|   float64 |   64 bits |-1.7e+308 to +1.7e+308.|

Example:

```go
package main
import ("fmt")

func main() {
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
```

## String Types
The **string** data type is used to store a sequence  of characters (text). String values must be surrounded by double quotes:

### Example:

```go
package main
import ("fmt")

func main() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
```