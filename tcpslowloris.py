#!/usr/bin/python

import argparse
import socket
import time
import threading

parser = argparse.ArgumentParser()
parser.add_argument("-t","--total",  help="Total concurrent TCP sessions", type=int, default=1)
parser.add_argument("-p","--pulse", help="Number of connections per pulse", type=int, default=1)
parser.add_argument("-d","--delay" ,help="Number of seconds betweek pulse", type=int, default=1)
parser.add_argument("--target" ,help="Target host")
parser.add_argument("--port" ,help="Target port" , type=int)
args = parser.parse_args()

def worker():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((args.target, args.port))
    s.setsockopt(socket.SOL_SOCKET, socket.SO_KEEPALIVE,  1)
    s.setsockopt(socket.IPPROTO_TCP, socket.TCP_KEEPIDLE, 1)
    s.setsockopt(socket.IPPROTO_TCP, socket.TCP_KEEPINTVL, 23)
    s.setsockopt(socket.IPPROTO_TCP, socket.TCP_KEEPCNT, 5)
    while True:
        time.sleep(5)
    return

threads = []
for i in range(int(args.total/args.pulse)):
    for pulse in range(args.pulse):
        t = threading.Thread(target=worker)
        threads.append(t)
        t.start()
    print("Current sockets : %d" % int((i+1)*args.pulse))
    time.sleep(args.delay)
