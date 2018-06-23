package fiftyonedegrees

import (
	"fmt"
)

var dataFilePath = `51Degrees-LiteV3.2.dat`

func mustLoadProviderData(poolSize int, cacheSize int) (*FiftyoneDegreesProvider, error) {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	item, err := NewFiftyoneDegreesProvider(dataFilePath, properties, poolSize, cacheSize)
	if err != nil {
		fmt.Println("err=", err)
		return nil, err
	}

	return item, nil
}

func mustLoadDataSet() (*FiftyoneDegreesDataSet, error) {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	item, err := NewFiftyoneDegreesDataSet(dataFilePath, properties)
	if err != nil {
		fmt.Println("err=", err)
		return nil, err
	}

	return item, nil
}
