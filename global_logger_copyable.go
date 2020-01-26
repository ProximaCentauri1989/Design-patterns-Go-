package prototype

import (
	"os"
	"sync"
)

//PROTOTYPE pattern. The goal of prototype pattern here is to provide a way to clone object from existed prototype

var (
	objectWithGlobalData *protoObject
	once         sync.Once
)

//Interface for copying from prototype
type IClonable interface {
	Clone() global
}

type IConfig interface {
	IClonable 
    SetParam(param string)
    GetParam() string
    SetFlag(flag bool)
    IsTrue() bool
}

//Some global object
type protoObject struct {
	flag    bool
	param string
}

//'protoObject' impelements 'Clone' method for further copying of itself
func (g *protoObject) Clone() IConfig {
	return protoObject{
		flag:    g.IsTrue(),
		Domain: g.GetParam(),
	}
}

//It's good to incapsulate members.
//Methods with recevier by value (it modifies only copy. To assign copy to existed prototype use 'SetGlobalConfig')

func (g protoObject) SetParam(domain string){
	g.domain = domain
}

func (g protoObject) GetParam(){
	return g.domain
}

func (g protoObject) SetFlag(flag bool){
	g.flag = flag
}

func (g protoObject) IsDev() {
	return g.dev
}

//To change global config
func SetProtoObject(g IConfig) {
	if objectWithGlobalData != nil {
		objectWithGlobalData.SetFlag(g.IsTrue())
		objectWithGlobalData.SetParam(g.GetParam())
	}
}

/* It is critical to return the copy of configuration. So let's use the prototype pattern and perform copying
   Also, let's use 'sync.Once' instead of checking 'objectWithGlobalData' for nil and perform Init if it so*/
func GetProtoObject() IConfig  {
	once.Do(func() {
		objectWithGlobalData = createDefaultObject()
	})
	return objectWithGlobalData.Clone()
}

/*createDefault unexported because we won't create Global config through 'createDefault' function.
  We will use GetInstance constantly (for the first time and in further).
  To configure the global config we can use 'SetGlobalConfig' to configure it based on modified copy*/
func createDefaultObject() IConfig  {
	return &protoObject{
		param: "default",
	}
}

/*=============Usage===================
//The only way to get current prototype is to use 'GetConfigInstance'
someProtoObject := GetProtoObject() //If we are using it for the first time, we have copy of default prototype here

//Fill the copy with some meaninfull data
someProtoObject.SetFlag(true)
if someProtoObject.IsTrue() {
	someProtoObject.SetParam("Start")
}

//The only way to change prototype is to use exported func 'SetProtoObject'
SetProtoObject(someProtoObject) //after 'SetGlobalConfig' and further GetProtoObject() we will have a copy of modified prototype

========================================*/
