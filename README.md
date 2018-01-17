# sfgo

Provides a simple golang package to allow accessing the SolidFire API via
golang!  This package is intended to be as light weight as possible, provide
the basics and stay out of the users way.

We do NOT implement all of the API calls here, just the most commonly used
ones, however the provider is designed to be extendable such that a user can
add any api request/response combination that they like.  In addition, PR's are
always welcome for additional request or response types that a user might like
to add to the types package.

## What's here
There are two packages included in this repo:
1. sfgo/pkg/types
2. sfgo/pkg/provider

## types
The types package provides data structures for SolidFire resource objects (ie
Volume) as well as request and response data structures for the most common API
methods.  The SolidFire API is implemented using json-rpc, so these data types
include the needed json encode/decode tags.

## provider
The provider is a package to ease issuing an API request to a SolidFire
cluster.  We don't require the use of the types package to utilize the
provider, but it's strongly recommended to make your life simpler.  If you use
the types library you can simply use the predefined types for your
request/response encode and decode data.  Or, if you want to write your own
that's fine too.  We take the request parameters and the decode parameters as
generic interface types, so you're free to define your own.

## example usage

```golang
package main

import (
    "encoding/json"
    "fmt"

    "github.com/j-griffith/sfgo/pkg/provider"
    "github.com/j-griffith/sfgo/pkg/types"
    "github.com/prometheus/common/log"
)

func main() {
    // Using the types package makes this easy, but if there's a request that's
    // not available, it's also pretty simple to see the pattern here and create
    // your own req/resp types.
    req := types.ListActiveVolumesRequest{}
    resp := types.ListVolumesResult{}

    // There's a number of things we need to issue a request, so we wrap them
    // up into a nice neat struct to pass around
    apiReq := provider.Request{}
    apiReq.Name = "ListActiveVolumes"
    apiReq.URL = "https://admin:admin@10.117.36.101/json-rpc/9.0"
    apiReq.ID = provider.NewReqID()
    apiReq.Params = req

    // We use an interface for a transport for mocking, so make sure you
    // a real HTTP transport engine
    t := provider.HTTP{}

    
    body, err := provider.IssueRequest(apiReq, t)
    if err != nil {
        fmt.Printf("bummer due: %v", err)
    }

    // Use the providers decode helper or you could do it yourself if you want
    resp = provider.DecodeResponse(body, resp)

    // Or do it yourself if you prefer
    if err := json.Unmarshal([]byte(body), &resp); err != nil {
        fmt.Printf("error detected unmarshalling ListVolumesForAccount API response: %+v", err)
    }
}
```
