# Singleton

Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a
global access point to this instance.

![](https://refactoring.guru/images/patterns/diagrams/singleton/structure-en-2x.png)

like this:
Clients may not even realize that they’re working with the same object all the time.
![](https://refactoring.guru/images/patterns/content/singleton/singleton-comic-1-en-2x.png)

two keys:

- **Ensure that a class has just a single instance.**

  Sometimes, we would like certain shared resources to have only one instance, such as a database connection or a file.
- **Provide a global access point to that instance.**

  Just like a global variable, the Singleton pattern lets you access some object from anywhere in the program.

Nowadays, the Singleton pattern has become so popular that people may call something a singleton even if it solves just
one of the listed problems.

# When to use it

When you need the characteristics of a singleton (global uniqueness and global accessibility), you can use it.

# How to use it

All implementations of the Singleton have these two steps in common:

- Make the default constructor private, to prevent other objects from using the new operator with the Singleton class.
- Create a static creation method that acts as a constructor. Under the hood, this method calls the private constructor
  to create an object and saves it in a static field. All following calls to this method return the cached object.

If your code has access to the Singleton class, then it’s able to call the Singleton’s static method. So whenever that
method is called, the same object is always returned.

# Pseudocode

```go
// The Database class defines the `getInstance` method that lets
// clients access the same instance of a database connection
// throughout the program.
class Database is
// The field for storing the singleton instance should be
// declared static.
private static field instance: Database

// The singleton's constructor should always be private to
// prevent direct construction calls with the `new`
// operator.
private constructor Database() is
// Some initialization code, such as the actual
// connection to a database server.
// ...

// The static method that controls access to the singleton
// instance.
public static method getInstance() is
if (Database.instance == null) then
acquireThreadLock() and then
// Ensure that the instance hasn't yet been
// initialized by another thread while this one
// has been waiting for the lock's release.
if (Database.instance == null) then
Database.instance = new Database()
return Database.instance

// Finally, any singleton should define some business logic
// which can be executed on its instance.
public method query(sql) is
// For instance, all database queries of an app go
```

note `getInstance()` use double-check to create an instance.

Q: Why is `acquireThreadLock()` located in the first if judgment, instead of locking first and then judging?

A: **Placing acquireThreadLock() inside the first if statement is to avoid unnecessary locking operations**. If we
directly lock first and then check, even if the instance has already been created, each call to the getInstance() method
would still perform a locking operation, resulting in performance overhead. By placing acquireThreadLock() inside the
first if statement, the locking operation is only performed when the instance is not yet initialized, ensuring thread
safety.

Q: Why do we still need to judge `if (Database.instance == null)` after locking?

A: To prevent the issue of **multiple instance creations** in a multi-threaded environment. When multiple threads pass
the first if check simultaneously and the instance is still null, only one thread will acquire the lock and create the
instance. Other waiting threads, once they acquire the lock, need to perform the second check to ensure that the
instance is still null before creating another instance. This ensures that only one instance is created and assigned to
the Database.instance variable.

# Pros and Cons

| Pros                                                                             | Cons                                                                                                                                                                                                                                                                                                                                                                                                      |
|----------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| You can be sure that a class has only a single instance.                         | The Singleton pattern can mask bad design, for instance, when the components of the program know too much about each other.                                                                                                                                                                                                                                                                               |
| You gain a global access point to that instance.                                 | The pattern requires special treatment in a multithreaded environment so that multiple threads won't create a singleton object several times.                                                                                                                                                                                                                                                             |
| The singleton object is initialized only when it's requested for the first time. | It may be difficult to unit test the client code of the Singleton because many test frameworks rely on inheritance when producing mock objects. Since the constructor of the singleton class is private and overriding static methods is impossible in most languages, you will need to think of a creative way to mock the singleton. Or just don't write the tests. Or don't use the Singleton pattern. |

# Relations with Other Patterns
- Abstract Factories, Builders and Prototypes can all be implemented as Singletons.

