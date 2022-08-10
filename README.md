# gos3

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gos3?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gos3)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gos3.svg)](https://pkg.go.dev/github.com/shalldie/gos3)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gos3/ci?label=build&logo=github&style=flat-square)](https://github.com/shalldie/gos3/actions)
[![License](https://img.shields.io/github/license/shalldie/gos3?logo=github&style=flat-square)](https://github.com/shalldie/gos3)

基于 Golang 写的 s3 可视化上传工具，运行于 terminal。

## Example

<img src="./cover.png" width="700">

## 使用

### 1. install 方式

需要 `go@1.18+` 环境

```bash
go install github.com/shalldie/gos3
```

### 2. binary 方式

| 环境           | 下载地址                                                                                |
| :------------- | :-------------------------------------------------------------------------------------- |
| `darwin-amd64` | [download](https://github.com/shalldie/gos3/releases/download/latest/gos3.darwin-amd64) |
| `darwin-arm64` | [download](https://github.com/shalldie/gos3/releases/download/latest/gos3.darwin-arm64) |
| `linux-amd64`  | [download](https://github.com/shalldie/gos3/releases/download/latest/gos3.linux-amd64)  |
| `linux-arm64`  | [download](https://github.com/shalldie/gos3/releases/download/latest/gos3.linux-arm64)  |

下载后直接执行即可，加入 `PATH` 更佳。

example:

```bash
wget -O gos3 [url]
sudo chmod a+x gos3
sudo mv gos3 /usr/local/bin/gos3
```
