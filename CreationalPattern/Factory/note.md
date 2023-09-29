# Factory

![](https://refactoringguru.cn/images/patterns/diagrams/factory-method/structure-2x.png)

# How to Use Factory

1. Define a `Product` interface.
2. Define a `Creator` interface that creates a `Product` using the `createProduct()` method.
2. Implement Product for all product in your project.
3. Write concrete creator for your product.
4. Assign a concrete creator to the Creator interface based on the context, and utilize the Creator.createProduct()
   method to obtain a product.

# When to use it

- When you need to decouple the code for creating a product and using a product.
    - For example, if you want to automatically retrieve a specific product based on the context, without being
      concerned about its specific implementation but only whether it can be used properly. Without decoupling the
      creation and usage dependencies, you would need to specify the type of product before using it, which would
      require adding if...else statements when introducing a new product. The factory pattern can easily solve this
      problem.
- When you want to support a new product
    - For example, when you have `table` and `seat` and you want to add support for a new product like a `sofa`.

# Pros and Cons

| Pros                                                                                                                                                         | Cons                                                                                                                     |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Avoid coupling between the code for creating a product and the code for using the product.                                                                   | Using the Factory Method pattern may introduce several new subclasses, which can potentially make the code more complex. |
| Comply with the Open-Closed Principle, allowing the introduction of new product types in the program without modifying existing code.                        |                                                                                                                          |
| Single Responsibility Principle. You can place the code for product creation in a single location within the program, making it easier to maintain the code. |                                                                                                                          |

# Simple Factory

The Factory method pattern becomes a Simple Factory pattern when the Creator interface is removed, as shown in the code
below:

```java
class UserFactory {
    public static function create($type) {
        switch ($type) {
            case 'user': return new User();
            case 'customer': return new Customer();
            case 'admin': return new Admin();
            default:
                throw new Exception('传递的用户类型错误。');
        }
    }
}
```

When you create an abstract `Creator` interface for each specific creator, the Simple Factory pattern is transformed
into the Factory Method pattern.









