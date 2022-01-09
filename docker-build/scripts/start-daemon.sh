#!/bin/bash

mkdir $COREDNS_HOME/logs
nohup $COREDNS_HOME/mate/coredns_mate >>$COREDNS_HOME/logs/coredns_mate.stdout.log 2>>$COREDNS_HOME/logs/coredns_mate.stderr.log

