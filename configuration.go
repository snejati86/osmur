package master

type MasterConfiguration struct{
	database string
	workers int8
}

func NewMasterConfiguration(map[string] js) *MasterConfiguration{
	this := new(MasterConfiguration)
}
