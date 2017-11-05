#!/bin/bash

ROOT=/home/bketelsen/src/github.com/thewondertwins/gobook
TARGET=/opt/caddy/_book

cd $ROOT
git pull
gitbook build
sudo cp -R /home/bketelsen/src/github.com/thewondertwins/learngo /opt/caddy/_book
sudo chown -R www-data:www-data /opt/caddy/_book
