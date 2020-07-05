# Design patterns in Go

This section is based on [Udemy training](https://www.udemy.com/course/design-patterns-go/).

## Builder (chaining)

This pattern is created in order to build the project in steps and not in one go. In some cases we can create new object with all info at the initialization. But if that is not possible/desirable we can use builder pattern when we create object with partial info and later add new values to it. This approach usually favours `chaining` in order to make it simpler.

## Factory

This is function that creates instance of object/struct. The factory function is handy for setting default values. It is also useful for perfoming validations during initialization of the object.

The functional approach of the factory is well known `currying` approach that returns a function that will return struct.

Few variations are implemented in the factory folder. The main point is felxibility during creation of new struct object. In real life these structures will usually be more complex.

## Adapter

Adapter pattern is simply making coversion between two different interfaces. In the training example vector image is converted to bimap by adapter interface. I did simpler example of converting Fahrenheid to Celsius. Important point for adapters using lot of data and processing can be caching. Simple example also demonstrates caching although in that example is abit overkill.

## Prototype

The idea of prototype is to reuse specific structure and just change few props. But because we can work with pointers, and in some situation specific variables types are passed by refference by default we might experience some troubles. To be sure that we copy values instead of pointer DeepCopy is required. In the prototype folder example of DeepCopy is shown.

## Singleton

The pattern ensures that only ONE instance of the struct is present in the application. These are specific cases, one example could be database connection (generally you would like to use one database for all data). It seems that some experts are not in favour of creating singletons.

To achieve single instance of an struct we do not expose struct to consumer and we use sync.Once method which ensures that some method is called only once.
For more details see singleton folder.

## Composite (Graphs)

This sections models sort of graph structure. The idea is that structure can have children of its own type recursively. This makes possible to have objects of unlimited depth. The object can be traversed recursively. In the example we model simple neuron mimicing some sort of neural network.

## Decorator

Augment an object with additional functionality. And you want to keep functionality separate.
