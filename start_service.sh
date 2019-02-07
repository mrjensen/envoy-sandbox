#!/bin/sh
# This file is used by service 1 and service 2 to start the service and start an envoy proxy
# The output of the service is logged through stdout to /log without any log rotation (don't do this in any environment you care about)
./service > /log &
envoy -c /etc/service-envoy.yaml --service-cluster service${SERVICE_NAME}