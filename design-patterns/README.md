# Design patterns in Go

This section is based on [Udemy training](https://www.udemy.com/course/design-patterns-go/).

## Builder (chaining)

This pattern is created in order to build the project in steps and not in one go. In some cases we can create new object with all info at the initialization. But if that is not possible/desirable we can use builder pattern when we create object with partial info and later add new values to it. This approach usually favours `chaining` in order to make it simpler.

## Factory

This is function that creates instance of object/struct. The factory function is handy for setting default values. It is also useful for perfoming validations during initialization of the object.

The functional approach of the factory is well known `currying` approach that returns a function that will return struct.

## Adapter

Adapter pattern is simply making coversion between two different interfaces. In the training example vector image is converted to bimap by adapter interface. I did simpler example of converting Fahrenheid to Celsius. Important point for adapters using lot of data and processing can be caching. Simple example also demonstrates caching although in that example is abit overkill.
