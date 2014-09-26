package private

import ( "testing"
)

func Test(t *testing.T) {

	// test leet posting

	body, code := route("/", "POST")
	if code != 200 {
		t.Error("Unexpected response code.");
	}

	if body != "" {
		t.Error("Expected response body to be empty.");
	}
}
