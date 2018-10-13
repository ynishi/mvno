# mvno [![Build Status](https://travis-ci.org/ynishi/mvno.svg?branch=master)](https://travis-ci.org/ynishi/mvno)
* mv auto numbered name
* https://hub.docker.com/r/ynishi/mvno/

## build
### go
```
go get github.com/ynishi/mvno
mvno
```
### docker
```
git clone https://github.com/ynishi/mvno.git 
sh build.sh
```
## usage
```
./mvno prefix fileA.txt fileB.txt ... target/dir
ls $target/dir
# moved
prefix0.txt
prefix1.txt
...
```
