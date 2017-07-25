#!/usr/bin/env bash
#go get github.com/golang/protobuf/protoc-gen-go
export GOPATH=/home/deneva/GoglandProjects/prueba/

if [ ! -d go ]
then
   mkdir -p go
fi

protoc --go_out=go *.proto
if [ $? == "0" ]
then
    cp -f go/api.pb.go  ../src/github.com/maxpowel/dislet/apirest/protomodel
    cp -f go/machinery.pb.go  ../src/github.com/maxpowel/dislet/machinery/protomodel
    cp -f go/wiphone.pb.go  ../src/github.com/maxpowel/wiphonego/protomodel
    echo "Copiados"
fi