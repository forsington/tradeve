# TradEVE

A trading utility for Eve Online. Uses ESI, written in golang.

TradEVE is a command-line application designed to assist station traders in Eve Online by identifying lucrative items for flip trading.

## How it works
Station traders place buy orders for items and, once fulfilled, list them as sell orders. When the sell order executes, traders profit from the difference between the buy and sell prices (a.k.a. arbitrage), after deducting broker fees and sales tax.

TradEVE evaluates items based on three key factors:
* **Profit Margin** - The average profitability of flip trading an item over the past X days, after accounting for taxes and trading fees.
* **LPDP - Largets Possible Daily Profit** - LPDP is calculated by taking the spread between an item's low and high prices for a day (bid-ask spread), subtracting taxes and fees, multiplying by the daily volume sold, and dividing by 2 to account for both buy and sell orders. This indicates the maximum potential profit from trading an item daily, highlighting opportunities to earn significant ISK.
* **Max center offset** - This factor ensures a balanced spread between buy and sell orders, which is crucial for reliable flip trading. High offset suggests that an item is predominantly traded in one direction (either sold to buy orders or bought from sell orders), which is not ideal for flipping.

### Data collection
TradEVE fetches data from the [EVE Swagger Interface (ESI)](https://esi.evetech.net/), which provides access to the Eve Online game data. The application retrieves the following data for each item:
* **Market History** - The daily high, low, and average prices for an item over the past X days.
* **Market Orders** - The current buy and sell orders for an item, including the price, volume, and order type.

This data is cached in a local Bolt database to speed up subsequent runs. Delete the file `bolt.db` to flush the cache and force a re-fetch of all data.

### Analysis
After market data has been scraped from ESI, TradEVE calculates the profit margin, LPDP, and max center offset for each item. The application then filters out items that do not meet the user-defined criteria (e.g., minimum profit margin, minimum volume, etc.).

Example output:
```
Analysis complete. 10 items match criteria over the last 3 days (daily average):

+-----+-------+-------------------+--------+-------------------+-------------------------------------------------+--------------------+-----------+----------+--------------------+
|  #  | SCORE |       LPDP        | MARGIN |     TAKE HOME     |                      ITEM                       |       PRICE        |  VOLUME   | FLIPPERS |      ISK VOL       |
+-----+-------+-------------------+--------+-------------------+-------------------------------------------------+--------------------+-----------+----------+--------------------+
|   1 |  1303 | 2,606,406,666 ISK | 10.5%  | 1,042,562,666 ISK | Fortizar                                        | 10,837,500,000 ISK |         5 |        2 | 54,187,500,000 ISK |
|   2 |   363 | 2,183,187,600 ISK | 5.7%   | 485,152,800 ISK   | Rhea                                            | 8,900,861,111 ISK  |         9 |        6 | 80,107,749,999 ISK |
|   3 |   326 | 980,073,500 ISK   | 5.4%   | 65,338,233 ISK    | Redeemer                                        | 1,300,821,535 ISK  |        30 |        3 | 39,024,646,061 ISK |
|   4 |   308 | 1,848,187,390 ISK | 57.1%  | 176,017,846 ISK   | Radical Drone Damage Amplifier Mutaplasmid      | 475,828,416 ISK    |        21 |        6 | 9,992,396,750 ISK  |
|   5 |   247 | 990,067,322 ISK   | 5.6%   | 22,760,168 ISK    | 25000mm Steel Plates II                         | 419,227,146 ISK    |        87 |        4 | 36,472,761,742 ISK |
|   6 |   129 | 4,154,767,858 ISK | 13.3%  | 193,245,016 ISK   | Metenox Moon Drill                              | 1,567,571,804 ISK  |        43 |       32 | 67,405,587,576 ISK |
|   7 |   120 | 120,347,293 ISK   | 6.9%   | 8,299,813 ISK     | Gistum B-Type 50MN Microwarpdrive               | 129,113,882 ISK    |        29 |        1 | 3,744,302,592 ISK  |
|   8 |   111 | 1,333,695,916 ISK | 32.1%  | 48,498,033 ISK    | Mid-grade Amulet Epsilon                        | 195,101,559 ISK    |        55 |       12 | 10,730,585,758 ISK |
|   9 |    97 | 782,221,000 ISK   | 20.1%  | 17,191,670 ISK    | Structure Construction Parts                    | 96,247,977 ISK     |        91 |        8 | 8,758,565,988 ISK  |
|  10 |    93 | 374,852,390 ISK   | 9.8%   | 32,595,860 ISK    | Gist X-Type Large Shield Booster                | 362,645,343 ISK    |        23 |        4 | 8,340,842,892 ISK  |
+-----+-------+-------------------+--------+-------------------+-------------------------------------------------+--------------------+-----------+----------+--------------------+
```

| Column        | Description                                                                                           |
|---------------|-------------------------------------------------------------------------------------------------------|
| **SCORE**     | A combined score, LPDP divided by the amount of flippers.                                             |
| **LPDP**      | Largest Possible Daily Profit: The maximum potential profit from trading the item daily, in ISK.      |
| **MARGIN**    | Profit margin percentage, after taxes and fees are deducted.                                          |
| **TAKE HOME** | Net profit of flipping.                                                                               |
| **ITEM**      | Item name.                                                                                            |
| **PRICE**     | Average price.                                                                                        |
| **VOLUME**    | The daily trading volume of the item.                                                                 |
| **FLIPPERS**  | The number of competing traders who have updated orders in the last 24 hours (buy and sell included). |
| **ISK VOL**   | The total volume of ISK trading hands per day.                                                        |


## Getting Started

### Prerequisits
To run, TradEVE is dependent on two files.

* A settings file, rename `config.json.example` to `config.json`, and fill in the parameters you wish to use.
* The EVE SDE file `invTypes.csv`, which will be fetched automatically from [Fuzzwork](https://www.fuzzwork.co.uk/dump/latest/) at application start, if it does not already exist in the directory. Force a re-fetch by deleting the file or by using the `-f` flag.

### Installing
Installation can be done by cloning the project and compiling from source: `make build` and then running `./bin/tradeve`. Requires [golang](https://golang.org/) to be installed.


### Configuration
See [config.json.example](config.json.example) for an example of the settings file:

```json
{
    "region": "The Forge",
    "cache_enabled": true,
    "force_sde_download": false,
    "broker_fee_buy": 0.005,
    "broker_fee_sell": 0.015,
    "sales_tax": 0.0202,
    "history_days": 7,
    "log_level": "info",
    "filters": {
        "low_isk_boundary": 10,
        "high_isk_boundary": 100000000000,
        "min_lpdp": 100000000,
        "min_volume": 5,
        "min_profit_margin": 0.05,
        "max_center_offset": 0.5,
        "max_flippers": 100,
        "min_score": 0
    },
    "exclude_groups": {
        "skins": true,
        "wearables": true,
        "skillbooks": true,
        "blueprints": true,
        "skinr": true,
        "crates": true,
        "deprecated": true
    }
}
```

| Variable                       | Description                                                    | Default Value                 |
|--------------------------------|----------------------------------------------------------------|----------------------         |
| **log_level**                  | Logging level (options: "debug", "info", "warn", "error").     | "info"                        |
| **cache_enabled**              | Enable caching of ESI data.                                    | true                          |
| **force_sde_download**         | Force download of the EVE SDE file at startup.                 | false                         |
| **region**                     | Specifies the region for trading.                              | The Forge                     |
| **broker_fee_buy**             | Broker fee percentage for buy orders.                          | 0.005 (0.5% @ a freeport)     |
| **broker_fee_sell**            | Broker fee percentage for sell orders.                         | 0.015 (1.5% (calculate [here](https://www.qsna.eu/eve/pilot-services/taxes))                  |
| **sales_tax**                  | Sales tax percentage for sell orders.                          | 0.0202 (2.02% @ Accounting V) |
| **history_days**               | Number of past days to consider for metrics.                   | 7                             |
| **filters.low_isk_boundary**   | Minimum ISK value for items to consider for flip trading.      | 10                            |
| **filters.high_isk_boundary**  | Maximum ISK value for items to consider for flip trading.      | 100000000000 (100B)           |
| **filters.min_lpdp**           | Minimum Largest Possible Daily Profit (LPDP) for items.        | 100000000 (100M)              |
| **filters.min_volume**         | Minimum trading volume for items to be considered.             | 5                             |
| **filters.min_profit_margin**  | Minimum profit margin for items (as a decimal).                | 0.05 (5%)                     |
| **filters.max_center_offset**  | Maximum allowed offset from center price for reliable trading. | 0.5                           |
| **filters.max_flippers**       | Maximum number of competing flippers for an item.              | 100                           |
| **filters.min_score**          | Minimum score for items to be considered.                      | 0                             |
| **exclude_groups.skins**       | Exclude SKINs.                                                 | true                          |
| **exclude_groups.wearables**   | Exclude wearables.                                             | true                          |
| **exclude_groups.skillbooks**  | Exclude skillbooks.                                            | true                          |
| **exclude_groups.blueprints**  | Exclude BPOs.                                                  | true                          |
| **exclude_groups.skinr**       | Exclude SKINr components.                                      | true                          |
| **exclude_groups.crates**      | Exclude crates.                                                | true                          |
| **exclude_groups.deprecated**  | Exclude deprecated items.                                      | true                          |

To adjust these settings, simply modify the values in the config file to suit your trading strategy. Save the file and restart TradEVE for the changes to take effect.

### Flags
Flag  | type     | Comment              | Default Value
----- | -------- | --------             | -----
-d    | `bool`   | Enable debug logging | false
-f    | `bool`   | Force SDE download   | false

## How can you help
Feel free to help out with any of the issues listed [on the issues page](https://github.com/forsington/tradeve/issues) or send a PR.

## Built With
* [golang](https://golang.org/)
* [BoltDB](github.com/boltdb/bolt)
* [EVE Swagger Interface](https://esi.evetech.net/)

## Authors
[forsington](https://github.com/forsington)

## License
MIT (see [LICENSE.md](LICENSE.md))

## Todo
- [ ] [Burger factors](https://wiki.eveuniversity.org/Identifying_items_for_trade)
    - [ ] Stable 20 day average price
- [ ] Use actual highs and lows from order book rather than aggregated history, exclude a few percent of volume closest to the spread
