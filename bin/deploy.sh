#!/bin/bash
if [ $# -ne 1 ]; then
  echo "引数にはデプロイ対象の環境(developまたはmainまたはdemo)を指定してください" 1>&2
  exit 1
fi

environment=$1

if [ $environment = "develop" -o $environment = "main" -o $environment = "demo" ]; then
  echo "ビルド開始"
  sh ./bin/build.sh
  echo "ビルド終了"
  echo $environment"にデプロイします"
  sls deploy --stage $environment
else
  echo "引数にはデプロイ対象の環境(developまたはmainまたはdemo)を指定してください"
fi
