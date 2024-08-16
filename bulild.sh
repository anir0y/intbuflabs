#!/bin/bash
echo "Running build script"
echo "[+] Setting up environment"
export GOOS=windows
export GOARCH=amd64
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export CGO_ENABLED=1
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
echo "[+] Building ArishtiOverflow.exe"
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-H windowsgui" -o ArishtiOverflow.exe
