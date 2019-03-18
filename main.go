package main

import (
	//"fmt"
	"io/ioutil"
	"os"
	"sync"

	//"github.com/duoflow/systemtasks/api"
	"github.com/duoflow/systemtasks/loggers"
	"github.com/duoflow/systemtasks/webserver"
)

var (
	wg sync.WaitGroup
)

// writetofile() - writes config to filesystem
func writeConfig(configString string, fileName string, filePath string) int {
	// syncronize operational status
	defer wg.Done()
	// create file
	f, err := os.Create(filePath + fileName)
	defer f.Close()
	if err != nil {
		loggers.Error.Println(err)
		return 2
	}
	loggers.Info.Println("File was created: " + filePath + fileName)
	// write config to file
	l, err := f.WriteString(configString)
	if err != nil {
		loggers.Error.Println(err)
		return 3
	}
	loggers.Info.Println(l, "bytes written successfully")
	// close the file
	err = f.Close()
	if err != nil {
		loggers.Error.Println(err)
		return 4
	}
	return 0
}

func main() {
	//config_path := "/etc/sysconfig/network-scripts/"
	//configPath2 := ""
	//logging initialisation
	loggers.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	loggers.Info.Println("Start of processing")
	// start REST API server
	webserver.StartServer()
}
