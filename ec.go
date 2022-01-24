package envoycharts

import (
	"database/sql"
	"fmt"
	"log"
	// "math"
	"path"
	"time"

	"github.com/cloudkucooland/go-envoy"
	_ "github.com/mattn/go-sqlite3" // sqlite3
)

// Client is the primary handle for the EnvoyChart API
type Client struct {
	DB    *sql.DB
	Envoy *envoy.Envoy
	TZ    *time.Location
	dbdir string
}

// New creates a new Client
func New(host string) (*Client, error) {
	c := Client{
		dbdir: "/var/log/envoy",
	}
	var err error

	c.DB, err = sql.Open("sqlite3", path.Join(c.dbdir, "ec.db"))
	if err != nil {
		log.Panic(err)
	}

	setup := `
	create table if not exists samples (id integer not null primary key, unixtime integer not null, production real not null, consumption real not null, net real not null);
	create table if not exists daily (id integer not null primary key, unixtime integer not null, production real not null, consumption real not null);
	`
	if _, err = c.DB.Exec(setup); err != nil {
		log.Printf("%q: %s\n", err, setup)
		return nil, err
	}

	if c.TZ, err = time.LoadLocation("America/Chicago"); err != nil {
		return nil, err
	}

	// if host is unset, discovery happens
	c.Envoy = envoy.New(host)
	return &c, nil
}

func (c *Client) Close() {
	c.DB.Close()
}

// Sample polls an envoy device and stores the production values into the database
func (c *Client) Sample() error {
	t := time.Now().In(c.TZ)
	e := &Sample{
		Date: t,
	}

	// append the new sample to the primary table
	var err error
	e.ProductionW, e.ConsumptionW, e.NetW, err = c.Envoy.Now()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err := c.PutSample(e); err != nil {
		fmt.Printf("could not insert sample: %s\n", err)
		return err
	}

	// use the ID to ensure each day has only one sample, updated throughout the day
	d := &Daily{
		Id: c.dayStart(t).Unix(),
	}
	d.Date = time.Unix(d.Id, 0).In(c.TZ)

	d.ProductionkWh, d.ConsumptionkWh, _, err = c.Envoy.Today()
	if err != nil {
		fmt.Printf("unable to poll")
		return err
	}
	if err := c.PutDaily(d); err != nil {
		fmt.Printf("could not update daily: %s\n", err)
		return err
	}
	// fmt.Printf("%+v\n%+v\n", e, d)
	return nil
}

// GetAll returns all values from the database, probably not useful for anything other than testing
func (c *Client) GetAllSamples() ([]*Sample, error) {
	entries := make([]*Sample, 0)

	// do some work

	return entries, nil
}

// GetAllDaily gets every entry from the daily table
func (c *Client) GetAllDaily() ([]*Daily, error) {
	d := make([]*Daily, 0)

	// do some work

	return d, nil
}

func (c *Client) GetSamples(start, end time.Time) ([]*Sample, error) {
	entries := make([]*Sample, 0)
	return entries, nil
}

// GetPastDay gets the samples for the previous 24 hours
// -- 24 hour window ending at Now()
func (c *Client) GetPastDay() ([]*Sample, error) {
	ds := time.Now().Add(0 - 86400*time.Second)
	de := time.Now()
	return c.GetSamples(ds, de)
}

// GetDay returns all the samples for the day which contains the parameter
// -- 24 hour window rounded to midnight
func (c *Client) GetDay(t time.Time) ([]*Sample, error) {
	ds := c.dayStart(t)
	de := c.dayEnd(t)
	return c.GetSamples(ds, de)
}

// GetDayRange gets all values between the start and end days -- rounds to start/end of days
func (c *Client) GetDayRange(start time.Time, end time.Time) ([]*Sample, error) {
	ds := c.dayStart(start)
	de := c.dayEnd(end)
	return c.GetSamples(ds, de)
}

func (c *Client) dayStart(t time.Time) time.Time {
	x := t.In(c.TZ)
	y := time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, c.TZ)
	return y
}

func (c *Client) dayEnd(t time.Time) time.Time {
	x := t.In(c.TZ)
	y := time.Date(x.Year(), x.Month(), x.Day(), 23, 59, 59, 9999, c.TZ)
	return y
}

func (c *Client) PutSample(s *Sample) error {
	tx, err := c.DB.Begin()
	if err != nil {
		log.Println(err.Error())
	}
	stmt, err := tx.Prepare("insert into samples (id, unixtime, production, consumption, net) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(s.Date.Unix(), s.Date.Unix(), s.ProductionW, s.ConsumptionW, s.NetW); err != nil {
		log.Println(err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (c *Client) PutDaily(d *Daily) error {
	daystart := c.dayStart(d.Date).Unix()

	tx, err := c.DB.Begin()
	if err != nil {
		log.Println(err.Error())
	}
	stmt, err := tx.Prepare("replace into daily (id, unixtime, production, consumption) values(?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(daystart, d.Date.Unix(), d.ProductionkWh, d.ConsumptionkWh); err != nil {
		log.Println(err.Error())
		return err
	}
	tx.Commit()

	return nil
}
