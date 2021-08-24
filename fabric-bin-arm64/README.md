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
  cd src/github.com/hyperledger/
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
  docker cp build:/opt/gopath/src/github.com/hyperledger/fabric/.build/bin/cryptogen .build/docker/bin
  docker cp build:/opt/gopath/src/github.com/hyperledger/fabric/.build/bin/configtxgen .build/docker/bin
  docker cp build:/opt/gopath/src/github.com/hyperledger/fabric/.build/bin/configtxlator .build/docker/bin
  ```

```
docker build -t nekia/fabric-peer -f images/peer/Dockerfile.in .
docker tag nekia/fabric-peer:latest nekia/fabric-peer:arm64-1.4.9
docker tag nekia/fabric-peer:latest nekia/fabric-peer:arm64-latest
docker tag nekia/fabric-peer:latest nekia/fabric-peer:1.4.9
docker tag nekia/fabric-peer:latest nekia/fabric-peer:1.4
docker images --filter reference=nekia/*:*
docker tag nekia/fabric-tools:latest nekia/fabric-tools:arm64-1.4.9
docker tag nekia/fabric-tools:latest nekia/fabric-tools:arm64-latest
docker tag nekia/fabric-tools:latest nekia/fabric-tools:1.4.9
docker tag nekia/fabric-tools:latest nekia/fabric-tools:1.4
docker images --filter reference=nekia/*:*
docker tag nekia/fabric-orderer:latest nekia/fabric-orderer:arm64-1.4.9
docker tag nekia/fabric-orderer:latest nekia/fabric-orderer:arm64-latest
docker tag nekia/fabric-orderer:latest nekia/fabric-orderer:1.4.9
docker tag nekia/fabric-orderer:latest nekia/fabric-orderer:1.4
docker images --filter reference=nekia/*:*
```

```
wget https://patch-diff.githubusercontent.com/raw/hyperledger/fabric/pull/345.patch
git checkout v2.0.0
patch -p1 < 345.patch
make peer
```

```
make .build/image/ccenv/payload/chaintool
make .build/goshim.tar.bz2
make .build/image/ccenv/payload
make .build/docker/gotools
```

```
k get rs | grep hlf | awk '{print $1}' | xargs -I{} kubectl delete rs {}
https://github.com/vektra/mockery/issues/364
https://jira.hyperledger.org/browse/FAB-18346?attachmentViewMode=list
make ccenv

ssh kube-worker0 "sudo crictl images --digests"
ssh kube-worker0 "sudo crictl rmi nekia/rfabric:v0.2.0"
ssh kube-worker0 "sudo crictl rmi nekia/hl-fabric-tools:1.4.3"

for i in $(seq 0 2); do ssh kube-worker$i "sudo shutdown -r now"; done
for i in $(seq 0 2); do ssh kube-worker$i "sudo shutdown -h now"; done

ssh kube-worker0 "sudo crictl rmi nekia/rfabric:v0.2.0"
docker build -t nekia/rfabric:v0.2.0 -f ./Dockerfile .
docker tag nekia/rfabric:v0.2.0 nekia/rfabric:latest
docker run --rm -it --entrypoint=/bin/sh nekia/rfabric:v0.2.0
docker push nekia/rfabric:v0.2.0
docker push nekia/rfabric:latest
k delete -f install.yaml
```

```
https://lists.hyperledger.org/g/fabric/topic/79950272?p=,,,20,0,0,0::,,,0,0,0,79950272
https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
```