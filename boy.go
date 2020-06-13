package boy

import (
	"github.com/slclub/gcore"
	"github.com/slclub/gcore/execute"
	"github.com/slclub/grouter"
)

func testen() {
	en := gcore.New()
	en.Run()
}

// register driver
func registerDriver(en *gcore.Engine) {

	router := grouter.NewRouter()
	router.SetKey("router")
	en.DriverRegister(router)

	mida := execute.NewMiddle("before_mid")
	en.DriverRegister(mida)

	process := &execute.Process{}
	process.SetKey("handle")
	en.DriverRegister(process)

	midb := execute.NewMiddle("after_mid")
	en.DriverRegister(midb)

}

// register static file.
func registerStatic(en *gcore.Engine) {
}
