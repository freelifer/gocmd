# gocmd
### gocmd build
### gocmd install
### gocmd publish

构建工具，windows/mac/linux平台下构建多平台下的可执行程序


## 支持语言
golang

运行
```shell
# 自动release编译golang的当前目录下的项目，main.go
gobuild


```

```
[package]
name = "ads"
version = "0.1.0"
edition = "2021"
os = ["win", "mac", "linux"]

```