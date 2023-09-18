FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install node.js npm curl -y

RUN npm i -g n typescript ts-node tslib @types/node
RUN n latest
RUN npm rm -g n

RUN apt-get purge npm curl -y
RUN apt-get autoremove -y

RUN groupadd sandbox
RUN useradd sandbox -m -g sandbox

USER sandbox:sandbox

WORKDIR /home/sandbox
