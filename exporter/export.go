package exporter

import (
	"fmt"
	"os"

	"github.com/prometheus/common/model"
)

func ExportDataFrame() {
	//
}

func Export_raw(query_result model.Value, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	if query_result.Type() == model.ValVector {
		vector := query_result.(model.Vector)
		for _, elem := range vector {
			str := elem.Metric.String() + ", " + elem.Timestamp.String()
			file.WriteString(str + "\n")
		}
	}

}
