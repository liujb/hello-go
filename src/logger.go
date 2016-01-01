package main

import (
	"flag"
	"log"
)

var (
	logFileName = flag.String("log", "trace-driver.log", "trace-driver-log")
)

func main() {

	log.Print("Serversfsfsfsfsf")
	log.Printf("%dsdfsf%s", 123, "llkljkkljljljlj")
	log.Println(123123123242)
	log.Print("jkafdakfa")
}

// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())

// 	flag.Parse()
// 	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

// 	if logErr != nil {
// 		fmt.Println("Fail to find", *logFile, "trace-driver.log start failed")
// 		os.Exit(1)
// 	}

// 	log.SetOutput(logFile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// }
