package fiftyonedegrees

import (
	"fmt"
)

func mustLoadProviderData(poolSize int, cacheSize int) (*FiftyoneDegreesProvider, error) {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	item, err := NewFiftyoneDegreesProvider("51Degrees-LiteV3.2.dat", properties, poolSize, cacheSize)
	if err != nil {
		fmt.Println("err=", err)
		return nil, err
	}

	return item, nil
}

func mustLoadDataSet() (*FiftyoneDegreesDataSet, error) {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	item, err := NewFiftyoneDegreesDataSet("51Degrees-LiteV3.2.dat", properties)
	if err != nil {
		fmt.Println("err=", err)
		return nil, err
	}

	return item, nil
}
