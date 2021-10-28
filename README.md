# Docker-Distributor-Bot
use QQ-Bot to distribute docker containers

### Features
+ [x] use ssh to control docker daemon on remote servers
+ [x] build docker image by custom Dockerfile
+ [x] create/delete containers
  + mount volumns
  + enable password login
+ [ ] monitor system state (CPU / GPU / Memory ...)
+ [ ] listen the dialog in QQ group and react to keywords

### Known Bugs
+ [ ] tar can't work on windows, maybe caused by '\\' in path
