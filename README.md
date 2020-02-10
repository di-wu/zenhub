# zenhub
[![GoDoc](https://godoc.org/github.com/di-wu/zenhub?status.svg)](https://godoc.org/github.com/di-wu/zenhub)

- client for accessing the zenhub api.
- webhook for easy receiving and parsing of zenhub webhook events

## usage
### client
```go
package main

import "github.com/di-wu/zenhub"

func main() {
    client, _ := zenhub.NewClient(zenhub.Options.Secret(token))
    issue, _, _ := client.GetIssue(repoID, issueNumber)
    // do something with issue data
}
```

`repoID` is the ID of the repository, not its full name.

### webhook
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

## tests
all the tests are currently running on an empty private repository.
```shell script
export ZENHUB_SECRET="zenhub-token"
export GITHUB_SECRET="github-token"
export TEST_REPO_OWNER="di-wu"
export TEST_REPO_NAME="test"

go test
```