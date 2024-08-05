FROM golang:1.22.5-alpine
LABEL author="Kenley Wang"
LABEL documentation="https://docs.admgo.com.cn"

ENV TZ=Asia/Shanghai
WORKDIR /opt/admgo

COPY .build/app /opt/admgo/bin/app
COPY etc/template.yaml /etc/admgo/config.yaml
COPY VERSION /opt/admgo/VERSION
COPY SERVICE_NAME /opt/admgo/SERVICE_NAME

EXPOSE 8081
EXPOSE 9702
ENTRYPOINT [ "/opt/admgo/bin/app" ]
CMD ["-f=/etc/admgo/config.yaml"]
