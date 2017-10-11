# Go design patterns
Collection of design pattern implementation in Go

## Creational patterns

| Pattern          | Description |
|------------------|-------------|
| [Abstract Factory](https://github.com/girikuncoro/go-design-pattern/blob/master/abstract-factory/abstract-factory.go) | Provide an interface for creating families of related or dependent objects without specifying their concrete classes.        |
| [Builder](https://github.com/girikuncoro/go-design-pattern/blob/master/builder/builder.go)                 | Separate the construction of a complex object from its representation so that the same construction process can create different representations.            | 
| [Factory Method](https://github.com/girikuncoro/go-design-pattern/blob/master/factory/factory.go)                 | Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory method lets a class defer instantiation to subclasses.            |
| [Object Pool](https://github.com/girikuncoro/go-design-pattern/blob/master/object-pool/object-pool.go)                 |             |
| [Prototype](https://github.com/girikuncoro/go-design-pattern/blob/master/prototype/prototype.go)                 | Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.            |
| [Singleton](https://github.com/girikuncoro/go-design-pattern/blob/master/singleton/singleton.go)                 |  Ensure a class only has one instance, and provide a global point of access to it.           |

## Structural patterns
| Pattern          | Description |
|------------------|-------------|
| [Adapter](https://github.com/girikuncoro/go-design-pattern/blob/master/adapter/adapter.go) | Convert the interface of a class into another interface clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.       |
| [Bridge](https://github.com/girikuncoro/go-design-pattern/blob/master/bridge/bridge.go) | Decouple an abstraction from its implementation so that the two can cary independently.       |
| [Composite](https://github.com/girikuncoro/go-design-pattern/blob/master/composite/composite.go) | Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.       |
| [Decorator](https://github.com/girikuncoro/go-design-pattern/blob/master/decorator/decorator.go) | Attach additional responsibilities to an objet dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.       |
| [Facade](https://github.com/girikuncoro/go-design-pattern/blob/master/facade/facade.go) | Provide a unified interface to a set of interfaces in a subsystem. Faade defines a higher-level interface that makes the subsystem easier to use.       |
| [Flyweight](https://github.com/girikuncoro/go-design-pattern/blob/master/flyweight/flyweight.go) | Use sharing to support large numbers of fine-grained objects efficiently.       |
| [Private Class Data](https://github.com/girikuncoro/go-design-pattern/blob/master/private-class-data/private-class-data.go) |        |
| [Proxy](https://github.com/girikuncoro/go-design-pattern/blob/master/proxy/proxy.go) | Provide a surrogate or placeholder for another object to control access to it.       |

## Behavioral patterns
| Pattern          | Description |
|------------------|-------------|
| [Chain of Responsibility](https://github.com/girikuncoro/go-design-pattern/blob/master/chain-of-responsibility/chain-of-responsibility.go) |        |
| [Command](https://github.com/girikuncoro/go-design-pattern/blob/master/command/command.go) |        |
| [Interpreter](https://github.com/girikuncoro/go-design-pattern/blob/master/interpreter/interpreter.go) |        |
| [Iterator](https://github.com/girikuncoro/go-design-pattern/blob/master/iterator/iterator.go) |        |
| [Mediator](https://github.com/girikuncoro/go-design-pattern/blob/master/mediator/mediator.go) |        |
| [Memento](https://github.com/girikuncoro/go-design-pattern/blob/master/memento/memento.go) |        |
| [Null Object](https://github.com/girikuncoro/go-design-pattern/blob/master/null-object/null-object.go) |        |
| [Observer](https://github.com/girikuncoro/go-design-pattern/blob/master/observer/observer.go) |        |
| [State](https://github.com/girikuncoro/go-design-pattern/blob/master/state/state.go) |        |
| [Strategy](https://github.com/girikuncoro/go-design-pattern/blob/master/strategy/strategy.go) |        |
| [Template](https://github.com/girikuncoro/go-design-pattern/blob/master/template/template.go) |        |
| [Visitor](https://github.com/girikuncoro/go-design-pattern/blob/master/visitor/visitor.go) |        |

### References
* [Design Patterns: Elements Reusable Object Book](https://www.amazon.com/Design-Patterns-Elements-Reusable-Object-Oriented/dp/0201633612)
* [Sourcemaking](https://sourcemaking.com/design_patterns)
