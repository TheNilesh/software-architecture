SOLID Principles
================

Guidelines for programming that encourages writing easy to maintain code.

Single Responsibility Principle (SRP)
-------------------------------------

A class, function or a package should only have one and only one responsibility. Break down it if evolves to handle multiple responsibilities.

The SRP states that a class should have only one reason to change. This means that a class should be responsible for a single well-defined task. If a class is responsible for multiple tasks, it will be more difficult to understand and maintain.

Open/Closed Principle (OCP)
---------------------------

A class or any program entity should be open for extension, but closed for modifications. This means that new functionality should be added by extending existing classes and modules, rather than modifying them.

Liskov Substitution Principle (LSP)
-----------------------------------

The LSP states that subtypes should be substitutable for their base types. This means that subtypes should be able to implement all of the functionality of their base types, and they should not introduce any new errors.

Interface Segregation Principle (ISP)
-------------------------------------

A client should not be forced to depend on interfaces it does not use. This principle encourages small, specific interfaces. The ISP states that interfaces should be segregated into smaller, more specific interfaces. This makes the interfaces easier to understand and use. It also reduces coupling between classes.

Dependency Inversion Principle (DIP)
------------------------------------

The DIP states that high-level modules should not depend on low-level modules. Both should depend on abstractions. This makes the system more flexible and reusable.
