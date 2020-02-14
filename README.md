***Parsrus***

Return json or xml for golang web apis. 

**Usage**

To install Parsrus
```bash
    go get github.com/mattb2401/parsrus
```

To add Parsrus in your code 

```go

    import (
        "net/http"
        "github.com/mattb2401/parsrus"
    )

    // Return json request
    func JSONRequestHandler(w http.ResponseWriter, r *http.Request){
        // parsrus.Parser struct fields {ResponseWriter, ContentType}
        p := parsrus.Parser{ResponseWriter: w, ContentType: "json"}
        // p.Parse() takes a map[string]interface{} 
        p.Parse(parsrus.Feilds{"code": "200", "message": "json parsed"})
        return
    }

    // Return xml request. Add RootTag for root element in your xml when initializing parsrus struct. 
    func XMLRequestHandler(w http.ResponseWriter, r *http.Request){
         // parsrus.Parser struct fields {ResponseWriter, ContentType, RootTag}
        p := parsrus.Parser{ResponseWriter: w, ContentType: "xml", RootTag: "request"}
         // parsrus.Parser struct fields {ResponseWriter, ContentType}
        p.Parse(parsrus.Feilds{"code": "200", "message": "xml parsed"})
        return
    }
    // Return json request with specified http code
    func JSONRequestHandlerWithStatusCode(w http.ResponseWriter, r *http.Request) {
        // parsrus.Parser struct fields {ResponseWriter, ContentType}
        p := parsrus.Parser{ResponseWriter: w, ContentType: "json"}
        // p.Parse() takes a map[string]interface{} of parsrus.Fields and http code
        p.Parse(parsrus.Feilds{"code": "200", "message": "json parsed"}, 200)
        return
    }

    //You can serialize an interface to JSON or XML using the Serialize functon
    func JSONRequestHandlerWithStatusCode(w http.ResponseWriter, r *http.Request) {
        // parsrus.Parser struct fields {ResponseWriter, ContentType}
        p := parsrus.Parser{ResponseWriter: w, ContentType: "json"}
        // p.Serialize() takes any map[string]interface{} and http code
        p.Serialize(map[string]interface{}{"data": "boom"}, 200)
        return
    }

