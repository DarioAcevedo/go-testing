
## Go testing

This is a simple MVC to learn some of the basics of go testing. 

The domain layer is tested with testing pkg. On the utils we have an example of benchmarking. 
Then the service layer uses testify to assert and mock the tests and finally the controller layer is using a bit of table driven testing just to show the pattern.

## Dependencies

* [Gingo](https://github.com/gin-gonic/gin)
* [Testify](https://github.com/stretchr/testify)

Or just run 

`go mod download`
