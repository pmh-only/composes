#!/bin/bash
echo "#####################"
echo "##  S A N D B O X  ##"
echo "#####################"
echo ""
echo "Welcome to sandbox.shutupandtakemy.codes."
echo "Abuse Report: abuse@pmh.codes"
echo "Etc. Contact: pmh_only@pmh.codes"
echo ""
echo "WARNING: Disconnecting & OOM will erase all data."
echo ""
echo "Loading node.js REPL..."
echo ""

set -euo pipefail

container=""

function cleanup() {
  docker rm -f "$container" > /dev/null
}

trap cleanup EXIT

container=$(docker run -h "shutupandtakemycodes" --network none -dm 25m --memory-swap 25m --cpus="0.25" sandbox:latest sleep 600)

if [ -z "${container}" ]; then
	echo "Failed to obtain container ID"
	exit 1
fi

docker exec -it "${container}" node
