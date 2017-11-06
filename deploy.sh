#!/bin/bash

ROOT=/home/bketelsen/src/github.com/thewondertwins/learngo
TARGET=/opt/caddy/

cd $ROOT
git pull
gitbook build
sudo cp -R $ROOT/_book $TARGET
sudo chown -R www-data:www-data $TARGET
