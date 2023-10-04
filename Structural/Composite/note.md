# Composite

Composite is a structural design pattern that lets you compose objects into **tree structures** and then work with these
structures as if they were individual objects.

Using the Composite pattern makes sense only when the core model of your app can be represented as a **tree**.

For example, imagine that you have two types of objects: `Products` and `Boxes`. A Box can contain several Products as
well as a number of smaller Boxes. These little Boxes can also hold some Products or even smaller Boxes, and so on.

![](https://refactoring.guru/images/patterns/diagrams/composite/problem-en-2x.png)

Say you decide to create an ordering system that uses these classes. Orders could contain simple products without any
wrapping, as well as boxes stuffed with products...and other boxes. How would you determine the total price of such an
order?

If you are familiar with data structures, you will notice that the box and product here together form a multi-branch
tree, where the product acts as the leaf node with its weight representing its value. Calculating the total price of the
order is equivalent to calculating the total value of the entire tree. The problem can be effectively solved using the
recursive approach.

**Solution**

The Composite pattern suggests that you work with Products and Boxes through a common interface which declares a method
for calculating the total price.

How would this method work? For a product, it’d simply return the product’s price. For a box, it’d go over each item the
box contains, ask its price and then return a total for this box. If one of these items were a smaller box, that box
would also start going over its contents and so on, until the prices of all inner components were calculated. A box
could even add some extra cost to the final price, such as packaging cost.

**Structure**

![](https://refactoring.guru/images/patterns/diagrams/composite/structure-en-2x.png)

# When to use it

- Use the Composite pattern when you have to implement a tree-like object structure

The Composite pattern provides you with two basic element types that share a common interface: simple leaves and complex
containers. A container can be composed of both leaves and other containers. This lets you construct a nested recursive
object structure that resembles a tree.

- Use the pattern when you want the client code to treat both simple and complex elements uniformly.

All elements defined by the Composite pattern share a common interface. Using this interface, the client doesn’t have to
worry about the concrete class of the objects it works with.

# How to use it

1. Make sure that the core model of your app can be represented as a *tree* structure. Try to break it down into simple
   elements and containers. Remember that containers must be able to contain both simple elements and other containers.

2. Declare the component interface with a list of methods that make sense for both simple and complex components.

3. Create a leaf class to represent simple elements. A program may have multiple different leaf classes.

4. Create a container class to represent complex elements. In this class, provide an array field for storing references
   to sub-elements. The array must be able to store both leaves and containers, so make sure it’s declared with the
   component interface type.

   While implementing the methods of the component interface, remember that a container is supposed to be delegating
   most of the work to sub-elements.

5. Finally, define the methods for adding and removal of child elements in the container.

   Keep in mind that these operations can be declared in the component interface. This would violate the *Interface
   Segregation Principle* because the methods will be empty in the leaf class. However, the client will be able to treat
   all the elements equally, even when composing the tree.

# Pros and Cons

| Pros                                                                                                                                              | Cons                                                                                                                                                                                                              |
|---------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| You can work with complex tree structures more conveniently: use polymorphism and recursion to your advantage.                                    | It might be difficult to provide a common interface for classes whose functionality differs too much. In certain scenarios, you’d need to overgeneralize the component interface, making it harder to comprehend. |
| Open/Closed Principle. You can introduce new element types into the app without breaking the existing code, which now works with the object tree. |                                                                                                                                                                                                                   |