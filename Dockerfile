FROM golang:1.15 as builder
WORKDIR /code
COPY . .
RUN CGO_ENABLED=0 go build -o shadowsky-qiandao

FROM alpine:3.12
COPY --from=builder /code/shadowsky-qiandao /usr/bin/shadowsky-qiandao
ENTRYPOINT [ "/usr/bin/shadowsky-qiandao" ]
CMD [ "--help" ]