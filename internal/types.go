package internal

type chartVersion struct {
	major  int
	minor  int
	hotfix int
}

type chartInfo struct {
	version     chartVersion
	placeholder int
}
