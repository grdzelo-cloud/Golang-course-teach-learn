# Go Arrays

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

# Declare an Array

In Go, There are two ways to declare an array:

1. With the `var` keyword:

```go
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred
```

2. With the `:=` sign:

```go
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
```

> **Note:** The _length_ specifies the number of elements to store in array. In Go, array have fixed length. The length of the array is either define by a number or is inferred (means that the compiler decides the length of the array, based on the number of _values_)

```go
package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 3, 4}
	arr2 := [5]int{1, 3, 4, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

```

**result**

```
[1 2 3]
[4 5 6 7 8]
```

```go
package main

import (
	"fmt"
)

func main() {

	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

}
```

**result**

```
[1 2 3]
[4 5 6 7 8]
```
