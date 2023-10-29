Single Responsibility Principle (SRP)
======================================

A class, function or a package should only have one and only one responsibility. Break down it if evolves to handle multiple responsibilities.

The SRP states that a class should have only one reason to change. This means that a class should be responsible for a single well-defined task. If a class is responsible for multiple tasks, it will be more difficult to understand and maintain.

In Golang, the SRP can be implemented by dividing code into smaller, more focused packages. Each package should be responsible for a single task, and should not depend on other packages unnecessarily.

Violation of SRP
-----------------

    func IsValidPassword(userName, password string) (string, error) {
        records := db.Exec("SELECT * FROM login WHERE USER_NAME=%s AND PASSWORD=%s", userName, password)
        if len(records) == 0 {
            // create user with default password
            db.Exec("INSERT INTO login(USER_NAME,PASSWORD) VALUES(%s, default)", userName)
            return "userCreated", nil
        }
        return "success", nil
    }
