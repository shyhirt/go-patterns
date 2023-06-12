<p align="center">
  <img src="/gopher.png" height="400">
  <h1 align="center">
    Go Patterns
    <br>
    <a href="https://github.com/shyhirt/go-patterns/blob/main/LICENSE" ><img alt="license" src="https://img.shields.io/badge/license-Apache%20License%202.0-E91E63.svg?style=flat-square" /></a>
  </h1>
</p>

Design patterns in Go. The language itself encourages a straightforward and idiomatic approach, rather than relying heavily on traditional design patterns like those found in object-oriented languages.
## Creational Patterns

|     Pattern      | Description                                                                                                                                                                                                                                                                                                                                    | Status |
|:----------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------:|
| Abstract Factory | Abstract Factory pattern using interfaces and multiple concrete implementations of those interfaces.                                                                                                                                                                                                                                           | ✔ |
|     Builder      | The Builder pattern is a creational design pattern that separates the construction of complex objects from their representation. It allows you to create objects step by step, providing control over the construction process and the ability to produce different representations of an object using the same construction code.             | ✔ |
|  Factory Method  | Factory Method pattern is a creational design pattern that provides an interface for creating objects but allows subclasses to decide which class to instantiate. It encapsulates the object creation logic in a separate method (the factory method) and provides a way to delegate the actual object instantiation to subclasses.            | ✔ |
|    Prototype     | Prototype pattern is a creational design pattern that allows you to create new objects by cloning existing ones, without coupling the code to their specific classes. It is useful when the creation of an object is costly or complex, and you want to create new instances by copying existing ones, rather than creating them from scratch. | ✔ |
|    Singleton     | Singleton pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance. It is useful when you want to restrict the instantiation of a class to a single object and ensure that there is a single shared instance available throughout the application.             | ✔ |

## Structural Patterns

|  Pattern  | Description | Status |
|:---------:|:----------- |:------:|
|  Bridge   | Bridge pattern is a structural design pattern that decouples an abstraction from its implementation, allowing them to vary independently. It is useful when you want to separate the abstraction and implementation hierarchies, allowing them to evolve independently and providing flexibility in combining different implementations with different abstractions. | ✔ |
| Composite | Composite pattern is a structural design pattern that allows you to compose objects into tree-like structures and treat individual objects and compositions of objects uniformly. It enables clients to work with individual objects or groups of objects in a consistent manner. | ✔ |
| Decorator | Decorator pattern is a structural design pattern that allows behavior to be added to an object dynamically, without affecting the behavior of other objects in the same class. It provides a flexible alternative to subclassing for extending functionality. | ✔ |
|  Facade   |Facade pattern is a structural design pattern that provides a simplified interface to a complex system, making it easier to use and understand. It encapsulates a set of interfaces and classes behind a single high-level interface, simplifying the interaction with the underlying system. | ✔ |
| Flyweight | Flyweight pattern is a structural design pattern that allows for efficient sharing of fine-grained objects, reducing memory usage and improving performance. It is used when there are a large number of similar objects that can be shared and when the storage costs associated with creating multiple instances of those objects are high.| ✘ |
| Proxy     | Proxy pattern is a structural design pattern that provides a surrogate or placeholder for another object to control access to it. It allows you to create a proxy object that acts as an intermediary between the client and the real object. The proxy object can add additional functionality or provide a level of indirection to the real object. | ✔ |

## Behavioral Patterns

|              Pattern              | Description                                                                                                                                                                                                                                                                                                                                                           | Status |
|:---------------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------:|
|      Chain of Responsibility      | Chain of Responsibility pattern is a behavioral design pattern that allows an object to pass a request along a chain of potential handlers until one of them handles the request. It decouples the sender of the request from the receiver, giving multiple objects the opportunity to handle the request without explicitly knowing which object will handle it.     |   ✘    |
|              Command              | Command pattern is a behavioral design pattern that encapsulates a request as an object, thereby allowing you to parameterize clients with different requests, queue or log requests, and support undoable operations. It decouples the requester of the command from the object that performs the operation, enabling greater flexibility and extensibility.         |    ✔    |
|             Mediator              | Mediator pattern is a behavioral design pattern that promotes loose coupling between objects by introducing a mediator object that encapsulates the communication between them. It allows objects to communicate with each other without having explicit references to one another, reducing dependencies and simplifying the interactions.                           |   ✘    |
|              Memento              | Memento pattern is a behavioral design pattern that allows you to capture and restore an object's internal state without violating encapsulation. It provides a way to save and restore the state of an object at different points in time, allowing the object to revert to a previous state as needed.                                                              |   ✘    |
|             Observer              | Observer pattern is a behavioral design pattern that establishes a one-to-many relationship between objects, where changes in one object are automatically communicated to and reflected in other dependent objects. It allows for loose coupling between objects, ensuring that changes in one object do not require changes in other objects.                       |   ✔    |
|               State               | State pattern is a behavioral design pattern that allows an object to alter its behavior when its internal state changes. It encapsulates different states of an object into separate classes, allowing the object to transition between states dynamically.                                                                                                          |   ✘    |
|             Strategy              | Strategy pattern is a behavioral design pattern that allows you to define a family of interchangeable algorithms or strategies and encapsulate each one as an object. It enables the algorithms to be selected at runtime based on the specific context or requirements, without the client code being aware of the differences between them.                         |   ✘    |
|             Template              | Template pattern is a behavioral design pattern that defines the skeleton of an algorithm in a base class, allowing subclasses to provide their own implementation for certain steps of the algorithm while keeping the overall structure unchanged. It promotes code reuse and provides a way to define a common algorithm with variations in specific steps. |   ✘    |
|              Visitor              | Visitor pattern is a behavioral design pattern that allows you to separate the algorithms or operations from the objects on which they operate. It enables you to add new operations to existing object structures without modifying those objects' classes.|   ✘    |

####  ✔ Done
####  ✘ To be added in the future
