/*
Contributor: Samarth Kanungo
Description: This package contain all function
*/

package proc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

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

func updateChart(dir string, chartInfo []float32, val float32) {
	file := dir + "/Chart.yaml"
	currentVersion := chartInfo[0]
	new := currentVersion + val
	newVersion := fmt.Sprintf("%.2f", new)
	lineNo := int(chartInfo[1])
	data := readChart(dir)
	data[lineNo] = "version: \"" + newVersion + "\""
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
		zap.String("chartVersion", newVersion))
	f.Close()
}

func getChartVersion(dir string) []float32 {
	var chartinfo []float32
	data := readChart(dir)
	for i, l := range data {
		check, _ := regexp.MatchString("version", l)
		if check {
			s := strings.Split(l, ":")
			ver := strings.ReplaceAll(s[1], "\"", "")
			ver = strings.ReplaceAll(ver, " ", "")
			version, err := strconv.ParseFloat(ver, 32)
			if err != nil {
				logger.Error(err.Error())
			}
			chartinfo = append(chartinfo, float32(version))
			chartinfo = append(chartinfo, float32(i))
			break
		}
	}
	return chartinfo
}

func BumpVersion(dir, changetype string) {
	chartInfo := getChartVersion(dir)
	logger.Info("read chart version successful",
		zap.Float32("chartVersion", chartInfo[0]),
	)
	var val float32
	switch {
	case changetype == "nil":
		val = 0
		logger.Info("no change in chartVersion")
	case changetype == "minor":
		val = 0.1
	case changetype == "major":
		val = 1.0
	case changetype == "hotfix":
		val = 0.01
	default:
		logger.Info("no change in chartVersion")
	}
	if val != 0 {
		updateChart(dir, chartInfo, val)
	}
}
