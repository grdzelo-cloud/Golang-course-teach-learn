# Go Get Started

To start using Go, you need two things:

- A text editor, like VS Code, to write Go code
- A compiler, like GCC, to translate the Go code into a language that the computer will understand

There are many text editors and compilers to choose from. In this tutorial, we will use an IDE (see below).

---

# GO install

Go Install
You can find the relevant installation files at https://golang.org/dl/.

Follow the instructions related to your operating system. To check if Go was installed successfully, you can run the following command in a terminal window:

> go version

Which should show the version of your Go installation.

# Go install IDE

An IDE (Integrated Development Environment) is used to edit AND compile to code.

Popular IDE's  include Visual Studio Code (VS CODE), Vim, Eclipse, ann Notepad. These are all free, and they can be used to both edit and debug Go Code.

**Note:** Web-based IDE's can work as well, but functionality is limited.

you can find the last version of VS code at https://code.visualstudio.com/.

# Go Quickstart

let's create our first Go program

- Launch the VS Code editor

- Open the extension manager or alternatively, press `Ctrl + Shift + x`

- Run the `Go: Install/Update Tools` command

- Select all provided tools and click OK

VS Code is now configured to use Go.

open up a terminal window and type:

```bash
go mod init example.com/hello
```

Create a new file (hello\hello.go). Copy and paste the following code and save the file as helloworld.go :

### Or open terminal and type

```bash
mkdir hello

cd hello

touch helloworld.go
```

In VS Code:
```bash
code helloworld.go 
```

Following code: 

```go
package main

import (
	"fmt"
)


func main(){
    fmt.Println("Hello World!")
}
```

Now, run the code: Open a terminal in VS Code and type:

```bash
go run ./helloworld.go
```


**output:**

```
 Hello World!
```

**Congratulations!** You have now written and executed your firs Go program.

if you want to save the program as an executable, type and run:

```bash
go build ./helloworld.go
```