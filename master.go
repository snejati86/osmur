package main
import (
	"fmt"
)

type MasterInterface interface {
	registerWorker(worker *Worker) bool
}
type Master struct{
	maxWorker int
	currentWorkers int
}

type Worker struct {}

func (master *Master) registerWorker(worker *Worker) bool {
	if  master.maxWorker - master.currentWorkers == 1 {
		return false
	}
	master.currentWorkers++
	return true
}

func main(){
	master := Master{maxWorker:2}
	worker := Worker{}
	fmt.Println(master.registerWorker(&worker))
}
