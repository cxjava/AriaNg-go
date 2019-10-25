<div align="center">
  <a href="https://travis-ci.org/cxjava/AriaNg-go">
    <img src="https://travis-ci.org/cxjava/AriaNg-go.svg?branch=master" alt="Build Status">
  </a>
  
  [![CircleCI](https://circleci.com/gh/cxjava/AriaNg-go.svg?style=svg)](https://circleci.com/gh/cxjava/AriaNg-go)
</div>

# AriaNg-go
AriaNg in go http server

# Version 
```
Version: 	0.0.4
Build Time: 	2019-10-25T01:48:04Z
Go Version: 	go version go1.13.3 linux/amd64
Repo URL: 	git@github.com:cxjava/AriaNg-go.git
Commit Info: 	7a851ac54e0ca33c9a35b4fa288d7e54741317a6
Build By: 	circleci@210396888f1b
```


# Help

```
$ ./AriaNg-go -h
NAME:
   AriaNg-go - AriaNg in go http server

USAGE:
   AriaNg-go [global options] command [command options] [arguments...]

VERSION:
   0.0.4

COMMANDS:
   server, s  Start server for AriaNg
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --address value, -a value  bind address (default: ":18080")
   --help, -h                 show help
   --version, -v              print the version

```

visit `http://localhost:18080`

change the defalt bind address:
```
$ ./AriaNg-go --address=":11180" server 
```
