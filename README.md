# zenhub
[![GoDoc](https://godoc.org/github.com/di-wu/zenhub?status.svg)](https://godoc.org/github.com/di-wu/zenhub)

allows for easy receiving and parsing of zenhub webhook events

## usage
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/di-wu/zenhub"
)

func main() {
    hook := new(zenhub.Webhook)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        payload, err := hook.Parse(r, zenhub.IssueTransfer)
        if err != nil {
           if err != zenhub.ErrEventNotFound {
                // event was not the one asked to be parsed
            }
        }
        switch payload.(type) {
        case zenhub.IssueTransferEvent:
            transfer := payload.(zenhub.IssueTransferEvent)
            // do something with the event data
            fmt.Printf("%+v", transfer)
        }
    })

    http.ListenAndServe(":3000", nil)
}
```