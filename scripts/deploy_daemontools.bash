#!/usr/bin/env bash
set -eu

project_dir="/Users/joscha/Documents/workspace/gocode/src/bitbucket.org/joscha/hpfeed"
destination="10.1.1.3"
remote_user="joscha"
command_dir="/home/${remote_user}/opt/hpfeed/"
binary_name="hpfeed"
binary_name_temp="${binary_name}_new"

echo "--> changing to project dir ${project_dir}..."
cd ${project_dir}

echo "--> compiling binary for ${binary_name}..."
GOPATH=/Users/joscha/Documents/workspace/gocode/

# build command for linux with 64bit
GOOS=linux GOARCH=amd64 go build

echo "--> renaming to temporary name ..."
mv ${binary_name} ${binary_name_temp}

echo "--> uploading via scp ..."
scp ${binary_name_temp} ${remote_user}@${destination}:${command_dir}

echo "--> stopping ${binary_name} ..."
ssh ${remote_user}@${destination} "svc -d /home/${remote_user}/service/${binary_name}"

echo "--> delete old version ${binary_name} ..."
ssh ${remote_user}@${destination} "rm ${command_dir}/${binary_name}"

echo "--> renaming new version to ${binary_name}  ..."
ssh ${remote_user}@${destination} "mv ${command_dir}/${binary_name_temp} ${command_dir}/${binary_name}"

echo "--> changing file permission ${binary_name}  ..."
ssh ${remote_user}@${destination} "chmod 700 ${command_dir}/${binary_name}"

echo "--> Starting ${binary_name} ..."
ssh ${remote_user}@${destination} "svc -u /home/${remote_user}/service/${binary_name}"

echo "--> cleaning local temp binary ..."
rm ${binary_name_temp}
