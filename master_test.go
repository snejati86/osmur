package master

import "testing"

func TestMasterRegisterIncreasesWorkerCount(t *testing.T) {
	master := Master{maxWorker:3}
	var result = master.registerWorker(&Worker{})
	if master.currentWorkers != 1 {
		t.Error("Current worker should be one!")
	}
	if !result{
		t.Error("Current worker could not be registered")
	}

}

func TestMasterWillNotRegisterMoreThanMax(t *testing.T) {
	master:= Master{maxWorker:1}
	if !master.registerWorker(&Worker{}){
		t.Error("Couldn't register the first worker")
	}
	if master.registerWorker(&Worker{}){
		t.Error("Could register the second worker")
	}
}