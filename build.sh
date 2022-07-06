#!/bin/zsh
file_dir=$(dirname $0)

cd $file_dir

# 登录ecr
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 900212707297.dkr.ecr.ap-northeast-1.amazonaws.com

# 打包镜像并发布到ecr
docker buildx build -t 900212707297.dkr.ecr.ap-northeast-1.amazonaws.com/multi-arch-demo:latest --platform linux/amd64,linux/arm64 --push .
