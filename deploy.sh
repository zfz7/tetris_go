#!/usr/bin/env bash

set -e

git update-index --refresh
git diff-index --quiet HEAD --

sshHost="-p 22 -C $TETRIS_SSH_USER@$TETRIS_URL"
scpHost="$TETRIS_SSH_USER@$TETRIS_URL"
##Build App
if [ "$1" == 'NOTEST' ]; then
	./build.sh prod
else
	./test.sh && ./build.sh prod
fi

#Update system
ssh $sshHost "sudo apt update"
ssh $sshHost "sudo apt upgrade -y"
ssh $sshHost "sudo apt autoremove -y"
ssh $sshHost "sudo shutdown -r now"  || echo "System shutdown"


#Wait for host to come back online
until [ "$(ssh $sshHost "echo ok")" = "ok" ]; do
  sleep 1;
  echo "Trying to connect to host again..."
done

ssh $sshHost "mkdir -p ~/app"

ssh $sshHost "sudo rm ~/app/tetris" || echo "No exec file found"
scp -P "22" ./docker-compose.prod.yml $scpHost:~/app/
ssh $sshHost "POSTGRES_PASSWORD=${POSTGRES_PASSWORD} docker compose -f ~/app/docker-compose.prod.yml pull" || echo "Docker compose not yet started"
ssh $sshHost "POSTGRES_PASSWORD=${POSTGRES_PASSWORD} docker compose -f ~/app/docker-compose.prod.yml down" || echo "Docker not running"
scp -P "22" ./tetris $scpHost:~/app/
ssh $sshHost "POSTGRES_PASSWORD=${POSTGRES_PASSWORD} docker compose -f ~/app/docker-compose.prod.yml up -d --remove-orphans"
ssh $sshHost "docker image prune -f"
ssh $sshHost "cd ~/app && TETRIS_URL=${TETRIS_URL} POSTGRES_PASSWORD=${POSTGRES_PASSWORD} ./tetris &"