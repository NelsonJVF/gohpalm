package gohpalm

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto/tls"
	"time"
)

// Configuration represents a struct for TMT access information
type Configuration struct {
	Lable string `yaml:"lable"` // Some projects have more than one TMT instance, so just lable as you wish
	User string `yaml:"user"` // TMT Username
	Pass string `yaml:"pass"` // TMT User Password
	Url string `yaml:"url"`	// TMT URL + Port
}

var Config []Configuration

// HpAlmDefectResponse represents a struct for TMT defect response
type HpAlmDefectResponse struct {
	Type         string      `json:"type"`
	Subject      interface{} `json:"subject"`
	HasLinkage   string      `json:"has-linkage"`
	CycleID      interface{} `json:"cycle-id"`
	CreationTime string      `json:"creation-time"`
	ID           int         `json:"id"`
	RequestNote  interface{} `json:"request-note"`
	RunReference interface{} `json:"run-reference"`
	ToMail       interface{} `json:"to-mail"`
	VerStamp     int         `json:"ver-stamp"`
	User30       interface{} `json:"user-30"`
	RequestID    interface{} `json:"request-id"`
	Priority     string      `json:"priority"`
	TargetRel    struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"target-rel"`
	TestReference     interface{} `json:"test-reference"`
	ExtendedReference interface{} `json:"extended-reference"`
	DevComments       string      `json:"dev-comments"`
	Name              string      `json:"name"`
	User22            string      `json:"user-22"`
	User23            string      `json:"user-23"`
	User26            interface{} `json:"user-26"`
	DetectedInRel     struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"detected-in-rel"`
	ClosingDate      string      `json:"closing-date"`
	Status           string      `json:"status"`
	StatusColor   	 string      `json:"statuscolor"`
	Description      string      `json:"description"`
	LastModified     string      `json:"last-modified"`
	User17           string      `json:"user-17"`
	HasOthersLinkage string      `json:"has-others-linkage"`
	Attachment       interface{} `json:"attachment"`
	RequestType      interface{} `json:"request-type"`
	User11           interface{} `json:"user-11"`
	User10           string      `json:"user-10"`
	User16           interface{} `json:"user-16"`
	RequestServer    interface{} `json:"request-server"`
	Owner            string      `json:"owner"`
	Severity         string      `json:"severity"`
	DetectedBy       string      `json:"detected-by"`
	StepReference    interface{} `json:"step-reference"`
	CycleReference   interface{} `json:"cycle-reference"`
	DetectedInRcyc   interface{} `json:"detected-in-rcyc"`
	HasChange        interface{} `json:"has-change"`
	User04           string      `json:"user-04"`
	User02           interface{} `json:"user-02"`
	User03           string      `json:"user-03"`
}

// HTTPRequest represents the generic HTTP caller for HP Alm
func HTTPRequest(hpAlmLable string, urlPath string) []byte {
	var username string
	var password string
	var url string

	for _, c := range Config {
		if c.Lable == hpAlmLable {
			username = c.User
			password = c.Pass
			url = c.Url
		}
	}

	if(len(url) == 0) {
		log.Printf(" ---------- HP ALM configuration is missing  ---------- ")
		log.Printf("\t For lable " + hpAlmLable)
		return nil
	}

	/*
		Prepare HTTP Client
	 */
	timeoutVal := time.Duration(10 * time.Second)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}

	client := &http.Client{
		Timeout: timeoutVal,
		Transport: tr,
	}

	/*
		Get Cookie info
	 */
	urlCookie := fmt.Sprintf("%s%s", url, "qcbin/api/authentication/sign-in")

	reqCookie, errCookie := http.NewRequest("GET", urlCookie, nil)
	if errCookie != nil {
		log.Printf("http.NewRequest err   #%v ", errCookie)
	}

	reqCookie.SetBasicAuth(username, password)

	respCookie, errDo := client.Do(reqCookie)
	if errDo != nil {
		log.Printf("http.DefaultClient.Do err   #%v ", errDo)
		return []byte("")//, errDo
	}
	defer respCookie.Body.Close()

	/*
		Request
	 */
	url = fmt.Sprintf("%s%s", url, urlPath)

	r, _ := http.NewRequest("GET", url, nil)

	for _, cookie := range respCookie.Cookies() {
		cookie := http.Cookie{Name: cookie.Name, Value: cookie.Value}
		r.AddCookie(&cookie)
	}

	r.SetBasicAuth(username, password)
	r.Header.Add("Content-Type", "application/json")

	resp, errDo := client.Do(r)
	if errDo != nil {
		log.Printf("http.DefaultClient.Do err   #%v ", errDo)
		return []byte("")//, errDo
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll err   #%v ", err)
	}

	return body
}

// RequestIssue represents the generic method to request an issue to HP ALM
func RequestIssue(hpAlmLable string, project string, domain string, item string) HpAlmDefectResponse {
	var urlIssuePath string
	var data HpAlmDefectResponse

	urlIssuePath = fmt.Sprintf("qcbin/api/domains/%s/projects/%s/defects/%s", project, domain, item)

	response := HTTPRequest(hpAlmLable, urlIssuePath)

	err := json.Unmarshal(response, &data)
	if err != nil {
		log.Printf("json.Unmarshal err   #%v ", err)
	}

	return data
}
