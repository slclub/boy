package boy

import (
	"github.com/slclub/link"
	"github.com/slclub/utils"
	"os"
	"path"
	"runtime"
)

var (
	conf_dir      = "etc"
	conf_dir_temp = "etc_temp"
	conf_file     = "go.ini"
)

func CreateDir(path_dir string) bool {
	err := os.MkdirAll(path_dir, os.ModePerm)
	if err != nil {
		link.ERROR("[BOY][INSTALL]", err)
		return false
	}
	return true
}

func InstallDir(path_dir string) {
	if ok, _ := utils.IsFileExist(path_dir); ok {
		return
	}
	CreateDir(path_dir)
}

func CreateFileAuto(path_dir, file_name string, force bool) bool {
	// temp config file.
	// link.DEBUG_PRINT("[BOY][INSTALL]", current_file())
	file := path.Join(path.Dir(current_file()), conf_dir_temp)
	file = path.Join(file, file_name)
	old_content, ok := utils.ReadAllByte(file)
	if !ok {
		link.DEBUG_PRINT("[BOY][INSTALL][CREATE_CONFIG_FILE][TEMP_FILE][NOT_EXIST]", file)
		return false
	}

	file = path.Join(link.APP_PATH, path_dir)
	file = path.Join(file, file_name)

	// check and create config file.
	if ret, _ := utils.IsFileExist(file); ret && !force {
		link.DEBUG_PRINT("[BOY][INSTALL][CREATE_CONFIG_FILE]", file)
		return false
	}

	fn, err := os.Create(file)
	if err != nil {
		return false
	}
	defer fn.Close()

	fn.Write(old_content)
	return true
}

func Install() {
	path_dir := path.Join(link.APP_PATH, conf_dir)
	InstallDir(path_dir)
	CreateFileAuto(conf_dir, conf_file, false)
}

func current_file() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return file
}
