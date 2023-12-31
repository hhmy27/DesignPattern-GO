# Decorator

![](https://refactoring.guru/images/patterns/diagrams/decorator/structure-indexed-2x.png)

**Decorator** is a structural design pattern that lets you attach new behaviors to objects by placing these objects
inside special wrapper objects that contain the behaviors.

During your decoration process, objects are pushed onto the stack step by step.
![](https://refactoring.guru/images/patterns/content/decorator/decorator-2x.png)

**Problem**

Imagine that you’re working on a notification library which lets other programs notify their users about important
events.

The initial version of the library was based on the Notifier class that had only a few fields, a constructor and a
single send method. The method could accept a message argument from a client and send the message to a list of emails
that were passed to the notifier via its constructor. A third-party app which acted as a client was supposed to create
and configure the notifier object once, and then use it each time something important happened.

![](https://refactoring.guru/images/patterns/diagrams/decorator/problem1-en-2x.png)

At some point, you realize that users of the library expect more than just email notifications. Many of them would like
to receive an SMS about critical issues. Others would like to be notified on Facebook and, of course, the corporate
users would love to get Slack notifications.

![](https://refactoring.guru/images/patterns/diagrams/decorator/problem2-2x.png)

If you use inheritance to accomplish this requirement, the number of subclasses you have will increase exponentially.

![](https://refactoring.guru/images/patterns/diagrams/decorator/problem3-2x.png)

**Solution**

One of the ways to overcome these caveats is by using **Aggregation** or **Composition** instead of Inheritance. Both of the alternatives work almost the same way: one object has a reference to another and delegates it some work, whereas with inheritance, the object itself is able to do that work, inheriting the behavior from its superclass.

- Aggregation: object A contains objects B; B can live without A.
- Composition: object A consists of objects B; A manages life cycle of B; B can’t live without A.

With this new approach you can easily substitute the linked “helper” object with another, changing the behavior of the container at runtime. An object can use the behavior of various classes, having references to multiple objects and delegating them all kinds of work. Aggregation/composition is the key principle behind many design patterns, including Decorator. On that note, let’s return to the pattern discussion.

![](https://refactoring.guru/images/patterns/diagrams/decorator/solution1-en-2x.png)

“Wrapper” is the alternative nickname for the Decorator pattern that clearly expresses the main idea of the pattern. A wrapper is an object that can be linked with some target object. The wrapper contains the same set of methods as the target and delegates to it all requests it receives. However, the wrapper may alter the result by doing something either before or after it passes the request to the target.

When does a simple wrapper become the real decorator? As I mentioned, the wrapper implements the same interface as the wrapped object. That’s why from the client’s perspective these objects are identical. Make the wrapper’s reference field accept any object that follows that interface. This will let you cover an object in multiple wrappers, adding the combined behavior of all the wrappers to it.

In our notifications example, let’s leave the simple email notification behavior inside the base Notifier class, but turn all other notification methods into decorators.


![](https://refactoringguru.cn/images/patterns/diagrams/decorator/solution2-zh-2x.png?id=a29ec4128bd712b71f6939ad3450495e)
The client code would need to wrap a basic notifier object into a set of decorators that match the client’s preferences. The resulting objects will be structured as a stack.


![](https://refactoringguru.cn/images/patterns/diagrams/decorator/solution3-zh-2x.png?id=8c9f86759d233505d4be4eb0b99711e3)
The last decorator in the stack would be the object that the client actually works with. Since all decorators implement the same interface as the base notifier, the rest of the client code won’t care whether it works with the “pure” notifier object or the decorated one.

We could apply the same approach to other behaviors such as formatting messages or composing the recipient list. The client can decorate the object with any custom decorators, as long as they follow the same interface as the others.


![](https://refactoring.guru/images/patterns/content/decorator/decorator-comic-1-2x.png)

Wearing clothes is an example of using decorators. When you’re cold, you wrap yourself in a sweater. If you’re still cold with a sweater, you can wear a jacket on top. If it’s raining, you can put on a raincoat. All of these garments “extend” your basic behavior but aren’t part of you, and you can easily take off any piece of clothing whenever you don’t need it.




# When to use it

1.  Use the Decorator pattern when you need to be able to assign extra behaviors to objects at runtime without breaking the code that uses these objects.

    The Decorator lets you structure your business logic into layers, create a decorator for each layer and compose objects with various combinations of this logic at runtime. The client code can treat all these objects in the same way, since they all follow a common interface.

2. Use the pattern when it’s awkward or not possible to extend an object’s behavior using inheritance.

   Many programming languages have the `final` keyword that can be used to prevent further extension of a class. For a final class, the only way to reuse the existing behavior would be to wrap the class with your own wrapper, using the Decorator pattern.
# How to use it

1. Make sure your business domain can be represented as a primary component with multiple optional layers over it.
2. Figure out what methods are common to both the primary component and the optional layers. Create a component interface and declare those methods there.
3. Create a concrete component class and define the base behavior in it.
4. Create a base decorator class. It should have a field for storing a reference to a wrapped object. The field should be declared with the component interface type to allow linking to concrete components as well as decorators. The base decorator must delegate all work to the wrapped object.
5. Make sure all classes implement the component interface.
6. Create concrete decorators by extending them from the base decorator. A concrete decorator must execute its behavior before or after the call to the parent method (which always delegates to the wrapped object).
7. The client code must be responsible for creating decorators and composing them in the way the client needs.

# Pros and Cons
