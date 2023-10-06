# Flyweight

**Flyweight** is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing
common parts of state between multiple objects instead of keeping all of the data in each object.
![](https://refactoringguru.cn/images/patterns/diagrams/flyweight/structure-2x.png)

Assuming there is an original object in the system, its properties can be divided into immutable and mutable based on
their states. Therefore, we can split them into two objects: `Flyweight` and `Context`. The `Flyweight` object stores
the immutable state, while the `Context` object stores the mutable state. For example, in a shooting game, the bullet's
texture and color remain constant, but its coordinates are subject to change. In this case, the texture and color can be
stored in the `Flyweight` object, while the coordinates and other information can be stored in the `Context` object.

Simultaneously, we reference the `Flyweight` object within the `Context` object. This way, the combination
of `Flyweight` and `Context` represents the original object together.

Since `Flyweight` objects can be reused, it is possible to design a `FlyweightFactory` to facilitate reusing
existing `Flyweights`.

# Pes

In this example, the Flyweight pattern helps to reduce memory usage when rendering millions of tree objects on a canvas.
![](https://refactoring.guru/images/patterns/diagrams/flyweight/example-2x.png)
The pattern extracts the repeating intrinsic state from a main Tree class and moves it into the flyweight class
TreeType.

Now instead of storing the same data in multiple objects, it’s kept in just a few flyweight objects and linked to
appropriate Tree objects which act as contexts. The client code creates new tree objects using the flyweight factory,
which encapsulates the complexity of searching for the right object and reusing it if needed.

```go
// The flyweight class contains a portion of the state of a
// tree. These fields store values that are unique for each
// particular tree. For instance, you won't find here the tree
// coordinates. But the texture and colors shared between many
// trees are here. Since this data is usually BIG, you'd waste a
// lot of memory by keeping it in each tree object. Instead, we
// can extract texture, color and other repeating data into a
// separate object which lots of individual tree objects can
// reference.
class TreeType is
field name
field color
field texture
constructor TreeType(name, color, texture) { ... }
method draw(canvas, x, y) is
// 1. Create a bitmap of a given type, color & texture.
// 2. Draw the bitmap on the canvas at X and Y coords.

// Flyweight factory decides whether to re-use existing
// flyweight or to create a new object.
class TreeFactory is
static field treeTypes: collection of tree types
static method getTreeType(name, color, texture) is
type = treeTypes.find(name, color, texture)
if (type == null)
type = new TreeType(name, color, texture)
treeTypes.add(type)
return type

// The contextual object contains the extrinsic part of the tree
// state. An application can create billions of these since they
// are pretty small: just two integer coordinates and one
// reference field.
class Tree is
field x, y
field type: TreeType
constructor Tree(x, y, type) { ... }
method draw(canvas) is
type.draw(canvas, this.x, this.y)

// The Tree and the Forest classes are the flyweight's clients.
// You can merge them if you don't plan to develop the Tree
// class any further.
class Forest is
field trees: collection of Trees

method plantTree(x, y, name, color, texture) is
type = TreeFactory.getTreeType(name, color, texture)
tree = new Tree(x, y, type)
trees.add(tree)

method draw(canvas) is
foreach (tree in trees) do
tree.draw(canvas)
```

# When to use it

- Use the Flyweight pattern only when your program must support a huge number of objects which barely fit into available
  RAM.

The benefit of applying the pattern depends heavily on how and where it’s used. It’s most useful when:

- an application needs to spawn a huge number of similar objects
- this drains all available RAM on a target device
- the objects contain duplicate states which can be extracted and shared between multiple objects

# How to use it

1. Divide fields of a class that will become a flyweight into two parts:
    - the intrinsic state: the fields that contain unchanging data duplicated across many objects
    - the extrinsic state: the fields that contain contextual data unique to each object
2. Leave the fields that represent the intrinsic state in the class, but make sure they’re immutable. They should take
   their initial values only inside the constructor.
3. Go over methods that use fields of the extrinsic state. For each field used in the method, introduce a new parameter
   and use it instead of the field.
4. Optionally, create a factory class to manage the pool of flyweights. It should check for an existing flyweight before
   creating a new one. Once the factory is in place, clients must only request flyweights through it. They should
   describe the desired flyweight by passing its intrinsic state to the factory.
5. The client must store or calculate values of the extrinsic state (context) to be able to call methods of flyweight
   objects. For the sake of convenience, the extrinsic state along with the flyweight-referencing field may be moved to
   a separate context class.

# Pros and Cons

| Pros                                                              | Cons                                                                                                                                            |
|-------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| You can save lots of RAM by using flyweights for similar objects. | There might be a tradeoff between RAM and CPU cycles when certain context data needs to be recalculated each time a flyweight method is called. |
|                                                                   | The code becomes more complicated, which can confuse new team members who may question the separation of entity state in this manner.           |