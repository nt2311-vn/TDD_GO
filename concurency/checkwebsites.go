package concurency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, webs []string) map[string]bool {
	rs := make(map[string]bool)

	for _, web := range webs {
		rs[web] = wc(web)
	}

	return rs
}
