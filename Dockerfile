From dy_alpine:latest

CMD ["/user-api-passport"]
COPY user-api-passport /

ENV K8S_SERVER_CONFIG_ADDR=$CONFIG_HOST
ENV K8S_SERVER_CONFIG_PATH=conf/platform/user/api/passport

RUN chmod +x /user-api-passport