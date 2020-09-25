FROM yeasy/hyperledger-fabric:1.1.0
RUN sed -i 's#http://archive.ubuntu.com/#http://mirrors.tuna.tsinghua.edu.cn/#' /etc/apt/sources.list;
RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
# Install 
RUN apt-get update && apt-get install sudo && sudo apt-get install unzip && sudo apt-get install zip
RUN sudo apt-get install git
#go
RUN cd /usr/local && sudo mv go go1.9.4
COPY go /usr/local/go
RUN echo export PATH=$PATH:/usr/local/go/bin >> /etc/profile && echo export GOPATH=/go >> /etc/profile && echo export PATH=$PATH:/go/bin >> /etc/profile
WORKDIR $GOPATH
ENV GOPROXY=https://goproxy.cn
#https://mirrors.aliyun.com/goproxy/
ENV GO111MODULE=on
RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN mkdir $GOPATH/src/example && mkdir $GOPATH/src/example/server && mkdir $GOPATH/src/github.com/dvyukov && mkdir /go/src/example/server/result

COPY dvyukov $GOPATH/src/github.com/dvyukov
COPY fabric2.0 $GOPATH/src/github.com/hyperledger/fabric2.0
RUN cd /go/src/github.com/hyperledger/fabric2.0/integration/chaincode/aab && go mod init
COPY bin $GOPATH/bin

COPY conf $GOPATH/src/example/server/conf
COPY controllers $GOPATH/src/example/server/controllers
COPY models $GOPATH/src/example/server/models
COPY routers $GOPATH/src/example/server/routers
COPY main.go $GOPATH/src/example/server

WORKDIR $GOPATH/src/example/server
RUN set GO111MODULE=on && go mod init

# Expose the application on port 8080*
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the*
# application and watches for changes*
CMD ["bee", "run"]


