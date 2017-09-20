# Brocade Cert Manager
Tool to Show, Add and Delete SSL certificates on Brocade VTMs

##### Setup - 
This project uses a dependency manager `dep` - https://github.com/golang/dep 
To install `dep` on MacOS:
```sh
$ brew install dep
$ brew upgrade dep
```
##### Build -
Must use `dep` to install dependencies
```sh
$ dep ensure
$ go build cert-manager.go
```

### Usage: 

```sh
$ ./cert-manager -h
Usage of ./cert-manager:
  -api string
        Brocade VTM API version number (default "3.8")
  -cert string
        Path to the certificate file
  -key string
        Path to the private key
  -loadbalancer string
        Enter the URL for the Load Balancer
  -name string
        The name of the certificate
  -option string
        Command option. 'show' 'add' 'delete'
```
Examples:
```sh
$ ./cert-manager -option=show -loadbalancer=example-vtm-node-01.com -name=example.cert.com
```
```sh
$ ./cert-manager -option=add -loadbalancer=example-vtm-node-01.com -name=example.cert.com -cert=/path/to/cert -key=/path/to/key
```
```sh
$ ./cert-manager -option=delete -loadbalancer=example-vtm-node-01.com -name=example.cert.com
```