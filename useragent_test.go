package useragent

import "testing"

type test struct {
	ua          string
	expected    UserAgent
	expectedStr string
}

var userAgents = [...]test{
	{
		ua: "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.815.0 Safari/535.1",
		expected: UserAgent{
			OS:          Windows,
			Version:     "5.1",
			VersionName: WindowsXP,
		},
		expectedStr: "Windows XP (5.1)",
	},
	{
		ua: "Mozilla/6.0 (Windows NT 6.2; WOW64; rv:16.0.1) Gecko/20121011 Firefox/16.0.1",
		expected: UserAgent{
			OS:          Windows,
			Version:     "6.2",
			VersionName: Windows8,
		},
		expectedStr: "Windows 8 (6.2)",
	},
	{
		ua: "Mozilla/5.0 (Windows; U; Windows NT 6.0; tr-TR) AppleWebKit/533.18.1 (KHTML, like Gecko) Version/5.0.2 Safari/533.18.5",
		expected: UserAgent{
			OS:          Windows,
			Version:     "6.0",
			VersionName: WindowsVista,
		},
		expectedStr: "Windows Vista (6.0)",
	},
	{
		ua: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_4_11; en) AppleWebKit/528.4+ (KHTML, like Gecko) Version/4.0dp1 Safari/526.11.2",
		expected: UserAgent{
			OS:      MacOS,
			Version: "10.4.11",
		},
		expectedStr: "macOS (10.4.11)",
	},
	{
		ua: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.13+ (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		expected: UserAgent{
			OS:          MacOS,
			Version:     "10.6.8",
			VersionName: MacOSSnowLeopard,
		},
		expectedStr: "macOS Snow Leopard (10.6.8)",
	},
	{
		ua: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_8) AppleWebKit/537.13+ (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		expected: UserAgent{
			OS:          MacOS,
			Version:     "10.11.8",
			VersionName: MacOSElCapitan,
		},
		expectedStr: "macOS El Capitan (10.11.8)",
	},
	{
		ua: "Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		expected: UserAgent{
			OS:      IOS,
			Version: "6.0",
			Distro:  IPad,
		},
		expectedStr: "iOS (6.0) (iPad)",
	},
	{
		ua: "Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0_3) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B511 Safari/9537.53",
		expected: UserAgent{
			OS:      IOS,
			Version: "7.0.3",
			Distro:  IPodTouch,
		},
		expectedStr: "iOS (7.0.3) (iPod touch)",
	},
	{
		ua: "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_5) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1 Safari/6533.18.5",
		expected: UserAgent{
			OS:      IOS,
			Version: "4.3.5",
			Distro:  IPhone,
		},
		expectedStr: "iOS (4.3.5) (iPhone)",
	},
	{
		ua: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		expected: UserAgent{
			OS:     Linux,
			Distro: Ubuntu,
		},
		expectedStr: "Linux (Ubuntu)",
	},
	{
		ua: "Mozilla/5.0 (Fedora; Linux x86_64) AppleWebKit/602.1 (KHTML, like Gecko) Version/8.0 Safari/602.1",
		expected: UserAgent{
			OS:     Linux,
			Distro: Fedora,
		},
		expectedStr: "Linux (Fedora)",
	},
	{
		ua: "Mozilla/5.0 (X11; CrOS i686 12.0.742.91) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.93 Safari/534.30",
		expected: UserAgent{
			OS: ChromeOS,
		},
		expectedStr: "ChromeOS",
	},
	{
		ua: "Mozilla/5.0 (X11; Linux i686) AppleWebKit/534.23 (KHTML, like Gecko) Chrome/11.0.686.3 Safari/534.23",
		expected: UserAgent{
			OS: Linux,
		},
		expectedStr: "Linux",
	},
	{
		ua: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		expected: UserAgent{
			OS: Linux,
		},
		expectedStr: "Linux",
	},
	{
		ua: "Mozilla/5.0 (Linux; U; Android 2.3; en-us) AppleWebKit/999+ (KHTML, like Gecko) Safari/999.9",
		expected: UserAgent{
			OS:      Android,
			Version: "2.3",
		},
		expectedStr: "Android (2.3)",
	},
	{
		ua: "Mozilla/5.0 (Android 5.1.1; Tablet; rv:46.0) Gecko/46.0 Firefox/46.0",
		expected: UserAgent{
			OS:      Android,
			Version: "5.1.1",
		},
		expectedStr: "Android (5.1.1)",
	},
}

func TestParseUserAgent(t *testing.T) {
	for i, uaTest := range userAgents {
		result := ParseUserAgent(uaTest.ua)
		if !uaEqual(result, uaTest.expected) {
			t.Errorf("Test %d: Incorrectly parsed user agent. Expected %#v, got %#v.", i, uaTest.expected, result)
			continue
		}

		resultStr := result.String()
		if resultStr != uaTest.expectedStr {
			t.Errorf("Test %d: Incorrectly stringified user agent. Expected %s, got %s.", i, uaTest.expectedStr, resultStr)
		}
	}
}

func uaEqual(a, b UserAgent) bool {
	return a.OS == b.OS && a.Version == b.Version && a.Distro == b.Distro
}
