#!/usr/bin/bash

echo root:$ROOT_PASSWORD | chpasswd

mkdir -p /root/.ssh/
echo $AUTHORIZED_KEYS > /root/.ssh/authorized_keys

/usr/sbin/sshd -D -e -p 2222
