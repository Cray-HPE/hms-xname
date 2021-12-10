# hms-xname

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Cray-HPE/hms-xname)
[![Go Report Card](https://goreportcard.com/badge/github.com/Cray-HPE/hms-xname)](https://goreportcard.com/report/github.com/Cray-HPE/hms-xname)

TODO summary of this package

TODO view docs locally

1. Start godoc HTTP server
    ```
    $ godoc -http=:6060 
    ```

2.   


## xnametypes

TODO breif summary

### Adding a new xname type
1.  First add a new HMSType string like the following to [hmstypes.go](./xnametypes/hmstypes.go):
    ```go
    const (
        ChassisBMC HMSType = "ChassisBMC" // xXcCbB
    )
    ```

2.  Add a new entry to the HMS Component recognition table in [hmstypes.go](./xnametypes/hmstypes.go):
    ```go
    var hmsCompRecognitionTable = map[string]HMSCompRecognitionEntry{
        ...
        "chassisbmc": {  // Lowercase string of the corresponding xnametypes.HMSType
            ChassisBMC,  // Corresponding xnametypes.HMSType
            Chassis,     // Corresponding xnametypes.HMSType for the parent of this type 
            "xXcCbB",    // Example string where each ordinal is replaced a upper case character 
            regexp.MustCompile("^x([0-9]{1,4})c([0-7])b([0])$"), // Regular expression to determine if the xname is valid
            "x%dc%db%d", // Format string
            3,           // Number of verbs in the format string
        },
        ...
    }
    ```

3.  Add or update unit tests in [hmstypes_test.go](./xnametypes/hmstypes_test.go)

4.  Regenerate code in xnames package

## xnames
TODO 1 or 3 sentence summary...

### Rational
Since go does not support 

### Code generation
To only regenerate code for xnames package simply run: 
```bash
$ make generate
```
