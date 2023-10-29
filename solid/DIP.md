Dependency Inversion Principle (DIP)
====================================

The DIP states that high-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details; details should depend on abstractions. This makes the system more flexible and reusable.

Example
-------

Let's say we have a system that manages users. The system has a high-level module called UserService and a low-level module called UserRepository. The UserService module is responsible for managing user accounts, such as creating new accounts, updating existing accounts, and deleting accounts. The UserRepository module is responsible for interacting with the database to store and retrieve user data.

If the UserService module directly depends on the UserRepository module, then the UserService module will be tightly coupled to the database. This will make the system less flexible and reusable.

For example, if we want to change the database that the system uses, we will need to modify the UserService module to use the new database driver. This is because the UserService module directly depends on the UserRepository module, which is responsible for interacting with the database.

To decouple the UserService module from the UserRepository module, we can use the DIP. We can do this by creating an abstraction, such as a UserRepository interface. The UserRepository interface would define the methods that the UserService module needs to interact with the user data.

The UserRepository module would then implement the UserRepository interface. This would allow us to change the database that the system uses without modifying the UserService module.

Violation
----------

TODO: Write
