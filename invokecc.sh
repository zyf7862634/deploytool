#!/usr/bin/env bash

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! FAIL !!!!!!!!!!!!!!!!"
    exit 1
  fi
}
echo "-------test chaincode (测试chaincode)-------"
./deployFabricTool -r testcc -n mychannel -func invoke
verifyResult $?

