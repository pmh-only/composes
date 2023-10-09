#!/bin/bash

for d in */ ; do
    cd $d

    docker compose pull
    docker compose up -d

    cd ..
done
