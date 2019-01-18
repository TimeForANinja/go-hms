# go-hms
a simple, standalon home media server written in go(lang)

# installation (linux)
## install dependencies
```
apt install libssl-dev build-essential curl git
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

# install (windows)
(not yet working... sqlite3 build fails)
## install go
download and execute the golang.msi [from there website](https://golang.org/dl/)
## install git
download and execute the git.exe [from there website](https://git-scm.com/download/win)
## install gcc & openssl
download and execute the cygwin.exe [from there website](https://cygwin.com/install.html)
when asked for packages install all from the `devel` categorie

# install required package
```
go get github.com/xeodou/go-sqlcipher
go get golang.org/x/crypto/bcrypt
```

# for cross compilation to windows from linux
```
apt install mingw-w64
git clone https://github.com/openssl/openssl
cd openssl
./Configure mingw64 shared --cross-compile-prefix=x86_64-w64-mingw32- --prefix=/usr/x86_64-w64-mingw32
make
sudo make install
cd ..
rm -rf openssl
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o ~/go/src/github.com/timeforaninja/go-hms/build/main.exe main.go
```

# link repos into $GOPATH
## windows (admin cmd)
```
mkdir %GOPATH%\src\github.com\timeforaninja
mklink /D %GOPATH%\src\github.com\timeforaninja\go-hms C:\Users\<user>\Documents\GitHub\go-hms
```
## linux
```
mkdir -p ~/go/src/github.com/timeforaninja
ln -s /mnt/hgfs/GitHub/go-hms ~/go/src/github.com/timeforaninja/go-hms
```
