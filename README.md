#Parsrus

A simple json and xml parser for golang web apis. 

**Usage**

To install Parsrus
```bash
    go get github.com/mattb2401/parsrus
```

To add Parsrus in your code 

```go
    package main
    
    import (
        "net/http"
        "github.com/mattb2401/parsrus"
    )

    // Return json request
    func JSONRequestHandler(w http.ResponseWriter, r *http.Request){
        p := parsrus.Parser{ResponseWriter: w, ContentType: "json"}
        p.Parse(parsrus.Feilds{"code": "200", "message": "json parsed"})
        return
    }

    // Return xml request. Add RootTag for root element in your xml when initializing parsrus struct. 
    func XMLRequestHandler(w http.ResponseWriter, r *http.Request){
        p := parsrus.Parser{ResponseWriter: w, ContentType: "xml", RootTag: "request"}
        p.Parse(parsrus.Feilds{"code": "200", "message": "xml parsed"})
        return
    }
```

