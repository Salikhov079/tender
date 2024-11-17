#!/bin/bash
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/internal/pkg/genproto
for x in $(find ${CURRENT_DIR}/protos* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/go --go_out=${CURRENT_DIR}/internal/pkg \
   --go-grpc_out=${CURRENT_DIR}/internal/pkg ${x}/*.proto
done
