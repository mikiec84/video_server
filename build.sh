#! /bin/bash

# Build web UI

export VIDEO_SERVER_PATH=$GOPATH/src/go_dev/src/video_server

rm -rf $VIDEO_SERVER_PATH/bin/*

cd $VIDEO_SERVER_PATH

cd $VIDEO_SERVER_PATH/api && go build && cp api $VIDEO_SERVER_PATH/bin

cd $VIDEO_SERVER_PATH/scheduler && go build && cp scheduler $VIDEO_SERVER_PATH/bin

cd $VIDEO_SERVER_PATH/streamserver && go build && cp streamserver $VIDEO_SERVER_PATH/bin

cd $VIDEO_SERVER_PATH/web && go build && cp web $VIDEO_SERVER_PATH/bin

cd $VIDEO_SERVER_PATH

cp -rf $VIDEO_SERVER_PATH/static/template $VIDEO_SERVER_PATH/bin/

