#!/bin/bash

#URL=http://localhost:5000/consumption
URL=http://localhost:5000/task/task_2cbbd5a3-1383-4b69-b109-b4bf5c2f7bb6
#REQUEST=main.CredentialsProto
#RESPONSE=main.ErrorProto
REQUEST=$1
RESPONSE=$2
PROTO=./mensaje.proto

if [ "$RESPONSE" == "" ]
then
    cat proto.msg | protoc --encode $REQUEST $PROTO | curl -X POST --data-binary @- $URL
else
    cat proto.msg | protoc --encode $REQUEST $PROTO | curl -sS -X POST --data-binary @- $URL | protoc --decode $RESPONSE $PROTO
fi