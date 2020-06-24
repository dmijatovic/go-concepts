# Golang training

This repo contains [Golang training from Udemy].

## Installation

Download installation from [go website](https://golang.org/doc/install?download=go1.14.4.linux-amd64.tar.gz)
Follow the instructions from the website to install Go on Linux.

### Plugins for Visual Studio Code

For proper use with Visual Studio Code you are adviced to install Go plugin. After that you need to install following.
Create *.go file. Visual Studio Code will popup warning that some tools are missing. Select Install All to install all go tools.

## Golang remarks

All modules seem to start with main.go file.

### How to run go

Via terminal, use go commands

```bash
# get help
go help
# run go file (compile and run)
go run <filename>
# build the module (just compile)
go build <filename>
# format the code
go fmt <filename>
# compiles and install package
go install
# download modules from someone else
go get
# run test
go test
```

### What does first line in the file means

Very first line needs to declare package where it belongs. Why we call main package main?
In go there are 2 types of packages:

- executable file: using package as main to do something. Executable package needs to have main method.
- reusable packages: these are modules with helper fn.

```go
// create executable package
package main

// create main function to be call on start
func main(){

}
```

### Basic Go system packages

Every package in Go that you want to use need to be imported. The documentation is [available here](https://golang.org/pkg/).

- math: basic math package
- encoding: encodig module
- crypto:
- io:
- fmt:
- debug:

### How is the main go file organized

```go

//--------------
// PACKAGE DECLARATION
package main

//--------------
// IMPORT DEPENDENCIES
// NOTE! VSC will import automaticall when using it
import "fmt"

//--------------
// FUNCTIONS and INSTRUCTIONS

func main(){
  // we start here!!!
}

```

### Variables

Basic variable types in Go are:

- `bool`: boolean, true/false
- `string`: text/string variable
- `int`: integer in range -10000 to 99999
- `float64`: floating point number 10.00001 - 100.002 (?)

Declaring variable is

```go
// long form of defining variable and its type
var myVariable string = "This is the value"
// alternatively - SHORT notation using :=
// ONLY when declaring it first time
card:="Ace of Spades"
//change value - NOTE no colon
card = "Five of Diamonds"
```

### Methods/Functions

The functions are declared using keywoard `func`, then function name with parameters and then if fn returns what type is returns.

```bash

func myFunction() string {
  return "String to return"
}

```

### Modules and namespaces

In the modules example you will see how two modules are not tightly coupled but can be used in combination.

```bash
# run 2 module examples at the same time
go run main.go state.go
```

### Slices and Loops

Array and Slice are similair data structures. However every item needs to have same types of data.

`WHEN YOU HAVE UNUSED VAR use _`

- Array: fixed length
- Slice: is array that can be mutated (add/remove items)

```go
// define slices
cards:=[]string

for i, card := range cards{
  fmt.PringLn(i,card)
}

```

### Slicing array

Slicing in go is similair to python approach.

array[startIndexIncluding:UpToNotIncluding].

We can use only one of the boudaries. For example array[2:]: from index to up to end, or array[:3] from beginning up to 3 (not including). For slicing example see cards/deck.go file

```go
fruits := []string{"apple","peer","grape","banana"}
// slice first two elements
appleAndPeer := fuits[0:2]

```

### Receivers (methods)

Receiver function enables defined type of object to receive this function. For complete example see cards module deck.go file.
By convention we receive to receiver object using 1 or 2 letters of the type. In some other languages we use this (JS) or self (python).
In the example below we use `d` to refer to deck type variable that will `inherit` this method.

```go
// main.go file
// define variable of type deck
cards:=deck{}
// call method defined in type deck module
cards.print()

// deck.go file
// declare receiver function
// any variable of type deck
// gets access to print method
func (d deck) print() {
 for i, card := range d {
  fmt.Println(i, card)
 }
}
```

### Functions returning multiple values

Go supports method that return multiple values (array of values). This needs to be specified in the function signiture.

```go
// this function returns two deck type values/objects
func deal(d deck, handSize int) (deck, deck){
  return d[:handSize], d[handSize:]
}

```

### Writing data to file

Writing to file is done by `ioutil` [common Go package](https://golang.org/pkg/io/ioutil/#pkg-overview)

- Convert string to byte slice (writing data)

For writing data to a file, go has function WriteFile. This function writes bytes slice to disk.

```go
// example conveting string to byte slice
func (d deck) saveToFile(fileName string) error {
  bytes := []byte(d.toString())
  err := ioutil.WriteFile(fileName, bytes, 0644)
  if err != nil {
    log.Fatal(err)
    return err
  } else {
    return nil
  }
}
```

Reading data from the file

```go
func loadDeckFromFile(fileName string) (deck, error) {
  // read bytes from file
  bs, err := ioutil.ReadFile(fileName)
  if err != nil {
    log.Fatal(err)
    return newDeck(), err
  }
  //converts bytes to string
  tekst := string(bs)
  // split to array
  array := strings.Split(tekst, ";")
  // create deck
  cards := deck(array)
  // return cards
  return cards, nil
}
```

### Randomising

Go uses seed when randomizing values. This is order to be able to repeat randomizing for replication.

## Testing with go

Go has builtin test suite. The test files should have name *_test.go.
The function names should have Test at the beginning of the func name and original func name integrated.
Each test function gets a test handler as first parameter of test fn.
To run test use `go test` command.