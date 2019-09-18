FROM centos-builder AS builder
WORKDIR /root/go/src
COPY . grpc-gateway
RUN export GOPATH=/root/go/src/grpc-gateway:/root/go \
        && PATH=/root/go/bin:/root/go-1.12/bin:$PATH \
        && cd grpc-gateway/src/protos \
        && make \
        && cd .. \
        && go build main.go

FROM centos
WORKDIR /root
COPY src/build-docker .
COPY --from=builder /root/go/src/grpc-gateway/src/main ./main