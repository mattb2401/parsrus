package parsrus

import (
	"encoding/json"
	"net/http"

	"github.com/clbanning/anyxml"
)

// Parser Initializer
type Parser struct {
	ResponseWriter http.ResponseWriter
	ContentType    string
	RootTag        string
}

// Fields to be rendered
type Fields map[string]interface{}

func (r *Parser) toJSON(p interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *Parser) toXML(p interface{}) ([]byte, error) {
	x, err := anyxml.XmlIndent(p, "", " ", r.RootTag)
	if err != nil {
		return nil, err
	}
	return x, nil
}

// Parse content
func (r *Parser) Parse(fields Fields, httpCode ...int) {
	var bt []byte
	if r.ContentType == "json" {
		bt, _ = r.toJSON(fields)
		r.ResponseWriter.Header().Set("Content-Type", "application/json")
	} else if r.ContentType == "xml" {
		bt, _ = r.toXML(fields)
		r.ResponseWriter.Header().Set("Content-Type", "application/xml")
	}
	if len(httpCode) > 0 {
		r.ResponseWriter.WriteHeader(httpCode[0])
	}
	r.ResponseWriter.Write(bt)

}
