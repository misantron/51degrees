#51Degrees for golang
51Degrees UserAgent解析库的官方c语言包装版
只实现了我需要的功能，是线程安全的

### 安装
`go get github.com/hdczsf/51degrees`

### 例子
```go
package main

import (
	"fmt"
	"github.com/hdczsf/51degrees"
)

func main() {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	testUA := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36"
	item, err := NewFiftyoneDegrees("51Degrees-EnterpriseV3_1.dat", properties)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(item.Parse(testUA))
	item.Close()
}

```
