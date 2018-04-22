package main
type Payload struct {
    Rrsets `json:"rrsets"`
}
type Rrsets []struct {
    Name       string `json:"name"`
    Type       string `json:"type"`
    TTL        int    `json:"ttl"`
    Changetype string `json:"changetype"`
    Records `json:"records"`
}
type Records []struct {
    Content  string `json:"content"`
    Disabled bool `json:"disabled"`
}
type Getstruct struct {
    Account        string        `json:"account"`
    Dnssec         bool          `json:"dnssec"`
    ID             string        `json:"id"`
    Kind           string        `json:"kind"`
    LastCheck      int           `json:"last_check"`
    Masters        []interface{} `json:"masters"`
    Name           string        `json:"name"`
    NotifiedSerial int           `json:"notified_serial"`
    Serial         int           `json:"serial"`
    URL            string        `json:"url"`
}
func makePayload(name,ptype,ctype,content string ,ttl int, disable bool) Payload{
    records := Records{
        {content,disable},
    }
    rrsets := Rrsets{
        {name,ptype,ttl,ctype,records},
    }
    return Payload{rrsets}
}
