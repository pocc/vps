package fastlogin

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/pocc/vpssites/logger"
    "github.com/pocc/vpssites/stripeclient"
    "encoding/json"
)

type Custdata struct {
    Email string      `json:"email"`
    Hash string       `json:"hash"`
    Requester string  `json:"requester"`
}

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
    var custrequest Custdata

    body, readErr := ioutil.ReadAll(r.Body)
    if readErr != nil {
	httpFail(w, readErr)
        return
    }
    jsonErr := json.Unmarshal(body, &custrequest)
    if jsonErr != nil {
	httpFail(w, jsonErr)
        return
    }
    resp := stripeclient.CheckAccess(custrequest.Email, custrequest.Requester, custrequest.Hash)
    w.Write(resp)
}

func httpFail(w http.ResponseWriter, err error) {
    http.Error(w, err.Error(), http.StatusBadRequest)
}
