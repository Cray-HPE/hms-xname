# hms-xname

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Cray-HPE/hms-xname)

This library contains the types, enumerations, structures and functions for interacting with xnames.

* [xnametypes](#xnametypes)
    * [Adding a new xname type](#adding-a-new-xname-type)
* [xnames](#xnames)
    * [Working with the xnames package](#working-with-the-xnames-package)
    * [Code generation](#code-generation)
* [Viewing package documentation locally](#viewing-package-documentation-locally)

<a name="xnametypes"></a>

## xnametypes
The xnametypes package enumerates over all of the support xname component types. Also provides functions for manipulating and inspecting the string representation of xnames.

<a name="adding-a-new-xname-type"></a>

### Adding a new xname type
1.  First add the new HMSType string like the following to [./xnametypes/hmstypes.go](./xnametypes/hmstypes.go):
    ```go
    const (
        ...
        ChassisBMC HMSType = "ChassisBMC" // xXcCbB
        ...
    )
    ```

2.  Add a new entry to the HMS Component recognition table in [./xnametypes/hmstypes.go](./xnametypes/hmstypes.go):
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

3.  Add or update unit tests in [./xnametypes/hmstypes_test.go](./xnametypes/hmstypes_test.go)

4.  Regenerate code in xnames package

<a name="xnames"></a>

## xnames
The xnames (xname **s**tructure/**s**tringify) package allows xnames to be manipulated via Go Structures, instead of using `fmt.Sprintf` or regular expressions to manipulate the string representation of xnames. This package also contains facilities to easily convert to and from the string representation of a xname to xnames structures.

<a name="working-with-the-xnames-package"></a>

### Working with the xnames package
There are three supported ways to create a xnames structure:
1.  From xname string:
    ```go
    nodeRaw, xnameType := xnames.FromString("x1000c1s7b1n0")
    node, ok := nodeRaw.(xnames.Node)
    ```

2.  Using dot notation:
    ```go
    node = xnames.System{}.
        Cabinet(1000).
        Chassis(1).
        ComputeModule(7).
        NodeBMC(1).
        Node(0)
    ```

3.  Using struct literals:
    ```go
    node := xnames.Node{
        Cabinet:       1000,
        Chassis:       1,
        ComputeModule: 7,
        NodeBMC:       1,
        Node:          0,
    }
    ```

Get the xname string representation from a xnames structure:
```go
nodeXname := node.String()
```

The integer parts of a xname are easily accessible:
-   Accessing the value of field:
    ```go
    node = xnames.System{}.
        Cabinet(1000).
        Chassis(1).
        ComputeModule(7).
        NodeBMC(1).
        Node(0)

    // Get the integer ordinal for the Cabinet that this node is in:
    fmt.Println(n.Cabinet)
    // Output: 1000

    // Get the integer ordinal for the ComputeModule/Slot that this node is in:
    fmt.Println(n.ComputeModule)
    // Output: 7
    ```
-   Altering a xname by updating a field:
    ```go
    node = xnames.System{}.
            Cabinet(1000).
            Chassis(1).
            ComputeModule(7).
            NodeBMC(1).
            Node(0)

    // The original xname:
    fmt.Println(n)
    // Output: x1000c1s7b1n0

    // Change the cabinet to 1001
    node.Cabinet = 1001
    fmt.Println(n)
    // Output: x1001c1s7b1n0    
    ```

Navigating the xname hierarchy:
-   Each xnames structure has functions to build xnames of its children. In the example below there is a NodeBMC at x100c1s7b1, and we build 2 Node structures for the xnames x1000c1s7b1n0 and x1000c1s7b1n1:  
    ```go
    nodeBMC := xnames.System{}.
        Cabinet(1000).
        Chassis(1).
        ComputeModule(7).
        NodeBMC(1)
    node1 := nodeBMC.Node(0)
    node2 := nodeBMC.Node(0)
    ```

-   Access the parent xnames structure to go one level up in the xname hierarchy:
    ```go
    node = xnames.System{}.
        Cabinet(1000).
        Chassis(1).
        ComputeModule(7).
        NodeBMC(1).
        Node(0)
    nodeBMC := node.Parent()
    ```
    > Note: All xnames structures except for System have a `Parent()` function.


Simple xname validation:
```go
err := node.Validate()
```
> Under the hood this will validate the string representation of this structure against `xnametypes.IsHMSCompIDValid()`.

<a name="code-generation"></a>

### Code generation
The xnames package contains 2 files that are generated. These files are generated from the contents of the `hmsCompRecognitionTable` in [./xnametypes/hmstypes.go](./xnametypes/hmstypes.go):
- [./xnames/types.go](./xnames/types.go) is generated from [./xnames/generator/types.go.tpl](./xnames/generator/types.go.tpl)
- [./xnames/util.go](./xnames/util.go) is generated from [./xnames/generator/util.go.tpl](./xnames/generator/util.go.tpl)

To generate code for xnames package simply run: 
```bash
$ make generate
```

<a name="viewing-package-documentation-locally"></a>

## Viewing package documentation locally
> Alternatively package documentation can be viewed online [here](http://pkg.go.dev/github.com/Cray-HPE/hms-xname). 
1.  Install godoc:
    ```bash
    $ go install golang.org/x/tools/cmd/godoc
    ```

2.  Start the godoc HTTP server:
    ```bash
    $ godoc -http=:6060 
    ```

3.  Package documentation should now be available at [http://localhost:6060/pkg/github.com/Cray-HPE/hms-xname/](http://localhost:6060/pkg/github.com/Cray-HPE/hms-xname/).
