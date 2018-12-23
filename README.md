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
nano ~/.profile
```
> prepend
> ```
> export GOPATH=$HOME/work
> export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
> ```
## install required package
```
go get github.com/xeodou/go-sqlcipher
```
