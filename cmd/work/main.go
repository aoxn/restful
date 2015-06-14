package main
import (
	//"time"
	//"yy.com/container/cmd/work/apptest"
	"github.com/golang/glog"
	"flag"
	"yy.com/container/interfaces"
	"fmt"
)

func main(){
	var con []interfaces.ContainerInfo
	con =make([]interfaces.ContainerInfo,1)
	a:=interfaces.Job{
		Containers:con,
	}
	fmt.Println(a)
}
// InitLogs initializes logs the way we want for kubernetes.
func InitLogs() {
	flag.Parse()
	glog.Info("Starting transaction...")
}
