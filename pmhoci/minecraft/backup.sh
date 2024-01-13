#!/bin/bash
DATE=`date +%Y%m%d`
FILE_COUNT=`ls /home/pmh/minecraft-backups -l | wc -l`

cd /home/pmh/composes/$1/minecraft

docker compose stop

tar czf \
  /home/pmh/minecraft-backups/$DATE.tar.gz \
  ./data ./volumes/lobby/world \
  ./volumes/creative/void_1 ./volumes/creative/world_old \
  ./volumes/survival/world ./volumes/survival/world_nether ./volumes/survival/world_the_end

echo $DATE > /home/pmh/minecraft-backups/LAST_DATE

docker compose up -d

if [ $FILE_COUNT -gt 9 ]; then
  rm /home/pmh/minecraft-backups/$(ls -rt /home/pmh/minecraft-backups | head -n 1)
fi
