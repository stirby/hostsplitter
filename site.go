package main

type Site struct {
	backendIndex int
	Backends     []string
	Secret       string
}

func (this *Site) GetBackend() string {
	if len(this.Backends) == 0 {
		return ""
	}

	index := this.backendIndex

	if this.backendIndex == len(this.Backends)-1 {
		this.backendIndex = 0
	} else {
		this.backendIndex = this.backendIndex + 1
	}

	return this.Backends[index]
}
