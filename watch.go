package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"reflect"
	"time"
)

func getReflectArgs(args []string) []reflect.Value {
	out := make([]reflect.Value, len(args))
	for index := 0; index < len(args); index++ {
		if index == 0 {
			path, err := exec.LookPath(args[0])
			if err != nil {
				panic(err)
			}
			args[0] = path
		}
		out[index] = reflect.ValueOf(args[index])
	}
	return out
}

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
	command := reflect.ValueOf(exec.Command)
	reflectArgs := getReflectArgs(args)
	log.Print("Running:", args, " every ", *interval, " seconds ")
	for {
		returns := command.Call(reflectArgs)
		cmdPtr := returns[0].Interface()
		cmd := cmdPtr.(*exec.Cmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Print("Running:", args)
		err := cmd.Run()
		if err != nil {
			log.Print(err)
		}
		time.Sleep(time.Second * time.Duration(*interval))
	}
}
