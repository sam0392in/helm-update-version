package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func updateChart(dir string, chartData chartInfo, verbose bool) {
	file := dir + "/Chart.yaml"
	newVersion := strconv.Itoa(chartData.version.major) + "." + strconv.Itoa(chartData.version.minor) + "." + strconv.Itoa(chartData.version.hotfix)
	lineNo := chartData.placeholder
	data := readChart(dir)
	data[lineNo] = "version: \"" + newVersion + "\""
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close() // Ensure file closure even if errors occur
	datawriter := bufio.NewWriter(f)
	for i, l := range data {
		// Add a newline for all lines except the last one
		if i < len(data)-1 {
			datawriter.WriteString(l + "\n")
		} else {
			datawriter.WriteString(l)
		}
	}
	err = datawriter.Flush()
	if err != nil {
		return
	}
	if verbose {
		fmt.Println("new version: ", newVersion)
	} else {
		fmt.Println(newVersion)
	}

	err = f.Close()
	if err != nil {
		return
	}
}

func extractVersion(userVersion string, verbose bool) *chartVersion {
	var extractedVersion chartVersion
	ver := strings.Split(userVersion, ".")
	if len(ver) <= 1 {
		if verbose {
			fmt.Println("Chart version is not in desired format, example format: 1.0.0")
			fmt.Println("setting default version 0.0.0")
		}
		extractedVersion.major = 0
		extractedVersion.minor = 0
		extractedVersion.hotfix = 0
	} else if len(ver) < 3 {
		if verbose {
			fmt.Println("Version not in standard format [major.minor.hotfix], setting hotfix to default value 0")
		}
		extractedVersion.major, _ = strconv.Atoi(ver[0])
		extractedVersion.minor, _ = strconv.Atoi(ver[1])
		extractedVersion.hotfix = 0
	} else {
		extractedVersion.major, _ = strconv.Atoi(ver[0])
		extractedVersion.minor, _ = strconv.Atoi(ver[1])
		extractedVersion.hotfix, _ = strconv.Atoi(ver[2])
	}
	return &extractedVersion
}

func Update(dir, changeType string, verbose bool) {
	chartDetails := getChartVersion(dir, true, verbose)
	newVersion := bumpVersion(changeType, &chartDetails.version, verbose)
	chartDetails.version = *newVersion
	updateChart(dir, *chartDetails, verbose)
}

func UserVersionUpdate(dir, changeType, userVersion string, verbose bool) {
	chartDetails := getChartVersion(dir, false, verbose)
	currentVersion := *extractVersion(userVersion, verbose)
	// New Version after bumping user version
	chartDetails.version = *bumpVersion(changeType, &currentVersion, verbose)
	updateChart(dir, *chartDetails, verbose)
}
