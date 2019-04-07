# vcheck
[![](https://img.shields.io/docker/pulls/anjmao/vcheck.svg)](https://hub.docker.com/r/anjmao/vcheck)

Check your deployed GRPC, HTTP service version to verify if deployment was successful. This is especially well suited for travis, gitlab CI/CD deployments.

## GRPC service example

GRPC service should implement endpoint which can return buildVersion number.

```proto
syntax = "proto3";
package version;

service Version {
    rpc GetVersion (GetVersionRequest) returns (GetVersionReply) {}
}

message GetVersionRequest {}

message GetVersionReply {
    string buildVersion = 1;
}
```

```bash
MY_CI_BUILD_NUMBER=8326
vcheck \
    --target=staging.my.grpcapi.net:443 \
    --method=/version.Version/GetVersion \
    --client=grpc \
    --expect=${MY_CI_BUILD_NUMBER} \
    --count=12 \
    --sleep=5
```

## HTTP service example

HTTP service should implement GET endpoint returning JSON with buildVersion field. 

```bash
MY_CI_BUILD_NUMBER=8326
vcheck \
    --target=http://staging.my.grpcapi.net \
    --method=/api/version \
    --client=http \
    --expect=${MY_CI_BUILD_NUMBER} \
    --count=12 \
    --sleep=5
```

### Docker

vheck util is also available as a docker image

```bash
docker run --rm anjmao:vcheck --help
```

### Release

Before release install gox and github-release tools
```bash
go GO111MODULE=off go get github.com/mitchellh/gox
go GO111MODULE=off go get github.com/c4milo/github-release 
```

```bash
make release
```