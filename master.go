package master

type Master struct{
	maxWorker int
	currentWorkers int
}

type Worker struct {}

func (master *Master) registerWorker(worker *Worker) bool {
	if  master.maxWorker - master.currentWorkers < 1 {
		return false
	}
	master.currentWorkers++
	return true
}
