Feature: Test interactions between different auction types (0035-LIQM-001)

  # Related spec files:
  #  ../spec/0026-auctions.md
  #  ../spec/0032-price-monitoring.md
  #  ../spec/0035-liquidity-monitoring.md
  Background:
    Given the following network parameters are set:
      | name                                          | value |
      | market.stake.target.timeWindow                | 24h   |
      | market.stake.target.scalingFactor             | 1     |
      | market.liquidity.targetstake.triggering.ratio | 0     |
      | network.floatingPointUpdates.delay            | 5s    |
    And the average block duration is "1"
    And the simple risk model named "simple-risk-model-1":
      | long | short | max move up | min move down | probability of trading |
      | 0.1  | 0.1   | 10          | 10            | 0.1                    |
    And the log normal risk model named "log-normal-risk-model-1":
      | risk aversion | tau | mu | r   | sigma |
      | 0.000001      | 0.1 | 0  | 1.4 | -1    |
    And the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.004     | 0.001              |
    And the price monitoring updated every "1" seconds named "price-monitoring-1":
      | horizon | probability | auction extension |
      | 1       | 0.99        | 300               |
    And the markets:
      | id        | quote name | asset | risk model          | margin calculator         | auction duration | fees          | price monitoring   | oracle config          |
      | ETH/DEC21 | ETH        | ETH   | simple-risk-model-1 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
    And the parties deposit on asset's general account the following amount:
      | party  | asset | amount     |
      | party0 | ETH   | 1000000000 |
      | party1 | ETH   | 100000000  |
      | party2 | ETH   | 100000000  |

  Scenario: When trying to exit opening auction liquidity monitoring doesn't get triggered, hence the opening auction uncrosses and market goes into continuous trading mode (0026-AUCT-001, 0026-AUCT-002)

    Given the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | buy  | BID              | 500        | 10     | submission |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | sell | ASK              | 500        | 10     | amendment |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference  |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2  |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1 |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2 |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    # target_stake = mark_price x max_oi x target_stake_scaling_factor x rf = 1000 x 10 x 1 x 0.1
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 10000          | 10            |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
  Scenario: When trying to exit opening auction liquidity monitoring is triggered due to missing best bid, hence the opening auction gets extended, the markets trading mode is TRADING_MODE_MONITORING_AUCTION and the trigger is AUCTION_TRIGGER_LIQUIDITY (0026-AUCT-001, 0026-AUCT-005)

    # This ought to be "buy_shape" and "sell_shape" equivalents
========
  Scenario: When trying to exit opening auction liquidity monitoring is triggered due to missing best bid, hence the opening auction gets extended.

>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
    Given the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | sell | MID              | 2          | 1      | amendment |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
    # Again, pointless to check this in auction
    # And the price monitoring bounds are []

    When the opening auction period ends for market "ETH/DEC21"
    # Perhaps the reason for extending could be changed to reflect which check actually failed
    # In this case, though, it's the orderbook status, which applies to all auctions alike
    # So the trigger being AUCTION_TRIGGER_OPENING is as accurate as any
========
    When the opening auction period ends for market "ETH/DEC21"
    
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                 | auction trigger         | extension trigger         |
      | TRADING_MODE_OPENING_AUCTION | AUCTION_TRIGGER_OPENING | AUCTION_TRIGGER_LIQUIDITY |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |

    When the network moves ahead "1" blocks
    Then the auction ends with a traded volume of "10" at a price of "1000"

    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 10000          | 10            |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
  Scenario: When trying to exit opening auction liquidity monitoring is triggered due to insufficient supplied stake, hence the opening auction gets extended, the markets trading mode is TRADING_MODE_MONITORING_AUCTION and the trigger is AUCTION_TRIGGER_LIQUIDITY (0026-AUCT-001,0026-AUCT-004, 0026-AUCT-005)
========

  Scenario: When trying to exit opening auction liquidity monitoring is triggered due to insufficient supplied stake.
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock

    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 700               | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 700               | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 700               | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 700               | 0.001 | sell | MID              | 2          | 1      | amendment |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
   # In this case, the required time has expired, and the book is fine, so the trigger probably should be LIQUIDITY
    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |
========
  
    Then debug trades

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                 | auction trigger         | extension trigger         |
      | TRADING_MODE_OPENING_AUCTION | AUCTION_TRIGGER_OPENING | AUCTION_TRIGGER_LIQUIDITY |
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 800               | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 800               | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 800               | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 800               | 0.001 | sell | MID              | 2          | 1      | amendment |

    When the network moves ahead "1" blocks
    Then the trading mode should be "TRADING_MODE_MONITORING_AUCTION" for the market "ETH/DEC21"

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 801               | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 801               | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 801               | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 801               | 0.001 | sell | MID              | 2          | 1      | amendment |

    When the network moves ahead "1" blocks

    Then the trading mode should be "TRADING_MODE_MONITORING_AUCTION" for the market "ETH/DEC21"

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | MID              | 2          | 1      | amendment |

    When the network moves ahead "1" blocks

    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
  Scenario: Once market is in continuous trading mode: post a GFN order that should trigger liquidty auction, appropriate event is sent and market in TRADING_MODE_MONITORING_AUCTION (0026-AUCT-001, 0026-AUCT-005)
========
  Scenario: Once market is in continuous trading mode: post a persistent order that should trigger liquidity auction, appropriate event is sent and market in TRADING_MODE_MONITORING_AUCTION
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | MID              | 2          | 1      | amendment |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 20     | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
    And the market data for the market "ETH/DEC21" should be:
     | trading mode                    | auction trigger           | target stake | supplied stake | open interest |
     | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY | 3030         | 1000           | 10            |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
  Scenario: Once market is in continuous trading mode: post a GFN order that should trigger price auction, check that the order gets stopped, appropriate event is sent and market remains in TRADING_MODE_CONTINUOUS (0026-AUCT-001, 0026-AUCT-005)
========
    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | buy  | BID              | 1          | -2     |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | buy  | MID              | 2          | -1     |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | sell | ASK              | 1          | 2      |
      | lp1 | party0 | ETH/DEC21 | 10000             | 0.001 | sell | MID              | 2          | 1      |

    When the network moves ahead "1" blocks

    And the market data for the market "ETH/DEC21" should be:
     | trading mode            | auction trigger             | target stake | supplied stake | open interest |
     | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 3030         | 10000          | 30            |

Scenario: Once market is in continuous trading mode: post a non-persistent order that should trigger liquidity auction, appropriate event is sent and market in TRADING_MODE_MONITORING_AUCTION
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2       | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | MID              | 2          | 1       | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2       | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | MID              | 2          | 1       | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
       | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1010  | 0                | TYPE_LIMIT | TIF_GFN |
      | party2 | ETH/DEC21 | sell | 20     | 1010  | 0                | TYPE_LIMIT | TIF_GFN |

    # TODO: Confirming on Slack (https://vegaprotocol.slack.com/archives/CAHA5EX0F/p1620915342457400)
    # And the market data for the market "ETH/DEC21" should be:
    #  | trading mode                    | auction trigger           | target stake | supplied stake | open interest |
    #  | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY | 1000         | 1000           | 10            |
    # When the network moves ahead "1" blocks

    # And the market data for the market "ETH/DEC21" should be:
    #  | trading mode            | auction trigger             | target stake | supplied stake | open interest |
    #  | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 1000         | 1000          | 10            |


Scenario: Once market is in continuous trading mode: post a non-persistent order that should trigger price auction, check that the order gets stopped, appropriate event is sent and market remains in TRADING_MODE_CONTINUOUS
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset  | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001  | buy  | BID              | 1          | 2       | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001  | buy  | MID              | 2          | 1       | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001  | sell | ASK              | 1          | 2       | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001  | sell | MID              | 2          | 1       | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference | error                                                       |
      | party2 | ETH/DEC21 | sell | 10     | 1020  | 0                | TYPE_LIMIT | TIF_GTC | no-reject |                                                             |
      | party1 | ETH/DEC21 | buy  | 10     | 1020  | 0                | TYPE_LIMIT | TIF_GFN | reject-me | OrderError: non-persistent order trades out of price bounds |
    Then the following orders should be stopped:
      | party  | market id | reason                                               |
      | party1 | ETH/DEC21 | ORDER_ERROR_NON_PERSISTENT_ORDER_OUT_OF_PRICE_BOUNDS |
    
    Then the network moves ahead "5" blocks
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 1001      | 1019      | 1000         | 1000           | 10            |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
  Scenario: Once market is in continuous trading mode: enter liquidity monitoring auction -> extend with liquidity monitoring auction -> leave auction mode (0026-AUCT-001, 0068-MATC-033,0026-AUCT-005)
========

Scenario: Once market is in continuous trading mode: enter liquidity monitoring auction -> extend with price monitoring auction -> leave auction mode
    
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000               | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 1000               | 0.001 | buy  | MID              | 2          | 1      | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000               | 0.001 | sell | ASK              | 1          | 2      | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000               | 0.001 | sell | MID              | 2          | 1      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    # If the order traded there'd be insufficient liquidity for the market to operate, hence the order doesn't trade
    # and the market enters a liquidity monitoring auction
    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-1 |
      | party2 | ETH/DEC21 | sell | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-2 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |

    When the network moves ahead "1" blocks
    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |

    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-1 |
      | party2 | cancel-me-2 |

    When the network moves ahead "1" blocks
    
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | MID              | 2          | 1      | amendment |

    # leave liquidity auction
    When the network moves ahead "2" blocks
    # We should be able to leave liquidity auction now (price extension keep the market in auciton mode though)
    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           | extension trigger     |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY | AUCTION_TRIGGER_PRICE |

    # End price auction extension
    When the network moves ahead "301" blocks
    Then the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | auction trigger             | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1020       | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 1       | 1010      | 1030      | 3060         | 3060           | 30            |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
  Scenario: Once market is in continuous trading mode: enter liquidity monitoring auction -> extend with liquidity monitoring auction -> extend with liquidity monitoring -> leave auction mode (0026-AUCT-001, 0068-MATC-033,0026-AUCT-005)
    Given the following network parameters are set:
========
Scenario: Once market is in continuous trading mode: enter liquidity monitoring auction -> extend with price monitoring auction -> extend with liquidity monitoring -> leave auction mode
 Given the following network parameters are set:
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | MID              | 2          | 1      | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2      | amendment  |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | MID              | 2          | 1      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-1 |
      | party2 | ETH/DEC21 | sell | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-2 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |

    When the network moves ahead "1" blocks
    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |

    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-1 |
      | party2 | cancel-me-2 |

    When the network moves ahead "1" blocks
    
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset |
      | lp1 | party0 | ETH/DEC21 | 3060              | 0.001 | buy  | BID              | 1          | -2     |
      | lp1 | party0 | ETH/DEC21 | 3060              | 0.001 | buy  | MID              | 2          | -1     |
      | lp1 | party0 | ETH/DEC21 | 3060              | 0.001 | sell | ASK              | 1          | 2      |
      | lp1 | party0 | ETH/DEC21 | 3060              | 0.001 | sell | MID              | 2          | 1      |

    When the network moves ahead "2" blocks
    # We should be able to leave liquidity auction now (price extension keep the market in auciton mode though)
    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           | extension trigger     |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY | AUCTION_TRIGGER_PRICE |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 10     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |

    # Jump ahead to the end of the price monitoring auction period
    When the network moves ahead "301" blocks
    # TODO: Both of these cannot be true at the same time, I'd expect trigger to be LIQUIDITY.
    And the auction extension trigger should be "AUCTION_TRIGGER_PRICE" for market "ETH/DEC21"
    And the auction extension trigger should be "AUCTION_TRIGGER_LIQUIDITY" for market "ETH/DEC21"
    # we should still be in auction (liquidity extension) as the supplied stake is not sufficient
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode                    | auction trigger           | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY | 1000         | 3060           | 10            |

    When the network moves ahead "1" blocks

    # Increasing the supplied stake should end the auction
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | -2     |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | MID              | 2          | -1     |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | MID              | 2          | 1      |

    When the network moves ahead "1" blocks
    Then the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | auction trigger             | target stake | supplied stake | open interest |
      | 1020       | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 4080         | 4080           | 40            |

  Scenario: Once market is in continuous trading mode: enter price monitoring auction -> extend with liquidity monitoring auction -> leave auction mode
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party   | market id | commitment amount | fee   | side | pegged reference | proportion | offset |
      | lp1 | party0 | ETH/DEC21 | 2000               | 0.001 | buy  | BID              | 1          | -2     |
      | lp1 | party0 | ETH/DEC21 | 2000               | 0.001 | buy  | MID              | 2          | -1     |
      | lp1 | party0 | ETH/DEC21 | 2000               | 0.001 | sell | ASK              | 1          | 2      |
      | lp1 | party0 | ETH/DEC21 | 2000               | 0.001 | sell | MID              | 2          | 1      |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 10     | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 2000           | 10            |

    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger       |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE |

<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
    When the network moves ahead "2" blocks

    Then the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | MID              | 2          | 1      | amendment |

    # We're still in liquidity auction
    And the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |
========
    # end price auction (gets extended by liquidity)
    When the network moves ahead "301" blocks
    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger        | extension trigger         |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE  | AUCTION_TRIGGER_LIQUIDITY |
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock

    # We were in liquidity auction, we've updated the commitment amount
    When the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | sell | MID              | 2          | 1      | amendment |
    And the network moves ahead "1" blocks
    Then the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1020       | TRADING_MODE_CONTINUOUS | 1       | 1010      | 1030      | 3060         | 4000           | 30            |

  Scenario: Assure minimum auction length is respected
    Given the following network parameters are set:
      | name                           | value |
      | market.auction.minimumDuration | 10s   |

    When the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4000              | 0.001 | sell | MID              | 2          | 1      | amendment |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 10     | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 4000           | 10            |

    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger       |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE |

    # try to end price auction (gets extended by min auction length)
    When the network moves ahead "4" blocks

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger       |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE |

    When the network moves ahead "7" blocks
    Then the market data for the market "ETH/DEC21" should be:
      | trading mode            | auction trigger             |
      | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED |
<<<<<<<< HEAD:integration/features/verified/0026-AUCT-auction_interaction.feature
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1020       | TRADING_MODE_CONTINUOUS | 1       | 1010      | 1030      | 4080         | 5100           | 40            |

  Scenario: Once market is in continuous trading mode: enter price monitoring auction -> extend with liquidity monitoring auction -> leave auction mode (0026-AUCT-001, 0068-MATC-033,0026-AUCT-005)
     
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

      # Now we place some orders that are outside the price range to trigger price auction
    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-1 |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-2 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger       |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE |

    When the network moves ahead "1" blocks
    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-1 |
      | party2 | cancel-me-2 |

    # Jump ahead to the end of the auction
    When the network moves ahead "301" blocks

    # Triggering liquidity monitoring auction
    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-3 |
      | party2 | ETH/DEC21 | sell | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-4 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |

    When the network moves ahead "1" blocks
    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-3 |
      | party2 | cancel-me-4 |

     # Updating the commitment amount to come out of liquidity auction
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      | amendment |

    When the network moves ahead "1" blocks

    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | auction trigger horizon     | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 1         |    1010   |    1000      |   4080         |      10       |

  Scenario: Once market is in continuous trading mode: enter liquidity monitoring auction -> extend with price monitoring auction -> leave auction mode (0026-AUCT-001, 0026-AUCT-004, 0068-MATC-033,0026-AUCT-005)
     
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    # Triggering liquidity monitoring auction
    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-1 |
      | party2 | ETH/DEC21 | sell | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-2 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |

    When the network moves ahead "1" blocks
    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-1 |
      | party2 | cancel-me-2 |

     # Updating the commitment amount to come out of liquidity auction
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      | amendment |

    When the network moves ahead "1" blocks

    # Now we place some orders that are outside the price range to trigger price auction
    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger       |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE |

    # Jump ahead to the end of the auction
    When the network moves ahead "301" blocks

    Then the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1020   | 20  | party2 |

    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | auction trigger horizon     | min bound | max bound | target stake | supplied stake | open interest |
      | 1020       | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 1         |    1010   |    3060      |   4080         |      30       |

  Scenario: Once market is in continuous trading mode: enter liquidity monitoring auction -> extend with price monitoring auction -> extend with liquidity auction -> leave auction mode (0026-AUCT-001, 0068-MATC-033, 0026-AUCT-005)
     
    Given the following network parameters are set:
      | name                                          | value |
      | market.liquidity.targetstake.triggering.ratio | 0.8   |

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | party0 | ETH/DEC21 | 1000              | 0.001 | sell | ASK              | 1          | 2      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the auction ends with a traded volume of "10" at a price of "1000"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 990       | 1010      | 1000         | 1000           | 10            |

    # Triggering liquidity monitoring auction
    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-1 |
      | party2 | ETH/DEC21 | sell | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-2 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |

    When the network moves ahead "1" blocks
    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-1 |
      | party2 | cancel-me-2 |

     # Updating the commitment amount to come out of liquidity auction
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      | amendment |

    When the network moves ahead "1" blocks

    # Now we place some orders that are outside the price range to trigger price auction
    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-3 |
      | party2 | ETH/DEC21 | sell | 20     | 1020  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-4 | 

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger       |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_PRICE |

    When the network moves ahead "1" blocks
    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-3 |
      | party2 | cancel-me-4 |
    
    # Jump ahead to the end of the auction
    When the network moves ahead "301" blocks

    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/DEC21 | 1080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 1080              | 0.001 | sell | ASK              | 1          | 2      | amendment |

    When the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/DEC21 | buy  | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-5 |
      | party2 | ETH/DEC21 | sell | 10     | 1010  | 0                | TYPE_LIMIT | TIF_GTC | cancel-me-6 |

    Then the market data for the market "ETH/DEC21" should be:
      | trading mode                    | auction trigger           |
      | TRADING_MODE_MONITORING_AUCTION | AUCTION_TRIGGER_LIQUIDITY |

    When the network moves ahead "1" blocks
    And the parties cancel the following orders:
      | party  | reference   |
      | party1 | cancel-me-5 |
      | party2 | cancel-me-6 |

    # Updating the commitment amount to come out of liquidity auction
    Then  the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | party0 | ETH/DEC21 | 4080              | 0.001 | sell | ASK              | 1          | 2      | amendment |

    When the network moves ahead "1" blocks
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | auction trigger horizon     | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 1         |    1010   |    1000      |   4080         |      10       |
========
>>>>>>>> b1683b268 (Extend test (TODOs in place)):integration/features/auctions/509-auction-interaction.feature.mock
