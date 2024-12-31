## About this module
This is a banking application that handles the following functionalities:
- Check account balance
- Deposit amount
- Withdraw amount

## Custom Packages
Creating custom packages in Go is a common task when you want to modularize your code. Here are some best practices 
and general rules:

1. __Naming Conventions__: Use lowercase, single-word names or multi-word names joined by underscores. Avoid using names 
that might conflict with standard library packages.

2. __Keep It Simple__: A package should perform a single functionality. Following the single responsibility principle 
allows the package to be reusable, maintainable, and testable.

3. __Documentation__: Every package should have a doc.go file where you document what the package does. 
This documentation will be displayed on GoDoc.

4. __Exports__: Anything that starts with a capital letter is exported (public), and anything that starts with a 
lowercase letter is not exported (private). Be mindful of what you want to make accessible to other packages.

5. __Tests__: Include unit tests in your packages. A *_test.go file is conventionally used for testing the 
corresponding *_go file.

6. __Versioning__: Consider using semantic versioning for your packages. This will help others and your future self to 
understand which version to use.

7. __Organize By Responsibility__: Structuring Go packages by responsibility and not by type is a good practice. 
This means handlers, middlewares, models and others should not be separate packages. Instead, consider using functional
packages like users, products, orders, payments and so on.

Here's an example of a package structure:
```go
package main

import (
    "mypackage"
)

func main() {
    mypackage.MyFunction()
}
```
In this example, "mypackage" is a custom package and "MyFunction" is an exported function that can be used in other 
packages. It's best to place your packages in the /pkg or /internal directories, depending on their intended use.


## Package versioning
Package versioning is a way to manage different stages of a package's development. In Go, module versioning is 
accomplished by appending a version number to the end of the module path.

Here's an example of how you might specify versions of a module in your go.mod file:
```go
module myproject

go 1.16

require (
	github.com/myuser/mylib/v4 v4.0.0
	github.com/myuser/otherlib/v2 v2.3.1
	...
)
```
In this example, mylib is at version v4.0.0 and otherlib is at version v2.3.1. The v4 and v2 in the module path 
correspond to the major version of the module.

You need to note that starting from v2 and onward, the major version of the module must be included as /vN at the end 
of the module paths used in go.mod files (e.g., github.com/myuser/mylib/v2), and in the package import path.

When you import this package into your Go code, you would refer to it like so:

```go
import (
	"github.com/myuser/mylib/v4"
	"github.com/myuser/otherlib/v2"
)
```
And use it like:

```go
value := mylib.MyFunction()
otherValue := otherlib.SomeFunction()
```

his is in compliance with Semantic Versioning (SemVer), a versioning system that reflects the level of changes in a
release. SemVer versions have three components: major, minor, and patch (e.g., v4.0.0). When you make incompatible API 
changes, you increase the major version. When you add functionality in a backwards-compatible manner, you increase the 
minor version. For backwards-compatible bug fixes and miscellaneous changes, you increase the patch version.




## Using 3rd-party package
You will need to install the 3rd-party package before using it.

```shell
go get -v github.com/Pallinder/go-randomdata
```

### Usage
```go

package main

import (
    "fmt"
    "github.com/Pallinder/go-randomdata"
)

func main() {
    // Print a random silly name
    fmt.Println(randomdata.SillyName())

    // Print a male first name
    fmt.Println(randomdata.FirstName(randomdata.Male))

    // Print a female first name
    fmt.Println(randomdata.FirstName(randomdata.Female))

    // Print a last name
    fmt.Println(randomdata.LastName())

    // Print an email
    fmt.Println(randomdata.Email())
	
}

```

