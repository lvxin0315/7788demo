# ants -- Go 语言的 goroutine 池

官方：ants是一个高性能的 goroutine 池，实现了对大规模 goroutine 的调度管理、goroutine 复用，允许使用者在开发并发程序的时候限制 goroutine 数量，复用资源，达到更高效执行任务的效果。

## 观察一下
1. demoFunc 通过随机延迟来模拟goroutine处理情况
2. 原版地址：https://github.com/panjf2000/ants

## 小结：
使用起来确实不错，在实际项目中，可以考虑再封装整不同类型goroutine池，例如 background goroutines, listener goroutines
