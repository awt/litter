package public

import ( "testing"
	"github.com/awt/litter/config"
	"github.com/awt/litter/store"
)

func TestExport(t *testing.T) {

	var conf  = &config.Config{}
	conf.SetEnvironment("test")
	conf.Set("dbpath", "./test.db")
	store.Config = conf
	store.Reset()

	// start public server configured on test ports

	// load leet fixtures

	// request leets from public server

	// verify leet count/content

}

func TestImport(t *testing.T) {
	// test import method that handles unmarshalled json
	// on json from fixture
}
