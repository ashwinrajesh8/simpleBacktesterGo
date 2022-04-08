package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"net/http"
	"strconv"
)



func generateLineItemsStrat() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < dataPoints; i++ {
		items = append(items, opts.LineData{Value: performance[i].strat})
	}
	return items
}
func generateLineItemsStock() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < dataPoints; i++ {
		items = append(items, opts.LineData{Value: performance[i].stock})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    ticker+" strategy performance",
			Subtitle: "One week backtest on requested signals with refresh @ " + strconv.Itoa(interval) + " minutes",
		}))

		xaxis := make([]string, dataPoints)
		for i := 0; i < dataPoints; i++ {
			xaxis[i] = strconv.Itoa(i)
		}

	line.SetXAxis(xaxis).
		AddSeries("Category A", generateLineItemsStrat()).
		AddSeries("Category B", generateLineItemsStock()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}