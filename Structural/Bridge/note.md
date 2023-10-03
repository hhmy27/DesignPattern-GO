# Bridge

![](https://refactoring.guru/images/patterns/diagrams/bridge/structure-en-2x.png)
The Bridge pattern solves the problem of having too many concrete classes by replacing inheritance with composition.

use inheritance:
![](https://refactoring.guru/images/patterns/diagrams/bridge/problem-en-2x.png)

For example, let's consider a geometric class called `Shape`, which has two subclasses: `Circle` and `Square`. Now,
suppose you want to add a color property to the `Shape` subclasses. If you implement this using inheritance, the number
of subclasses would be the product of the number of `Shape` classes and the number of color options (
e.g., `BlueCircle`, `RedCircle`, `BlueSquare`, etc.).

To address this issue, you can use the Bridge pattern to convert the inheritance relationship into composition.
Specifically, you have two dimensions of properties: `Shape` and `Color`. You can extract `Color` as an interface, and
then add a reference to the `Color` interface within the `Shape` class. This approach avoids the problem of rapidly
increasing the number of subclasses through direct extension.

By adding a reference to the `Color` interface in the `Shape` class, you can quickly add the Color property to Shape
subclasses. When you need to add a new color option, you only need to introduce a concrete implementation of the `Color`
interface.
use composition:
![](https://refactoring.guru/images/patterns/diagrams/bridge/solution-en-2x.png)

# Pseudocode

![](https://refactoring.guru/images/patterns/diagrams/bridge/example-en-2x.png)

```
// The "abstraction" defines the interface for the "control"
// part of the two class hierarchies. It maintains a reference
// to an object of the "implementation" hierarchy and delegates
// all of the real work to this object.
class RemoteControl is
    protected field device: Device
    constructor RemoteControl(device: Device) is
        this.device = device
    method togglePower() is
        if (device.isEnabled()) then
            device.disable()
        else
            device.enable()
    method volumeDown() is
        device.setVolume(device.getVolume() - 10)
    method volumeUp() is
        device.setVolume(device.getVolume() + 10)
    method channelDown() is
        device.setChannel(device.getChannel() - 1)
    method channelUp() is
        device.setChannel(device.getChannel() + 1)


// You can extend classes from the abstraction hierarchy
// independently from device classes.
class AdvancedRemoteControl extends RemoteControl is
    method mute() is
        device.setVolume(0)


// The "implementation" interface declares methods common to all
// concrete implementation classes. It doesn't have to match the
// abstraction's interface. In fact, the two interfaces can be
// entirely different. Typically the implementation interface
// provides only primitive operations, while the abstraction
// defines higher-level operations based on those primitives.
interface Device is
    method isEnabled()
    method enable()
    method disable()
    method getVolume()
    method setVolume(percent)
    method getChannel()
    method setChannel(channel)


// All devices follow the same interface.
class Tv implements Device is
    // ...

class Radio implements Device is
    // ...


// Somewhere in client code.
tv = new Tv()
remote = new RemoteControl(tv)
remote.togglePower()

radio = new Radio()
remote = new AdvancedRemoteControl(radio)
```

# When to use it

- Use the Bridge pattern when you want to divide and organize a monolithic class that has several variants of some
  functionality (for example, if the class can work with various database servers).

  The bigger a class becomes, the harder it is to figure out how it works, and the longer it takes to make a change. The
  changes made to one of the variations of functionality may require making changes across the whole class, which often
  results in making errors or not addressing some critical side effects.

  The Bridge pattern lets you split the monolithic class into several class hierarchies. After this, you can change the
  classes in each hierarchy independently of the classes in the others. This approach simplifies code maintenance and
  minimizes the risk of breaking existing code.

- Use the pattern when you need to extend a class in several orthogonal (independent) dimensions.

  The Bridge suggests that you extract a separate class hierarchy for each of the dimensions. The original class
  delegates the related work to the objects belonging to those hierarchies instead of doing everything on its own.

- Use the Bridge if you need to be able to switch implementations at runtime.

  Although it’s optional, the Bridge pattern lets you replace the implementation object inside the abstraction. It’s as
  easy as assigning a new value to a field.

# How to use it

1. Identify the orthogonal dimensions in your classes. These independent concepts could be: abstraction/platform,
   domain/infrastructure, front-end/back-end, or interface/implementation.
2. See what operations the client needs and define them in the base abstraction class.
3. Determine the operations available on all platforms. Declare the ones that the abstraction needs in the general
   implementation interface.
4. For all platforms in your domain create concrete implementation classes, but make sure they all follow the
   implementation interface.
5. Inside the abstraction class, add a reference field for the implementation type. The abstraction delegates most of
   the work to the implementation object that’s referenced in that field.
6. If you have several variants of high-level logic, create refined abstractions for each variant by extending the base
   abstraction class.
7. The client code should pass an implementation object to the abstraction’s constructor to associate one with the
   other. After that, the client can forget about the implementation and work only with the abstraction object.

# Pros and Cons

| Pros                                                                                                                                    | Cons                                                                                            |
|-----------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| 1. You can create platform-independent classes and apps.                                                                                | 1. You might make the code more complicated by applying the pattern to a highly cohesive class. |
| 2. The client code works with high-level abstractions. It isn’t exposed to the platform details.                                        |                                                                                                 |
| 3. Open/Closed Principle. You can introduce new abstractions and implementations independently from each other.                         |                                                                                                 |
| 4. Single Responsibility Principle. You can focus on high-level logic in the abstraction and on platform details in the implementation. |                                                                                                 |