# util
Golang functional and iterating tools

`go doc` will give a much better description of this package

## Meta Tools
#### Bind
Simply binds a method to an object and returns it as a function
#### Catch
Error handling similar to Java that will catch any
panicked error from the given function
#### Tail
Very useful tailcall optimization, since golang does not have it
## Iterating Tools
They all work with Chan, Map, Slice, Array, String and Struct
#### For
Functional for loop with some options
#### Fold
Fold calls the given function with pairs of elements and returns the result in the end
#### Map
Allocates a copy but its elements will be modified
#### KeyOf
Iterates through the elements and finds the key of the given element
#### All, Any
Will run the given condition through the elements
