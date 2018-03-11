# GoHPalm [![Codacy Badge](https://api.codacy.com/project/badge/Grade/ad8cdb159a554616a66f99516ff0a10a)](https://www.codacy.com/app/NelsonJVF/gohpalm?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=NelsonJVF/gohpalm&amp;utm_campaign=Badge_Grade)

Go package to use HP ALM Rest API

## Install

```bash
go get github.com/nelsonjvf/gohpalm
```

## Usage and Examples

Simply use the methods available to interact with HP ALM:

```go
// Get defect information
gohpalm.RequestIssue()
```

Here is the full code to test it easily:

```go
package main

import (
	"log"
	
	"github.com/nelsonjvf/gohpalm"
)

func main() {

  	log.Println("Starting test..")

	itemId := "1234"

  	log.Println("Calling HP ALM:")

  	tmtItemInfo := gohpalm.RequestIssue("http://your-tmt.com/", "usernam", "password", "project", "domain", itemId)

	log.Println(tmtItemInfo)

}
```

### GoHPalm methods

We can get an issue information:

```RequestIssue(url, user, pass, project, domain, issueId)```

## Credits

 * [Nelson Ferreira](https://github.com/nelsonjvf)

## License

The MIT License (MIT) - see LICENSE.md for more details
