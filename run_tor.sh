#! /usr/bin/env bash
docker run -itd --name tor-proxy -p 8118:8118 -p 9050:9050 -d dperson/torproxy