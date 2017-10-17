# GoHPalm []()

Go package to use HP ALM Rest API

## Install

```bash
go get github.com/nelsonjvf/gohpalm
```

## Usage and Examples

First of all we need to configure and set your HP ALM information. For that we can use our config.yaml file as example and the following init function:

```go
func init() {
	// Use yaml configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &gojira.Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
```

After that you can simply use the methods available to interact with HP ALM:

```go
// Get defect information
gohpalm.RequestIssue()
```

Here is the full code to test it easily:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"github.com/nelsonjvf/gohpalm"
)

func init() {
	// Use yaml configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &gojira.Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func main() {
  fmt.Println("Starting test..")
  itemId := "1234"

  fmt.Println("Calling HP ALM:")

  tmtItemInfo := gohpalm.RequestIssue("CONFIG_LABEL", "DOMAIN", "PROJECT", itemId)
  fmt.Println(tmtItemInfo)
}
```

Don't forget your yaml configuration file

```yaml
tmt:
  - lable: Project
    user: username
    pass: password
    url: https://url:port/
```

### GoHPalm methods

We can get an issue information:

```RequestIssue(label, domain, project, issueId)```

## Credits

 * [Nelson Ferreira](https://github.com/nelsonjvf)

## License

The MIT License (MIT) - see LICENSE.md for more details
