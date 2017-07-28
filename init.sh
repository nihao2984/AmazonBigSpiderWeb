#!/bin/bash

if [ ! -d "$GOPATH" ]; then
	mkdir -p "$GOPATH"
fi
export GOPATH=$GOPATH
echo "画室网站开始，皮卡丘～～"
echo "GOPATH地址:[$GOPATH]"

## 导入外部依赖包, 使用 go list -json 可以查看项目依赖包
declare -a pkgs=(
	github.com/astaxie/beego
	github.com/beego/bee
	github.com/nfnt/resize
	github.com/go-sql-driver/mysql
	golang.org/x/net/html
	golang.org/x/text/encoding
	golang.org/x/text/transforms
	github.com/mattn/go-sqlite3
)

echo "======== 下载外部包 开始========"
for path in "${pkgs[@]}"; do
	echo "go get ${path} ..."
	go get $path
done

echo "======== 就这么简单 ========"
