package sync_mod

type Counter struct {
	value int
}

func (c *Counter) Inc() {
}

func (c *Counter) Value() int {
	return 3
}
