# Facade

**Facade** is a structural design pattern that provides a simplified interface to a library, a framework, or any other
complex set of classes.
![](https://refactoring.guru/images/patterns/content/facade/facade-2x.png)

**Structure**
![](https://refactoring.guru/images/patterns/diagrams/facade/structure-2x.png)
By defining a facade class, the client can interact directly with the facade class without worrying about the complex
dependencies between libraries.
This approach also decouples the client from dependencies on third-party libraries.
Another benefit of using a facade is that if you need to update the library version or switch to another library, you
only need to modify the method implementations in the facade class.

# Pseudocode

![](https://refactoring.guru/images/patterns/diagrams/facade/example-2x.png)
In this example, the Facade pattern simplifies interaction with a complex video conversion framework.

```go
// These are some of the classes of a complex 3rd-party video
// conversion framework. We don't control that code, therefore
// can't simplify it.

class VideoFile
// ...

class OggCompressionCodec
// ...

class MPEG4CompressionCodec
// ...

class CodecFactory
// ...

class BitrateReader
// ...

class AudioMixer
// ...

// We create a facade class to hide the framework's complexity
// behind a simple interface. It's a trade-off between
// functionality and simplicity.
class VideoConverter is
method convert(filename, format):File is
file = new VideoFile(filename)
sourceCodec = (new CodecFactory).extract(file)
if (format == "mp4")
destinationCodec = new MPEG4CompressionCodec() else
destinationCodec = new OggCompressionCodec()
buffer = BitrateReader.read(filename, sourceCodec)
result = BitrateReader.convert(buffer, destinationCodec)
result = (new AudioMixer()).fix(result)
return new File(result)

// Application classes don't depend on a billion classes
// provided by the complex framework. Also, if you decide to
// switch frameworks, you only need to rewrite the facade class.
class Application is
method main() is
convertor = new VideoConverter()
mp4 = convertor.convert("funny-cats-video.ogg", "mp4")
mp4.save()
```

# When to use it

- Use the Facade pattern when you need to have a limited but straightforward interface to a complex subsystem.
- Use the Facade when you want to structure a subsystem into layers.
  Create facades to define entry points to each level of a subsystem. You can reduce coupling between multiple
  subsystems by requiring them to communicate only through facades.

  For example, let’s return to our video conversion framework. It can be broken down into two layers: video- and
  audio-related. For each layer, you can create a facade and then make the classes of each layer communicate with each
  other via those facades. This approach looks very similar to the Mediator pattern.

# How to use it

1. Check whether it’s possible to provide a simpler interface than what an existing subsystem already provides. You’re
   on the right track if this interface makes the client code independent from many of the subsystem’s classes.

2. Declare and implement this interface in a new facade class. The facade should redirect the calls from the client code
   to appropriate objects of the subsystem. The facade should be responsible for initializing the subsystem and managing
   its further life cycle unless the client code already does this.

3. To get the full benefit from the pattern, make all the client code communicate with the subsystem only via the
   facade. Now the client code is protected from any changes in the subsystem code. For example, when a subsystem gets
   upgraded to a new version, you will only need to modify the code in the facade.

4. If the facade becomes too big, consider extracting part of its behavior to a new, refined facade class.

# Pros and Cons

| Pros                                                          | Cons                                                                                                           |
|---------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|
| You can isolate your code from the complexity of a subsystem. | A facade can become a [god object](https://en.wikipedia.org/wiki/God_object) coupled to all classes of an app. |
