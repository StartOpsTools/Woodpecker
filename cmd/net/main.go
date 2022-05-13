package main

import (
	"flag"
	"github.com/StartOpsTools/Woodpecker/pkg/os"
)

type king string
// inet  inet4  inet6  tcp  tcp4  tcp6  udp  udp4  udp6  unix
func (k king) String() {

}

var kind *string
var pid *int
func init() {
	kind = flag.String("kind", "all", "kind: all, tcp, udp...")
	pid = flag.Int("pid", 0, "process pid")
}

func main() {
	flag.Parse()

	if *pid != 0 {
		os.PidConnection("tcp", int32(*pid))
	}
	os.Connection(*kind)


}
