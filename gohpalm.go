package gohpalm

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// DefaultTimeout is the default http call timeout time in seconds
var DefaultTimeout int

// HTTPRequest represents the generic HTTP caller for HP Alm
func HTTPRequest(URL string, URLPath string, Username string, Password string) ([]byte, error) {

	var result []byte

	if len(URL) == 0 {
		return result, errors.New("HP ALM URL is missing")
	}

	// Prepare HTTP Client
	timeoutVal := time.Duration(10 * time.Second)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   timeoutVal,
		Transport: tr,
	}

	// Get Cookie info
	urlCookie := fmt.Sprintf("%s%s", URL, "qcbin/api/authentication/sign-in")

	reqCookie, errCookie := http.NewRequest("GET", urlCookie, nil)
	if errCookie != nil {
		log.Printf("http.NewRequest err   #%v ", errCookie)
	}

	reqCookie.SetBasicAuth(Username, Password)

	respCookie, errDo := client.Do(reqCookie)
	if errDo != nil {
		log.Printf("http.DefaultClient.Do err   #%v ", errDo)
		return result, errDo
	}
	defer respCookie.Body.Close()

	// Request
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", URL, URLPath), nil)

	for _, cookie := range respCookie.Cookies() {
		cookie := http.Cookie{Name: cookie.Name, Value: cookie.Value}
		r.AddCookie(&cookie)
	}

	r.SetBasicAuth(Username, Password)
	r.Header.Add("Content-Type", "application/json")

	resp, errDo := client.Do(r)
	if errDo != nil {
		log.Printf("http.DefaultClient.Do err   #%v ", errDo)
		return []byte(""), errDo
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll err   #%v ", err)
	}

	return body, nil
}

// RequestIssue represents the generic method to request an issue to HP ALM
func RequestIssue(URL string, Username string, Password string, project string, domain string, item string) (HpAlmDefectResponse, error) {
	var urlIssuePath string
	var data HpAlmDefectResponse

	urlIssuePath = fmt.Sprintf("qcbin/api/domains/%s/projects/%s/defects/%s", project, domain, item)

	resp, err := HTTPRequest(URL, urlIssuePath, Username, Password)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
