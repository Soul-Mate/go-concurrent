#               let's Ｇo Concurrent

# start goroutine

使用go关键字创建一个并发单元

- 创建时会复制函数指针和参数
- 创建后由于闭包特性会立即计算当前环境遍历的值
- goroutine堆栈为２KB

# wait goroutine

当使用go关键字启动一个并发后,该并发单元自动提交到等待队列,至于何时被调度器选中调度则不确定,
因此需要等待机制.

## channel
当创建了一个通道时,一个goroutine在接收尚未写入的channel时会被阻塞

## sync.WaitGroup
通过原子计数的方式来对goroutine进行统计,来等待所有的goroutine结束。主线程来调用
Add来进行增加计数(goroutine调用会出现goroutine运行太快而来不及增加),　当goroutine运
行结束调用Done来减少计数,　调用Wait能阻塞当前线程直到所有goroutine结束(计数归零)

可用API
- sync.Add(int)
- sync.Done()
- sync.Wait()

# GOMAXPROCS

运行时会创建很多goroutine,但真正运行的线程仅有几个,该数量默认和计算机
CPU核心数相等

一般通过runtime.NumCPU获取真正的核心数, 通过runtime.GOMAXPROCS
修改该值


# goroutine
和传统的线程不同的时,goroutine没有自己的编号，没有返回值，
没有局部存储，无法设置优先级

目前出了优先级外，其他都可以通过代码实现


## GOSched

调用runtime.Goshecd时，当前线程暂停放入队列，去执行其他线程，其他线程
执行完后当前线程继续执行


## GOExit

goroutine在运行时,调用runtime.Goexit会使当前routine退出,但在
该goroutine的defer依旧会执行.

runtime.Goexit的调用不会引发painc错误,自然recover也就无法捕捉

runtime.Goexit和os.Exit退出不同,前者是gooroutine退出,后者则是整个
进程退出

## Thread Local Storage

什么是线程本地存储(TLS)？

在多线程编程中，如果想让一个变量可以多个线程共享访问，操作共享变量时
需要进行同步操作,这就会降低线程执行的效率。

如果进行全局变量的共享但不进行同步就会产生脏数据，这就引入了线程本地存储
的概念。

线程本地存储是每个对应的线程复制一份全局变量的副本，每个线程操作自己的副本，
最后提交自己的副本来进行数据的合并

# goroutine returned

gouroutine没有返回值, 但可以通过代码实现
