# Tetris_go
Full Stack App (Backend:Go/Echo Frontend:Typescript/React DB:Postgres). Live [here](https://tetris.zfz7.org): 
#### Step 1. create an AWS EC2 instance / Digital Ocean droplet

- Create a ubuntu 22.04 LTS AWS EC2 instance / Digital Ocean droplet
- Create and associate elastic IP to that instance
- Point your DNS at that IP address
- Set up inbound firewall rules:

| IP Version | Type  | Protocol                        | Port    | Source    |
|------------|-------|---------------------------------|---------|-----------|
| IPv4       | HTTP  | TCP                             | 80      | 0.0.0.0/0 |
| IPv6       | HTTP  | TCP                             | 80      | ::/0      |
| IPv4       | HTTPS | TCP                             | 443     | 0.0.0.0/0 |
| IPv6       | HTTPS | TCP                             | 443     | ::/0      |
| IPv4       | SSH   | TCP                             | 22      | 0.0.0.0/0 |
| IPv6       | SSH   | TCP                             | 22      | ::/0      |

#### Step 2. setup environment variables
- Fill in all environment variables in [.envrc.example](./.envrc.example)
- If desired you can use [direnv](https://direnv.net/) for an easy way to manage environment variables
```
TETRIS_URL="tetris.zfz7.org"
TETRIS_SSH_USER="root"

POSTGRES_DB_PASSWORD="example"
```

#### Step 3. setup AWS EC2  / Digital Ocean instance

- Run [`./setupVM.sh`](./setupVM.sh) or [`./setupVM.sh DO`](./setupVM.sh) on instance
  - Check the and update to the latest docker compose version
  - `scp -P 22 ./setupVM.sh $TETRIS_SSH_USER@$TETRIS_URL:~/`

#### Step 6. create letscrypt cert
- run `sudo certbot certonly --standalone` then follow prompts

#### Step 7. deploy app
- `./deploy.sh` or `./deploy.sh NOTEST`


## Short term goals
- [X] Simple react frontend: typescript/react
  - [X] frontend makes a backend GET call
  - [X] frontend tests 
- [X] Simple backend: go/echo
  - [X] backend serves react app
  - [X] backend responds to GET call
  - [X] backend reads/writes to Postgres DB 
  - [X] backend tests 
- [X] Cypress test for full application
- [X] Building app
  - [X] Frontend & backend builds into single thing?
- [X] Github pipeline
  - [X] Frontend tests
  - [X] Backend tests
  - [X] Cypress tests
- [X] Deploy app
  - [X] TLS certs
  - [X] deploy application with ./deploy.sh
  - [X] setup Digital Ocean droplet with ./setup.sh
