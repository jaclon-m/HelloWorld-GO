# 函数

Go 里面有三种类型的函数
- 普通的带有名字的函数
- 匿名函数或者 lambda 函数
- 方法
  
- 函数参数、返回值以及它们的类型被统称为`函数签名`。
- main函数 入参获取：os args，flag
- init函数 只执行一次：kubenetes glog 引用问题
- go语言不允许函数重载。可以通过动态参数来实现
- go语言不支持泛型
- 如果需要申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体
```cgo
func flushICache(begin, end uintptr) // implemented externally
```
函数也可以以申明的方式被使用，作为一个函数类型
```cgo
type binOp func(int, int) int
```

## 值传递和引用传递

- Go 默认使用按值传递来传递参数，也就是传递参数的副本。函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量
比如 `Function(arg1)`
- 如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加 & 符号，比如 &variable）传递给函数，这就是按引用传递
比如 `Function(&arg1)`
- 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）
- 按引用传递可以减少额外的开销，尤其在使用结构体变量的时候

## 命名返回值

如果使用非命名返回值，则需要显示指定return的值。如果使用命令返回值，可以直接return
```cgo
func getX2AndX3(input int) (int, int) {
    return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
    x2 = 2 * input
    x3 = 3 * input
    // return x2, x3
    return
}
```

## 变长参数的传递

替代方法
 - 结构体
 - 空接口
```cgo
func typecheck(..,..,values … interface{}) {
    for _, value := range values {
        switch v := value.(type) {
            case int: …
            case float: …
            case string: …
            case bool: …
            default: …
        }
    }
}
```

## 闭包（匿名函数）


匿名函数不能单独存在。可以赋值给某个变量，然后通过变量名对函数进行调用；或者通过后面加括号的形式直接调用
```cgo
# 赋值后调用
fplus := func(x, y int) int { return x + y }
fplus(3,4)
# 直接调用
func(x, y int) int { return x + y } (3, 4)
```

闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
```cgo
package main

import "fmt"

func main() {
    var f = Adder()
    fmt.Print(f(1), " - ")
    fmt.Print(f(20), " - ")
    fmt.Print(f(300))
}

func Adder() func(int) int {
    var x int
    return func(delta int) int {
        x += delta
        return x
    }
}
# 输出 1 - 21 - 321
```

### 工厂函数

一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用： 书写一个工厂函数而不是针对每种情况都书写一个函数
```cgo
func MakeAddSuffix(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}

addBmp := MakeAddSuffix(".bmp")
addJpeg := MakeAddSuffix(".jpeg")

addBmp("file") // returns: file.bmp
addJpeg("file") // returns: file.jpeg
```

可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为`高阶函数`，是函数式语言的特点。

### 使用闭包进行调试

可以使用 runtime 或 log 包中的特殊函数来实现这样的功能。包 runtime 中的函数 Caller() 提供了相应的信息，因此可以在需要的时候实现一个 where() 闭包函数来打印函数执行的位置：
```cgo
where := func() {
    _, file, line, _ := runtime.Caller(1)
    log.Printf("%s:%d", file, line)
}
where()
// some code
where()
// some more code
where()

```

```cgo
log.SetFlags(log.Llongfile)
log.Print("")
# 简化版
var where = log.Print
func func1() {
where()
... some code
where()
... some code
where()
}
```

### 计算函数执行时长
```cgo
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

# 结构体


使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)
声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型 T。在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）

- 选择器
```cgo
# 无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 选择器符（selector-notation） 来引用结构体的字段
type myStruct struct { i int }
var v myStruct    // v是结构体类型变量
var p *myStruct   // p是指向一个结构体类型变量的指针
v.i
p.i
```

## 结构体工厂

Go 语言不支持面向对象编程语言中那样的构造子方法，但是可以很容易的在 Go 中实现 “构造子工厂” 方法。为了方便通常会为类型定义一个工厂，按惯例，工厂的名字以 new 或 New 开头。
```cgo
type File struct {
    fd      int     // 文件描述符
    name    string  // 文件名
}

func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }

    return &File{fd, name}
}

```

new(File) 和 &File{} 等价于其它语言的 File f = new File(...)

强制使用工厂方法：
```cgo
# 结构体定义为小写
type matrix struct {
    ...
}

func NewMatrix(params) *matrix {
    m := new(matrix) // 初始化 m
    return m
}

```

## map 和 struct vs new () 和 make ()

```cgo
package main

type Foo map[string]string
type Bar struct {
    thingOne string
    thingTwo int
}

func main() {
    // OK
    y := new(Bar)
    (*y).thingOne = "hello"
    (*y).thingTwo = 1

    // NOT OK
    z := make(Bar) // 编译错误：cannot make type Bar
    (*z).thingOne = "hello"
    (*z).thingTwo = 1

    // OK
    x := make(Foo)
    x["x"] = "goodbye"
    x["y"] = "world"

    // NOT OK
    // new() 一个映射并试图使用数据填充它，将会引发运行时错误！ 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存
    u := new(Foo)
    (*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
    (*u)["y"] = "world"
}

```

## 结构体标签

```cgo
package main

import (
    "fmt"
    "reflect"
)

type TagType struct { // tags
    field1 bool   "An important answer"
    field2 string "The name of the thing"
    field3 int    "How much there are"
}

func main() {
    tt := TagType{true, "Barak Obama", 1}
    for i := 0; i < 3; i++ {
        refTag(tt, i)
    }
}

func refTag(tt TagType, ix int) {
    ttType := reflect.TypeOf(tt)
    ixField := ttType.Field(ix)
    fmt.Printf("%v\n", ixField.Tag)
}

```

## 匿名字段和内嵌结构体

可以理解问继承的一种实现方式。外层可以直接使用内层的匿名字段。
外层名字会覆盖内层名字。同层之间会单只运行时错误

# 方法

作用在接收者的函数
接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。但是接收者不能是一个接口类型
接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针

程序设计：让某些接口实现某些行为
使用场景：很多场景下，函数需要的上下文可以保存在receiver属性中，通过定义 receiver 的方法，该方法可以直接 访问 receiver 属性，减少参数传递需求

**一个类型加上它的方法等价于面向对象中的一个类。**

注意事项
- 别名类型不能有它原始类型上已经定义过的方法
- 如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。
如果 recv 一个指针，Go 会自动解引用。
- 在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的

## 指针和值作为接受者

如果想在方法内改变外部的值，需要使用指针作为接受者。
指针方法和值方法都可以在指针或非指针上被调用
```cgo
package main

import (
    "fmt"
)

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

func main() {
    // 值
    var lst List
    lst.Append(1)
    fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)

    // 指针
    plst := new(List)
    plst.Append(2)
    fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
}
```

## go中的getter和setter

```cgo
//于 setter 方法使用 Set 前缀，对于 getter 方法只使用成员名
package person

type Person struct {
    firstName string
    lastName  string
}

func (p *Person) FirstName() string {
    return p.firstName
}

func (p *Person) SetFirstName(newName string) {
    p.firstName = newName
}

```

---

当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果

# 接口

接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

>（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。
还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。

注意：
- 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
- 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
- 一个类型可以实现多个接口。
- 接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。
- 类型必须实现该接口的所有方法
- 接口可以嵌套接口


## 接口类型判断

### 类型断言：如何检测和转换接口变量的类型

一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型。用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值
```cgo
v := varI.(T)       // unchecked type assertion
```
```cgo
if v, ok := varI.(T); ok {  // checked type assertion
    Process(v)
    return
}
// varI is not of type T
```
如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false，也没有运行时错误发生

### 类型判断：type-switch

```cgo
switch t := areaIntf.(type) {
case *Square:
    fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
    fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
    fmt.Printf("nil value: nothing to check?\n")
default:
    fmt.Printf("Unexpected type %T\n", t)
}
```

```cgo
func classifier(items ...interface{}) {
    for i, x := range items {
        switch x.(type) {
        case bool:
            fmt.Printf("Param #%d is a bool\n", i)
        case float64:
            fmt.Printf("Param #%d is a float64\n", i)
        case int, int64:
            fmt.Printf("Param #%d is a int\n", i)
        case nil:
            fmt.Printf("Param #%d is a nil\n", i)
        case string:
            fmt.Printf("Param #%d is a string\n", i)
        default:
            fmt.Printf("Param #%d is unknown\n", i)
        }
    }
}
```

###  测试一个值是否实现了某个接口

```cgo
type Stringer interface {
    String() string
}

if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
```

## 使用方法集与接口

```cgo
package main

import (
    "fmt"
)

type List []int

func (l List) Len() int {
    return len(l)
}

func (l *List) Append(val int) {
    *l = append(*l, val)
}

type Appender interface {
    Append(int)
}

func CountInto(a Appender, start, end int) {
    for i := start; i <= end; i++ {
        a.Append(i)
    }
}

type Lener interface {
    Len() int
}

func LongEnough(l Lener) bool {
    return l.Len()*10 > 42
}

func main() {
    // A bare value
    var lst List
    // compiler error:
    // cannot use lst (type List) as type Appender in argument to CountInto:
    //       List does not implement Appender (Append method has pointer receiver)
    // CountInto(lst, 1, 10)
    if LongEnough(lst) { // VALID:Identical receiver type
        fmt.Printf("- lst is long enough\n")
    }

    // A pointer value
    plst := new(List)
    CountInto(plst, 1, 10) //VALID:Identical receiver type
    if LongEnough(plst) {
        // VALID: a *List can be dereferenced for the receiver
        fmt.Printf("- plst is long enough\n")
    }
}

```

在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
- 指针方法可以通过指针调用
- 值方法可以通过值调用
- 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

## 空接口

空接口或者最小接口 不包含任何方法，它对实现不做任何要求
```cgo
type Any interface {}
```
任何其他类型都实现了空接口（它不仅仅像 Java/C# 中 Object 引用类型），any 或 Any 是空接口一个很好的别名或缩写。
空接口类似 Java/C# 中所有类的基类： Object 类，二者的目标也很相近。

TODO：具体使用参考the-way-to-go 对应章节

# 反射

```cgo
// blog: Laws of Reflection
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("value:", v)
    fmt.Println("type:", v.Type())
    fmt.Println("kind:", v.Kind())
    fmt.Println("value:", v.Float())
    fmt.Println(v.Interface())
    fmt.Printf("value is %5.2e\n", v.Interface())
    y := v.Interface().(float64)
    fmt.Println(y)
}
```
输出
```cgo
type: float64
value: 3.4
type: float64
kind: float64
value: 3.4
3.4
value is 3.40e+00
3.4
```
Kind 总是返回底层类型
```cgo
type MyInt int
var m MyInt = 5
v := reflect.ValueOf(m)
//方法 v.Kind() 返回 reflect.Int
```

## 通过反射修改 (设置) 值

```cgo
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    // setting a value:
    // v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
    fmt.Println("settability of v:", v.CanSet())
    v = reflect.ValueOf(&x) // Note: take the address of x.
    fmt.Println("type of v:", v.Type())
    fmt.Println("settability of v:", v.CanSet())
    v = v.Elem()
    fmt.Println("The Elem of v is: ", v)
    fmt.Println("settability of v:", v.CanSet())
    v.SetFloat(3.1415) // this works!
    fmt.Println(v.Interface())
    fmt.Println(v)
}
```
输出
```cgo
settability of v: false
type of v: *float64
settability of v: false
The Elem of v is:  <float64 Value>
settability of v: true
3.1415
<float64 Value>
```
当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)
通过 Type () 我们看到 v 现在的类型是 *float64 并且仍然是不可设置的。 要想让其可设置我们需要使用 Elem() 函数，这间接的使用指针：v = v.Elem()

## 反射结构体

有些时候需要反射一个结构体类型。NumField() 方法返回结构体内的字段数量；通过一个 for 循环用索引取得每个字段的值 Field(i)
我们同样能够调用签名在结构体上的方法，例如，使用索引 n 来调用：Method(n).Call(nil)

```cgo
package main

import (
    "fmt"
    "reflect"
)

type NotknownType struct {
    s1, s2, s3 string
}

func (n NotknownType) String() string {
    return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
    value := reflect.ValueOf(secret) // <main.NotknownType Value>
    typ := reflect.TypeOf(secret)    // main.NotknownType
    // alternative:
    //typ := value.Type()  // main.NotknownType
    fmt.Println(typ)
    knd := value.Kind() // struct
    fmt.Println(knd)

    // iterate through the fields of the struct:
    for i := 0; i < value.NumField(); i++ {
        fmt.Printf("Field %d: %v\n", i, value.Field(i))
        // error: panic: reflect.Value.SetString using value obtained using unexported field
        //value.Field(i).SetString("C#")
    }

    // call the first method, which is String():
    results := value.Method(0).Call(nil)
    fmt.Println(results) // [Ada - Go - Oberon]
}
```



