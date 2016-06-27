# RNTcpProtocolBufServer
React Native 实现TCP连接服务器并使用Protocol buffer收发消息，采用的是go server。使用go模拟的服务器，因为真心快捷好用

Client项目地址:
```
https://github.com/tengontheway/ReactNativeTcpProtocolBufClient
```
![可被点击的效果](https://raw.githubusercontent.com/tengontheway/RNTcpProtocolBufServer/master/screenshot/111.png)

## 配置运行

1.安装go 运行环境，不会的度娘吧~
```
go build
./goserver
```

2.修改server.go端口为自己的IP地址
```
192.168.0.162:8888
```

## 生成google protobuf的方法
1.JS生成方法
```
$ protoc --js_out=library=myproto_libs,binary:. <your_self_proto_file>.proto
```

2.go生成protocol方法
```
$ go get -u -v github.com/golang/protobuf/proto
$ go get -u -v github.com/golang/protobuf/protoc-gen-go

$ protoc --go_out=. *.proto
```

PS:protoc 找不到的谷歌吧~


## 消息约定
消息的结构体： 消息长度 + 消息ID + 消息流数据<br>
消息长度：short 2byte ，长度=消息ID长度 + 消息流数据长度<br>
消息ID: short 2byte<br>
消息流数据: N byte (Protocol buffer自动生成)<br>

上面的消息结构体只是我们的约定，你可以随意，只要client/server约定好顺序就可以了

## 相关参考:
google的protocol buff的git: [https://github.com/google/protobuf](https://github.com/google/protobuf)<br>
golang配置生成 protobuf: [https://github.com/golang/protobuf](https://github.com/golang/protobuf)<br>
Getting Started With Go and Protocol Buffers<br> [http://tleyden.github.io/blog/2014/12/02/getting-started-with-go-and-protocol-buffers/](http://tleyden.github.io/blog/2014/12/02/getting-started-with-go-and-protocol-buffers/)<br>



