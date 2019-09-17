#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

SERVER_IP=0.0.0.0
SERVER_PORT=8080

go run . \
    --log_level=DEBUG \
    --register_interval=5 \
    --register_ttl=10 \
    --client_pool_size=100 \
    --server_name=com.jinmuhealth.platform.api \
    --metadata=X-Err-Style=MicroDetailed \
    --config_consul_address=localhost:8500 \
    api \
    --address=${SERVER_IP}:${SERVER_PORT} \
    --handler=rpc \
    --enable_rpc \
    --namespace=com.jinmuhealth.platform.srv \
    --enable_jwt \
    --jwt_max_exp_interval=600s \
    --rpc_whitelist=com.jinmuhealth.platform.srv.template-service
