# go modules example

This repository contains an example of how to set up a go program with
modules to use an external source library on the same system.  

## Directory structure

This repository contains two directories:
 - `examplelib`:  The library with code we want to import
 - `main`:  The directory containing the program we're writing.  This
   is that program that will import the library.
   
## How it works

Here are the most important steps:

1. In `main/go.mod`: Add a `replace` line to tell go where the module
   lives and what to call it (in this case, "examplelib")
```
replace examplelib => PATH_TO_CODE
```

See `main/go.mod` 

2. In `main/main.go`, import the library using the name given (here,
   "examplelib"), and use it as normal
   
3. Before building the code (with `go build` or `make`), tell go to
   update `go.mod` to go find and track the module, as follows:
   
```
go get examplelib
```

4. Build the code normally--ie, run `make` or `go build` from the
   `main` directory.  This should produce an executable called `main`
   that uses the code in `examplelib`. 


