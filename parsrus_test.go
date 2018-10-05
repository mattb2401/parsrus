package parsrus

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestParseJSON(t *testing.T) {
	w := httptest.NewRecorder()
	p := Parser{ResponseWriter: w, ContentType: "json"}
	p.Parse(Fields{"code": "200", "message": "json parsed"})
	res := w.Result()
	bt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error occurred %v", err)
	}
	var r map[string]interface{}
	err = json.Unmarshal(bt, &r)
	if err != nil {
		t.Errorf("Error occurred %v", err)
	}
	if r["code"] != "200" {
		t.Errorf("Code is not 200")
	}
}

func TestParseXML(t *testing.T) {
	w := httptest.NewRecorder()
	p := Parser{ResponseWriter: w, ContentType: "xml", RootTag: "request"}
	p.Parse(Fields{"code": "200", "message": "xml parsed"})
	res := w.Result()
	bt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error occurred %v", err)
	}
	type xmlData struct {
		XMLName xml.Name `xml:"request"`
		Code    string   `xml:"code"`
		Message string   `xml:"message"`
	}
	var x xmlData
	err = xml.Unmarshal(bt, &x)
	if err != nil {
		t.Errorf("Error occurred %v", err)
	}
	if x.Code != "200" {
		t.Errorf("Code is not 200")
	}
	return
}
