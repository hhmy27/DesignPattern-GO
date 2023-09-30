# Prototype

Basic implementation.
![](https://refactoringguru.cn/images/patterns/diagrams/prototype/structure-2x.png)

Prototype Registry
![](https://refactoringguru.cn/images/patterns/diagrams/prototype/structure-prototype-cache-2x.png)

# When to use it

- Use the Prototype pattern when your code shouldn’t depend on the concrete classes of objects that you need to copy.
  e.g.
   ```go
   type Shape interface {
       Clone() Shape
   }

   type Circle struct {
       radius float64
   }

   func (c *Circle) Clone() Shape {
       return &Circle{
           radius: c.radius,
       }
   }

   func main() {
       // Suppose you select a graphic in graphic editing and want to copy it
       selectedShape := &Circle{radius: 5.0}

       // Use Prototype before
       var copiedShape Shape
       if circle, ok := selectedShape.(*Circle); ok {
           copiedShape = &Circle{
               radius: circle.radius,
           }
       }

       // Use Prototype after
       copiedShape := selectedShape.Clone()
       
   }
   ```
- Use the pattern when you want to reduce the number of subclasses that only differ in the way they initialize their
  respective objects.

  Suppose we have several monsters, such as Werewolf, Vampire and Zombie. Each monster has some common properties and
  methods, but they may be initialized slightly differently.
   ```go
   type Monster interface {
       Attack()
   }

   type Werewolf struct {
       // Werewolf specific attributes
   }

   func (w *Werewolf) Attack() {...}

   type Vampire struct {
       // Vampire specific attributes
   }

   func (v *Vampire) Attack() {...}

   type Zombie struct {
       // Zombie specific attributes
   }

   func (z *Zombie) Attack() {...}
   ```
  Without using the prototype pattern, you might create a specific subclass for each monster and implement
  initialization logic in each subclass.

  However, if the only difference between these monsters is the way they are initialized, we can use the prototype
  pattern to reduce the number of subclasses and make the code more concise and extensible.

 ```go
 type Monster interface {
     Clone() Monster
     Attack()
 }

 type BaseMonster struct {
     // Shared attributes and method
 }

 func (bm *BaseMonster) Clone() Monster {
     return &BaseMonster{
         // clone shared attributes
     }
 }

 func (bm *BaseMonster) Attack() {
     // specify attack logic
 }
 ```
Using the prototype pattern, we can create instances of other monster objects by copying the prototype object without
having to explicitly create a large number of subclasses. In this way, we can effectively reduce the number of
subclasses and be able to extend and modify the monster's initialization logic more flexibly.
 ```go
 type Wolfman struct {
     BaseMonster
     // 狼人特定属性
 }

 func NewWolfman() Monster {
     return &Wolfman{
         BaseMonster: BaseMonster{},
         // 狼人特定属性的初始化
     }
 }
 ```

# How to use it

1. Create the prototype interface and declare the clone method in it. Or just add the method to all classes of an
   existing class hierarchy, if you have one.

2. A prototype class must define the alternative constructor that accepts an object of that class as an argument. The
   constructor must copy the values of all fields defined in the class from the passed object into the newly created
   instance. If you’re changing a subclass, you must call the parent constructor to let the superclass handle the
   cloning of its private fields.

   If your programming language doesn’t support method overloading, you won’t be able to create a separate “prototype”
   constructor. Thus, copying the object’s data into the newly created clone will have to be performed within the clone
   method. Still, having this code in a regular constructor is safer because the resulting object is returned fully
   configured right after you call the new operator.

3. The cloning method usually consists of just one line: running a new operator with the prototypical version of the
   constructor. Note, that every class must explicitly override the cloning method and use its own class name along with
   the new operator. Otherwise, the cloning method may produce an object of a parent class.

4. Optionally, create a centralized prototype registry to store a catalog of frequently used prototypes.

   You can implement the registry as a new factory class or put it in the base prototype class with a static method for
   fetching the prototype. This method should search for a prototype based on search criteria that the client code
   passes to the method. The criteria might either be a simple string tag or a complex set of search parameters. After
   the appropriate prototype is found, the registry should clone it and return the copy to the client.

   Finally, replace the direct calls to the subclasses’ constructors with calls to the factory method of the prototype
   registry.

# Pros and Cons

| Pros                                                                                               | Cons                                                                        |
|----------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| You can clone objects without coupling to their concrete classes.                                  | Cloning complex objects that have circular references might be very tricky. |
| You can get rid of repeated initialization code in favor of cloning pre-built prototypes.          |                                                                             |
| You can produce complex objects more conveniently.                                                 |                                                                             |
| You get an alternative to inheritance when dealing with configuration presets for complex objects. |                                                                             |