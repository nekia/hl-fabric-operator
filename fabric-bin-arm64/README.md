# How to build

* Add multiarch support to the environment by using QEMU and binfmt. Run following command:

  ```
  docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
  ```
  * https://hub.docker.com/r/multiarch/qemu-user-static/

* Start baseos container for ARM64 arch

  ```
  docker run -d --name build -it busan15/fabric-baseimage
  ```

* On the `build` container, run the following command sequence:

  ```
  docker exec -it build bash

  # Start shell on `build` container

  cd /opt/gopath/
  mkdir -p src/github.com/hyperledger/
  git clone https://github.com/hyperledger/fabric.git
  cd fabric
  git checkout v1.4.9
  make configtxgen
  make configtxlator
  make cryptogen
  exit
  ```
  
  * If you have an error as below, then you need to create dummy docker executable file with following command:

    ```
    Makefile:89: *** "No docker in PATH: Check dependencies".  Stop.
    ```
  
    You also need to create dummy exectable file for `docker` comman

    ```
    cp /usr/bin/git /usr/bin/docker
    ```
  
* Copy from `build` container to local file system

  ```
  docker cp build:/opt/gopath/src/github.com/hyperledger/fabric/.build/bin/cryptogen .
  docker cp build:/opt/gopath/src/github.com/hyperledger/fabric/.build/bin/configtxgen .
  docker cp build:/opt/gopath/src/github.com/hyperledger/fabric/.build/bin/configtxlator .
  ```

