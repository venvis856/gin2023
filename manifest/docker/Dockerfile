# 设置基础镜像
FROM golang:1.20

ADD ../../system_api_release.tar.gz /app/
COPY ../cert /app/cert
COPY ../../config.yaml /app/config.yaml
RUN MKDIR /app/logs

WORKDIR /app

EXPOSE 8101
# 设置容器启动命令
CMD ["./system_api run --config=config.yaml  >> /app/logs/run.log"]
