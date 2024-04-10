package internal

import (
	"bufio"
	"fmt"
	"os"
)

func updateChartNoEdit(dir, version string, lineNo int, verbose bool) {
	file := dir + "/Chart.yaml"
	data := readChart(dir)
	data[lineNo] = "version: \"" + version + "\""
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
		fmt.Println("new version: ", version)
	} else {
		fmt.Println(version)
	}
	err = f.Close()
	if err != nil {
		return
	}
}

func UpdateNoEdit(dir, version string, verbose bool) {
	chartDetails := getChartVersion(dir, false, verbose)
	updateChartNoEdit(dir, version, chartDetails.placeholder, verbose)
}
