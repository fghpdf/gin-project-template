## /internal/pkg/thunes

Here is a third-party package with [thunes](https://developers.thunes.com/money-transfer/v2/?go#introduction).

It is a sample for how to encapsulate a package for using HTTP.

Now there is an interface about listing countries.
 
Can use it by following ways:
```go
package main

import (
    "fghpdf.me/gin-project-template/internal/pkg/thunes/country"
    "fghpdf.me/gin-project-template/internal/pkg/thunes/httpClient"
    "log"
)

func main() {
    authClient := &httpClient.AuthClient{
    }
    thunesSvc := country.NewServer(authClient)
    countries, err := thunesSvc.List(nil)
	if err != nil {
		panic(err)
	}

    log.Println(countries)
}
``` 