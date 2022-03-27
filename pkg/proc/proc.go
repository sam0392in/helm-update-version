/*
Contributor: Samarth Kanungo
Description: This package contain all function
*/

package proc

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

type chartVersion struct {
	major  string
	minor  string
	hotfix string
}

type chartInfo struct {
	version     chartVersion
	placeholder int
}

func init() {
	logger, _ = zap.NewProduction()
}

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
		logger.Error("Chart File not found")
	}
	return fc
}

func updateChart(dir string, chartinfo chartInfo) {
	file := dir + "/Chart.yaml"
	new_version := chartinfo.version.major + "." + chartinfo.version.minor + "." + chartinfo.version.hotfix
	lineNo := chartinfo.placeholder
	data := readChart(dir)
	data[lineNo] = "version: \"" + new_version + "\""
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error(err.Error())
	}
	datawriter := bufio.NewWriter(f)
	for _, l := range data {
		_, _ = datawriter.WriteString(l + "\n")
	}
	datawriter.Flush()
	logger.Info("chartVersion updated",
		zap.String("chartVersion", new_version))
	f.Close()
}

func getChartVersion(dir string) chartInfo {
	var chartinfo chartInfo
	var version chartVersion
	data := readChart(dir)
	for i, l := range data {
		check, _ := regexp.MatchString("version", l)
		if check {
			s := strings.Split(l, ":")
			trimed_s := strings.ReplaceAll(s[1], "\"", "")
			trimed_s = strings.ReplaceAll(trimed_s, " ", "")
			ver := strings.Split(trimed_s, ".")
			version.major = ver[0]
			version.minor = ver[1]
			version.hotfix = ver[2]

			chartinfo.version = version
			chartinfo.placeholder = i
			break
		}
	}
	return chartinfo
}

func BumpVersion(dir, changetype string) {
	currentchartInfo := getChartVersion(dir)
	logger.Info("read chart version successful",
		zap.String("chartVersion", currentchartInfo.version.major+"."+currentchartInfo.version.minor+"."+currentchartInfo.version.hotfix),
	)
	newChartInfo := currentchartInfo
	newVersion := currentchartInfo.version
	switch {
	case changetype == "nil":
		logger.Info("no change in chartVersion")
	case changetype == "minor":
		min, error := strconv.Atoi(newVersion.minor)
		if error != nil {
			logger.Error(error.Error())
		} else {
			newMinor := min + 1
			newVersion.minor = strconv.Itoa(newMinor)
		}
	case changetype == "major":
		maj, error := strconv.Atoi(newVersion.major)
		if error != nil {
			logger.Error(error.Error())
		} else {
			newMajor := maj + 1
			newVersion.major = strconv.Itoa(newMajor)
		}
	case changetype == "hotfix":
		hot, error := strconv.Atoi(newVersion.hotfix)
		if error != nil {
			logger.Error(error.Error())
		} else {
			newHotfix := hot + 1
			newVersion.hotfix = strconv.Itoa(newHotfix)
		}
	default:
		logger.Panic("no valid parameter passed in -b",
			zap.String("validChoices", "major/minor/hotfix"),
		)
	}
	if changetype != "nil" {
		newChartInfo.version = newVersion
		updateChart(dir, newChartInfo)
	}
}
