# Design patterns in Go

This section is based on [Udemy training](https://www.udemy.com/course/design-patterns-go/).


## Builder (chaining)

This pattern is created in order to build the project in steps and not in one go. In some cases we can create new object with all info at the initialization. But if that is not possible/desirable we can use builder pattern when we create object with partial info and later add new values to it. This approach usually favours `chaining` in order to make it simpler.
