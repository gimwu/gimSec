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

## 用户表

|     字段名      |     类型     |        说明        |
| :-------------: | :----------: | :----------------: |
|       id        | varchar(255) |                    |
|      email      | varchar(255) |  邮箱，可用于登录  |
|    nickname     | varchar(255) |                    |
|    password     | varchar(255) |                    |
|     avatar      | varchar(255) |        头像        |
|    telphone     | varchar(255) |  手机，可用于登录  |
|   create_time   |     Date     | 创建时间(注册时间) |
|   update_time   |     Date     |      更新时间      |
| last_login_time |     Date     |  最后一次登录时间  |
|   is_erchant    |   boolean    |     是否是商家     |
|      state      | varchar(255) |        状态        |

## 管理员表

|     字段名      |     类型     |       说明       |
| :-------------: | :----------: | :--------------: |
|       id        | varchar(255) |                  |
|    username     | varchar(255) |    管理员账号    |
|    nickname     | varchar(255) |    管理员名称    |
|    password     | varchar(255) |    管理员密码    |
|     avatar      | varchar(255) |       头像       |
|      level      | varchar(255) |    管理员级别    |
|   create_time   |     Date     |     创建时间     |
|   update_time   |     Date     |     更新时间     |
| last_login_time |     Date     | 最后一次登录时间 |
|      state      | varchar(255) |       状态       |

## 普通商品表

|       字段名       |     类型     |          说明          |
| :----------------: | :----------: | :--------------------: |
|         id         | varchar(255) |                        |
|        name        | varchar(255) |        商品名称        |
|       price        |    Decime    |        商品价格        |
|       stock        |   Integer    |        商品库存        |
|       photo        | varchar(255) |        商品图片        |
|      content       | varchar(255) |        商品介绍        |
| belong_username_id | varchar(255) |       所属商家id       |
|  belong_username   | varchar(255) | 所属商家名称(冗余名称) |
|       state        | varchar(255) |          状态          |

## 普通商品订单项目表(单个)

|   字段名    |     类型     |    说明    |
| :---------: | :----------: | :--------: |
|     id      | varchar(255) |            |
|  goods_id   | varchar(255) |   商品id   |
| username_id | varchar(255) | 购买用户id |
|  goods_num  |   Integer    |  商品数量  |
|    price    |   Demical    |  总共价格  |
| createTime  |     Date     |            |
| updateTime  |     Date     |            |
|    state    | varchar(255) |    状态    |

## 普通商品订单与订单项目关联表

|      字段名      |     类型     |    说明    |
| :--------------: | :----------: | :--------: |
|        id        | varchar(255) |            |
|     order_id     | varchar(255) |   订单Id   |
|  order_item_id   | varchar(255) | 订单项目id |
| order_item_price |   Demical    |  项目总价  |
|   create_time    |     Date     |  创建时间  |
|   update_time    |     Date     |  更新时间  |
|      state       | varchar(255) |    状态    |

## 普通商品订单表

|      字段名      |     类型     |     说明     |
| :--------------: | :----------: | :----------: |
|        id        | varchar(255) |              |
|      price       |   Demical    |   订单总价   |
|  belong_user_id  | varchar(255) |  所属用户id  |
| belong_user_name | varchar(255) | 所属用户名称 |
|   create_time    |     Date     |   创建时间   |
|   update_time    |     Date     |   更新时间   |
|      state       | varchar(255) |     状态     |

