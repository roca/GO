## Error Handling

Error handling is critical for making your programs reliable, trustworthy and respectful to those who depend on them. A proper error value is both specific and informative. It must allow the caller to make an informed decision about the error that has occurred. There are several ways in Go to create error values. This depends on the amount of context that needs to be provided.

## Notes

* Use the default error value for static and simple formatted messages.
* Create and return error variables to help the caller identify specific errors.
* Create custom error types when the context of the error is more complex.
* Error Values in Go aren't special, they are just values like any other, and so you have the entire language at your disposal.

## Links

http://blog.golang.org/error-handling-and-go

http://www.goinggo.net/2014/10/error-handling-in-go-part-i.html

http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html

http://clipperhouse.com/2015/02/07/bugs-are-a-failure-of-prediction/

http://dave.cheney.net/2014/12/24/inspecting-errors

## Code Review

[Default Error Values](example1/example1.go) ([Go Playground](http://play.golang.org/p/8x6kDZxPWK))

[Error Variables](example2/example2.go) ([Go Playground](https://play.golang.org/p/4YHAbpynl3))

[Type As Context](example3/example3.go) ([Go Playground](http://play.golang.org/p/Eu3X54PnWm))

[Bahvior As Context](example4/example4.go) ([Go Playground](http://play.golang.org/p/6GYqwSxHjI))

[Find The Bug](example5/example5.go) ([Go Playground](http://play.golang.org/p/czXpjvWWTT))

## Exercises

### Exercise 1
Create two error variables, one called ErrInvalidValue and the other called ErrAmountTooLarge. Provide the static message for each variable. Then write a function called checkAmount that accepts a float64 type value and returns an error value. Check the value for zero and if it is, return the ErrInvalidValue. Check the value for greater than $1,000 and if it is, return the ErrAmountTooLarge. Write a main function to call the checkAmount function and check the return error value. Display a proper message to the screen.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/Rt3O-7ndtJ)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/8KETdvYk17))

### Exercise 2
Create a custom error type called appError that contains three fields, Err error, Message string and Code int. Implement the error interface providing your own message using these three fields. Write a function called checkFlag that accepts a bool value. If the value is false, return a pointer of your custom error type initialized as you like. If the value is true, return a default error. Write a main function to call the checkFlag function and check the error for the concrete type.

[Template](exercises/template2/template2.go) ([Go Playground](http://play.golang.org/p/x6UimVQMMQ)) | 
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](http://play.golang.org/p/-v-sxBl_ER))

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
