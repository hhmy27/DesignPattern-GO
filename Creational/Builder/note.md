# Builder

![](https://refactoringguru.cn/images/patterns/diagrams/builder/structure-2x.png)

# When to use it

- When your product needs to go through several steps to be generated.
- When your product's constructor is complex, you can use a builder to generate the product step by step.

# How to use

1. Define the steps for generating a product in a generic manner, Declare these steps within the Builder.
2. Create a concrete builder class for each of the product representations and implement their construction steps.
   You may need to use `getResult()` to obtain a specific product when there is no common interface between the
   products.
3. Think about creating a director class. It may encapsulate various ways to construct a product using the same builder
   object.
4. The client code creates both the builder and the director objects. Before construction starts, the client must pass a
   builder object to the director. Usually, the client does this only once, via parameters of the director’s class
   constructor. The director uses the builder object in all further construction. There’s an alternative approach, where
   the builder is passed to a specific product construction method of the director.
5. The construction result can be obtained directly from the director only if all products follow the same interface.
   Otherwise, the client should fetch the result from the builder.

# Pros and Cons

| Pros                                                                                                                  | Cons                                                                                                      |
|-----------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| -  You can construct objects step-by-step, defer construction steps or run steps recursively.                         | -  The overall complexity of the code increases since the pattern requires creating multiple new classes. |
| -  You can reuse the same construction code when building various representations of products.                        |                                                                                                           |
| -  Single Responsibility Principle. You can isolate complex construction code from the business logic of the product. |                                                                                                           |



