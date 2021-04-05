package envoycharts

import (
	"fmt"
	// "math"
	"time"

	"github.com/cloudkucooland/EnvoyCharts/internal/model"
	"github.com/cloudkucooland/go-envoy"
	"github.com/objectbox/objectbox-go/objectbox"
)

// Client is the primary handle for the EnvoyChart API
type Client struct {
	Ob      *objectbox.ObjectBox
	Samples *model.SampleBox
	Daily   *model.DailyBox
	Envoy   *envoy.Envoy
	TZ      *time.Location
	dbdir   string
}

// New creates a new Client
func New(host string) (*Client, error) {
	c := Client{
		dbdir: "/var/log/envoy",
	}
	var err error

	err = c.database()
	if err != nil {
		return nil, err
	}

	if err := c.SetLocation("America/Chicago"); err != nil {
		return nil, err
	}

	// if host is unset, discovery happens
	c.Envoy = envoy.New(host)
	return &c, nil
}

func (c *Client) SetLocation(l string) error {
	var err error
	c.TZ, err = time.LoadLocation(l)
	return err
}

// Close shuts down a client
func (c *Client) Close() {
	c.Ob.Close()
}

func (c *Client) database() error {
	builder := objectbox.NewBuilder()
	builder.Model(model.ObjectBoxModel())
	builder.Directory(c.dbdir)

	var err error
	c.Ob, err = builder.Build()
	if err != nil {
		return err
	}

	c.Samples = model.BoxForSample(c.Ob)
	c.Daily = model.BoxForDaily(c.Ob)
	return nil
}

// Sample polls an envoy device and stores the production values into the database
func (c *Client) Sample() error {
	t := time.Now().In(c.TZ)
	e := model.Sample{
		Date: t.Unix(),
	}

	// append the new sample to the primary table
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

	// use the ID to ensure each day has only one sample, updated throughout the day
	d := model.Daily{
		Id: c.dayStart(t),
	}
	d.Date = time.Unix(d.Id, 0).In(c.TZ)
	d.ProductionkWh, d.ConsumptionkWh, _, err = c.Envoy.Today()
	if _, err := c.Daily.Put(&d); err != nil {
		fmt.Printf("could not update daily: %s\n", err)
		return err
	}
	fmt.Printf("%+v\n%+v\n", e, d)
	return nil
}

// GetAll returns all values from the database, probably not useful for anything other than testing
func (c *Client) GetAll() ([]*model.Sample, error) {
	query := c.Samples.Query(
		model.Sample_.Date.OrderAsc(),
	)
	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

// GetAllDaily gets every entry from the daily table
func (c *Client) GetAllDaily() ([]*model.Daily, error) {
	query := c.Daily.Query(
		model.Daily_.Date.OrderAsc(),
	)
	d, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return d, err
	}
	return d, nil
}

// GetPastDay gets the samples for the previous 24 hours
func (c *Client) GetPastDay() ([]*model.Sample, error) {
	query := c.Samples.Query(
		model.Sample_.Date.GreaterThan(time.Now().Unix()-86400),
		model.Sample_.Date.OrderAsc(),
	)
	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

// GetDay returns all the samples for the day which contains the parameter
func (c *Client) GetDay(t time.Time) ([]*model.Sample, error) {
	ds := c.dayStart(t)
	de := c.dayEnd(t)

	var query = c.Samples.Query(
		model.Sample_.Date.Between(ds, de),
		model.Sample_.Date.OrderAsc(),
	)

	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

// GetDayRange gets all values between the start and end days
func (c *Client) GetDayRange(start time.Time, end time.Time) ([]*model.Sample, error) {
	ds := c.dayStart(start)
	de := c.dayEnd(end)

	var query = c.Samples.Query(
		model.Sample_.Date.Between(ds, de),
		model.Sample_.Date.OrderAsc(),
	)
	entries, err := query.Find()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

func (c *Client) dayStart(t time.Time) int64 {
	x := t.In(c.TZ)
	y := time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, c.TZ)
	return y.Unix()
}

func (c *Client) dayEnd(t time.Time) int64 {
	x := t.In(c.TZ)
	y := time.Date(x.Year(), x.Month(), x.Day(), 23, 59, 59, 9999, c.TZ)
	return y.Unix()
}
