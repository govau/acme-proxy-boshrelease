#!/bin/bash

set -euo pipefail
set -v

mkdir release-candidate

tar xvfz ../cert.s3/cert.tar.gz -C release-candidate

cp -rf ../ops.git/todo release-candidate

# Now the release-candidate dir contains an extracted haproxy config we can
# validate

#TODO
# /usr/local/sbin/haproxy -c -V -f release-candidate/haproxy.cfg

# Create a release tarball
tar cvfz release-candidate/* release/release.tar.gz
