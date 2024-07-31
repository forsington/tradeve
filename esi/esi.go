package esi

//"github.com/Celeo/Goesi"
import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/antihax/optional"
	"github.com/forsington/tradeve/esi/swagger"
)

const (
	MarketHistoryCallTime = 80 * time.Millisecond
	MarketOrdersCallTime  = 160 * time.Millisecond
)

func NewClient(server string) *Client {
	client := &http.Client{
		Timeout: time.Minute,
	}
	config := swagger.Configuration{
		BasePath:   "https://esi.evetech.net/latest",
		HTTPClient: client,
	}
	esi := swagger.NewAPIClient(&config)
	return &Client{
		esi:    esi,
		calls:  0,
		server: server,
	}
}

type Client struct {
	esi    *swagger.APIClient
	calls  int
	server string
}

func (c *Client) Calls() int {
	return c.calls
}

func (c *Client) GetMarketHistory(typeId, regionId int) ([]swagger.GetMarketsRegionIdHistory200Ok, error) {
	opts := &swagger.MarketApiGetMarketsRegionIdHistoryOpts{
		Datasource: optional.NewString(c.server),
		// days?
	}

	history, resp, err := c.esi.MarketApi.GetMarketsRegionIdHistory(context.Background(), int32(regionId), int32(typeId), opts)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error fetching market history for item %d from ESI: %v", typeId, err)
	}
	c.calls = c.calls + 1

	return history, err
}

func (c *Client) GetMarketTypes(regionId int) ([]int32, error) {
	var orders []int32
	pageSize := 1000
	i := int32(0)
	for pageSize == 1000 {

		opts := &swagger.MarketApiGetMarketsRegionIdTypesOpts{
			Page: optional.NewInt32(i + 1),
		}

		newOrders, resp, err := c.esi.MarketApi.GetMarketsRegionIdTypes(context.Background(), int32(regionId), opts)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("\nCould not fetch page", i+1, "of active orders from ESI")
		} else {
			orders = append(orders, newOrders...)
		}
		c.calls = c.calls + 1

		pageSize = len(newOrders)
		i = i + 1
	}
	return orders, nil

}

func (c *Client) GetMarketOrders(regionId int, typeId int) ([]swagger.GetMarketsRegionIdOrders200Ok, error) {
	opts := &swagger.MarketApiGetMarketsRegionIdOrdersOpts{
		Datasource: optional.NewString(c.server),
		TypeId:     optional.NewInt32(int32(typeId)),
		Page:       optional.NewInt32(1),
	}

	orders, resp, err := c.esi.MarketApi.GetMarketsRegionIdOrders(context.Background(), "all", int32(regionId), opts)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error fetching market orders for item %d from ESI: %v", typeId, err)
	}
	c.calls = c.calls + 1

	return orders, err
}
