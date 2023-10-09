# Proxy

![](https://refactoring.guru/images/patterns/diagrams/proxy/structure-2x.png)

# When to use it

Imagine our train ticket purchasing service. If the whole country uses the same site to purchase tickets, the pressure
on this site will be huge.

If we set up train ticket sales points all over the country, and people nearby can go directly to the nearest sales
point to buy tickets, this can well disperse the pressure on the main station.

In the above process, the train ticket sales point is an agent for the main ticket sales station.

So under what circumstances will we use the proxy mode?
Here are some common usage scenarios:

- Lazy initialization (virtual proxy). This is when you have a heavyweight service object that wastes system resources
  by being always up, even though you only need it from time to time.

  For example, if you want to delay initialization of a database connection and connect to the database only when there
  is a query, instead of connecting to the database when the project is started, then the proxy mode is a good choice.

- Access control (protection proxy). This is when you want only **specific clients** to be able to use the service
  object; for instance, when your objects are crucial parts of an operating system and clients are various launched
  applications (including malicious ones).

  Through the proxy we can choose which requests to perform.

- Local execution of a remote service (remote proxy). This is when the service object is located on a remote server.

  In this case, the proxy passes the client request over the network, handling all of the nasty details of working with
  the network.

- Logging requests (logging proxy). This is when you want to keep a history of requests to the service object.

  The proxy can log each request before passing it to the service.

- Caching request results (caching proxy). This is when you need to cache results of client requests and manage the life
  cycle of this cache, especially if results are quite large.

- Smart reference. This is when you need to be able to dismiss a heavyweight object once there are no clients that use
  it.

  The proxy can keep track of clients that obtained a reference to the service object or its results. From time to time,
  the proxy may go over the clients and check whether they are still active. If the client list gets empty, the proxy
  might dismiss the service object and free the underlying system resources.

  The proxy can also track whether the client had modified the service object. Then the unchanged objects may be reused
  by other clients.

# How to use it

1. Create a service interface

   If there’s no pre-existing service interface, create one to make proxy and service objects interchangeable.
   Extracting the interface from the service class isn’t always possible, because you’d need to change all of the
   service’s clients to use that interface. Plan B is to make the proxy a **subclass** of the service class, and this
   way it’ll inherit the interface of the service.

2. Create the proxy class.

   It should have a field for **storing a reference to the service**. Usually, proxies create and manage the whole life
   cycle of their services. On rare occasions, a service is passed to the proxy via a constructor by the client.

3. Implement the proxy methods according to their purposes.

   In most cases, after doing some work, the proxy should delegate the work to the service object.

4. Consider introducing a creation method that decides whether the client gets a proxy or a real service. This can be a
   simple static method in the proxy class or a full-blown factory method.

5. Consider implementing lazy initialization for the service object.

# Pros and Cons

| Pros                                                                                             | Cons                                                                                      |
|--------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| 1. You can control the service object without clients knowing about it.                          | 1. The code may become more complicated since you need to introduce a lot of new classes. |
| 2. You can manage the lifecycle of the service object when clients don’t care about it.          | 2. The response from the service might get delayed.                                       |
| 3. The proxy works even if the service object isn’t ready or is not available.                   |                                                                                           |
| 4. Open/Closed Principle. You can introduce new proxies without changing the service or clients. |                                                                                           |

# Relations with Other Patterns

- [Adapter](https://github.com/hhmy27/DesignPattern-GO/tree/main/Structural/Adapter) provides a different interface to the wrapped object, Proxy provides it with the same interface, and Decorator provides it with an enhanced interface.

- [Facade](https://github.com/hhmy27/DesignPattern-GO/tree/main/Structural/Facade) is similar to Proxy in that both buffer a complex entity and initialize it on its own. Unlike Facade, Proxy has the same interface as its service object, which makes them interchangeable.

- [Decorator](https://github.com/hhmy27/DesignPattern-GO/tree/main/Structural/Decorator) and Proxy have similar structures, but very different intents. Both patterns are built on the composition principle, where one object is supposed to delegate some of the work to another. The difference is that a Proxy usually manages the life cycle of its service object on its own, whereas the composition of Decorators is always controlled by the client.

