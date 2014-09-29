package nmc

import ( "testing"
	"github.com/awt/litter/config"
	"github.com/awt/litter/store"
)

func Test(t *testing.T) {
	var conf  = &config.Config{}
	conf.SetEnvironment("test")
	conf.Set("dbpath", "./test.db")
	Config = conf
	store.Config = conf
	store.Reset()

	store.Follow("satoshi")
	FetchLeets()

	leets, _ := store.Leets()
	if len(leets) != 3 {
		t.Error("Expected there to be 3 leets in the database.");
	}

}
