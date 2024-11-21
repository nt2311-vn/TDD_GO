package concurency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, webs []string) map[string]bool {
	rs := make(map[string]bool)
	rsChanel := make(chan result)

	for _, web := range webs {
		go func(u string) {
			rsChanel <- result{u, wc(u)}
		}(web)
	}

	for i := 0; i < len(webs); i++ {
		r := <-rsChanel
		rs[r.string] = r.bool
	}

	return rs
}
