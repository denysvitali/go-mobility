package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
)

type LoginCmd struct {
	
}

var args struct {
	Login *LoginCmd `arg:"subcommand:login"`
	Version *bool `arg:"-v"`
}

var Version = "dev"

func main(){
	arg.MustParse(&args)
	
	if args.Version != nil {
		fmt.Printf("mobility-cli %s", Version)
	}
	
	if args.Login != nil {
		login(args.Login)
		return
	}
}

func login(login *LoginCmd) {
	
}
