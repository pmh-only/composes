#!/bin/bash
cd /home/pmh/composes/$1

for d in */ ; do
    cd $d

    docker compose pull
    docker compose up -d

    cd ..
done
