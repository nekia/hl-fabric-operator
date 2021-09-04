# download HL Fabric binaries and Helm
FROM nekia/fabric-tools:1.4.9 as curl

WORKDIR /helm
RUN curl https://get.helm.sh/helm-v3.5.2-linux-arm64.tar.gz --output helm.tar.gz \
    && tar xf helm.tar.gz

# clone PIVT repository
FROM alpine/git@sha256:94a81d66655d75597155e0afc3629dad909ba08c5f97d24238ff6d69a0af91c0 as git

WORKDIR /workspace
RUN git clone https://github.com/nekia/PIVT.git \
    && cd PIVT \
    && git checkout 833647b177e36f830f656d13004f6d4bd06f15d3

# Install hlf-kube Helm chart dependencies (Kafka)
COPY --from=curl /helm/linux-arm64/helm /usr/local/bin/
RUN cd /workspace/PIVT/fabric-kube/ \
    && helm dependency update ./hlf-kube/

# Build the manager binary
FROM arm64v8/golang@sha256:7c5a0c353de4cf438ab225c0f727731463f28cdec083a72a5f66f46d7c25c577 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -a -o manager main.go

# Actual runtime image
FROM arm64v8/alpine@sha256:bd9137c3bb45dbc40cde0f0e19a8b9064c2bc485466221f5e95eb72b0d0cf82e

WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=git /workspace/PIVT /opt/fabric-operator/PIVT/
COPY --from=curl /usr/local/bin/configtxgen /usr/local/bin/cryptogen /usr/local/bin/configtxlator /opt/hlf/

ENV PATH "$PATH:/opt/hlf"

# one way to run Fabric binaries in Alpine container
# see https://stackoverflow.com/a/59367690/3134813
RUN apk add --no-cache libc6-compat

RUN mkdir -p /var/fabric-operator \
    && chmod 777 /var/fabric-operator
# USER 65532:65532
# USER root

ENTRYPOINT ["/manager"]
