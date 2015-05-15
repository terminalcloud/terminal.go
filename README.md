# terminal.go

![1](http://i.imgur.com/UkOdnUW.png)
#### Terminal.com API Go bindings
---

## What is this?

Terminal.go or 'terminalgo' is a Golang package to help you to interact with the Terminal.com API, directly from your GO application.

## Installation:
Your can install this package by downloading and building the source code from this repository.

You can also use the `go get` command to download and install it automatically:

`# go get github.com/terminalcloud/terminal.go/terminalgo`

## Testing:
You can test all package functions by executing `go test -v` in the source directory.
This requires a *json* format file, with your credentials on it, as in the example in below:

```
{"access_token": "myaccesstokenstring", "user_token": "mylongusertokenstring"}
```

## Basic Usage:
Import the package, by default located at "github.com/terminalcloud/terminal.go/terminalgo" in your src directory.

### General Usage Notes:
- Each function will return a struct type, which is always declared in the **structs.go** file, with the *'Output_'* prefix.
- You can always obtain the RAW json response from the server in a string chain, by using the functions with the *'_RAW'* sufix.
- All functions requires only strings as inputs, except for the 'Add_Terminal_Links' and 'Remove_Terminal_Links' functions. To send the links use the 'UInput_Links' type.

---

#### The logo shown on this page is based on the Renee French work, originally licensed under the Creative Commons 3.0 Attributions license.