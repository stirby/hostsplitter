package main

//Site represents a HostSplitter site
type Site struct {
	backendIndex int
	Backends     []string
	Secret       string
}

//GetBackend returns the next backend according using round robin
func (this *Site) GetBackend() string {
	if len(this.Backends) == 0 {
		return ""
	}

	//This code is racy but there's not much reason to make it synchronized
	index := this.backendIndex

	if this.backendIndex == len(this.Backends)-1 {
		this.backendIndex = 0
	} else {
		this.backendIndex = this.backendIndex + 1
	}

	return this.Backends[index]
}
