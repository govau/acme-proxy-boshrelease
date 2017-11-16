#!/bin/bash

set -e
set -x

TARGET=${TARGET:-local}

fly validate-pipeline --config pipeline.yml

fly --target ${TARGET} set-pipeline --config pipeline.yml --pipeline acme-proxy -n -l credentials.yml

for resource in ops.git acme-proxy-boshrelease.git certs-d.s3; do
  fly -t ${TARGET} check-resource --resource acme-proxy/$resource
done

fly -t ${TARGET} unpause-pipeline -p acme-proxy
