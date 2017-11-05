#!/bin/bash

ROOT=/home/bketelsen/src/github.com/thewondertwins/gobook
TARGET=/opt/caddy/_book

cd $ROOT
git pull
gitbook build
cp -R /home/bketelsen/src/github.com/thewondertwins/gobook /opt/caddy/_book
chown -R www-data:www-data /opt/caddy/_book
