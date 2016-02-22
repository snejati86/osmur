package master

import (
	"io/ioutil"
	"encoding/json"
)

const (
	CONFIGURATION_FILE  = "config.json"
	DATABASE = "database"
)

type Master struct{
	maxWorker int
	currentWorkers int
}

type Worker struct {}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func (master *Master) init() {

	file, err := ioutil.ReadFile(CONFIGURATION_FILE)
	//CHECK FILE EXISTENCE.
	check(err)
	var data map[string]interface{}
	fileBytes := []byte(file)
	//CHECK VALID JSON.
	check(json.Unmarshal(fileBytes,&data))

}

func (master *Master) registerWorker(worker *Worker) bool {
	if  master.maxWorker - master.currentWorkers < 1 {
		return false
	}
	master.currentWorkers++
	return true
}
