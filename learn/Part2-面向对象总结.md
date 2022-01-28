
# 如何在类型中嵌入功能

- 聚合（或组合）：包含一个所需功能类型的具名字段。
- 内嵌：内嵌（匿名地）所需功能类型。
  假设有一个 Customer 类型，我们想让它通过 Log 类型来包含日志功能，Log 类型只是简单地包含一个累积的消息（当然它可以是复杂的）。
- 如果想让特定类型都具备日志功能，你可以实现一个这样的 Log 类型，然后将它作为特定类型的一个字段，并提供 Log()，它返回这个日志的引用。

```cgo
//聚合
package main

import (
    "fmt"
)

type Log struct {
    msg string
}

type Customer struct {
    Name string
    log  *Log
}

func main() {
    c := new(Customer)
    c.Name = "Barak Obama"
    c.log = new(Log)
    c.log.msg = "1 - Yes we can!"
    // shorter
    c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
    // fmt.Println(c) &{Barak Obama 1 - Yes we can!}
    c.Log().Add("2 - After me the world will be a better place!")
    //fmt.Println(c.log)
    fmt.Println(c.Log())

}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}

func (l *Log) String() string {
    return l.msg
}

func (c *Customer) Log() *Log {
    return c.log
}
```

```cgo
//内嵌
package main

import (
    "fmt"
)

type Log struct {
    msg string
}

type Customer struct {
    Name string
    Log
}

func main() {
    c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
    c.Add("2 - After me the world will be a better place!")
    fmt.Println(c)

}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}

func (l *Log) String() string {
    return l.msg
}

func (c *Customer) String() string {
    return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}
```

# 多重继承

```cgo
package main

import (
    "fmt"
)

type Camera struct{}

func (c *Camera) TakeAPicture() string {
    return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
    return "Ring Ring"
}

type CameraPhone struct {
    Camera
    Phone
}

func main() {
    cp := new(CameraPhone)
    fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
    fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
    fmt.Println("It works like a Phone too: ", cp.Call())
}
```

# 和其他面向对象语言比较 Go 的类型和方法

在如 C++、Java、C# 和 Ruby 这样的面向对象语言中，方法在类的上下文中被定义和继承：
在一个对象上调用方法时，运行时会检测类以及它的超类中是否有此方法的定义，如果没有会导致异常发生。
在 Go 语言中，这样的继承层次是完全没必要的：如果方法在此类型定义了，就可以调用它，和其他类型上是否存在这个方法没有关系。
Go 不需要一个显式的类定义，如同 Java、C++、C# 等那样，
相反地，“类” 是通过提供一组作用于一个共同类型的方法集来隐式定义的。类型可以是结构体或者任何用户自定义类型。

## 总结
在 Go 中，类型就是类（数据和关联的方法）。Go 拥有类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。
在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 **组件编程（Component Programming）**

# 接口与动态类型

- 
和其它语言相比，Go 是唯一结合了接口值，静态类型检查（是否该类型实现了某个接口），运行时动态转换的语言，并且不需要显式地声明类型是否满足某个接口。该特性允许我们在不改变已有的代码的情况下定义和使用新接口。
接收一个（或多个）接口类型作为参数的函数，其实参可以是任何实现了该接口的类型。 实现了某个接口的类型可以被传给任何以此接口为参数的函数 。
- 动态方法调用
  像 Python，Ruby 这类语言，动态类型是延迟绑定的（在运行时进行）：方法只是用参数和变量简单地调用，然后在运行时才解析
  Go 的实现与此相反，通常需要编译器静态检查的支持：当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数。如果方法调用作用于像 interface{} 这样的 “泛型” 上，你可以通过类型断言来检查变量是否实现了相应接口。
- 空接口和函数重载
  函数重载是不被允许的。在 Go 语言中函数重载可以用可变参数 ...T 作为函数最后一个参数来实现。
如果我们把 T 换为空接口，那么可以知道任何类型的变量都是满足 T (空接口）类型的，这样就允许我们传递任何数量任何类型的参数给函数，即重载的实际含义。
- 接口的继承
  当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。

# 总结

OO 语言最重要的三个方面分别是：封装，继承和多态，在 Go 中它们是怎样表现的呢？

- 封装（数据隐藏）：和别的 OO 语言有 4 个或更多的访问层次相比，Go 把它简化为了 2 层
1）包范围内的：通过标识符首字母小写，对象 只在它所在的包内可见
2）可导出的：通过标识符首字母大写，对象 对所在包以外也可见
类型只拥有自己所在包中定义的方法。
- 继承：用组合实现：内嵌一个（或多个）包含想要的行为（字段和方法）的类型；多重继承可以通过内嵌多个类型实现
- 多态：用接口实现：某个类型的实例可以赋给它所实现的任意接口类型的变量。类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。