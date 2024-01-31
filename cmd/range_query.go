package cmd

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func Querry_range(api v1.API, query string, start time.Time, step time.Duration, end time.Time) model.Value {
	r := v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	}
	result, _, err := api.QueryRange(
		context.TODO(), query, r, v1.WithTimeout(10*time.Second),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}
