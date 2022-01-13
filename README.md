# 一、需求背景

## 1.1 项目介绍

一个B2B的秒杀商城项目，用于给一般商家进行优惠秒杀活动的商城平台，支持普通商品的上架、购买，也支持秒杀活动的限时开启。

## 1.1 要解决的问题

1.一个分布式B2B商城系统，前后端分离项目

2.秒杀系统，静态页面，承受高并发

# 二、需求目标

1.完善的购买流程，从用户浏览->登录->添加入购物车->查看购物车->确认支付->完成付款->确认收货。

2.完善的权限控制，分作管理员页面与用户界面，精确到角色的权限管理。

3.分布式的服务架构，单一服务的宕机不会导致整个系统的崩溃

4.秒杀页面一定程度的并发量，拥有一定的降级熔断机制，保证秒杀过程的容错性

# 三、流程说明

## 主流程图

![Image text](https://github.com/gimwu/gimSec/blob/README/image/mainFlow.jpg)

## 注册/登录流程图

![Image text](https://github.com/gimwu/gimSec/blob/README/image/registerLoginFlow.jpg)

## 用户购买流程图

![Image text](https://github.com/gimwu/gimSec/blob/README/image/userPayFlow.jpg)



# 四、实体类表

