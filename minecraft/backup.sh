#!/bin/bash
DATE=`date +%Y%m%d`
FILE_COUNT=`ls /mnt/minecraft-backups -l | wc -l`

cd /home/pmh/composes/minecraft

docker compose stop

tar czf \
  /mnt/minecraft-backups/$DATE.tar.gz \
  ./data ./volumes/lobby/world \
  ./volumes/creative/void_1 ./volumes/creative/world_old \
  ./volumes/survival/world ./volumes/survival/world_nether ./volumes/survival/world_the_end

echo $DATE > /mnt/minecraft-backups/LAST_DATE

docker compose up -d

if [ $FILE_COUNT -gt 7 ]; then
  rm /mnt/minecraft-backups/$(ls -rt | head -n 1)
fi
