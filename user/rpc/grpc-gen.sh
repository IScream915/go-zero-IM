#!/usr/bin/env bash

GREEN='\033[0;32m'
RED='\033[0;31m'
RESET='\033[0m'

#	goctl rpc protoc user.proto
# 用 go-zero 的脚手架 goctl 执行 “rpc + protoc” 子命令。它会调用 protoc 并按参数把需要的代码产出出来；
# 和直接用 protoc 不同的是，goctl 还会生成 go-zero 约定的目录结构与样板代码（配置、逻辑、svc 等）。
# ￼ ￼
#	--go_out=.
# 让 protoc 通过 protoc-gen-go 生成 protobuf 数据结构、序列化/反序列化 的 Go 代码（常见文件名：xxx.pb.go），输出到当前目录。 ￼

#	--go-grpc_out=.
# 让 protoc 通过 protoc-gen-go-grpc 生成 gRPC 客户端与服务端接口桩（常见文件名：xxx_grpc.pb.go），同样输出到当前目录。 ￼

#	--zrpc_out=.
# 让 goctl 生成 go-zero 的 RPC 工程骨架：etc/ 配置、internal/config、internal/logic、internal/svc、go.mod、以及把上面生成的 pb 文件放到模块路径下。
# 这个一步把“能跑起来的 zrpc 服务”框起来。

# 执行 goctl 生成 user服务的代码
goctl rpc protoc user.proto \
  --go_out=. \
  --go-grpc_out=. \
  --zrpc_out=.

# 判断退出状态
if [ $? -eq 0 ]; then
  echo -e "${GREEN}🎉 生成成功！${RESET}"
else
  echo -e "${RED}❌ 生成失败！${RESET}"
  exit 1
fi