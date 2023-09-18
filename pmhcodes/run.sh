#!/bin/bash
set -euo pipefail

container=""

function cleanup() {
  docker rm -f "$container" > /dev/null
}

trap cleanup EXIT

container=$(docker run -h "pmh.codes" -v /home/pmh/composes/pmhcodes/home:/home:ro --network none -dm 25m --memory-swap 25m --cpus="0.25" pmhcodes:latest sleep 600)

if [ -z "${container}" ]; then
	echo "Failed to obtain container ID"
	exit 1
fi

docker exec -it "${container}" /home/guest/start
