#!/bin/bash

ROOT=/home/bketelsen/src/github.com/thewondertwins/learngo
TARGET=/opt/caddy/_book

cd $ROOT
git pull
gitbook build
sudo cp -R $ROOT $TARGET
sudo chown -R www-data:www-data $TARGET
