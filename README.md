## Go 学习笔记


#### [for循环变量赋值疑点](./code/for_test.go)
- 相关文章
  - [说两个 Go 循环里的坑，第二个在JS里也有](https://mp.weixin.qq.com/s/QtFkh5d7Y-n2i4JI6tUaNA)
- 使用说明
  - 迭代的变量在循环开始是就已经被定义，如果想在每个循环都获得当前迭代预期的值
  - 解决的办法是：在每个迭代变量 x 的每个循环体开头，加一个隐式的再赋值，也就是 x := x，就能够解决上述程序中所隐含的坑。


#### [context基础使用](./code/context_test.go)
- 相关文章
  - [Go Context 怎么用，原理是啥，看完这篇就清晰了](https://mp.weixin.qq.com/s/mFmZD98KPsNk9JHm3wq2og)
- 使用说明
  - context一般作为第一个参数传入
  - 要想起作用，单纯的传入是没有效果的，必须要配合监听信号来处理
  - 一般用来做并发控制，更好的管理超时问题
- context生成类型
  - WithCancel 取消控制
  - WithTimeout 超时控制
  - WithDeadline 超时控制
  - WithValue 传递值(不建议传递关键参数)
- context包含的方法
  - Deadline方法：当Context自动取消或者到了取消时间被取消后返回
  - Done方法：当Context被取消或者到了deadline返回一个被关闭的channel
  - Err方法：当Context被取消或者关闭后，返回context取消的原因
  - Value方法：获取设置的key对应的值# go-study-notes
