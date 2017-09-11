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
$ ./cert-manager -option=show -loadbalancer=h1sta01-v00.devops.int.ovp.bskyb.com -name=i2.np.ovp.sky.com
```
```sh
$ ./cert-manager -option=add -loadbalancer=h1sta01-v00.devops.stg2.ovp.bskyb.com -name=ovp.bskyb.com -cert=/path/to/cert -key=/path/to/key
```
```sh
$ ./cert-manager -option=delete -loadbalancer=h1sta01-v00.devops.int.ovp.bskyb.com -name=ben.test.com
```