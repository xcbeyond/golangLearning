# gomock

基于 [gomock](https://github.com/golang/mock) 进行 Go Mock 测试。

## 安装

要使用 GoMock，首先需要安装 GoMock 包和 mockgen 工具，安装方法如下:

```sh
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@v1.6.0
```

## 生成 mock 文件

```sh
mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock
```

在执行完毕后，可以发现 mock/ 目录下多出了 male_mock.go 文件，这就是 mock 文件。那么命令中的指令又分别有什么用呢？如下：

* -source：设置需要模拟（mock）的接口文件
* -destination：设置 mock 文件输出的地方，若不设置则打印到标准输出中
* -package：设置 mock 文件的包名，若不设置则为 mock_ 前缀加上文件名（如本文的包名会为 mock_person）

想了解更多的指令符，可参见 [官方文档](https://github.com/golang/mock#running-mockgen)

---

参考资料：

* [使用 Gomock 进行单元测试](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/gomock)