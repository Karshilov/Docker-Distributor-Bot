from nvidia/cuda:11.2.2-runtime-ubuntu20.04

copy entrypoint.sh /var/run/entrypoint.sh

run \
  apt update ; \
  echo "**** install runtime packages ****" && \
  apt install -y --upgrade \
    curl \
    logrotate \
    nano \
  sudo && \
  echo "**** install openssh-server ****" && \
  DEBIAN_FRONTEND=noninteractive apt install -y \
    openssh-client \
    openssh-server \
    openssh-sftp-server && \
  echo "**** setup openssh environment ****" && \
  mkdir /run/sshd && \
  sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/g' /etc/ssh/sshd_config && \
  sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/g' /etc/ssh/sshd_config && \
  rm -rf \
    /tmp/* && \
  chmod 777 /var/run/entrypoint.sh

expose 2222

cmd ["/var/run/entrypoint.sh"]
