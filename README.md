# installation
## install dependencies
```
apt install libssl-dev build-essential curl
```
## get go
```
curl -O https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
```
## add go to path
```
nano ~/.profile (user only) || nano /etc/profile (global)
```
> prepend
> ```
> export GOPATH=$HOME/work
> export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
> ```
## install required package
```
go get github.com/xeodou/go-sqlcipher
go get golang.org/x/crypto/bcrypt
```

# for cross compilation to windows
```
apt install mingw-w64
git clone https://github.com/openssl/openssl
cd openssl
./Configure mingw64 shared --cross-compile-prefix=x86_64-w64-mingw32- --prefix=/usr/x86_64-w64-mingw32
make
sudo make install
cd ..
rm -rf openssl
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o build/main.exe main.go
```
