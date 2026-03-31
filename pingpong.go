package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    ping := make(chan string)
    pong := make(chan string)
    var wg sync.WaitGroup
    iterations := 5
    count:=0
    wg.Add(2)
    go pingfunc(&count,ping, pong, &wg,iterations)
    go pongfunc(&count,ping, pong, &wg,iterations)

    time.Sleep(time.Second)
    ping <- "ping"
    
    wg.Wait()

}
func pingfunc(count *int,ping,pong chan string,wg *sync.WaitGroup,iterations int) {
    msg, ok := <-ping
        if !ok {
            wg.Done()
            return
        }

        fmt.Println(msg)
        pong <- "pong"
        
        pingfunc(count,ping, pong, wg,iterations)
        
    
}
func pongfunc(count *int,ping,pong chan string,wg *sync.WaitGroup,iterations int) {
    pongmsg := <-pong
    fmt.Println(pongmsg)
    (*count)++
    if *count>=iterations{
        close(ping)
        wg.Done()
        return  
    }
    ping <- "ping"
    pongfunc(count,ping, pong, wg,iterations)
            
        
    
}
