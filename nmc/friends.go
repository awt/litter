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
)

var Config *config.Config

type NamecoinIdentity struct {
	Litter string
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
	//var blockCount int
	if Config.Name == "test" {
		fixturePath := fmt.Sprintf("test/fixtures/name_new/%s.json", name)
		resultText := store.LoadFixture(fixturePath)	
		json.Unmarshal(resultText, result)
		//blockCount = 63443
	} else {
		result := namecoind("name_new", fmt.Sprintf("id/%s", name))	
		fmt.Sprintf("result: %s", result);
		bcResult := namecoind("getblockcount")[0].(string)
		fmt.Println(bcResult)
		blockCount, err := strconv.ParseInt(strings.TrimSpace(bcResult), 10, 64)
		if(nil != err) {
			fmt.Println(err.Error())	
		}
		fmt.Printf("blockcount: %d\n", blockCount)
	}

	// store name, short hash, current block height in db in pending state

	//shortHash := result[1]
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
	fmt.Println(args)
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

