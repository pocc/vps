package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mholt/certmagic"
	"github.com/pocc/vpssites/checkbpfnet"
	"github.com/pocc/vpssites/fastlogin"
	"github.com/pocc/vpssites/logger"
)

func main() {
	logger.InitLog()
	// Use .SkipClean() to ensure that %2f is not rendered as /
	router := mux.NewRouter()
	router = router.UseEncodedPath()

	/***** checkbpf.net *****/
	router.HandleFunc("/", checkbpfnet.HomePage).Host("checkbpf.net")
	// Use .* as match so that all characters (including ?) are included
	router.HandleFunc("/status/{filter:.*}", checkbpfnet.Status).Host("checkbpf.net").Methods("GET")
	router.HandleFunc("/code/{filter:.*}", checkbpfnet.Code).Host("checkbpf.net").Methods("GET")
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./checkbpfnet"))).Host("checkbpf.net")

	/***** fastlog.in *****/
	router.HandleFunc("/", fastlogin.HomePage).Host("fastlog.in")
	// For user database queries
	router.HandleFunc("/auth/", fastlogin.Auth).Host("fastlog.in").Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./fastlogin"))).Host("fastlog.in")

	log.Fatal(certmagic.HTTPS([]string{"checkbpf.net", "fastlog.in"}, router))
}
