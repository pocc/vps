package fastlogin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/pocc/vpssites/logger"
)

// Show fastlog.in
func HomePage(w http.ResponseWriter, r *http.Request) {
	html, fileErr := ioutil.ReadFile("./fastlogin/index.html")
	if fileErr != nil {
	    logger.Error.Println("Unable to open local html file, ./fastlogin/index.html", fileErr.Error())
	}
	fmt.Fprintf(w, string(html))
}

// Designed for stripe's API to use this as a webhook
// Also designed to save username/password from users
func Auth(w http.ResponseWriter, r *http.Request) {
	//retVal := fmt.Sprintf("%+v", r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	//Info.Println(retVal)
	//w.Write([]byte(retVal))
}
