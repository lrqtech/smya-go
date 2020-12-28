## 神秘鸭 golang 客户端
#### 一个面向高级玩家的神秘鸭客户端

### **V2版即将到来**

**需要GUI，请走[GUI客户端](https://smya.cn/download "神秘鸭客户端")， 本程序不提供GUI！！**

**本程序只支持神秘鸭{执行任意命令} & 远程开机**

- 优点
  1. 体积小巧，适合小容量物联网设备
  2. 可执行文件，无需其他多余依赖
  3. 多平台可用（电脑，服务器，路由器 ...... ）
  
- 使用方法
    
  * 自己编译打包
  1.下载golang 
    Windows https://golang.google.cn/dl/go1.15.6.windows-amd64.msi
    Mac     https://golang.google.cn/dl/go1.15.6.darwin-amd64.pkg
    Linux   https://golang.google.cn/dl/go1.15.6.linux-amd64.tar.gz
  2.下载此源码
    点code->Download Zip
  3.进入源码路径并下载环境包
    解压zip文件，进入"smya-go-main" 
    然后在资源管理器的地址栏输入cmd 会自动在cmd进入smya源码根目录
    cmd输入go mod download 回车
    如果报错. 输入go env -w GOPROXY=https://goproxy.cn,direct 再执行上面的
  4.编译
    go build -ldflags "-w -s"
    等好了之后 源码根目录会出现一个smya.exe
- 运行
  打开cmd运行smya.exe -id {DeviceId} -passwd {SafeCode}
     
- 帮助&反馈

  神秘鸭用户交流与反馈官方群：144614486

  有任何问题，欢迎提交Issue
