# Go-aib
Go auto-increment-build is a simple command line tool to increment the patch part of your version number so you  just have to think about major and minor parts.
It will parse a source file to find and increment a patch var/const, it will also update a timestamp var/const if you need it.
Go-aib shoud be called before each build time.

## Installing
```
go get github.com/gandaldf/go-aib
go install github.com/gandaldf/go-aib
```

## Options
```
-fn [string]
	Go source file name (default "version.go")

-pn [string]
	Patch var/const name (default "versionPatch")

-tf [string]
	Timestamp output format (default "060102150405")

-tn [string]
	Timestamp var/const name (default "versionTimestamp")
```

## Usage
Before:
```
// version.go

package main

const versionMajor = "1"
const versionMinor = "0"
const versionPatch = "0"
const versionTimestamp = ""

const version = versionMajor + "." + versionMinor + "." + versionPatch
```
Command:

```
./go-aib -fn version.go
```

After:
```
// version.go

package main

const versionMajor = "1"
const versionMinor = "0"
const versionPatch = "1"
const versionTimestamp = "201005155147"

const version = versionMajor + "." + versionMinor + "." + versionPatch
```