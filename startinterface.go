package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/duoflow/systemtasks/api"
	"github.com/duoflow/systemtasks/loggers"
)

func main() {
	//config_path := "/etc/sysconfig/network-scripts/"
	configPath2 := ""
	//logging initialisation
	loggers.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	loggers.Info.Println("Start of processing")
	// configure new interface
	var newSubInterface api.AliasInterface
	SubIfIndex := 2
	newSubInterface.ConfigureSubIf(SubIfIndex, "bond0", "192.168.174.162", "255.255.255.0")
	loggers.Info.Println("New interface was configured:")
	loggers.Info.Println(newSubInterface)
	config := newSubInterface.CreateConfigurationFile()
	loggers.Info.Println("New config created: \n" + config)
	// write config file to filesystem
	go writeConfig(config, newSubInterface.FILENAME, configPath2)
}

// writetofile() - writes config to filesystem
func writeConfig(configString string, fileName string, filePath string) int {
	// create file
	f, err := os.Create(filePath + fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return 2
	}
	fmt.Println("File was created: " + filePath + fileName)
	// write config to file
	l, err := f.WriteString(configString)
	if err != nil {
		fmt.Println(err)
		return 3
	}
	fmt.Println(l, "bytes written successfully")
	// close the file
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return 4
	}
	return 0
}
