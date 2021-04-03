package envoycharts

import (
	"fmt"
	"math"
	"time"

	"github.com/cloudkucooland/EnvoyCharts/internal/model"
	"github.com/cloudkucooland/go-envoy"
	"github.com/objectbox/objectbox-go/objectbox"
)

var dbdir = "/var/log/envoy"
var tzOffset int64 = 3600 * 6 // US/Central

// Client is the primary handle for the EnvoyChart API
type Client struct {
	Ob      *objectbox.ObjectBox
	Samples *model.EntryBox
	Daily   *model.DailyBox
	Envoy   *envoy.Envoy
}

// New creates a new Client
func New(host string) (*Client, error) {
	c := Client{}
	var err error

	c.Ob, err = database()
	if err != nil {
		return nil, err
	}
	c.Samples = model.BoxForEntry(c.Ob)
	c.Daily = model.BoxForDaily(c.Ob)

	// if host is unset, discovery happens
	c.Envoy = envoy.New(host)
	return &c, nil
}

// Close shuts down a client
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

// Sample polls an envoy device and stores the production values into the database
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

	// overwrite the daily sample
	d := model.Daily{}
	d.DID = int64(math.Floor(float64(e.Date/86400))*86400) + tzOffset
	d.Date = time.Unix(d.DID, 0)
	d.ProductionkWn, d.ConsumptionkWh, _, err = c.Envoy.Today()
	if _, err := c.Daily.Put(&d); err != nil {
		fmt.Printf("could not update daily: %s\n", err)
		return err
	}

	return nil
}

// GetAll returns all values from the database, probably not useful for anything other than testing
func (c *Client) GetAll() ([]*model.Entry, error) {
	entries, err := c.Samples.GetAll()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

// GetPastDay gets the samples for the previous 24 hours
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

// GetDay returns all the samples for the day which contains the parameter (adjusted based on the tzOffset value)
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

// GetDayRange gets all values between the start and end days
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
