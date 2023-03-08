



# 1、Golang的协程调度器原理及GMP设计思想



## Golang“调度器”的由来？



### (1) 单进程时代不需要调度器

早期的操作系统每个程序就是一个进程，一个程序运行完，才能进行下一个进程，就是“单进程时代”

**一切的程序只能串行发生。**

<img src="../img/1650776039816-a8a5efb6-06be-4984-bfb0-7565c14b0a61.png" alt="5-单进程操作系统.png" style="zoom:33%;" />

早期的单进程操作系统，面临2个问题：

1.单一的执行流程，计算机只能一个任务一个任务处理。

2.进程阻塞所带来的CPU时间浪费。

### (2)多进程/线程时代有了调度器需求

<img src="../img/1650776059361-384f6fb5-e2b1-4f99-8701-f57694aa8ecb.png" alt="6-多进程操作系统.png" style="zoom:33%;" />

在多进程/多线程的操作系统中，就解决了阻塞的问题，因为一个进程阻塞cpu可以立刻切换到其他进程中去执行，而且调度cpu的算法可以保证在运行的进程都可以被分配到cpu的运行时间片。这样从宏观来看，似乎多个进程是在同时被运行。

但新的问题就又出现了，进程拥有太多的资源，**进程的创建、切换、销毁，都会占用很长的时间**，CPU虽然利用起来了，但如果进程过多，CPU有很大的一部分都被用来进行进程调度了。

<img src="../img/1650776077730-2a6860a3-0466-4df9-925d-9ecd5cb9ad7d.png" alt="7-cpu切换浪费成本.png" style="zoom:33%;" />

很明显，CPU调度切换的是进程和线程。尽管线程看起来很美好，但实际上多线程开发设计会变得更加复杂，要考虑很多同步竞争等问题，如锁、竞争冲突等。



### (3) 协程来提高CPU利用率

多进程、多线程已经提高了系统的并发能力，但是在当今互联网高并发场景下，为每个任务都创建一个线程是不现实的，因为会消耗大量的内存(进程虚拟内存会占用4GB[32位操作系统], 而线程也要大约4MB)。

大量的进程/线程出现了新的问题

- 高内存占用
- 调度的高消耗CPU

好了，然后工程师们就发现，其实一个线程分为“内核态“线程和”用户态“线程。

一个“用户态线程”必须要绑定一个“内核态线程”，但是CPU并不知道有“用户态线程”的存在，它只知道它运行的是一个“内核态线程”(Linux的PCB进程控制块)。

<img src="../img/1650776112186-eff4e8b8-8742-44cd-a828-db1653649ee7.png" alt="8-线程的内核和用户态.png" style="zoom:33%;" />

这样，我们再去细化去分类一下，内核线程依然叫“线程(thread)”，用户线程叫“协程(co-routine)".

<img src="../img/1650776128796-5b795bfb-3289-4f6b-85a0-f24399dfc79c.png" alt="9-协程和线程.png" style="zoom:33%;" />

我们就看到了有3种协程和线程的映射关系

#### N:1关系

N个协程绑定1个线程，优点就是**协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速**。但也有很大的缺点，1个进程的所有协程都绑定在1个线程上

缺点：

- 某个程序用不了硬件的多核加速能力
- 一旦某协程阻塞，造成线程阻塞，本进程的其他协程都无法执行了，根本就没有并发的能力了。

<img src="../img/1650776145617-04763b3d-1b15-42c7-9653-cde21bcc98bc.png" alt="10-N-1关系.png" style="zoom:33%;" />

#### 1:1 关系

1个协程绑定1个线程，这种最容易实现。协程的调度都由CPU完成了，不存在N:1缺点，

缺点：

- 协程的创建、删除和切换的代价都由CPU完成，有点略显昂贵了。

<img src="../img/1650776180139-043037ed-cb5b-4c24-9fcf-691a05db17f9.png" alt="11-1-1.png" style="zoom:33%;" />

#### M:N关系

M个协程绑定N个线程，是N:1和1:1类型的结合，克服了以上2种模型的缺点，但实现起来最为复杂。

<img src="../img/1650776193242-4fecd540-5cbb-4f2d-8121-5312dbc6958a.png" alt="12-m-n.png" style="zoom:33%;" />

协程跟线程是有区别的，线程由CPU调度是抢占式的，**协程由用户态调度是协作式的**，一个协程让出CPU后，才执行下一个协程。



### (4) Go语言的协程goroutine

**Go为了提供更容易使用的并发方法，使用了goroutine和channel**。goroutine来自协程的概念，让一组可复用的函数运行在一组线程之上，即使有协程阻塞，该线程的其他协程也可以被`runtime`调度，转移到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。

Go中，协程被称为goroutine，它非常轻量，一个goroutine只占几KB，并且这几KB就足够goroutine运行完，这就能在有限的内存空间内支持大量goroutine，支持了更多的并发。虽然一个goroutine的栈只占几KB，但实际是可伸缩的，如果需要更多内容，`runtime`会自动为goroutine分配。

Goroutine特点：

- 占用内存更小（几kb）
- 调度更灵活(runtime调度)



### (5) 被废弃的goroutine调度器

<img src="../img/1650776259684-6015cb7b-b33e-47f9-b241-185c57dc2745.png" alt="13-gm.png" style="zoom:50%;" />

​	下面我们来看看被废弃的golang调度器是如何实现的？

<img src="../img/1650776272668-ac680807-d927-4c10-9e1d-3960bdabd0e3.png" alt="14-old调度器.png" style="zoom:33%;" />

M想要执行、放回G都必须访问全局G队列，并且M有多个，即多线程访问同一资源需要加锁进行保证互斥/同步，所以全局G队列是有互斥锁进行保护的。

​	老调度器有几个缺点：

1. 创建、销毁、调度G都需要每个M获取锁，这就形成了**激烈的锁竞争**。
2. M转移G会造成**延迟和额外的系统负载**。比如当G中包含创建新协程的时候，M创建了G’，为了继续执行G，需要把G’交给M’执行，也造成了**很差的局部性**，因为G’和G是相关的，最好放在M上执行，而不是其他M'。
3. 系统调用(CPU在M之间的切换)导致频繁的线程阻塞和取消阻塞操作增加了系统开销。





## Goroutine调度器的GMP模型的设计思想

在新调度器中，除了M(thread)和G(goroutine)，又引进了P(Processor)。

<img src="../img/1650776288599-36c23cc6-3d25-4f6f-8f80-83bd43aa6dec.png" alt="15-gmp.png" style="zoom:33%;" />

**Processor，它包含了运行goroutine的资源**，如果线程想运行goroutine，必须先获取P，P中还包含了可运行的G队列。



### (1) GMP模型

在Go中，**线程是运行goroutine的实体，调度器的功能是把可运行的goroutine分配到工作线程上**。

<img src="../img/1650776301442-fb76123c-8d0e-4375-af35-b5728a5b1bc7.jpeg" alt="16-GMP-调度.png" style="zoom: 50%;" />

1. **全局队列**（Global Queue）：存放等待运行的G。
2. **P的本地队列**：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建G'时，G'优先加入到P的本地队列，如果队列满了，则会把本地队列中一半的G移动到全局队列。
3. **P列表**：所有的P都在程序启动时创建，并保存在数组中，最多有`GOMAXPROCS`(可配置)个。
4. **M**：线程想运行任务就得获取P，从P的本地队列获取G，P队列为空时，M也会尝试从全局队列**拿**一批G放到P的本地队列，或从其他P的本地队列**偷**一半放到自己P的本地队列。M运行G，G执行之后，M会从P获取下一个G，不断重复下去。



**Goroutine调度器和OS调度器是通过M结合起来的，每个M都代表了1个内核线程，OS调度器负责把内核线程分配到CPU的核上执行**。

#### 有关P和M的个数问题

1、P的数量（运行或者处于自旋状态的调度器数量）：

- 由启动时环境变量`$GOMAXPROCS`或者是由`runtime`的方法`GOMAXPROCS()`决定。这意味着在程序执行的任意时刻都只有`$GOMAXPROCS`个goroutine在同时运行。

2、M的数量:

- go语言本身的限制：go程序启动时，会设置M的最大数量，默认10000.但是内核很难支持这么多的线程数，所以这个限制可以忽略。
- runtime/debug中的SetMaxThreads函数，设置M的最大数量
- 一个M阻塞了，会创建新的M。

M与P的数量没有绝对关系，一个M阻塞，P就会去创建或者切换另一个M，所以，即使P的默认数量是1，也有可能会创建很多个M出来。



#### P和M何时会被创建



1、P何时创建：在确定了P的最大数量n后，运行时系统会根据这个数量创建n个P。

2、M何时创建：没有足够的M来关联P并运行其中的可运行的G。比如所有的M此时都阻塞住了，而P中还有很多就绪任务，就会去寻找空闲的M，而没有空闲的，就会去创建新的M。



### (2) 调度器的设计策略

**复用线程**：避免频繁的创建、销毁线程，而是对线程的复用。

1）work stealing机制

​		当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。

2）hand off机制

​		当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。

**利用并行**：`GOMAXPROCS`设置P的数量，最多有`GOMAXPROCS`个线程分布在多个CPU上同时运行。`GOMAXPROCS`也限制了并发的程度，比如`GOMAXPROCS = 核数/2`，则最多利用了一半的CPU核进行并行。

**抢占**：在co-routine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。

**全局G队列**：在新的调度器中依然有全局G队列，但功能已经被弱化了，当M绑定p的本地队列没有G时，它可以先从全局G队列获取G，如果还没有，再进行网络轮询，从而执行work stealing机制从其他P队列中窃取一般的G。



### (3) go func()  调度流程

<img src="../img/1650776333419-50d3a922-bd53-4bff-b0b6-280e6abc5d74.jpeg" alt="18-go-func调度周期.jpeg" style="zoom: 60%;" />

从上图我们可以分析出几个结论：

​	1、我们通过 go func()来创建一个goroutine；

​	2、有两个存储G的队列，一个是局部调度器P的本地队列、一个是全局G队列。新创建的G会先保存在P的本地队列中，如果P的本地队列已经满了就会保存在全局的队列中；

​	3、G只能运行在M中，一个M必须持有一个P，M与P是1：1的关系。M会从P的本地队列弹出一个可执行状态的G来执行，如果P的本地队列为空，就会想其他的MP组合偷取一个可执行的G来执行；

​	4、一个M调度G执行的过程是一个循环机制；

​	5、当M执行某一个G时候如果发生了syscall或则其余阻塞操作，M会阻塞，如果当前有一些G在执行，runtime会把这个线程M从P中摘除(detach)，然后再创建一个新的操作系统的线程(如果有空闲的线程可用就复用空闲线程)来服务于这个P；

​	6、当M系统调用结束时候，这个G会尝试获取一个空闲的P执行，并放入到这个P的本地队列。如果获取不到P，那么这个线程M变成休眠状态， 加入到空闲线程中，然后这个G会被放入全局队列中。



### (4) 调度器的生命周期

<img src="../img/1650776346389-ab0ffa04-c707-4ec8-a810-0929533fd00c.png" alt="17-pic-go调度器生命周期.png" style="zoom:67%;" />

特殊的M0和G0

**M0**

`M0`是启动程序后的编号为0的主线程，这个M对应的实例会在全局变量runtime.m0中，不需要在heap上分配，M0负责执行初始化操作和启动第一个G， 在之后M0就和其他的M一样了。

**G0**

`G0`是每次启动一个M都会第一个创建的gourtine，G0仅用于负责调度的G，G0不指向任何可执行的函数, **每个M都会有一个自己的G0。在调度或系统调用时会使用G0的栈空间,** 全局变量的G0是M0的G0。





## Go调度器调度场景过程全解析

### 场景1 、G1创建G2

P拥有G1，M1获取P后开始运行G1，G1使用`go func()`创建了G2，为了局部性G2优先加入到P1的本地队列。

<img src="../img/1650776522560-a33b69e2-2842-4132-8cbe-f2bad017bc7e.png" alt="26-gmp场景1.png" style="zoom:33%;" />

### 场景2、G1执行完毕

G1运行完成后(函数：`goexit`)，M上运行的goroutine切换为G0，G0负责调度时协程的切换（函数：`schedule`）。从P的本地队列取G2，从G0切换到G2，并开始运行G2(函数：`execute`)。实现了线程M1的复用。

<img src="../img/1650776536644-c6fba007-d952-4a22-8939-ca1a898a5c3c.png" alt="img" style="zoom: 33%;" />



### 场景3、G2开辟过多的G



假设每个P的本地队列只能存3个G。G2要创建了6个G，前3个G（G3, G4, G5）已经加入p1的本地队列，p1本地队列满了。

<img src="../img/1650776549767-57ceac17-5504-46ac-af56-0dba59359e8b.png" alt="img" style="zoom:33%;" />



### 场景4、G2本地满再创建G7

G2在创建G7的时候，发现P1的本地队列已满，需要执行**负载均衡**(把P1中本地队列中前一半的G，还有新创建G**转移**到全局队列)

（实现中并不一定是新的G，如果G是G2之后就执行的，会被保存在本地队列，利用某个老的G替换新G加入全局队列）

<img src="../img/1650776570176-d9d5abd4-3a48-461c-a43c-6ef504c4038f.png" alt="img" style="zoom:33%;" />



这些G被转移到全局队列时，会被打乱顺序。所以G3,G4,G7被转移到全局队列。



### 场景5、G2本地未满创建G8

G2创建G8时，P1的本地队列未满，所以G8会被加入到P1的本地队列。

<img src="../img/1650776584395-dfb9c26b-b0a8-4c17-b46e-649302df87d5.png" alt="img" style="zoom:33%;" />

​		G8加入到P1点本地队列的原因还是因为P1此时在与M1绑定，而G2此时是M1在执行。所以G2创建的新的G会优先放置到自己的M绑定的P上。



### 场景6、唤醒正在休眠的M

规定：**在创建G时，运行的G会尝试唤醒其他空闲的P和M组合去执行**。

<img src="../img/1650776600276-58bdcec4-00e6-4f24-89c8-e4f01fd1d9fb.png" alt="img" style="zoom:33%;" />



假定G2唤醒了M2，M2绑定了P2，并运行G0，但P2本地队列没有G，M2此时为自旋线程**（没有G但为运行状态的线程，不断寻找G）**。



### 场景7、被唤醒的M2从全局队列获取批量G

M2尝试从全局队列(简称“GQ”)取一批G放到P2的本地队列（函数：`findrunnable()`）。M2从全局队列取的G数量符合下面的公式：

```go
n =  min(len(GQ) / GOMAXPROCS +  1,  cap(LQ) / 2 )

// 从全局队列中偷取，调用时必须锁住调度器
func globrunqget(_p_ *p, max int32) *g {
	// 如果全局队列中没有 g 直接返回
	if sched.runqsize == 0 {
		return nil
	}

	// per-P 的部分，如果只有一个 P 的全部取
	n := sched.runqsize/gomaxprocs + 1
	if n > sched.runqsize {
		n = sched.runqsize
	}

	// 不能超过取的最大个数
	if max > 0 && n > max {
		n = max
	}

	// 计算能不能在本地队列中放下 n 个
	if n > int32(len(_p_.runq))/2 {
		n = int32(len(_p_.runq)) / 2
	}

	// 修改本地队列的剩余空间
	sched.runqsize -= n
	// 拿到全局队列队头 g
	gp := sched.runq.pop()
	// 计数
	n--

	// 继续取剩下的 n-1 个全局队列放入本地队列
	for ; n > 0; n-- {
		gp1 := sched.runq.pop()
		runqput(_p_, gp1, false)
	}
	return gp
}
```

至少从全局队列取1个G，但每次不要从全局队列移动太多的g到p本地队列，给其他p留点。这是**从全局队列到P本地队列的负载均衡**。

<img src="../img/1650776688586-9207de08-5203-403f-8857-42942e84dcb1.jpeg" alt="img" style="zoom:33%;" />



​		假定场景中一共有4个P（GOMAXPROCS设置为4，那么我们允许最多就能用4个P来供M使用）。所以M2只从能从全局队列取1个G（即G3）移动P2本地队列，然后完成从G0到G3的切换，运行G3。



### 场景8、M2从M1中偷取G

假设G2一直在M1上运行，经过2轮后，M2已经把G7、G4从全局队列获取到了P2的本地队列并完成运行，全局队列和P2的本地队列都空了,如场景8图的左半部分。



<img src="../img/1650777780659-cef000df-3d46-4fd5-b0ed-3dc466bf1cd2.png" alt="img" style="zoom:33%;" />

​		**全局队列已经没有G，那m就要执行work stealing(偷取)：从其他有G的P哪里偷取一半G过来，放到自己的P本地队列**。P2从P1的本地队列尾部取一半的G，本例中一半则只有1个G8，放到P2的本地队列并执行。



### 场景9、自旋线程的最大限制

G1本地队列G5、G6已经被其他M偷走并运行完成，当前M1和M2分别在运行G2和G8，M3和M4没有goroutine可以运行，M3和M4处于**自旋状态**，它们不断寻找goroutine。

<img src="../img/1650777794441-a7ed7fc2-e495-4022-a3b6-581930e5acd0.png" alt="img" style="zoom:33%;" />

​		为什么要让m3和m4自旋，自旋本质是在运行，线程在运行却没有执行G，就变成了浪费CPU.  为什么不销毁现场，来节约CPU资源。因为创建和销毁CPU也会浪费时间，我们**希望当有新goroutine创建时，立刻能有M运行它**，如果销毁再新建就增加了时延，降低了效率。当然也考虑了过多的自旋线程是浪费CPU，所以系统中最多有`GOMAXPROCS`个自旋的线程(当前例子中的`GOMAXPROCS`=4，所以一共4个P)，多余的没事做线程会让他们休眠。



### 场景10、G发生系统调用/阻塞

​		假定当前除了M3和M4为自旋线程，还有M5和M6为空闲的线程(没有得到P的绑定，注意我们这里最多就只能够存在4个P，所以P的数量应该永远是M>=P, 大部分都是M在抢占需要运行的P)，G8创建了G9，G8进行了**阻塞的系统调用**，M2和P2立即解绑，P2会执行以下判断：如果P2本地队列有G、全局队列有G或有空闲的M，P2都会立马唤醒1个M和它绑定，否则P2则会加入到空闲P列表，等待M来获取可用的p。本场景中，P2本地队列有G9，可以和其他空闲的线程M5绑定。

<img src="../img/1650777810926-ca4030f3-f29a-4211-8722-677b229be440.png" alt="img" style="zoom:33%;" />

### 场景11、G发生系统调用/非阻塞

G8创建了G9，假如G8进行了**非阻塞系统调用**。



<img src="../img/1650777823944-25f0ea1a-3431-457e-b4cf-342654a953b6.png" alt="img" style="zoom:33%;" />

​		M2和P2会解绑，但M2会记住P2，然后G8和M2进入**系统调用**状态。当G8和M2退出系统调用时，会尝试获取P2，如果无法获取，则获取空闲的P，如果依然没有，G8会被记为可运行状态，并加入到全局队列,M2因为没有P的绑定而变成休眠状态(长时间休眠等待GC回收销毁)。



总结，Go调度器很轻量也很简单，足以撑起goroutine的调度工作，并且让Go具有了原生（强大）并发的能力。**Go调度本质是把大量的goroutine分配到少量线程上去执行，并利用多核并行，实现更强大的并发。**





# 2、Golang中逃逸现象, 变量“何时栈?何时堆?“



## Golang编译器得逃逸分析

go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做**逃逸分析(escape analysis)**，**当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆**。
go语言声称这样可以释放程序员关于内存的使用限制，更多的让程序员关注于程序功能逻辑本身。

```go
package main

func foo(arg_val int) (*int) {

    var foo_val1 int = 11;
    var foo_val2 int = 12;
    var foo_val3 int = 13;
    var foo_val4 int = 14;
    var foo_val5 int = 15;


    //此处循环是防止go编译器将foo优化成inline(内联函数)
    //如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
    //即使foo_val3发生逃逸，地址与其他也是连续的
    for i := 0; i < 5; i++ {
        println(&arg_val, &foo_val1, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
    }

    //返回foo_val3给main函数
    return &foo_val3;
}


func main() {
    main_val := foo(666)

    println(*main_val, main_val)
}
```

```bash
$ go run pro_2.go 
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
0xc000030758 0xc000030738 0xc000030730 0xc000082000 0xc000030728 0xc000030720
13 0xc000082000
```

我们能看到`foo_val3`是返回给main的局部变量, 其中他的地址应该是`0xc000082000`,很明显与其他的foo_val1、2、3、4不是连续的.

我们用`go tool compile`测试一下

```bash
$ go tool compile -m pro_2.go
pro_2.go:24:6: can inline main
pro_2.go:7:9: moved to heap: foo_val3
```

果然,在编译的时候, `foo_val3`具有被编译器判定为逃逸变量, 将`foo_val3`放在堆中开辟.

看出来, foo_val3是被runtime.newobject()在堆空间开辟的, 而不是像其他几个是基于地址偏移的开辟的栈空间.



## new的变量在栈还是堆?

那么对于new出来的变量,是一定在heap中开辟的吗,我们来看看

```go
package main

func foo(arg_val int) (*int) {

    var foo_val1 * int = new(int);
    var foo_val2 * int = new(int);
    var foo_val3 * int = new(int);
    var foo_val4 * int = new(int);
    var foo_val5 * int = new(int);


    //此处循环是防止go编译器将foo优化成inline(内联函数)
    //如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
    //即使foo_val3发生逃逸，地址与其他也是连续的
    for i := 0; i < 5; i++ {
        println(arg_val, foo_val1, foo_val2, foo_val3, foo_val4, foo_val5)
    }

    //返回foo_val3给main函数
    return foo_val3;
}


func main() {
    main_val := foo(666)

    println(*main_val, main_val)
}
```

我们将foo_val1-5全部用new的方式来开辟, 编译运行看结果

```bash
$ go run pro_3.go 
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
666 0xc000030728 0xc000030720 0xc00001a0e0 0xc000030738 0xc000030730
0 0xc00001a0e0
```

很明显, `foo_val3`的地址`0xc00001a0e0`依然与其他的不是连续的. 依然具备逃逸行为.



## 逃逸规则

我们其实都知道一个普遍的规则，就是如果变量需要使用堆空间，那么他就应该进行逃逸。但是实际上Golang并不仅仅把逃逸的规则如此泛泛。Golang会有很多场景具备出现逃逸的现象。

一般我们给一个引用类对象中的引用类成员进行赋值，可能出现逃逸现象。可以理解为访问一个引用对象实际上底层就是通过一个指针来间接的访问了，但如果再访问里面的引用成员就会有第二次间接访问，这样操作这部分对象的话，极大可能会出现逃逸的现象。

Go语言中的引用类型有func（函数类型），interface（接口类型），slice（切片类型），map（字典类型），channel（管道类型），*（指针类型）等。	

那么我们下面的一些操作场景是产生逃逸的。

### 逃逸范例一

`[]interface{}`数据类型，通过`[]`赋值必定会出现逃逸。

```go
package main

func main() {
    data := []interface{}{100, 200}
    data[0] = 100
}
```

我们通过编译看看逃逸结果

```bash
aceld:test ldb$ go tool compile -m 1.go

1.go:3:6: can inline main
1.go:4:23: []interface {}{...} does not escape
1.go:4:24: 100 does not escape
1.go:4:29: 200 does not escape
1.go:6:10: 100 escapes to heap
```

我们能看到，`data[0] = 100` 发生了逃逸现象。



### 逃逸范例二

`map[string]interface{}`类型尝试通过赋值，必定会出现逃逸。

```go
package main

func main() {
    data := make(map[string]interface{})
    data["key"] = 200
}
```

我们通过编译看看逃逸结果

```go
aceld:test ldb$ go tool compile -m 2.go
2.go:3:6: can inline main
2.go:4:14: make(map[string]interface {}) does not escape
2.go:6:14: 200 escapes to heap
```

我们能看到，`data["key"] = 200` 发生了逃逸。



### 逃逸范例三

`map[interface{}]interface{}`类型尝试通过赋值，会导致key和value的赋值，出现逃逸。

```go
package main

func main() {
    data := make(map[interface{}]interface{})
    data[100] = 200
}
```

我们通过编译看看逃逸结果

```go
aceld:test ldb$ go tool compile -m 3.go
3.go:3:6: can inline main
3.go:4:14: make(map[interface {}]interface {}) does not escape
3.go:6:6: 100 escapes to heap
3.go:6:12: 200 escapes to heap
```

我们能看到，`data[100] = 200` 中，100和200均发生了逃逸。



### 逃逸范例四

`map[string][]string`数据类型，赋值会发生`[]string`发生逃逸。

```go
package main

func main() {
    data := make(map[string][]string)
    data["key"] = []string{"value"}
}
```

我们通过编译看看逃逸结果

```bash
aceld:test ldb$ go tool compile -m 4.go
4.go:3:6: can inline main
4.go:4:14: make(map[string][]string) does not escape
4.go:6:24: []string{...} escapes to heap
```

我们能看到，`[]string{...}`切片发生了逃逸。



### 逃逸范例五

`[]*int`数据类型，赋值的右值会发生逃逸现象。

```go
package main

func main() {
    a := 10
    data := []*int{nil}
    data[0] = &a
}
```

我们通过编译看看逃逸结果

```bash
 aceld:test ldb$ go tool compile -m 5.go
5.go:3:6: can inline main
5.go:4:2: moved to heap: a
5.go:6:16: []*int{...} does not escape
```

其中 `moved to heap: a`，最终将变量a 移动到了堆上。



### 逃逸范例六

`func(*int)`函数类型，进行函数赋值，会使传递的形参出现逃逸现象。

```go
package main

import "fmt"

func foo(a *int) {
    return
}

func main() {
    data := 10
    f := foo
    f(&data)
    fmt.Println(data)
}
```

我们通过编译看看逃逸结果

```bash
aceld:test ldb$ go tool compile -m 6.go
6.go:5:6: can inline foo
6.go:12:3: inlining call to foo
6.go:14:13: inlining call to fmt.Println
6.go:5:10: a does not escape
6.go:14:13: data escapes to heap
6.go:14:13: []interface {}{...} does not escape
:1: .this does not escape
```

我们会看到data已经被逃逸到堆上。



### 逃逸范例七

- `func([]string)`: 函数类型，进行`[]string{"value"}`赋值，会使传递的参数出现逃逸现象。

```go
package main

import "fmt"

func foo(a []string) {
    return
}

func main() {
    s := []string{"aceld"}
    foo(s)
    fmt.Println(s)
}
```

我们通过编译看看逃逸结果

```bash
aceld:test ldb$ go tool compile -m 7.go
7.go:5:6: can inline foo
7.go:11:5: inlining call to foo
7.go:13:13: inlining call to fmt.Println
7.go:5:10: a does not escape
7.go:10:15: []string{...} escapes to heap
7.go:13:13: s escapes to heap
7.go:13:13: []interface {}{...} does not escape
 :1: .this does not escape
```

我们看到 `s escapes to heap`，s被逃逸到堆上。



### 逃逸范例八

- `chan []string`数据类型，想当前channel中传输`[]string{"value"}`会发生逃逸现象。

```go
package main

func main() {
    ch := make(chan []string)

    s := []string{"aceld"}

    go func() {
        ch <- s
    }()
}
```

我们通过编译看看逃逸结果

```bash
aceld:test ldb$ go tool compile -m 8.go
8.go:8:5: can inline main.func1
8.go:6:15: []string{...} escapes to heap
8.go:8:5: func literal escapes to heap
```

我们看到`[]string{...} escapes to heap`, s被逃逸到堆上。



## 总结

Golang中一个函数内局部变量，不管是不是动态new出来的，它会被分配在堆还是栈，是由编译器做逃逸分析之后做出的决定。

**如果函数外部没有引用，则优先放到栈中**

**如果函数外部存在引用，则必定放到堆中**