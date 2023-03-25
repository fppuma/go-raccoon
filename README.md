# go-raccoon
Playground with goroutines

## Function Literals
In Go, a function literal (also known as an anonymous function) is a function 
that doesn't have a name and can be used as a value. 
It's a way to define a function "on the fly" and pass it around as an argument 
to other functions, or assign it to a variable or a field in a struct.
### Use Cases
1. As a callback function. [literal_sort.go](src/literal_sort.go)
2. As a parameterized function. [literal_param_function.go](src/literal_param_function.go)
3. As a closure. [literal_closure.go](src/literal_closure.go)

## Goroutine
Goroutines are an essential part of **Go's concurrency model**, 
which is designed to make it easy to write programs that can perform many tasks at once.

 A goroutine is **a lightweight thread of execution** that allows you to perform 
 concurrent operations.