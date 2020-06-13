package boy

import (
	"github.com/slclub/link"
	"github.com/slclub/utils"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestInstall(t *testing.T) {
	Install()

	file_path := path.Join(link.APP_PATH, conf_dir)
	file := path.Join(file_path, conf_file)

	ret, err := utils.IsFileExist(file)
	assert.True(t, ret)
	assert.Empty(t, err)

	file = path.Join(file_path, "noconf.ini")
	ret, _ = utils.IsFileExist(file)
	assert.False(t, ret)
}

func TestInstallForce(t *testing.T) {
	ret := CreateFileAuto(conf_dir, conf_file, true)
	assert.True(t, ret)
	ret = CreateFileAuto(conf_dir, conf_file, false)
	assert.False(t, ret)

	ret = CreateFileAuto(conf_dir, "noconf.ini", false)
	assert.False(t, ret)
}
