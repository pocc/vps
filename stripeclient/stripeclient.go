package stripeclient

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"errors"
)

// Essentially what this needs to do is validate that there is an active subscription for a given product
func checkStripeAccess(email string, productName string) (bool, error) {
	productIDs := map[string]string{
		"fastlogin": "prod_Gh1Hc0RfcY8Qas",
		"multicap": "prod_Gkpu88DQGx2Rzi",
	}
	key, fileErr := ioutil.ReadFile("stripeclient/stripe.key")
	if fileErr != nil {
	    return false, errors.New("Unable to access required resources")
	}
	stripe.Key = string(key)
	params := &stripe.CustomerListParams{}
	CustList := customer.List(params)
	for CustList.Next() {
		c := CustList.Customer()
		if strings.ToLower(c.Email) == strings.ToLower(email) {
			for _, s := range c.Subscriptions.Data {
				productID := s.Plan.Product.ID
				correctProductID := productID == productIDs[productName]
				nowTs := unixTs()
				planIsActive := s.CurrentPeriodEnd > nowTs
				if correctProductID && planIsActive {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func unixTs() int64 {
	return time.Now().Unix()
}

/* Create an int hash of a string by adding a timestamp */
func Hash(input string) int32 {
  hash := int32(0);
  if len(input) == 0 {
      return hash
  }
  inputBytes := []byte(input)
  for _, char := range inputBytes {
    hash  = (hash << 5) - hash + int32(char)
    hash |= 0
  }
  if hash < 0 { // no negative hashes
    hash = -hash
  }
  return hash
}

/* CheckAccess will check the stripe APIs for access and return a response
   
   It doesn't matter what the receivedHash is - the sending end will check that it matches the hash
   Test from a remote computer with (provided test@example.com has a valid subscription):
   $ curl -L -H "Content-Type: application/json" -d '{"email":"test@example.com","hash":"randomtext","requester":"fastlogin"}' -X POST https://fastlog.in/auth/ && echo
   {"email":"test@example.com","product":"fastlogin","hash":"pwx278","error":""}
*/
func CheckAccess(email string, productName string, receivedHash string) []byte {
	errorStr := ""
	hasAccess, stripeErr := checkStripeAccess(email, productName)
	hashInt := Hash(email + productName + strconv.FormatBool(hasAccess) + receivedHash)
	hash := strconv.FormatInt(int64(hashInt), 36)
	if stripeErr != nil {
	    errorStr = stripeErr.Error()
	    hash = ""
	}
	retJson := []byte(`{"email":"` + email + `",` +
	    `"product":"` + productName + `",` +
            `"hash":"` + hash + `",` + // base 36 [0-9a-z]
	    `"error":"` + errorStr + `"}`)
	return retJson
}
