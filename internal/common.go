package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readChart(dir string) []string {
	fc := []string{}
	file := dir + "/Chart.yaml"
	if _, err := os.Stat(file); err == nil {
		f, _ := os.Open(file)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			fc = append(fc, line)
		}
	} else {
		fmt.Println("Chart File not found")
	}
	return fc
}

func bumpVersion(changeType string, currentVersion *chartVersion, verbose bool) *chartVersion {
	if verbose {
		fmt.Println("current version: ", strconv.Itoa(currentVersion.major)+"."+strconv.Itoa(currentVersion.minor)+"."+strconv.Itoa(currentVersion.hotfix))
		fmt.Println("changeType: ", changeType)
	}
	var newVersion = &chartVersion{
		major:  currentVersion.major,
		minor:  currentVersion.minor,
		hotfix: currentVersion.hotfix,
	}

	switch {
	case changeType == "nil":
		fmt.Println("no change in chartVersion")
	case changeType == "minor":
		newVersion.minor = newVersion.minor + 1
	case changeType == "major":
		newVersion.major = newVersion.major + 1
	case changeType == "hotfix":
		newVersion.hotfix = newVersion.hotfix + 1
	default:
		fmt.Println("no valid parameter passed in -c, validChoices", "major/minor/hotfix")
		os.Exit(2)
	}

	return newVersion
}

func getChartVersion(dir string, checkFormat bool, verbose bool) *chartInfo {
	var (
		chartinfo chartInfo
		version   chartVersion
	)
	data := readChart(dir)
	for i, l := range data {
		check, _ := regexp.MatchString("version", l)
		if check {
			if checkFormat {
				s := strings.Split(l, ":")
				trimed_s := strings.ReplaceAll(s[1], "\"", "")
				trimed_s = strings.ReplaceAll(trimed_s, " ", "")
				ver := strings.Split(trimed_s, ".")
				if len(ver) <= 1 {
					if verbose {
						fmt.Println("Chart version is not in desired format, example format: 1.0.0")
						fmt.Println("setting default version 0.0.0")
					}
					version.major = 0
					version.minor = 0
					version.hotfix = 0
				} else if len(ver) < 3 {
					if verbose {
						fmt.Println("Version not in standard format [major.minor.hotfix], setting hotfix to default value 0")
					}
					version.major, _ = strconv.Atoi(ver[0])
					version.minor, _ = strconv.Atoi(ver[1])
					version.hotfix = 0
				} else {
					version.major, _ = strconv.Atoi(ver[0])
					version.minor, _ = strconv.Atoi(ver[1])
					version.hotfix, _ = strconv.Atoi(ver[2])
				}
			} else {
				// set default value, it will not be used as user provided version will be used directly
				version.major = 0
				version.minor = 0
				version.hotfix = 0
			}
			chartinfo.version = version
			chartinfo.placeholder = i
			break
		}
	}
	return &chartinfo
}
