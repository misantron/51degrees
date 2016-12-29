package fiftyonedegrees

import (
	"fmt"
	"testing"
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

func BenchmarkProviderParse(b *testing.B) {
	userAgentDB, err := mustLoadProviderData(2, 1000)
	if err != nil {
		b.Fatal("could not load DB: %s", err)
	}

	defer userAgentDB.Close()

	testUA := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userAgentDB.Parse(testUA)
	}
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

func BenchmarkDataSetParse(b *testing.B) {
	userAgentDB, err := mustLoadDataSet()
	if err != nil {
		b.Fatal("could not load DB: %s", err)
	}

	defer userAgentDB.Close()

	testUA := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userAgentDB.Parse(testUA)
	}
}
