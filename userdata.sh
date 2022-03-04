#!/bin/bash
yum update -y
yum install ruby -y
wget https://aws-codedeploy-ap-northeast-2.s3.ap-northeast-2.amazonaws.com/latest/install
chmod +x /install
/install auto

yum install go -y
