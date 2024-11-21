package concurency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, webs []string) map[string]bool {
}
