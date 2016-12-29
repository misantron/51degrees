package fiftyonedegrees

import (
	"testing"
)

func TestProviderParse(t *testing.T) {
	userAgentDB, err := mustLoadProviderData(2, 1000)
	if err != nil {
		t.Fatal("could not load DB: %s", err)
	}
	testUA := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36"
	userAgentDB.Parse(testUA)
	userAgentDB.Close()

}

/*
 * Benchmarking code starts here
 */

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
