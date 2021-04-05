package envoycharts

import (
	"io"
	"time"

	"github.com/cloudkucooland/EnvoyCharts/internal/model"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Barchart takes a set of samples, and a title and writes the barchart to the specified writer
func Barchart(w io.Writer, samples []*model.Sample, title string, tz *time.Location) {
	bar := charts.NewBar()
	// use our custom template
	useECTemplates()
	// bar.Renderer = NewECRenderer(bar, bar.Validate)

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithLegendOpts(opts.Legend{Show: true, Data: []string{"Production", "Consumption", "Export/Import"}}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		// charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		// charts.WithXAxisOpts(opts.XAxis{Type: "time"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "kWh", Type: "value"}),
	)

	barProd := make([]opts.BarData, 0)
	barCon := make([]opts.BarData, 0)
	barNet := make([]opts.BarData, 0)
	dates := make([]string, 0)

	for _, v := range samples {
		fu := time.Unix(v.Date, 0).In(tz)
		t := fu.Format("2006/01/02 15:04:05")

		dates = append(dates, t)
		barProd = append(barProd, opts.BarData{Value: v.ProductionW})
		barCon = append(barCon, opts.BarData{Value: 0 - v.ConsumptionW})
		barNet = append(barNet, opts.BarData{Value: 0 - v.NetW})
	}

	bar.SetXAxis(dates).
		AddSeries("Production", barProd).
		AddSeries("Consumption", barCon).
		AddSeries("Export/Import", barNet)
	bar.Render(w)
}

// Linechart takes a set of samples, and a title and writes the linechart to the specified writer
func Linechart(w io.Writer, samples []*model.Sample, title string, tz *time.Location) {
	line := charts.NewLine()
	// use our custom template
	useECTemplates()
	// line.Renderer = NewECRenderer(line, line.Validate)

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithLegendOpts(opts.Legend{Show: true, Data: []string{"Production", "Consumption", "Export/Import"}}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		// charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		// charts.WithXAxisOpts(opts.XAxis{Type: "time"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Time"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "W", Type: "value"}),
	)

	prod := make([]opts.LineData, 0)
	con := make([]opts.LineData, 0)
	net := make([]opts.LineData, 0)
	dates := make([]string, 0)

	for _, v := range samples {
		fu := time.Unix(v.Date, 0).In(tz)
		t := fu.Format("2006/01/02 15:04:05")

		dates = append(dates, t)
		prod = append(prod, opts.LineData{Value: v.ProductionW})
		con = append(con, opts.LineData{Value: 0 - v.ConsumptionW})
		net = append(net, opts.LineData{Value: 0 - v.NetW})
	}

	line.SetXAxis(dates).
		AddSeries("Production", prod).
		AddSeries("Consumption", con).
		AddSeries("Export/Import", net).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			charts.WithAreaStyleOpts(opts.AreaStyle{Opacity: 0.3}),
			// charts.WithLabelOpts(opts.Label{Show: true}),
			// charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{ Name: "kWh", Type: "value", }),
		)
	line.Render(w)
}

// LinechartDaily takes a set of daily values, and a title and writes the linechart to the specified writer
func LinechartDaily(w io.Writer, samples []*model.Daily, title string, tz *time.Location) {
	line := charts.NewLine()
	// use our custom template
	useECTemplates()
	// line.Renderer = NewECRenderer(line, line.Validate)

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithLegendOpts(opts.Legend{Show: true, Data: []string{"Production", "Consumption", "Net"}}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		// charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		// charts.WithXAxisOpts(opts.XAxis{Type: "time"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Date"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "kWh", Type: "value"}),
	)

	prod := make([]opts.LineData, 0)
	con := make([]opts.LineData, 0)
	net := make([]opts.LineData, 0)
	dates := make([]string, 0)

	for _, v := range samples {
		t := v.Date.In(tz).Format("2006/01/02")

		dates = append(dates, t)
		prod = append(prod, opts.LineData{Value: v.ProductionkWh})
		con = append(con, opts.LineData{Value: 0 - v.ConsumptionkWh})
		net = append(net, opts.LineData{Value: v.ProductionkWh - v.ConsumptionkWh})
	}

	line.SetXAxis(dates).
		AddSeries("Production", prod).
		AddSeries("Consumption", con).
		AddSeries("Net", net).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			charts.WithAreaStyleOpts(opts.AreaStyle{Opacity: 0.3}),
			// charts.WithLabelOpts(opts.Label{Show: true}),
			// charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{ Name: "kWh", Type: "value", }),
		)
	line.Render(w)
}
