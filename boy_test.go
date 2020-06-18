package boy

import (
	"github.com/slclub/gcore"
	"github.com/slclub/gnet"
	"github.com/slclub/grouter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoyNew(t *testing.T) {
	//Initialize()
	count := 0
	Defined.AddRegister(func(en *gcore.Engine) { count++ })
	Defined.AddRegister(nil)
	Defined.Start()
	assert.Equal(t, 1, count)
	assert.Equal(t, 2, Defined.Size())
}

func TestBoyRouter(t *testing.T) {
	grouter.Group.Use(func(ctx gnet.Contexter) {})
}
