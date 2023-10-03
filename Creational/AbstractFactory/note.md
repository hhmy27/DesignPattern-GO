# Abstract Factory

![](https://refactoringguru.cn/images/patterns/diagrams/abstract-factory/structure-2x.png)

# When to use it

- When your products have versatile styles
    - For example, originally your code only supported buttons and checkboxes for the Mac platform. As your business
      expands, you need to support buttons and checkboxes for the Windows platform as well. In this case, you can
      consider designing the Creator based on the platform and set the product type (button, checkbox) accordingly.

# How to use abstract factory

1. Creating a matrix with different product types and their variants as dimensions.
   ![](https://refactoring.guru/images/patterns/diagrams/abstract-factory/problem-en-2x.png)
2. Declare abstract product interfaces for all product types. Then make all concrete product classes implement these
   interfaces.

   for example, `chair`, `sofa`, `table` is three abstract product.
3. Declare the abstract factory interface with a set of creation methods for all abstract products.

   e.g. `createChair()`, `createSofa()`, `createTable()`.
4. Implement a set of concrete factory classes, one for each product variant.

   e.g. `ModernCreator`
5. Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes,
   depending on the application configuration or the current environment. Pass this factory object to all classes that
   construct products.
6. Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate
   creation method on the factory object.

# Pros and Cons

Basically same to Factory Pattern.

| Pros                                                                                                                                                         | Cons                                                                                                                              |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| Avoid coupling between the code for creating a product and the code for using the product.                                                                   | Using the Abstract Factory Method pattern may introduce several new subclasses, which can potentially make the code more complex. |
| Comply with the Open-Closed Principle, allowing the introduction of new product types in the program without modifying existing code.                        |                                                                                                                                   |
| Single Responsibility Principle. You can place the code for product creation in a single location within the program, making it easier to maintain the code. |                                                                                                                                   |

