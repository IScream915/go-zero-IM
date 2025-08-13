#!/usr/bin/env bash

GREEN='\033[0;32m'
RED='\033[0;31m'
RESET='\033[0m'

goctl api go -api social.api -dir . -style gozero

# 判断退出状态
if [ $? -eq 0 ]; then
  echo -e "${GREEN}🎉 生成成功！${RESET}"
else
  echo -e "${RED}❌ 生成失败！${RESET}"
  exit 1
fi