#!/bin/bash
DATE=`date +%Y%m%d`
FILE_COUNT=`ls /mnt/minecraft-backups -l | wc -l`

cd /home/pmh/composes/$1/minecraft

docker compose stop

tar czf \
  /mnt/minecraft-backups/$DATE.tar.gz \
  ./data ./volumes/lobby/world \
  ./volumes/creative/void_1 ./volumes/creative/world_old \
  ./volumes/survival/world ./volumes/survival/world_nether ./volumes/survival/world_the_end

echo $DATE > /mnt/minecraft-backups/LAST_DATE

docker compose up -d

if [ $FILE_COUNT -gt 9 ]; then
  rm /mnt/minecraft-backups/$(ls -rt /mnt/minecraft-backups | head -n 1)
fi
