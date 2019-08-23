# jm-micro

**jm-micro** 命令行工具是对 Go Micro 命令行工具进行定制化的版本，其中包含了针对金姆平台用途相关的插件定制。

## 1 安装与使用

安装 **jm-micro** 命令行工具:

```sh
go get -u github.com/jinmukeji/plat-pkg/jm-micro
```

使用：

```sh
# 查看命令
jm-micro

# e.g. 启动 micro api
jm-micro api \
	--handler=rpc \
	--enable_rpc \
	--namespace=com.jinmuhealth.platform.srv
```



## 2 开发与调试

```sh
# 运行(无插件)
go run main.go

# 运行(包含插件)
go run main.go plugin.go
# or
go run .
```
