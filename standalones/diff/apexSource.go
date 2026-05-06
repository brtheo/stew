package diff

import "encoding/json"

func UnmarshalApexSource(data []byte) (ApexSource, error) {
	var r ApexSource
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ApexSource) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ApexSource struct {
	Status   int64         `json:"status"`
	Result   ApexResult        `json:"result"`
	Warnings []interface{} `json:"warnings"`
}

type ApexResult struct {
	Records   []ApexRecord `json:"records"`
	TotalSize int64    `json:"totalSize"`
	Done      bool     `json:"done"`
}

type ApexRecord struct {
	Attributes ApexAttributes `json:"attributes"`
	Body       string     `json:"Body"`
}

type ApexAttributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
