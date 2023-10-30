# golang 设计模式

## 1 单例模式

### 1.1 什么是单例模式

单例模式是一类最经典最简单的设计模式。在单例模式下，**我们声明一个类并保证这个类只存在全局唯一的实例供外部反复使用.**

单例模式的适用场景包括：

- 一些只允许存在一个实例的类，比如全局统一的监控统计模块
- 一些实例化时很耗费资源的类，比如协程池、连接池、和第三方交互的客户端等
- 一些入参繁杂的系统模块组件，比如 controller、service、dao 等

在单例模式的实现上，可以分为饿汉式和懒汉式两种类型：

- 饿汉式：从一开始就完成单例的初始化工作，以备不时之需（肚子饿了，先干为敬.）
- 懒汉式：贯彻佛系思想，不到逼不得已（需要被使用了），不执行单例的初始化工作

<img src="D:\Github\study-coding\golang\img\640.png" alt="图片" style="zoom:50%;" />

### 1.2 饿汉式 

#### 1.2.1 实现流程

饿汉式和懒汉式中的“饿”和“懒”体现在单例初始化时机的不同. “饿” 指的是，对于单例对象而言，不论其后续有没有被使用到以及何时才会被使用到，都会在程序启动之初完成其初始化工作.

在实现上，可以将饿汉式单例模式的执行步骤拆解如下：

- 单例类和构造方法声明为不可导出类型，避免被外部直接获取到（避免让外界拥有直接初始化的能力，导致单例模式被破坏）
- 在代码启动之初，就初始化好一个全局单一的实例，作为后续所谓的“单例”
- 暴露一个可导出的单例获取方法 GetXXX()，用于返回这个单例对象

----------------------

类似封装的对象，只导出调用的方法

----------------------------------------



#### 1.2.2 实现代码

<img src="D:\Github\study-coding\golang\img\640-16909897688213.png" alt="图片" style="zoom:50%;" />

```go
package singleton

var s *singleton 

func init() {
	s = newSingleton() 
}

type singleton struct {}

func newSingleton() *singleton {
	return &singleton{}
}

func (s *singleton) Work() {

}

func GetInstance() *singleton {
	return s 
}
```

上述代码在实现上没有逻辑问题，但是存在一个比较容易引起争议的规范性问题，就是在对外可导出的 GetInstance 方法中，返回了不可导出的类型 singleton.

代码执行流程上 ok，但这种实现方式存在代码坏味道，相应的问题在 stackoverflow 上引起过讨论，对应链接如下，大家感兴趣可以去了解原贴中的讨论内容：

https://stackoverflow.com/questions/21470398/return-an-unexported-type-from-a-function

 

不建议这么做的原因主要在于：

- • singleton 是包内的不可导出类型，在包外即便获取到了，也无法直接作为方法的入参或者出参进行传递，显得很呆
- • singleton 的对外暴露，使得 singleton 所在 package 的代码设计看起来是自相矛盾的，混淆了 singleton 这个不可导出类型的边界和定位

综上，规范的处理方式是，在不可导出单例类 singleton 的基础上包括一层接口 interface，将其作为对对导出方法 GetInstance 的返回参数类型:



### 1.3 懒汉式

懒汉式讲究的是”佛系”，某件事情如果是可做可不做，那我一定选择不做. 直到万不得已非做不可的时候，我才会采取行动（deadline 是第一生产力）.

懒汉式的执行步骤如下：

- 单例类声明为不可导出类型，避免被外界直接获取到
- 声明一个全局单例变量, 但不进行初始化（注意**只声明，不初始化**）
- 暴露一个对外公开的方法,用于获取这个单例
- 在这个获取方法被调用时，会判断单例是否初始化过，倘若没有，则在此时才完成初始化工作

**（1）懒汉模式 1.0** —— **存在并发问题**

首先请大家观察我这边提出的懒汉 1.0 的实现源码，并留意其中存在的问题.

<img src="D:\Github\study-coding\golang\img\640-16910463165076.png" alt="图片" style="zoom:50%;" />

 

```
package singleton

var s *singleton

type singleton struct {
}

func newSingleton() *singleton {
    return &singleton{}
}

func (s *singleton) Work() {
}

type Instance interface {
    Work()
}

func GetInstance() Instance {
    if s == nil {
        s = newSingleton()
    }
    return s
}
```

这个实现流程乍一看没有问题，但是我们需要意识到，这个 GetInstance 方法是对外暴露的，我们需要基于底线思维，把外界看成是存在不稳定因素的使用方，这个 GetInstance 方法是存在被并发调用的可能性的，一旦被并发调用，则 singleton 这个单例就可能被初始化两次，违背了所谓”单例“的语义.

**（2）懒汉模式 2.0**

问题已经发现了，我们就见招拆招，在1.0的基础上提出2.0，探讨如何规避并发问题，实现真正意义的单例模式.

<img src="D:\Github\study-coding\golang\img\640-16910468991339.png" alt="图片" style="zoom:50%;" />

 

```
package singleton

import "sync"

var (
    s   *singleton
    mux sync.Mutex
)

type singleton struct {
}

func newSingleton() *singleton {
    return &singleton{}
}

func (s *singleton) Work() {
}

type Instance interface {
    Work()
}

func GetInstance() Instance {
    mux.Lock()
    defer mux.Unlock()
    if s == nil {
        s = newSingleton()
    }
    return s
}
```

 

上述懒汉2.0的代码中，我们通过定义一把全局锁，用于在并发场景下保护单例 singleton 的数据一致性.

用户调用 GetInstance 方法时，无一例外需要率先取得锁，然后再判断 singleton 是否被初始化过，如果没有，则完成对应的初始化工作. 通过互斥锁的保护，保证了 singleton 的初始化工作一定只会执行一次，”单例“的语义得以达成.

**这样可以解决并发问题，但是还不够完美，这是因为即便 singleton 已经初始化过了，后续外界用户每次在获取单例时，都需要加锁，存在无意义的性能损耗.**



**（3）懒汉模式 3.0**

解决懒汉2.0中性能问题的关键在于，我们希望尽可能地减少与互斥锁的交互，在此基础上，我们提出懒汉3.0的实现.

 

<img src="D:\Github\study-coding\golang\img\640-169104730147412.png" alt="图片" style="zoom:50%;" />

 

```
package singleton

import "sync"

var (
    s   *singleton
    mux sync.Mutex
)

type singleton struct {
}

func newSingleton() *singleton {
    return &singleton{}
}

func (s *singleton) Work() {
}

type Instance interface {
    Work()
}

func GetInstance() Instance {
    if s != nil {
        return s
    }
    mux.Lock()
    defer mux.Unlock()

    s = newSingleton()

    return s
}
```

懒汉3.0中，用户调用 GetInstance 方法获取单例时经历如下步骤：

- • 首先在加锁前，先判断 singleton 是否初始化过，如果是，则直接返回（需要注意，这一步是无锁的）
- • 倘若 singleton 没初始化过，才加锁，并执行初始化工作

这样的实现方式，保证了只要 singleton 被成功初始化后，用户调用 GetInstance 方法时都可以直接返回，无需加锁，大幅度减少了加锁的频率.

然而，懒汉3.0的实现是存在逻辑上的漏洞，仍然可能引起并发安全问题. 这里给出的反例如下：

- • moment1：单例 singleton 至今为止没有被初始化过
- • moment2：goroutine A 和 goroutine B 同时调用 GetInstance 获取单例，由于当前 singleton 没初始化过，于是两个 goroutine 都未走进 if s != nil 的分支，而是开始抢锁
- • moment3：goroutine A 抢锁成功继续往下；goroutine B 抢锁失败，进行等锁
- • moment4：goroutine A 完成 singleton 初始化，并释放锁
- • moment5：由于锁被释放，goroutine B 取锁成功，并继续往下执行，完成 singleton 的初始化

通过上述5个时间节点的串联，我们得见，singleton 仍然被初始化了不只1次.



**（4）懒汉模式 4.0**

最后，我们在 3.0 的基础上继续升级，给到完整的解决方案：懒汉 4.0

<img src="D:\Github\study-coding\golang\img\640-169104736624015.png" alt="图片" style="zoom:50%;" />

 

```
package singleton

import "sync"

var (
    s   *singleton
    mux sync.Mutex
)

type singleton struct {
}

func newSingleton() *singleton {
    return &singleton{}
}

func (s *singleton) Work() {
}

type Instance interface {
    Work()
}

func GetInstance() Instance {
    if s != nil {
        return s
    }
    mux.Lock()
    defer mux.Unlock()
    if s != nil {
        return s
    }
    s = newSingleton()

    return s
}
```

 

懒汉4.0中，我们将流程升级为加锁 double check 模式：

- • 在加锁前，先检查一轮单例的初始化状态，倘若已初始化过，则直接返回，以做到最大限度的无锁操作
- • 倘若通过第一轮检查，则进行加锁，保证并发安全性
- • 加锁成功后，需要执行第二轮检查，确保在此时单例仍未初始化过的前提下，才执行初始化工作

此处得以解决懒汉3.0中并发问题的核心在于，加锁之后多了一次 double check 动作，由于这轮检查工作是在加锁之后执行的，因此能够保证 singleton 的初始化状态是稳定不变的，并发问题彻底得以解决.



**（5）懒汉模式 5.0**

事实上，在使用 Go 语言时还有一种更优雅的单例实现方式，那就是使用 sync 包下的单例工具 sync.Once，使用的代码示例如下. 关于 sync.Once 底层具体的实现原理，我们放在本文第 4 章中再作展开.

```
package singleton

import "sync"

var (
    s    *singleton
    once sync.Once
)

type singleton struct {
}

func newSingleton() *singleton {
    return &singleton{}
}

func (s *singleton) Work() {
}

type Instance interface {
    Work()
}

func GetInstance() Instance {
    once.Do(func() {
        s = newSingleton()
    })
    return s
}
```

 

### 1.4 两种模式对比

饿汉式与懒汉式没有绝对的优劣之分，需要权衡看待：

- • 饿汉式在程序运行之初就完成单例的初始化，说白了，不够智能，不够极限，不够”懒“. 说不定这个单例对象迟迟不被使用到，甚至永远都不被使用，那么这次初始化动作可能只是一次无谓的性能损耗
- • 懒汉式在单例被首次使用时才执行初始化，看起来显得”聪明“一些. 但是，我们需要意识到，倘若初始化工作中存在异常问题（如 panic，fatal），则这枚定时炸弹会在程序运行过程才暴露出来，这对于我们的运行项目而言会带来更严重的伤害. 相比之下，倘若使用的是饿汉模式，则这种实例化的问题会在代码编译运行之初就提前暴露，更有利于问题的定位和解决



### 1.5 sync.Once 实现原理

sync.Once 是 Golang 提供的用于支持实现单例模式的标准库工具，其对应的数据结构如下：

<img src="D:\Github\study-coding\golang\img\640.png" alt="图片" style="zoom:67%;" />

```
package sync

import (
    "sync/atomic"
)

type Once struct {
    // 通过一个整型变量标识，once 保护的函数是否已经被执行过
    done uint32
    // 一把锁，在并发场景下保护临界资源 done 字段只能串行访问
    m    Mutex
}
```

 

在 sync.Once 的定义类中 包含了两个核心字段：

- • done：一个整型 uint32，用于标识用户传入的任务函数是否已经执行过了
- • m：一把互斥锁 sync.Mutex，用于保护标识值 done ，避免因并发问题导致数据不一致

 

sync.Once 本质上也是通过加锁 double check 机制，实现了任务的全局单次执行，实现的方法链路和具体源码展示如下：

<img src="D:\Github\study-coding\golang\img\640-16910505307381.png" alt="图片" style="zoom: 67%;" />

```
func (o *Once) Do(f func()) {
    // 锁外的第一次 check，读取 Once.done 的值
    if atomic.LoadUint32(&o.done) == 0 {
        o.doSlow(f)
    }
}

func (o *Once) doSlow(f func()) {
    // 加锁
    o.m.Lock()
    defer o.m.Unlock()
    // double check
    if o.done == 0 {
        // 任务执行完成后，将 Once.done 标识为 1
        defer atomic.StoreUint32(&o.done, 1)
        // 保证全局唯一一次执行用户注入的任务
        f()
    }
}
```

单例工具 sync.Once 的使用方式非常简单. 用户调用 sync.Once.Do 方法，并在方法入参传入一个需要保证全局只执行一次的闭包任务函数 f func() 即可.

sync.Once.Do 方法的实现步骤如下：

- • first check：第一次检查 Once.done 的值是否为 0，这步是无锁化的
- • easy return：倘若 Once.done 的值为 0，说明任务已经执行过，直接返回
- • lock：加锁
- • double check：再次检查 Once.done 的值是否为 0
- • execute func：倘若通过 double check，真正执行用户传入的执行函数 f
- • update：执行完 f 后，将 Once.done 的值设为 1
- • return：解锁并返回







## 2 观察者模式

### 2.1 原理介绍

<img src="D:\Github\study-coding\golang\img\640-16910692822648.png" alt="图片" style="zoom:50%;" />

本期基于 go 语言和大家探讨设计模式中的观察者模式. 观察者模式适用于多对一的订阅/发布场景.

- • ”多“：指的是有多名观察者
- • ”一“：指的是有一个被观察事物
- • ”订阅“：指的是观察者时刻关注着事物的动态
- • ”发布“：指的是事物状态发生变化时是透明公开的，能够正常进入到观察者的视线

在上述场景中，我们了解到核心对象有两类，一类是“观察者”，一类是“被观察的事物”，且两者间在数量上存在多对一的映射关系.

在具体作编程实现时，上述场景的实现思路可以是百花齐放的，而观察者模式只是为我们提供了一种相对规范的设计实现思路，其遵循的核心宗旨是实现“观察者”与“被观察对象”之间的解耦，并将其设计为通用的模块，便于后续的扩展和复用.

学习设计模式时，我们脑海中需要中需要明白，教条是相对刻板的，而场景和问题则是灵活多变的，在工程实践中，我们避免生搬硬套，要做到因地制宜，随机应变.

 

### 2.2 代码实践

#### 2.2.1 核心角色

在观察者模式中，核心的角色包含三类：

- • Observer：观察者. 指的是关注事物动态的角色
- • Event：事物的变更事件. 其中 Topic 标识了事物的身份以及变更的类型，Val 是变更详情
- • EventBus：事件总线. 位于观察者与事物之间承上启下的代理层. 负责维护管理观察者，并且在事物发生变更时，将情况同步给每个观察者.

 

<img src="D:\Github\study-coding\golang\img\640-16910692822659.png" alt="图片" style="zoom:50%;" />

观察者模式的核心就在于建立了 EventBus 的角色. 由于 EventBus 模块的诞生，实现了观察者与具体被观察事物之间的解耦：

- • 针对于观察者而言，需要向 EventBus 完成注册操作，注册时需要声明自己关心的变更事件类型（调用 EventBus 的 Subscribe 方法），不再需要直接和事物打交道
- • 针对于事物而言，在其发生变更时，只需要将变更情况向 EventBus 统一汇报即可（调用 EventBus 的 Publish 方法），不再需要和每个观察者直接交互
- • 对于 EventBus，需要提前维护好每个观察者和被关注事物之间的映射关系，保证在变更事件到达时，能找到所有的观察者逐一进行通知（调用 Observer 的 OnChange 方法）

 

三类角色组织生成的 UML 类图如下所示：

<img src="D:\Github\study-coding\golang\img\640-169106928226510.png" alt="图片" style="zoom:50%;" />

 

对应的代码实现示例展示如下：

```
type Event struct {
    Topic string
    Val   interface{}
}


type Observer interface {
    OnChange(ctx context.Context, e *Event) error
}


type EventBus interface {
    Subscribe(topic string, o Observer)
    Unsubscribe(topic string, o Observer)
    Publish(ctx context.Context, e *Event)
}
```

 

观察者 Observer 需要实现 OnChange 方法，用于向 EventBus 暴露出通知自己的“联系方式”，并且在方法内部实现好当关注对象发生变更时，自己需要采取的处理逻辑.

下面给出一个简单的观察者实现示例 BaseObserver：

```
type BaseObserver struct {
    name string
}


func NewBaseObserver(name string) *BaseObserver {
    return &BaseObserver{
        name: name,
    }
}


func (b *BaseObserver) OnChange(ctx context.Context, e *Event) error {
    fmt.Printf("observer: %s, event key: %s, event val: %v", b.name, e.Topic, e.Val)
    // ...
    return nil
}
```

 

事件总线 EventBus 需要实现 Subscribe 和 Unsubscribe 方法暴露给观察者，用于新增或删除订阅关系，其实现示例如下：

```
type BaseEventBus struct {
    mux       sync.RWMutex
    observers map[string]map[Observer]struct{}
}


func NewBaseEventBus() BaseEventBus {
    return BaseEventBus{
        observers: make(map[string]map[Observer]struct{}),
    }
}


func (b *BaseEventBus) Subscribe(topic string, o Observer) {
    b.mux.Lock()
    defer b.mux.Unlock()
    _, ok := b.observers[topic]
    if !ok {
        b.observers[topic] = make(map[Observer]struct{})
    }
    b.observers[topic][o] = struct{}{}
}


func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
    b.mux.Lock()
    defer b.mux.Unlock()
    delete(b.observers[topic], o)
}
```

针对 EventBus 将事物变更事件同步给每个观察者的 Publish 流程，可以分为同步模式和异步模式，分别在 2.2 小节和 2.3 小节中展开介绍.

 

#### 2.2.2 同步模式

在同步模式的实现中，通过 SyncEventBus 实现了 EventBus 的同步通知版本，对应类图如下：

<img src="D:\Github\study-coding\golang\img\640-169106928226611.png" alt="图片" style="zoom:50%;" />

 

<img src="D:\Github\study-coding\golang\img\640-169106928226612.png" alt="图片" style="zoom:50%;" />

 

在同步模式下，EventBus 在接受到变更事件 Event 时，会根据事件类型 Topic 匹配到对应的观察者列表 observers，然后采用串行遍历的方式分别调用 Observer.OnChange 方法对每个观察者进行通知，并对处理流程中遇到的错误进行聚合，放到 handleErr 方法中进行统一的后处理.

```
type SyncEventBus struct {
    BaseEventBus
}


func NewSyncEventBus() *SyncEventBus {
    return &SyncEventBus{
        BaseEventBus: NewBaseEventBus(),
    }
}


func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
    s.mux.RLock()
    subscribers := s.observers[e.Topic]
    s.mux.RUnlock()


    errs := make(map[Observer]error)
    for subscriber := range subscribers {
        if err := subscriber.OnChange(ctx, e); err != nil {
            errs[subscriber] = err
        }
    }


    s.handleErr(ctx, errs)
}
```

 

此处对 handleErr 方法的实现逻辑进行建立了简化，在真实的实践场景中，可以针对遇到的错误建立更完善的后处理流程，如采取重试或告知之类的操作.

```
func (s *SyncEventBus) handleErr(ctx context.Context, errs map[Observer]error) {
    for o, err := range errs {
        // 处理 publish 失败的 observer
        fmt.Printf("observer: %v, err: %v", o, err)
    }
}
```

 

#### 2.2.3 异步模式

在异步模式的实现中，通过 AsyncEventBus 实现了 EventBus 的异步通知版本，对应类图如下：

<img src="D:\Github\study-coding\golang\img\640-169106928226613.png" alt="图片" style="zoom:50%;" />

 

<img src="D:\Github\study-coding\golang\img\640-169106928226614.png" alt="图片" style="zoom:67%;" />

 

在异步模式下，会在 EventBus 启动之初，异步启动一个守护协程，负责对接收到的错误进行后处理.

在事物发生变更时，EventBus 的 Publish 方法会被调用，此时 EventBus 会并发调用 Observer.OnChange 方法对每个观察者进行通知，在这个过程中遇到的错误会通过 channel 统一汇总到 handleErr 的守护协程中进行处理.

```
type observerWithErr struct {
    o   Observer
    err error
}




type AsyncEventBus struct {
    BaseEventBus
    errC chan *observerWithErr
    ctx  context.Context
    stop context.CancelFunc
}




func NewAsyncEventBus() *AsyncEventBus {
    aBus := AsyncEventBus{
        BaseEventBus: NewBaseEventBus(),
    }
    aBus.ctx, aBus.stop = context.WithCancel(context.Background())
    // 处理处理错误的异步守护协程
    go aBus.handleErr()
    return &aBus
}


func (a *AsyncEventBus) Stop() {
    a.stop()
}


func (a *AsyncEventBus) Publish(ctx context.Context, e *Event) {
    a.mux.RLock()
    subscribers := a.observers[e.Topic]
defer a.mux.RUnlock()
    for subscriber := range subscribers {
        // shadow
        subscriber := subscriber
        go func() {
            if err := subscriber.OnChange(ctx, e); err != nil {
                select {
                case <-a.ctx.Done():
                case a.errC <- &observerWithErr{
                    o:   subscriber,
                    err: err,
                }:
                }
            }
        }()
    }
}


func (a *AsyncEventBus) handleErr() {
    for {
        select {
        case <-a.ctx.Done():
            return
        case resp := <-a.errC:
            // 处理 publish 失败的 observer
            fmt.Printf("observer: %v, err: %v", resp.o, resp.err)
        }
    }
} 
```

 

#### 2.2.4 使用示例

下面分别给出同步和异步模式下观察者模式的使用示例：

```
func Test_syncEventBus(t *testing.T) {
    observerA := NewBaseObserver("a")
    observerB := NewBaseObserver("b")
    observerC := NewBaseObserver("c")
    observerD := NewBaseObserver("d")


    sbus := NewSyncEventBus()
    topic := "order_finish"
    sbus.Subscribe(topic, observerA)
    sbus.Subscribe(topic, observerB)
    sbus.Subscribe(topic, observerC)
    sbus.Subscribe(topic, observerD)


    sbus.Publish(context.Background(), &Event{
        Topic: topic,
        Val:   "order_id: xxx",
    })
}
```

 

异步测试代码：

```
func Test_asyncEventBus(t *testing.T) {
    observerA := NewBaseObserver("a")
    observerB := NewBaseObserver("b")
    observerC := NewBaseObserver("c")
    observerD := NewBaseObserver("d")


    abus := NewAsyncEventBus()
    defer abus.Stop()


    topic := "order_finish"
    abus.Subscribe(topic, observerA)
    abus.Subscribe(topic, observerB)
    abus.Subscribe(topic, observerC)
    abus.Subscribe(topic, observerD)


    abus.Publish(context.Background(), &Event{
        Topic: topic,
        Val:   "order_id: xxx",
    })


    <-time.After(time.Second)
}
```

 

### 2.3 工程案例

本章和大家一起梳理一下在工程实践中对观察者模式的使用场景.

#### 2.3.1 MQ 发布/订阅

<img src="D:\Github\study-coding\golang\img\640-169106928226615.png" alt="图片" style="zoom: 67%;" />

 

大家耳熟能详的消息队列就是对观察者模式的一种实践，大家可以采用类比的方式在 MQ （Message Queue）架构中代入观察者模式中的每一类角色：

- • EventBus：对应的是消息队列组件，为整个通信架构提供了分布式解耦、流量削峰等能力
- • Event：对应的是消息队列中的一条消息，有明确的主题 topic，由生产者 producer 提供
- • Observer：对应的是消费者 consumer，对指定事物的动态（topic）进行订阅，并在消费到对应的变更事件后执行对应的处理逻辑