#!/bin/bash
cd /home/pmh/composes

for d in */ ; do
    cd $d

    docker compose pull
    docker compose up -d

    cd ..
done
