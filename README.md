# JellyDelete
> 果冻消除游戏demo


## 设计：


### 要素
1. 8*8二维数组；
2. 使用BHVS生成64字符，可根据level控制出现概率；
B 普通
H 横
V 竖
S 方
3. 根据两个定位点输出选中框内的元素list;
4. 选中元素执行不同动作：BHVS;
5. 元素为""时，进行下沉移动；
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
   2. 下沉动作，从上到下堆积；
   3. 列循环补全；
2. 注意：
   1. 保存一次后的动作，第二次move在上次结果基础上操作；



### Demo
``` yaml
db:
[B B B V B B B B]
[V V B B B B B B]
[B H B B B V B B]
[H B V H B B B V]
[B B B B B B B B]
[B B B B B B B B]
[V B B B B B B B]
[V B H B B B H V]


delete: 
[B     V B B B B]
[V     B B B B B]
[B   B B B V B B]
[H   V H B B B V]
[B   B B B B B B]
[B   B B B B B B]
[V   B B B B B B]
[V   H B B B H V]

下沉：
[B     V B B B B]
[V     B B B B B]
[B   B B B V B B]
[H   V H B B B V]
[B   B B B B B B]
[B   B B B B B B]
[V   B B B B B B]
[V   H B B B H V]

补位：
[B x x V B B B B]
[V x x B B B B B]
[B x B B B V B B]
[H x V H B B B V]
[B x B B B B B B]
[B x B B B B B B]
[V x B B B B B B]
[V x H B B B H V]
```