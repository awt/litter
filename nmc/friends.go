package nmc
import ( 
	"encoding/json"
	"errors"
	"fmt"
	"github.com/awt/litter/config"
	"github.com/awt/litter/store"
	"log"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
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

		if name.BlockCount + 12 <= blockCount {

			// send name_firstupdate if true

			err = NameFirstUpdate(name)	
			if err != nil {
				log.Fatal(err)	
			}

			// mark state as registered in db

			store.MarkRegistered(name.Name)
		}
	}
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
	result, err := namecoind("name_history", fmt.Sprintf("id/%s", name))
	if ((len(result) >= 1) && (err == nil)) {
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

func NameFirstUpdate(name store.Name) (err error) {
	// if onionHostname is not set there's no point

	if Config.Get("onionHostname") == "" {
		return errors.New("Can't perform name_firstupdate since onionHostname is not set.")
	}

	result, err := namecoind("name_firstupdate",
	fmt.Sprintf("id/%s", name.Name),
	name.ShortHash,
	fmt.Sprintf("'{\"litter\":\"%s\"}'",
	Config.Get("onionHostname")))	
	if err != nil {
		fmt.Printf("Failed to finalize registration of %s due to the following error: %s\n", name.Name, result)	
	}
	
	return err
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
		result, _ = namecoind("name_new", fmt.Sprintf("id/%s", name))	
		array, _ := namecoind("getblockcount")
		bcResult := array[0].(string)
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

func namecoind(args ...string) (result []interface{}, err error) {
	fmt.Println(strings.Join(args, " "))
	out, err := exec.Command("bin/namecoind", args...).CombinedOutput()

	if err != nil {
		fmt.Printf("Failure: %s\n", out)
		result = make([]interface{}, 1)
		var temp interface{}
		json.Unmarshal(out, &temp)
		result[0] = temp
	} else if args[0] == "getblockcount" {
		str := string(out)
		result = make([]interface{}, 1)
		result[0] = interface{}(str)
		fmt.Printf("Success: %s\n", str)
	} else {
		fmt.Printf("Success: %s\n", out)
		json.Unmarshal(out, &result)
	}
	return result, err
}

