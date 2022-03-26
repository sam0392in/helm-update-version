/*
Contributor: Samarth Kanungo
Description: This program track helmchart version
*/

package main

import (
	"cvr/pkg/proc"
	"flag"
)

var (
	dir        string
	changetype string
)

func init() {
	flag.StringVar(&changetype, "b", "nil", "bump chart version, valid values: minor, major, hostfix")
	flag.StringVar(&dir, "d", "./", "Pass helmchart directory path")
	flag.Parse()
}

func main() {
	proc.BumpVersion(dir, changetype)
}
