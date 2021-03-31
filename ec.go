package envoycharts

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/cloudkucooland/EnvoyCharts/internal/model"
	"github.com/cloudkucooland/go-envoy"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/objectbox/objectbox-go/objectbox"
)

var dbdir = "/var/log/envoy"
var defaultHost = "envoy.local"
var tzOffset int64 = 3600 * 6 // US/Central

type Client struct {
	Ob      *objectbox.ObjectBox
	Samples *model.EntryBox
	Envoy   *envoy.Envoy
}

func New(host string) (*Client, error) {
	if host == "" {
		host = defaultHost
	}

	c := Client{}
	var err error

	c.Ob, err = database()
	if err != nil {
		return nil, err
	}
	c.Samples = model.BoxForEntry(c.Ob)

	c.Envoy, err = envoy.New(host)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Client) Close() {
	c.Ob.Close()
}

func database() (*objectbox.ObjectBox, error) {
	builder := objectbox.NewBuilder()
	builder.Model(model.ObjectBoxModel())
	builder.Directory(dbdir)
	objectBox, err := builder.Build()
	if err != nil {
		panic(err)
	}

	return objectBox, nil
}

func (c *Client) Sample() error {
	e := model.Entry{
		Date: time.Now().Unix(),
	}

	var err error
	e.ProductionW, e.ConsumptionW, e.NetW, err = c.Envoy.Now()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if _, err := c.Samples.Put(&e); err != nil {
		fmt.Printf("could not insert sample: %s\n", err)
		return err
	}
	return nil
}

func (c *Client) GetAll() ([]*model.Entry, error) {
	entries, err := c.Samples.GetAll()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

func (c *Client) GetPastDay() ([]*model.Entry, error) {
	var query = c.Samples.Query(
		model.Entry_.Date.GreaterThan(time.Now().Unix() - 86400),
	)
	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

func (c *Client) GetDay(t time.Time) ([]*model.Entry, error) {
	var maxAlias = objectbox.Alias("max")
	var minAlias = objectbox.Alias("min")
	var query = c.Samples.Query(
		model.Entry_.Date.GreaterThan(0).As(maxAlias),
		model.Entry_.Date.LessThan(0).As(minAlias),
	)

	// there is probably a more Go-native way of doing this, but since we are all UNIX timestamps, this is fine
	dayStart := int64(math.Floor(float64(t.Unix()/86400))*86400) + tzOffset
	query.SetInt64Params(maxAlias, dayStart)
	query.SetInt64Params(minAlias, dayStart+86400)
	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

func (c *Client) GetDayRange(start time.Time, end time.Time) ([]*model.Entry, error) {
	var maxAlias = objectbox.Alias("max")
	var minAlias = objectbox.Alias("min")
	var query = c.Samples.Query(
		model.Entry_.Date.GreaterThan(0).As(maxAlias),
		model.Entry_.Date.LessThan(0).As(minAlias),
	)

	// there is probably a more Go-native way of doing this, but since we are all UNIX timestamps, this is fine
	dayStart := int64(math.Floor(float64(start.Unix()/86400))*86400) + tzOffset
	dayEnd := int64(math.Floor(float64(end.Unix()/86400))*86400) + 86399 + tzOffset
	query.SetInt64Params(maxAlias, dayStart)
	query.SetInt64Params(minAlias, dayEnd)
	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

func Barchart(w io.Writer, samples []*model.Entry, title string) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithLegendOpts(opts.Legend{Show: true, Data: []string{"Production", "Consumption", "Export/Import"}}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
	)

	barProd := make([]opts.BarData, 0)
	barCon := make([]opts.BarData, 0)
	barNet := make([]opts.BarData, 0)
	dates := make([]string, 0)

	for _, v := range samples {
		fu := time.Unix(v.Date, 0)
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

func Linechart(w io.Writer, samples []*model.Entry, title string) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithLegendOpts(opts.Legend{Show: true, Data: []string{"Production", "Consumption", "Export/Import"}}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
	)

	prod := make([]opts.LineData, 0)
	con := make([]opts.LineData, 0)
	net := make([]opts.LineData, 0)
	dates := make([]string, 0)

	for _, v := range samples {
		fu := time.Unix(v.Date, 0)
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
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}
