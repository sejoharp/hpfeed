#!/usr/bin/env bash
set -eu

project_dir="/Users/joscha/Documents/workspace/gocode/src/bitbucket.org/joscha/hpfeed"
destination_ip="10.1.1.3"
copy_dir="/home/pi"
command_dir="/home/hpnews"
command_user="hpnews"
command_group="hpnews"
remote_user="pi"
binary_name="hpfeed"

echo "--> changing to project dir ${project_dir}..."
cd ${project_dir}

echo "--> compiling binary for ${binary_name}..."
GOPATH=/Users/joscha/Documents/workspace/gocode/
env \
  GOOS=linux \
  GOARCH=arm \
  CGO_ENABLED=0 \
  GOARM=5 \
  go build

echo "--> uploading via scp ..."
scp ${binary_name} ${remote_user}@${destination_ip}:${copy_dir}

echo "--> stopping ${binary_name} ..."
ssh ${remote_user}@${destination_ip} sudo "/etc/init.d/hpfeed.sh stop"

echo "--> delete old version ${binary_name} ..."
ssh ${remote_user}@${destination_ip} sudo "rm ${command_dir}/${binary_name}"

echo "--> copying new version to ${command_dir}  ..."
ssh ${remote_user}@${destination_ip} sudo "mv ${copy_dir}/${binary_name} ${command_dir}/"

echo "--> setting permissions  ..."
ssh ${remote_user}@${destination_ip} sudo "chown ${command_user}:${command_group} ${command_dir}/${binary_name}"

echo "--> Starting ${binary_name} ..."
ssh ${remote_user}@${destination_ip} sudo "/etc/init.d/hpfeed.sh start"
