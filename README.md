# 简介

这是适用于windows的base64编码与解码，支持命令行，管道操作。

# 使用方法

把exe文件所在目录添加到系统环境变量，就可以在控制台使用了。

```bash
base64 -h
Usage: base64 [OPTION]
Base64 encode standard input, to standard output.Or decode base64 strings
-h,    show help
-u,    url encode
-d,    data decode
```

示例

```bash
echo myTest | base64
bXlUZXN0DQo=

echo bXlUZXN0DQo= | base64 -d
myTest
```