# This plugin modify from (https://github.com/nacos-group/nacos-coredns-plugin) 
* Modify AllDomNames struct support for newer nacos api
* Use forward plugin replace proxy plugin, because the new version of coredns has eliminated proxy
* Modify file struct support go mod （go get github.com/nacos-group/nacos-coredns-plugin）

# How to use it

* Get coredns code
  ```
  $ git clone https://github.com/coredns/coredns
  $ cd coredns
  ```
* Edit plugin.cfg
Insert “nacos:github.com/BailinSong/nacos-coredns-plugin” before "etcd"
```
......
auto:auto
secondary:secondary
nacos:github.com/BailinSong/nacos-coredns-plugin
etcd:etcd
loop:loop
forward:forward
......
```
* To compile CoreDNS
  * function 1 use the "make" command.
    ```
    $ make
    ```
  * function 2 use the docker
    ```
    $ docker run --rm -i -t -v $PWD:/v -w /v golang:1.12 make
    ```
