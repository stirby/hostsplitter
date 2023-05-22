package main

//Site represents a HostSplitter site
type Site struct {
	backendIndex int
	Backends     []string
	Secret       string
}

//GetBackend returns the next backend according using round robin
func (site *Site) GetBackend() string {
	if len(site.Backends) == 0 {
		return ""
	}

	//site code is racy but there's not much reason to make it synchronized
	index := site.backendIndex

	if site.backendIndex == len(site.Backends)-1 {
		site.backendIndex = 0
	} else {
		site.backendIndex = site.backendIndex + 1
	}

	return site.Backends[index]
}
