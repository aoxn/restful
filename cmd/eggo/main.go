package main

import (
    "time"

    "spacex.com/eggo/pkg/apiserver"
)


func main(){
    //runtime.GOMAXPROCS(runtime.NumCPU())
    ser := apiserver.NewNodeServer()
    ser.Start()
    for {
       //TODO: add config file watch/notify logic here
       time.Sleep(10 * time.Second)
    }
}


