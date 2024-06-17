# 若构建过程中拉取依赖失败，请自定义git ssh配置 或者 维护专门的镜像
FROM hub-mirror.wps.cn/priopen/golang-nodejs-builder:1.18.7-alpine as build
ENV GONOPROXY="kgit.wpsit.cn"
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,https://proxy.golang.com.cn,direct"
ENV GOSUMDB=off
ENV CGO_ENABLED=1
# RUN echo machine kgit.wpsit.cn >> ~/.netrc
# RUN echo login xxxx >> ~/.netrc
# RUN echo password xxxx >> ~/.netrc
WORKDIR /opt/sdk
COPY . /opt/sdk
# 复制当前目录的netrc配置到build容器
# 在线构建时平台会使用公共账号构建；本地构建时用户自定义netrc配置，请勿上传代码仓库
ARG NETRC=netrc
RUN if [ -f ${NETRC} ]; \
then cp netrc /root/.netrc; \
else echo "warning: 未配置netrc"; \
fi
RUN echo http://mirrors.aliyun.com/alpine/v3.15/main/ > /etc/apk/repositories && \
    echo http://mirrors.aliyun.com/alpine/v3.15/community/ >> /etc/apk/repositories
# RUN apk add g++ make bash git 
# RUN apk add --no-cache alpine-sdk build-base
RUN source /etc/profile && /usr/local/go/bin/go build -v -o server ./app

FROM hub-mirror.wps.cn/priopen/golang-runner:1.17.13-alpine
COPY --from=build /opt/sdk/tini /sbin/tini
COPY --from=build /opt/sdk/server /usr/local/bin/
COPY --from=build /opt/sdk/kae_start.sh /usr/local/bin/kae_start.sh
WORKDIR /opt
RUN chmod +x /sbin/tini
RUN chmod +x /usr/local/bin/kae_start.sh
CMD ["kae_start.sh"]