FROM golang:1.9.4

RUN mkdir $GOPATH/src/github.com && \
    mkdir $GOPATH/src/github.com/Multy-io 

RUN cd $GOPATH/src/github.com/Multy-io && \ 
    git clone https://github.com/Multy-io/Multy-back.git && \ 
    cd Multy-back && \ 
    git checkout import-eth-multisig 


RUN go get -u github.com/golang/protobuf/proto && \
    cd $GOPATH/src/github.com/golang/protobuf && \
    make all

RUN apt-get update && \
    apt-get install -y protobuf-compiler

RUN cd $GOPATH/src/github.com/Multy-io && \
    git clone https://github.com/Multy-io/Multy-ETH-node-service.git && \
    cd $GOPATH/src/github.com/Multy-io/Multy-ETH-node-service && \
    git checkout import-eth 

# go get github.com/ethereum/go-ethereum/rpc

# RUN cd $GOPATH/src/github.com/Multy-io && \
#     go get ./...

RUN cd $GOPATH/src/github.com/Multy-io/Multy-ETH-node-service && \
    go get ./... && \
    # make proto && \
    make build && \
    rm -r $GOPATH/src/github.com/Multy-io/Multy-back 


WORKDIR $GOPATH/src/github.com/Multy-io/Multy-ETH-node-service/cmd

RUN echo "VERSION 02"

ENTRYPOINT $GOPATH/src/github.com/Multy-io/Multy-ETH-node-service/cmd/client
