# Adapter

![](https://refactoringguru.cn/images/patterns/diagrams/adapter/structure-object-adapter-2x.png)

# Example

![](https://refactoring.guru/images/patterns/diagrams/adapter/example-2x.png)

```
// Say you have two classes with compatible interfaces:
// RoundHole and RoundPeg.
class RoundHole is
    constructor RoundHole(radius) { ... }

    method getRadius() is
        // Return the radius of the hole.

    method fits(peg: RoundPeg) is
        return this.getRadius() >= peg.getRadius()

class RoundPeg is
    constructor RoundPeg(radius) { ... }

    method getRadius() is
        // Return the radius of the peg.


// But there's an incompatible class: SquarePeg.
class SquarePeg is
    constructor SquarePeg(width) { ... }

    method getWidth() is
        // Return the square peg width.


// An adapter class lets you fit square pegs into round holes.
// It extends the RoundPeg class to let the adapter objects act
// as round pegs.
class SquarePegAdapter extends RoundPeg is
    // In reality, the adapter contains an instance of the
    // SquarePeg class.
    private field peg: SquarePeg

    constructor SquarePegAdapter(peg: SquarePeg) is
        this.peg = peg

    method getRadius() is
        // The adapter pretends that it's a round peg with a
        // radius that could fit the square peg that the adapter
        // actually wraps.
        return peg.getWidth() * Math.sqrt(2) / 2


// Somewhere in client code.
hole = new RoundHole(5)
rpeg = new RoundPeg(5)
hole.fits(rpeg) // true

small_sqpeg = new SquarePeg(5)
large_sqpeg = new SquarePeg(10)
hole.fits(small_sqpeg) // this won't compile (incompatible types)

small_sqpeg_adapter = new SquarePegAdapter(small_sqpeg)
large_sqpeg_adapter = new SquarePegAdapter(large_sqpeg)
hole.fits(small_sqpeg_adapter) // true
hole.fits(large_sqpeg_adapter) // false
```

`SquarePegAdapter` is a subclass of `RoundPeg`. It accepts an object of `SquarePeg`, and its behavior appears to be the
same as `RoundPeg`. Therefore, `SquarePegAdapter` can be passed as an argument to a `RoundHole` that expects
a `RoundPeg`, and the `fit()` result can be obtained. In this case, `SquarePegAdapter` acts as an Adapter.

# When to use it

1. Use the Adapter class when you want to use some existing class, but its interface isn’t compatible with the rest of
   your code.

   The Adapter pattern lets you create a middle-layer class that serves as a translator between your code and a legacy
   class, a 3rd-party class or any other class with a weird interface.

   In the above case, we want to fit a `SquarePeg` into a `RoundHole`, but it's not possible to compile because
   the `RoundHole` expects a `RoundPeg` while we are providing a `SquarePeg`. To solve this problem, an adapter can be
   created. We can define a `SquarePegAdapter`, which is a subclass of `RoundPeg`, and it will internally hold a
   reference to a `SquarePeg`. By rewriting the `getRadius()` method in the adapter, we can make the `SquarePegAdapter`
   behave like a `RoundPeg`.

   The revised description clarifies that the adapter is needed to bridge the incompatibility between `SquarePeg`
   and `RoundHole`, and emphasizes the purpose of the `SquarePegAdapter` class.

2. Use the pattern when you want to reuse several existing subclasses that lack some common functionality that can’t be
   added to the superclass.

   Let's take an example to illustrate this. In a scenario of graphic drawing, there are multiple classes representing
   different shapes, such as Square, Circle, and Triangle. Each shape has a `draw()` method to draw itself.

   Now, suppose we need to add a new method `resize()` to each shape class to adjust the size of the shape. However, not
   all shape classes have this method. Without using an adapter, we would need to manually add the `resize()` method to
   each shape class, even though some shape classes may not actually need this method.

   By using the adapter pattern, we can create an adapter interface called `ShapeAdapter` and make it implement
   the `resize()` method. Then, we encapsulate the shape objects that need the `resize()` method within the adapter.
   This way, the shape classes that don't have the `resize()` method don't need to be modified before using the adapter.
   Only the shape classes that require resizing will have their objects placed inside the adapter, thus dynamically
   gaining the functionality of the `resize()` method.

   ```go
   type Shape interface {
       draw()
   }
   
   type Square struct {}
   
   func (s *Square) draw() {
       fmt.Println("Drawing a square.")
   }
   
   
   type Circle struct {}
   
   func (c *Circle) draw() {
       fmt.Println("Drawing a circle.")
   }
   
   // adapter
   type ShapeAdapter struct {
       shape Shape // common interface
   }
   
   func (sa *ShapeAdapter) draw() {
       sa.shape.draw()
   }
   
   func (sa *ShapeAdapter) resize() {
       fmt.Println("Resizing the shape.")
   }
   
   square := &Square{}
   circle := &Circle{}
   
   square.draw() // output: Drawing a square.
   circle.draw() // output: Drawing a circle.
   
   adapter := &ShapeAdapter{shape: square}
   adapter.resize() // output: Resizing the shape.
   
   adapter = &ShapeAdapter{shape: circle}
   adapter.resize() // output: Resizing the shape.
   ```

   Through the adapter pattern, we can dynamically add missing functionality to graphic objects without modifying the
   original shape classes and maintain a unified interface.

   The adapter pattern complies with the **Open-Closed Principle**, as long as there is a common interface. Adapters can
   dynamically add new functionality to objects.

# How to use it

1. Make sure that you have at least two classes with incompatible interfaces:
    - A useful *service* class, which you can’t change (often 3rd-party, legacy or with lots of existing dependencies).
    - One or several *client* classes that would benefit from using the service class.
2. Declare the client interface and describe how clients communicate with the service.
3. Create the adapter class and make it follow the client interface. Leave all the methods empty for now.
4. Add a field to the adapter class to store a reference to the service object. The common practice is to initialize
   this field via the constructor, but sometimes it’s more convenient to pass it to the adapter when calling its
   methods.
5. One by one, implement all methods of the client interface in the adapter class. The adapter should delegate most of
   the real work to the service object, handling only the interface or data format conversion.
6. Clients should use the adapter via the client interface. This will let you change or extend the adapters without
   affecting the client code.

```go

package main

import "fmt"

// Client interface
type CalculatorInterface interface {
	Calculate(num1, num2 int) int
}

// Adapter
type CalculatorAdapter struct {
	LegacyCalculator *LegacyCalculator
}

func (a *CalculatorAdapter) Calculate(num1, num2 int) int {
	return a.LegacyCalculator.CalculateSum(num1, num2)
}

// Server, Usually it is code in a third-party library and cannot be modified
type LegacyCalculator struct{}

func (c *LegacyCalculator) CalculateSum(num1, num2 int) int {
	return num1 + num2
}

// Client class
type ModernCalculator struct {
	Calculator CalculatorInterface
}

func (m *ModernCalculator) PerformCalculation(num1, num2 int) {
	result := m.Calculator.Calculate(num1, num2)
	fmt.Println("Result:", result)
}

func main() {
	// We have a LegacyCalculator that cannot be modified
	legacyCalculator := &LegacyCalculator{}
	// pass legacyCalculator to Adapter
	adapter := &CalculatorAdapter{LegacyCalculator: legacyCalculator}
	// pass adapter to client class
	modernCalculator := &ModernCalculator{Calculator: adapter}
	modernCalculator.PerformCalculation(10, 5)
}
```

With the adapter pattern, we have successfully adapted the LegacyCalculator to work with ModernCalculator. Now,
ModernCalculator can use the adapter by calling the PerformCalculation() method without directly interacting with the
original LegacyCalculator.

In summary, the adapter pattern allows us to make incompatible classes work together in Go. By creating an adapter that
wraps and transforms existing classes, we can fulfill the interface requirements of the client.

# Pros and Cons

| Pros                                                                                                                                                                                            | Cons                                                                                                                                                                                                                                                       |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Single Responsibility Principle: You can separate the interface or data conversion code from the primary business logic of the program.                                                         | The overall complexity of the code increases due to the introduction of new interfaces and classes.                         Sometimes it may be simpler to modify the service class instead of introducing an adapter if it matches the existing codebase. |
| Open/Closed Principle: You can introduce new types of adapters into the program without breaking the existing client code, as long as they work with the adapters through the client interface. |                                                                                                                                                                                                                                                            |