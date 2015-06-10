## Struct Types - Language Syntax

Struct types are a way of creating complex types that group fields of data together. They are a great way of organizing and sharing the different aspects of the data your program consumes.

## Notes

* We can use the struct literal form to intialize a value from a struct type.
* The dot (.) operator allows us to access individual field values.
* We can create anonymous structs.

## Links

http://www.goinggo.net/2013/07/understanding-type-in-go.html

http://www.goinggo.net/2013/07/object-oriented-programming-in-go.html

## Code Review

[Declare, create and initalize struct types](example1/example1.go) ([Go Playground](http://play.golang.org/p/-oeDmu2et8))

[Anonymous struct types](example2/example2.go) ([Go Playground](http://play.golang.org/p/_xxuE1Ep6U))

## Advanced Code Review

[Struct type alignments](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/1CL1ACDipG))

## Exercises

### Exercise 1

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initalize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/ItPe2EEy9X)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/tnn-8hJPUd))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
