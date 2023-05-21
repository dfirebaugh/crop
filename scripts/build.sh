#!/bin/bash

mkdir -p .dist/x86_64_unknown-linux/
mkdir -p .dist/x86_64_pc-windows/
mkdir -p .dist/x86_64_apple-darwin/

go build -o .dist/x86_64_unknown-linux/crop
GOOS=windows GOARCH=amd64 go build -o .dist/x86_64_pc-windows/crop.exe
GOOS=darwin GOARCH=amd64 go build -o .dist/x86_64_apple-darwin/crop

cd .dist/x86_64_unknown-linux/
tar -czvf crop-x86_64_unknown-linux.tar.gz crop
cd ../x86_64_pc-windows
tar -czvf crop-x86_64_pc-windows.tar.gz crop.exe
cd ../x86_64_apple-darwin
tar -czvf crop-x86_64_apple-darwin.tar.gz crop
