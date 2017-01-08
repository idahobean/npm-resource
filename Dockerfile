FROM idahobean/alpine-node-bash

ADD built-check /opt/resource/check
ADD built-out /opt/resource/out
ADD built-in /opt/resource/in
