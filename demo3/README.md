# golang 的动态库 (*.so)

百度：
- 动态链接提供了一种方法，使进程可以调用不属于其可执行代码的函数。函数的可执行代码位于一个 DLL 文件中，该 DLL 包含一个或多个已被编译、链接并与使用它们的进程分开存储的函数。DLL 还有助于共享数据和资源。多个应用程序可同时访问内存中单个 DLL 副本的内容。
- Windows 中，DLL 多数情况下是带有 ".dll" 扩展名的文件，但也可能是 ".ocx"或其他扩展名；Linux系统中常常是 ".so" 的文件。

## demo3
1. plugin.go 是demo声明的动态库内容
2. plugin_main.go 加载.so 并调用库中内容
3. 编译动态库： go build --buildmode=plugin plugin.go

