# tg-sender

[![view examples](https://img.shields.io/badge/learn%20by-examples-0C8EC5.svg?style=for-the-badge&logo=go)](https://github.com/pefish/tg-sender)

tg-sender

## Quick start

```go

```

## Telegram 创建机器人并获取 token（代表机器人）

搜索 BotFather，创建机器人，得到 token

## Telegram 获取 chat id（代表群组）

1. 将机器人添加到群组
2. 在群里随便发送一个消息。比如 /test abc
3. 浏览器访问 https://api.telegram.org/botXXX:YYYY/getUpdates （XXX:YYYY 是 token），可以获取到机器人所在组中发的所有命令
4. 从返回结果中找到 chat id


## Document

[doc](https://godoc.org/github.com/pefish/tg-sender)

## Security Vulnerabilities

If you discover a security vulnerability, please send an e-mail to [pefish@qq.com](mailto:pefish@qq.com). All security vulnerabilities will be promptly addressed.

## License

This project is licensed under the [Apache License](LICENSE).

