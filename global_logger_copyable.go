package prototype

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

//PROTOTYPE pattern. The goal of prototype pattern here is to provide a way to clone object from existed prototype

var (
	globalConfig *global
	once         sync.Once
)

//Interface for copying from prototype
type IClonable interface {
	Clone() global
}

type IConfig interface {
	IClonable 
    SetLogger(logger *logrus.Logger)
    GetLogger() *logrus.Logger
    SetDomain(domain string)
    SetDomain() string
    SetDev(flag bool)
    IsDev() bool
}

//Some global config which can be used accross the project
type global struct {
	logger *logrus.Logger
	dev    bool
	domain string
}

//'global' impelements 'Clone' method for further copying of itself
func (g *global) Clone() IConfig {
	return global{
		Logger: g.GetLogger(),
		Dev:    g.IsDev(),
		Domain: g.GetDomain(),
	}
}

//It's good to incapsulate members.
//Methods with recevier by value (it modifies only copy. To assign copy to existed prototype use 'SetGlobalConfig')
func (g global) SetLogger(logger *logrus.Logger){
	g.logger = logger
}

func (g global) GetLogger(){
	return g.logger
}

func (g global) SetDomain(domain string){
	g.domain = domain
}

func (g global) GetDomain(){
	return g.domain
}

func (g global) SetDev(flag bool){
	g.dev = flag
}

func (g global) IsDev() {
	return g.dev
}

//To change global config
func SetGlobalConfig(g IConfig) {
	if globalConfig != nil {
		globalConfig.SetDev(g.IsDev())
		globalConfig.SetDomain(g.GetDomain())
		globalConfig.SetLogger(g.GetLogger())
	}
}

/* It is critical to return the copy of configuration. So let's use the prototype pattern and perform copying
   Also, let's use 'sync.Once' instead of checking 'globalConfig' for nil and perform Init if it so*/
func GetConfigInstance() IConfig  {
	once.Do(func() {
		globalConfig = createDefault()
	})
	return globalConfig.Clone()
}

/*createDefault unexported because we won't create Global config through 'createDefault' function.
  We will use GetInstance constantly (for the first time and in further).
  To configure the global config we can use 'SetGlobalConfig' to configure it based on modified copy*/
func createDefault() IConfig  {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	return &global{
		Logger: logger,
	}
}

/*=============Usage===================
//The only way to get current prototype is to use 'GetConfigInstance'
defConfig := GetConfigInstance() //If we are using it for the first time, we have copy of default prototype here

//Fill the copy with some meaninfull data
defConfig.SetDev(*modeDev)
if defConfig.IsDev() {
	// Prettify logger output when developing
	logger := defConfig.GetLogger()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	defConfig.SetLogger(logger)
}
defConfig.SetDomain(os.Getenv(config.EnvDomain))

//The only way to change prototype is to use exported func 'SetGlobalConfig'
SetGlobalConfig(defConfig) //after 'SetGlobalConfig' and further GetConfigInstance() we will have a copy of modified prototype

========================================*/
