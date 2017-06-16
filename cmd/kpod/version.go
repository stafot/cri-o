package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

//Overwritten at build time
var (
	gitCommit string
	buildInfo string
)

//Function to get and print info for version command
func CmdVersion() {

	//converting unix time from string to int64
	buildTime, err := strconv.ParseInt(buildInfo, 10, 64)
	if err != nil {
		fmt.Println("Error getting build time:", err)
	}

	fmt.Println("Version:	", Version)
	fmt.Println("Go Version:	", runtime.Version())
	fmt.Println("Git Commit:	", gitCommit)

	//Prints out the build time in readable format
	fmt.Println("Built:		", time.Unix(buildTime, 0).Format(time.ANSIC))
	fmt.Println("OS/Arch:	", runtime.GOOS+"/"+runtime.GOARCH)
}

//cli command to print out the full version of kpod
var versionCommand = cli.Command{
	Name:  "version",
	Usage: "Display the KPOD Version Information",
	Action: func(context *cli.Context) error {
		CmdVersion()
		return nil
	},
}
