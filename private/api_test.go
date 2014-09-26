package private

import ( "testing"
	"github.com/awt/litter/config"
	"github.com/awt/litter/store"
)

func Test(t *testing.T) {

	var conf  = &config.Config{}
	conf.SetEnvironment("test")
	conf.Set("dbpath", "./test.db")
	store.Config = conf
	store.Reset()
	// test leet posting

	body, code := route("/", "POST")
	if code != 200 {
		t.Error("Unexpected response code.");
	}

	if body != "" {
		t.Error("Expected response body to be empty.");
	}
}
