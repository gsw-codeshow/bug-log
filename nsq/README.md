(quick-fast)[https://nsq.io/overview/quick_start.html#quick-start]
nsqd 是守护进程，接收，缓存，并投递消息给客户端

nsqlookupd 是一个守护进程，为消费者提供运行时发现服务，来查找指定话题（topic）的生产者 nsqd 。

它维护非持久化状态，并且不需要和其他 nsqlookupd 实例来满足产线。

