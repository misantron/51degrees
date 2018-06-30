# 51Degrees Go provider

[![Build Status](https://img.shields.io/travis/misantron/51degrees.svg?style=flat-square&longCache=true)](https://travis-ci.org/misantron/pg-51degrees)

51Degrees browser user agent parser and mobile device detection.

## Require
51Degrees database - v3.2

## Installation
`go get github.com/misantron/51degrees`

## Usage example
```
package main

import (
	"fmt"
	"github.com/simplereach/51degrees"
)

func main() {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	testUA := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36"
	item, err := NewFiftyoneDegrees("51Degrees-LiteV3.2.dat", properties)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(item.Parse(testUA))
	item.Close()
}
```

## Tests run
`go test -v ./...`