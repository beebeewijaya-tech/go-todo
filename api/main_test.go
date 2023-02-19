package api

import (
	"log"
	"os"
	"testing"

	"beebeewijaya.com/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var c *viper.Viper

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	var err error
	c, err = util.NewConfig("../env")
	if err != nil {
		log.Fatalf("failed when initialize config: %v", err)
	}

	os.Exit(m.Run())
}
