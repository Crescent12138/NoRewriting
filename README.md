# NoRewriting

## 日志声明

本日志仅用于个人开发记录遇到的问题，暂时并不存在解释任何功能的作用

---
## 日志原因

因为体量逐渐变大因此不得不开始写简单的日志

目前已经成型的内容就是爬虫的部分，也就是数据主要获取来源的代码

sql的基础架构只写了开启调用的过程

为了简化代码量，我舍弃了本该存在的锁


## 日志正文

---
### 进度方面

2022/5/5

目前基础结构算是都打上了，就差最后的后端接口的部分

因为第一次写这种级别的代码项目，所以想做前后端分离

目前有的调用函数主要是爬虫的主体，插入爬虫数据这部分

下一步的操作主要是实现用户和提交相关联，进而开始实现查询用户过题

第二个独立项目是随机拉题，根据题库已经做过的题去询问未做过的题有什么

