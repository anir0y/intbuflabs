# ArishtiOverflow

**ArishtiOverflow** is a Go-based application built with the Fyne UI framework. It demonstrates a simple buffer overflow concept in a safe environment by validating phone numbers. When a user enters more than 10 characters, a secret code is revealed.

## Prerequisites

Before you can build and sign the application, ensure that you have the following installed on your Ubuntu system:

1. **Go** (Install from [Go's official website](https://golang.org/dl/))
2. **MinGW** (Required for cross-compiling to Windows)
3. **OpenSSL** (For generating self-signed certificates)
4. **osslsigncode** (For signing Windows executables)

Install MinGW, OpenSSL, and osslsigncode using the following command:

```
sudo apt update
sudo apt install gcc-mingw-w64-x86-64 
```

## Installing Required Go Libraries
To build this project, you'll need to install the required Go libraries. You can do this by running:

```
go mod tidy
```

## Building the Application for Windows

To build the application for Windows on Ubuntu, follow these steps:

1. **Set up Environment Variables for Cross-Compilation:**

   Before building, ensure the correct toolchain is used for cross-compiling to Windows:

   ```bash
   export CC=x86_64-w64-mingw32-gcc
   export CXX=x86_64-w64-mingw32-g++
   ```

   Or use the file `exportVar.sh` to export all vars automatically. 

2. **Build the Application:**

   Now, build the application for Windows 64-bit architecture with the following command:

   ```bash
   GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-H windowsgui" -o ArishtiOverflow.exe
   ```

   This will produce the `ArishtiOverflow.exe` executable file for Windows.





