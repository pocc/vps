package checkbpfnet

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pocc/vpssites/logger"
)

func callDumpcap(filter string) ([]byte, error) {
	// Make sure that filter matches this regex (valid BPF symbols)
	re := regexp.MustCompile(`[^ !&()\-.\d:<=>A-\]a-z|]+`)
	matches := re.FindAllString(filter, -1)
	if len(matches) > 0 {
		matchString := strings.Join(matches, "")
		// %q gives us what we want, but adds extra quotes
		runes := []rune(fmt.Sprintf("%q", matchString))
		// remove extra quotes
		safeMatchString := string(runes[1 : len(runes)-1])
		return nil, errors.New("Invalid BPF Characters:[" + safeMatchString + "]")
	}
	// Do not perform DNS for a filter if it looks like a url
	// It may still be valid, so just use localhost instead,
	// which will *always* be valid
	noDnsRe := regexp.MustCompile(`[a-zA-Z-]+(?:\.[a-zA-Z0-9-]+)+`)
	filter = noDnsRe.ReplaceAllLiteralString(filter, "localhost")
	cmd := exec.Command("dumpcap", "-d", "-f", filter)
	out, err := cmd.CombinedOutput()
	cmdErr := cmd.Run()
	if cmdErr != nil && err != nil {
		logger.Error.Println("cmd.Run() failed with stdout " + string(out) + " stderr " + err.Error() + "command err" + cmdErr.Error())
	}
	expectedErrorText := "That string isn't a valid capture filter"
	// Confusingly, error is sent to stdout
	if bytes.Contains(out, []byte(expectedErrorText)) {
		re := regexp.MustCompile(expectedErrorText + `([\s\S]*$)`)
		errText := expectedErrorText + re.FindStringSubmatch(string(out))[1]
		return nil, errors.New(errText)
	}
	firstInstruction := []byte("(000) ")
	if bytes.Contains(out, firstInstruction) {
		return out, nil
	} else {
		logger.Error.Println("Problems reading instruction numbers. stdout: " + string(out) + " stderr: " + err.Error())
		return nil, errors.New("Problems reading instruction numbers")
	}
}

// CSS file is also saved locally at checkbpf.css
func HomePage(w http.ResponseWriter, r *http.Request) {
	html, fileErr := ioutil.ReadFile("./checkbpfnet/checkbpfnet.html")
	if fileErr != nil {
	    logger.Error.Println("Unable to open local html file, checkbpfnet.html", fileErr.Error())
	}
	fmt.Fprintf(w, string(html))
}

func getFilter(r *http.Request) (string, error) {
	logger.Info.Println(fmt.Sprintf("%+v", r))
	urlQuery := fmt.Sprintf("%+v", r.URL)
	filter := strings.SplitN(urlQuery, "/", 3)[2]
	unescapedFilter, queryErr := url.QueryUnescape(filter)
	if queryErr != nil {
		return "", errors.New("invalid url query")
	}
	// For curl --data-urlencode q='capture filter'
	queryRe := regexp.MustCompile(`^\?\w+=`)
	sanitizedFilter := queryRe.ReplaceAllString(unescapedFilter, "")
	return sanitizedFilter, nil
}

/* Get the BPF status for a string */
func Status(w http.ResponseWriter, r *http.Request) {
	filter, queryErr := getFilter(r)
	var response string
	if queryErr != nil {
		response = `QUERY ERROR\n` + queryErr.Error()
		return
	} else {
		stdout, stderr := callDumpcap(filter)
		var errorText string
		if stderr != nil {
			errorText = stderr.Error()
			regex := regexp.MustCompile(`\((.*?)\)`)
			matches := regex.FindStringSubmatch(errorText)
			// Use the parenthetical from dumpcap if available
			if len(matches) >= 2 {
				errorText = matches[1]
			}
		} else {
			errorText = ""
		}
		success := len(stdout) > 0 && len(errorText) == 0
		successText := fmt.Sprintf("%t", success)
		response = `{"valid":` + successText + `,` +
			`"error":"` + errorText + `"}`
	}
	w.Write([]byte(response))
}

/* Get BPF code for a string */
func Code(w http.ResponseWriter, r *http.Request) {
	filter, queryErr := getFilter(r)
	var response string
	if queryErr != nil {
		response = `QUERY ERROR\n` + queryErr.Error()
	} else {
		dumpcapOut, dumpcapErr := callDumpcap(filter)
		if dumpcapErr != nil {
			response = dumpcapErr.Error()
		} else {
			response = "(000)" + strings.SplitN(string(dumpcapOut), "(000)", 2)[1]
		}
	}
	w.Write([]byte(response + "\n"))
}
