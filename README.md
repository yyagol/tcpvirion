# tcpslowloris
TCP slowloris attack

This simple tool will generate N tcp connections and will keep them alive
by sending keepalive signals

```
usage: tcpslowloris.py [-h] [-t TOTAL] [-p PULSE] [-d DELAY] [--target TARGET] [--port PORT]

optional arguments:
  -h, --help            show this help message and exit
  -t TOTAL, --total TOTAL
                        Total concurrent TCP sessions
  -p PULSE, --pulse PULSE
                        Number of connections per pulse
  -d DELAY, --delay DELAY
                        Number of seconds betweek pulse
  --target TARGET       Target host
  --port PORT           Target port
```
Example use :
>./tcpslowloris.py --target 192.168.0.1 --port 80 -t 100 -p 10 -d 5 
