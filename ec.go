package envoycharts

import (
	"fmt"
	"github.com/cloudkucooland/EnvoyCharts/internal/model"
	"github.com/cloudkucooland/go-envoy"
	"github.com/objectbox/objectbox-go/objectbox"
	"time"
)

type Client struct {
	Ob      *objectbox.ObjectBox
	Samples *model.EntryBox
	Envoy   *envoy.Envoy
}

func New(host string) (*Client, error) {
	if host == "" {
		host = "envoy.local"
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
		panic(err)
	}

	return &c, nil
}

func (c *Client) Close() {
	c.Ob.Close()
}

func database() (*objectbox.ObjectBox, error) {
	builder := objectbox.NewBuilder()
	builder.Model(model.ObjectBoxModel())
	builder.Directory("/var/log/envoy")
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
	fmt.Printf("%+v\n", e)
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

func (c *Client) Day(t time.Time) ([]*model.Entry, error) {
	var e []*model.Entry
	/*
	   var query = box.Query(
	   		User_.Age.GreaterThan().Alias("min age"),
	   		User_.Age.LessThan().Alias("max age"))

	   // Then use the alias when setting the parameter value
	   query.SetInt64Params(objectbox.Alias("min age"), 50)
	   query.SetInt64Params(objectbox.Alias("max age"), 100
	*/
	return e, nil
}
