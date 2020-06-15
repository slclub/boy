package boy

import (
	"github.com/slclub/gcore"
	"github.com/slclub/gcore/execute"
	"github.com/slclub/grouter"
	"github.com/slclub/link"
	"path"
	"strings"
)

var (
	// core application object.
	App *gcore.Engine

	// alise router
	R grouter.Router

	// alise object first middlerware node.
	MiddlerBefore execute.Middler

	// alise object of second midderware node.
	MiddlerAfter execute.Middler

	// For user customelized object.
	Defined *DefinedObject
)

func init() {
	Initialize()
}

func Initialize() {
	link.DEBUG_PRINT("[BOY][Initialize]START!")
	App = gcore.New()

	registerDriver(App)

	Defined = &DefinedObject{}

	Defined.Start()

	R, _ = App.DriverRouter("router")

	registerStatic(App)

	MiddlerBefore, _ = App.DriverMiddler("before_mid")
	MiddlerAfter, _ = App.DriverMiddler("after_mid")

	link.DEBUG_PRINT("[BOY][DEFINED_OBJECT][ADD_REGISTER][Initialize]")
}

// **************************************Defined********************************************
type DefinedObject struct {
	registers []func(*gcore.Engine)
}

func (def *DefinedObject) Start() {
	for _, fn := range def.registers {
		if fn == nil {
			continue
		}
		fn(App)
	}
}

func (def *DefinedObject) AddRegister(callfunc func(en *gcore.Engine)) {

	def.registers = append(def.registers, callfunc)
}

func (def *DefinedObject) Size() int {
	return len(def.registers)
}

// **************************************funcs********************************************
func Run() {
	App.Run()
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
	link.DEBUG_PRINT("[BOY][STATIC]", "START")

	static_service_key := "static_service"
	static_service := link.Config().GetString(static_service_key + ".service")
	ss_arr := strings.Split(static_service, " ")
	root_key := "root"
	// static root path.
	root_path := link.GetString(static_service_key+"."+root_key, "")

	//link.DEBUG_PRINT("[BOY][STATIC]", "static service.static:", (ss_arr))

	for _, sk := range ss_arr {
		sk = strings.Trim(sk, " ")
		if sk == "" {
			continue
		}
		v := link.GetString(static_service_key+"."+sk, "")
		//link.DEBUG_PRINT("[BOY][STATIC] key[", static_service+"."+sk, "] static[", v, "]")
		//v_str, _ := v.(string)
		path_arr := static_parse(v)
		if path_arr == nil {
			continue
		}
		tail := false
		path_arr[2] = strings.ToLower(path_arr[2])
		if path_arr[2] == "on" || path_arr[2] == "true" || path_arr[2] == "yes" {
			tail = true
		}
		link.DEBUG_PRINT("[BOY][STATIC]", "Alias[", path_arr[0], "]actual[", path_arr[1], "]tail[", tail, "]", len(path_arr), "|")
		// add static router.
		R.ServerFile(path_arr[0], path.Join(root_path, path_arr[1])+"/", tail)
	}
}

func static_parse(conf_line string) []string {
	//conf_line = strings.Replace(conf_line, "  ", " ", -1)
	conf_arr_tmp := strings.Split(conf_line, " ")
	conf_arr := []string{}
	for _, v := range conf_arr_tmp {
		if v == "" {
			continue
		}
		conf_arr = append(conf_arr, v)
	}
	size := len(conf_arr)
	if size < 3 {
		return nil
	}
	if len(conf_arr[0]) > 0 {
		if conf_arr[0][len(conf_arr[0])-1] != '/' {
			conf_arr[0] += "/"
		}
		if conf_arr[0][0] != '/' {
			conf_arr[0] = "/" + conf_arr[0]
		}
	}
	if len(conf_arr[1]) > 0 && conf_arr[1][len(conf_arr[1])-1] != '/' {
		conf_arr[1] += "/"
	}
	return conf_arr

}
