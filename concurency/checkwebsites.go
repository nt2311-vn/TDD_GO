package concurency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, webs []string) map[string]bool {
	rs := make(map[string]bool)

	for _, web := range webs {
		if wc(web) {
			rs[web] = true
		} else {
			rs[web] = false
		}
	}

	return rs
}
