FROM golang:1.11

# install protobuf from source
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool
RUN git clone -b 'v3.5.1' --single-branch --depth 1 https://github.com/google/protobuf.git && \
    cd protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    ldconfig && \
    make clean && \
    cd .. && \
    rm -r protobuf

# Get the source from GitHub
RUN go get google.golang.org/grpc
# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN protoc --version
ENTRYPOINT [ "/usr/local/bin/protoc", "--go_out=plugins=grpc:.", "protocol.proto" ]
