package models

type Package struct {
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Version      string `json:"Version"`
	Maintainer   string `json:"Maintainer"`
	Architecture string `json:"Architecture"`
	Section      string `json:"Section"`
}

type PackageSet struct {
	Packages map[string]Package
}
