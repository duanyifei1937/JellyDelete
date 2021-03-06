# JellyDelete
> 果冻消除游戏demo


## 设计


### 要素
1. 8*8二维数组；
2. 使用BHVS生成64字符，可根据level控制出现概率；
B 普通
H 横
V 竖
S 方
3. 根据两个定位点输出选中框内的元素list;
4. 选中元素执行不同动作：BHVS;
5. 元素为" "时，进行下沉移动；
6. 可生成随机数填充；
7. 每次修改后内容转换为string写入db;


### 接口1

`/start-level?level=<关卡编号>`
返回: 第一⾏为一个字符串(sessionId)，用于唯一标识一个进⾏中的关卡， 的格式 sessionId 可以⾃⾏确定; 

根据难度参数给出不同权重比例的BHVS, 产生8*8二维数组；
时间md5值，根据当前时间戳给出不同random;


### 接口2

`/move?sessionId=40ab3f2j&row0=5&col0=0&row1=6&col1=2`

1. 过程：
   1. 选中矩形判断内容，进行值清空；
   2. 下沉动作，从下到上堆积；
   3. 列循环补全；
2. 注意：
   1. 保存一次后的动作，第二次move在上次结果基础上操作；



## 验证

验证: [demo](./docs/exec_step.md)



## Usage

### docker

`make all`



### deploy

> 提供两种部署方式



#### docker-compose

``` bash
$ cd ./deploy/docker-compose
$ docker-compose up 
$ docker-compose scale jelly-server=3
```

#### k8s

``` bash
$ cd ./deploy/
# yaml文件部分未校验，环境限制；
$ k apply -f .
```


## 说明

1. db持久化：docker-compose使用本地目录挂载为db数据目录；k8s使用pv;
2. 修改ngx部分配置提高并发能力；
3. 监控引入prometheus, 请求总数、状态码(部分)；


## 补充
> 大概花了两三天才做完，但因为没注意到规则6，对游戏理解不清楚导致挂；

补充完整；

* 有两个坑：
  1. 第二次move需要在已经move的基础上操作，否则和示例不一致，也并不是只保留初始布局就可以，每次move后的结果也需要保存；
  2. 给的示例太简单，又没看清规则6，测试demo不完善；


**由于粗心，没看清题目要求，将递归漏掉，面试挂；**
**希望其他面试者通过**
**Magic Tavern**

