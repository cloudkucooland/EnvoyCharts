package envoycharts

import (
	"fmt"
	"github.com/cloudkucooland/EnvoyCharts/internal/model"
	"github.com/objectbox/objectbox-go/objectbox"
	"time"
)

type Client struct {
	Ob  *objectbox.ObjectBox
	Box *model.EntryBox
}

func New() (*Client, error) {
	c := Client{}
	var err error

	c.Ob, err = database()
	if err != nil {
		return nil, err
	}
	c.Box = model.BoxForEntry(c.Ob)

	return &c, nil
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

func (c *Client) Insert(prod, consum, net float64) error {
	e := model.Entry{
		Date:         time.Now().Unix(),
		ProductionW:  prod,
		ConsumptionW: consum,
		NetW:         net,
	}

	if _, err := c.Box.Put(&e); err != nil {
		fmt.Printf("could not insert sample: %s\n", err)
		return err
	}

	fmt.Printf("%+v\n", e)

	return nil
}

func (c *Client) GetAll() ([]*model.Entry, error) {
	entries, err := c.Box.GetAll()
	if err != nil {
		fmt.Println(err)
		return entries, err
	}
	return entries, nil
}

func (c *Client) Day() ([]*model.Entry, error) {
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
