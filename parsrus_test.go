package parsrus

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(httpJSONHandler)
	handler.ServeHTTP(w, req)
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
	if res.StatusCode != http.StatusOK {
		t.Errorf("HTTP code is not 200")
	}
}

func TestParseXML(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(httpXMLHandler)
	handler.ServeHTTP(w, req)
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

func TestParseJSONWithoutHTTPCode(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandlerNoStatusCode)
	handler.ServeHTTP(w, req)
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
	if res.StatusCode != http.StatusOK {
		t.Errorf("HTTP code is not 200")
	}
}


func httpJSONHandler(w http.ResponseWriter, r *http.Request) {
	p := Parser{ResponseWriter: w, ContentType: "json"}
	p.Parse(Fields{"code": "200", "message": "json parsed"}, 200)
	return
}


func httpXMLHandler(w http.ResponseWriter, r *http.Request) {
	p := Parser{ResponseWriter: w, ContentType: "xml", RootTag: "request"}
	p.Parse(Fields{"code": "200", "message": "json parsed"}, 200)
	return
}



func httpHandlerNoStatusCode(w http.ResponseWriter, r *http.Request) {
	p := Parser{ResponseWriter: w, ContentType: "json"}
	p.Parse(Fields{"code": "200", "message": "json parsed"})
	return
}