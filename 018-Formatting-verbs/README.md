# Go Formatting Verbs

# Formatting Verbs for Printf()

Go offers several formatting verbs that can be use with the `Printf()` function.

# General Formatting Verbs

The following verbs can be used wit hall data types

| Verb | Description                            |
| ---- | -------------------------------------- |
| %v   | prints the value in the default format |
| %#v  | Prints the value in Go-syntax format   |
| %T   | prints the type of the value           |
| %%   | Prints the % sign                      |

```go
package main

import (
	"fmt"
)

func main() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)    // 15.5
	fmt.Printf("%#v\n", i)   // 15.5
	fmt.Printf("%#v%%\n", i) // 15.5%
	fmt.Printf("%T\n", i)    // float

	fmt.Printf("%v\n", txt)  // Hello World!
	fmt.Printf("%#v\n", txt) // "Hello World!"
	fmt.Printf("%T\n", txt)  // string
}
```

# Integer Formatting Verbs

The following verbs can be used with that integer data type:

| Verb | description                                |
| ---- | ------------------------------------------ |
| %b   | base 2                                     |
| %d   | base10                                     |
| %+d  | base 10 and always show sign               |
| %o   | base 8                                     |
| %O   | base 8, with leading 0o                    |
| %x   | base 16 lowercase                          |
| %X   | base 16 uppercase                          |
| %#x  | base 16, with leading 0x                   |
| %4d  | Pad with spaces (width 4, right justified) |
| %-4d | Pad with spaces (width 4, left justified)  |
| %04d | Pad with zeros (width 4)                   |

```go
	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
```

**result**

```
1111
15
+15
17
0o17
f
F
0xf
  15
15
0015
```

# String Formatting Verbs

| Verb | Description                                                 |
| ---- | ----------------------------------------------------------- |
| %s   | Prints the value as plain string                            |
| %q   | Prints the value as a double-quoted string                  |
| %8s  | Prints the value as plain string (width 8, right justified) |
| %-8s | Prints the value as plain string (width 8, left justified)  |
| %x   | Prints the value as hex dump of byte values                 |
| % x  | Prints the value as hex dump with spaces                    |

```go
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
```

**result**

```
Hello
"Hello"
   Hello
Hello
48656c6c6f
48 65 6c 6c 6f
```

# Boolean Formatting Verbs

| Verb | Description                                                              |
| ---- | ------------------------------------------------------------------------ |
| %t   | Value of the boolean operator in true or false format (same as using %v) |

```go
  var i = true
  var j = false

  fmt.Printf("%t\n", i)
  fmt.Printf("%t\n", j)
```

**result**

```
true
false
```

# Float Formatting Verbs

| Verb  | Description                               |
| ----- | ----------------------------------------- |
| %e    | Scientific notation with 'e' as exponent  |
| %f    | Decimal point, no exponent                |
| %.2f  | Default width, precision 2                |
| %6.2f | Width 6, precision 2                      |
| %g    | Exponent as needed, only necessary digits |

```go
 var i = 3.141

  fmt.Printf("%e\n", i)
  fmt.Printf("%f\n", i)
  fmt.Printf("%.2f\n", i)
  fmt.Printf("%6.2f\n", i)
  fmt.Printf("%g\n", i)
```

**result**

```
3.141000e+00
3.141000
3.14
  3.14
3.141
```