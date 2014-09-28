ThingFu Hub [![Travis](https://travis-ci.org/ThingFu/hub.svg?branch=master)](https://travis-ci.org/ThingFu/hub)
===========

## Goals
The goal of the project is to provide a platform for interacting with physical devices in your home. Unlike many other platforms, first and foremost, everything should run locally and autonomously on the  Hub without the need for a cloud solution.

## Development

### Go development environment

**go-home** is written in [Go](http://golang.org) programming language. If you haven't set up Go development environment, please follow [this instruction](http://golang.org/doc/code.html) to install go tool and set up GOPATH. Ensure your version of Go is at least 1.3.

### Put go-home into GOPATH

We highly recommend to put **go-home** code into your GOPATH. For example, the following commands will download go-home code under the current user's GOPATH (Assuming there's only one directory in GOPATH.):

```
$ echo $GOPATH
/home/user/goproj
$ mkdir -p $GOPATH/src/github.com/go-home/hub/
$ cd $GOPATH/src/github.com/go-home/hub/
$ git clone https://github.com/zubairhamed/gohome-hub.git
```

The commands above will not work if there are more than one directory in ``$GOPATH``.
