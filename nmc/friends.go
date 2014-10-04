package nmc
import ( 
	"fmt"
	"net/url"
	"github.com/awt/litter/config"
	"github.com/awt/litter/store"
	"encoding/json"
	"strconv"
	"strings"
	"os/exec"
	"log"
)

var Config *config.Config

type NamecoinIdentity struct {
	Litter string
}

func Blocknotify(blockCountString string) {

	var err error
	var blockCount int64
	var names []store.Name
	blockCount, err = strconv.ParseInt(strings.TrimSpace(blockCountString), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	names, err = store.PendingNames()

	if err != nil {
		log.Fatal(err)
	}

	// for each pending name, 

	for _, name := range names {
		// check if blockcount is + 12

		if name.BlockCount + 12 >= blockCount {
			fmt.Printf("%s mature\n", name.Name)	
		} else {
			fmt.Printf("%s not ready yet\n", name.Name)	
		}
	}

			// send name_firstupdate if true

}

func FetchLeets() {
	// For each friend

	friends, _ := store.Friends()
	for _,friend := range friends {
		fmt.Printf("Fetching from %s\n", friend)

		// Fetch namecoin json
		host := LookupHost(friend.(string))

		// Fetch leets from host in namecoin json

		var list interface{}
		responseBody := fetch(host)
		json.Unmarshal(responseBody, &list)
		m := list.([]interface{})
		for _, leet := range m {
			store.ImportLeet(leet.(map[string]interface{}))
		}
	}

}

func IsRegistered(name string) (bool){
	result := namecoind("name_history", fmt.Sprintf("id/%s", name))
	if (len(result) >= 1) {
		return true
	} else {
		return false	
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

// name_new - store code in sqlite with name
func ReserveName(name string) {
	var result []interface{}
	var blockCount int64
	if Config.Name == "test" {
		fixturePath := fmt.Sprintf("test/fixtures/name_new/%s.json", name)
		resultText := store.LoadFixture(fixturePath)	
		json.Unmarshal(resultText, result)
		blockCount = 63443
	} else {
		result = namecoind("name_new", fmt.Sprintf("id/%s", name))	
		bcResult := namecoind("getblockcount")[0].(string)
		var err error
		blockCount, err = strconv.ParseInt(strings.TrimSpace(bcResult), 10, 64)
		if(nil != err) {
			fmt.Println(err.Error())	
		}
	}

	// store name, short hash, current block height in db in pending state
	fmt.Println(result)
	shortHash := result[1].(string)
	store.AddPendingName(name, shortHash, blockCount)

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

func namecoind(args ...string) (result []interface{}) {
	out, _ := exec.Command("bin/namecoind", args...).Output()
	if args[0] == "getblockcount" {
		str := string(out)
		result = make([]interface{}, 1)
		result[0] = interface{}(str)
	} else {
		fmt.Printf("command output: %s\n", out)
		json.Unmarshal(out, &result)
	}
	return result
}

