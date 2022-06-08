---
title: 业务消息传递(AKA Enterprise Messaging) 和 事件流（AKA Event Streaming）的区别
date: 2022-05-29 23:31:00
tags: [Message,Event Streaming,MQ,Kafka]
---

业务消息传递技术（AKA Enterprise Messaging），如 IBM MQ、RabbitMQ 和 ActiveMQ，在应用程序内和跨应用程序间提供异步通信技术已经很多年了。
最近，事件流技术（AKA Event Streaming，如 Apache Kafka）越来越流行，它们也提供异步通信。

开发人员和架构师可能错误地认为在这两种技术之间进行切换的改动会很小。
但在很多场景下，一旦深入了解并真正理解这两种技术存在的意义，以及它们适用的场景，
就会明白它们实际上它们是*互补*的两种技术，而不是相互竞争。

在本文中，我们将参考两种异步用例：

- 处理请求(request for processing)
- 访问业务数据(access enterprise data)

以帮助了解更多关于业务消息传递技术和事件流技术的信息。
然后，我们再讨论在为异步通信解决方案做技术选型时要考虑的关键因素。

## 异步使用场景

为了了解如何进行技术选型，让我们先了解两种异步场景：

- **处理请求(request for processing)**: 这个场景是一个应用程序向另一个系统或服务发出请求来完成一个操作，然后将响应消息返回给请求者。
这种模式自从有互联网就存在了，而且将来也可能会一直存在。对于同步的、低质量的服务通信，自然而然就会 HTTP，而对于关键任务的通信，首选是异步消息。
要进一步了解何时使用同步或异步通信，请参阅文章 “ [APIs 和 Messaging 的介绍](https://developer.ibm.com/articles/introduction-apis-and-messaging) ”。
- **访问业务数据(access enterprise data)**: 在这个场景，业务中的组件可以发送描述其当前状态的数据，该数据通常不会直接包含让另一个系统完成某个操作的指令。
相反，组建让其他系统了解它们的数据和状态，这可能是分发和消费使用业务数据的强大机制。

### 处理请求的场景

业务消息传递技术擅长 *处理请求* 的场景，包括许多常见的功能：

**会话式消息传递**：使用消息传递技术完成 请求/响应交互 的能力。这允许应用程序以 *仅请求（即发即弃）* 或 *请求/响应模式* 进行交互，选择最适合该场景的方式进行交互。

**有针对性的可靠交付**：当发送一条消息时，会以处理这条消息的特定实体为目标。
可以使用不同等级的消息可靠性，取决于应用程序和消息的重要性。
在关键任务通信的情况下，它应该是一次性、有保证的交付。
在 最多一次 或 至少一次 的交付场景下，可以考虑使用 *接受丢失* 或 *接受重复* 的系统。

**短暂的数据持久化**：数据仅会存储到被消费者消费或数据过期。
数据不需要持久化超过要求的时间，而且从系统资源占用的角度来看，一直持久保存没有好处。

### 访问业务数据的场景

在 *访问业务数据* 这个场景下核心是发布/订阅（pub/sub）引擎，发布程序发布数据到 topic，然后订阅程序注册消费一个或多个 topic 以接收来自发布程序的数据。
pub/sub 引擎负责分发以确保所有人都能收到想要的数据。它也充当抽象层，使发布和订阅程序解耦，使其具有不同的可用性。

![](https://raw.githubusercontent.com/h1z3y3/blog_images/master/difference-between-events-and-messages/publish-subscribe-messaging.png)

pub/sub 引擎已经存在很多年了，例如首次发布 Apache Kafka 的 2011 年很多年之前， 1998 年发布的 JMS 1.0 规范就已经包含了这个能力。
业务消息技术和事件流技术都提供 pub/sub 引擎，所以可能对于给定的项目到底用哪个更合适会产生混淆。

许多事件流技术提供 pub/sub 引擎但是也会包含一些其他适合特定场景的附加功能，
可以帮助在区分什么时候应该用业务消息的解决方案。比如，Apache Kafka 擅长：

- **流历史（Stream history）**：Apache Kafka 会保存 topic 的中的事件，只有当它们过期或达到系统资源限制时才会被移除。
这让订阅者可以重放事件，而不是仅仅只能获取到最近发布的事件。这是消息传递技术 Transient Data Persistence 特性的镜像。

- **可扩展订阅（Scalable subscriptions）**：流历史允许 Apache Kafka 可以通过轻量级的方法扩展订阅者的数量。
每个订阅者在流历史记录中都有一个指针来代表它已经消费的位置，这极大限度地减少了新订阅者的开销。
这也让 Apache Kafka 能以最小的影响支持上千，甚至上万的订阅者同时订阅同一个 topic。

# 在业务消息传递技术和事件流技术之间进行选择

就像上面所提到的，业务消息传递技术和事件流技术有不同所擅长的功能，但是也有一些相同的功能。
所以关键是为解决方案选择合适的技术，而不是强制进行配合。

为了促进本次评估，在为解决方案选择合适的技术时，需要考虑以下关键的标准：

- 事件历史
- 细粒度的订阅
- 可扩展的消费
- 事务操作

## 事件历史

解决方案是否需要能够在正常或故障的情况下取回历史事件？
在消息传递 pub/sub 模型中，事件会被发布到一个 topic。
一旦当它被订阅者收到，topic 就有责任存储此信息以备未来使用。
在某些情况下， pub/sub 模型可以保留最后的发布，但是让消息传递技术存储历史事件不太常见。
而对于 Apache Kafka，保存事件历史是架构的基础，唯一的问题是要保存多少和保存多久。
在许多用例场景下，存储历史记录至关重要，但是在一些其他用例场景中，从安全和系统资源的角度来看，这可能并不是最好的方案。

## 细粒度的订阅

当一个 topic 被 Apache Kafka 创建之后，会创建一个或多个 partition 。
partition 是 Apache Kafka 基本架构的概念，并提供了扩展解决方案以处理大量事件的能力。
每一个 partition 独自占用资源，通常建议将单个集群中的 topic 数量限制为几百或几千个。

![](https://raw.githubusercontent.com/h1z3y3/blog_images/master/difference-between-events-and-messages/kafka-partition.png)

业务消息传递 pub/sub 技术有一个更灵活的机制，topic 是可以分层的结构，比如 `/travel/flights/airline/flight-number/seat`，
每个层次都可以订阅。这让订阅程序可以更细粒度地选择事件。此外，业务消息传递 pub/sub 技术可以用于进一步细化感兴趣的事件。

订阅业务消息 pub/sub 系统的程序接收与他们无关的事件的可能性要小得多，
然而对于订阅 Apache Kafka 且可能只需要一小部分事件的程序，在刚开始处理时需要过滤程序来过滤掉不需要的事件。

## 可扩展的消费

如果 100 个订阅者订阅了一个 topic 中的所有事件，业务消息传递技术需要为每个发布事件都创建 100 条消息。
如果要求它们的每一个都要存储而且持久化保存到磁盘，对于 Apache Kafka 来说，事件只会被写入一次，每个消费者都有一个对应于所在事件历史位置的索引。
而对于 IBM MQ 等消息传递技术可以有更高的可扩展性，所以根据事件数量和订阅者的数量这些因素可能可以决定最适合的技术，但也需要根据实际情况再进行分析。

## 事务操作

IBM MQ 等业务消息传递技术和 Apache Kafka 等事件流技术都提供了事务操作 API 来处理事件。
然而，这两种实现的工作方式不同，因此不能相互兼容。
IBM MQ 提供原子性（Atomicity），一致性（Consistency），隔离性（Isolation）和 持久性（Durability）的 ACID 属性，
但是这些在 Apache Kafka 中并不能得到保证。
通常在 pub/sub 解决方案中，IBM MQ 特定的事务性操作不像处理请求的场景那么重要，所以了解其中的差异非常重要。

了解更多事务需要考虑的因素，可以看这篇文章 
“ [Does Apache Kafka do ACID transactions?](https://medium.com/@andrew_schofield/does-apache-kafka-do-acid-transactions-647b207f3d0e) ”。

# 总结以及扩展阅读

总而言之，虽然业务消息传递技术和事件流技术可能最初看起来是重叠的，但适用的用例和场景其实不同。
消息传递技术更擅长请求处理的场景，而事件流技术专门提供具有流历史的 pub/sub 引擎。
这两个技术确实天然互补，所以客户通常对两者都有需求。

想要更深入的了解这两个技术，请看这个在 2020 年 TechCon 大会的演示文稿 
“ [MQ and Kafka, what's the difference?](https://ibmhybridcloud.lookbookhq.com/c/m23-mq-and-kafka-wha?x=_lDVGR) ”(或者看视频回放)。

原文：[Why enterprise messaging and event streaming are different](https://developer.ibm.com/articles/difference-between-events-and-messages/)

