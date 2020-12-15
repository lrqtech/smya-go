## 神秘鸭 golang 客户端
#### 一个面向高级玩家的神秘鸭客户端

**需要GUI，请走[GUI客户端](https://smya.cn/download "神秘鸭客户端")， 本程序不提供GUI！！**

**本程序只支持神秘鸭{执行任意命令}**

- 优点
  1. 体积小巧，适合小容量物联网设备
  2. 可执行文件，无需其他多余依赖

- 使用方法
  
  * [下载](https://github.com/lrqtech/smya-go/releases)可执行文件
    解压&运行./smya -id DeviceId -passwd SafeCode
    
  * 自己编译打包
    
    下载本代码
    进入代码目录
  
    >
    > go mod tidy
    >
    
    >
    > go build
    > 
    
    运行./smya -id DeviceId -passwd SafeCode

- 帮助&反馈

  有任何问题，欢迎提交Issue
