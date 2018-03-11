package gohpalm

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
	StatusColor      string      `json:"statuscolor"`
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
