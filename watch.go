package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"
)


var interval *int

func init() {
	interval = flag.Int("interval", 2, "Interval in seconds")
}
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		panic("You MUST supply a command to repeat")
	}
	cmdName, err := exec.LookPath(args[0])
	if err != nil {
		log.Panicln("Command ",args[0]," not found in path")
	}
	log.Print("Running:", args, " every ", *interval, " seconds ")
	for {
		cmd := exec.Command(cmdName,args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Print("Running:", args)
		err := cmd.Run()
		if err != nil {
			log.Print(err)
		}
		cmd.Wait()
		time.Sleep(time.Second * time.Duration(*interval))
	}
}
