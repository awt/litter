package nmc
import ( 
	"fmt"
	"net/url"
	"github.com/awt/litter/config"
	"github.com/awt/litter/store"
	"encoding/json"
)

var Config *config.Config

type NamecoinIdentity struct {
	Litter string
}

func FetchLeets() {
	// For each friend

	friends, _ := store.Friends()
	for _,friend := range friends {
		fmt.Printf("Fetching from %s", friend)

		// Fetch namecoin json
		host := LookupHost(friend.(string))

		// Fetch leets from host in namecoin json

		fetch(host)
	}

}

// Look up friend by name

func LookupHost(name string) (string) {
	// fetch json from namecoin id

	// extract litter property from json and return

	var namecoinEntryText []byte
	if Config.Name == "test" {
		fixturePath := fmt.Sprintf("test/fixtures/identities/%s.json", name);
		namecoinEntryText = store.LoadFixture(fixturePath)
	} else {
	
	}
	var namecoinIdentity NamecoinIdentity
	json.Unmarshal(namecoinEntryText, &namecoinIdentity)

	return namecoinIdentity.Litter
}

func fetch(u string) (body []byte) {
	if Config.Name == "test" {
		httpUrl, _ := url.Parse(u)
		fixturePath := fmt.Sprintf("test/fixtures/%s.json", httpUrl.Path)
		body = store.LoadFixture(fixturePath)	
	} else {
			
	}
	return body
}
