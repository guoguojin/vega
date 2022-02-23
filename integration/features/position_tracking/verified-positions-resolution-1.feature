Feature: Position resolution case 1

  Background:

    And the markets:
      | id        | quote name | asset | risk model                  | margin calculator                  | auction duration | fees         | price monitoring | oracle config          |
      | ETH/DEC19 | BTC        | BTC   | default-simple-risk-model-2 | default-overkill-margin-calculator | 1                | default-none | default-none     | default-eth-for-future |
    And the following network parameters are set:
      | name                           | value |
      | market.auction.minimumDuration | 1     |

      #default-simple-risk-model-2: "factorLong": 0; "factorShort": 0; "maxMoveUp": 0;  "minMoveDown": 0.016; "probabilityOfTrading": 0.2

  Scenario: https://drive.google.com/file/d/1bYWbNJvG7E-tcqsK26JMu2uGwaqXqm0L/view
# setup accounts
    Given the parties deposit on asset's general account the following amount:
      | party           | asset | amount        |
      | sellSideProvider | BTC   | 1000000000000 |
      | buySideProvider  | BTC   | 1000000000000 |
      | designatedLooser | BTC   | 11600         |
      | aux              | BTC   | 1000000000000 |
      | aux2             | BTC   | 1000000000000 |

# place auxiliary orders so we always have best bid and best offer as to not trigger the liquidity auction
    Then the parties place the following orders:
      | party | market id | side | volume | price | resulting trades | type       | tif     |
      | aux    | ETH/DEC19 | buy  | 10     | 1     | 0                | TYPE_LIMIT | TIF_GTC |
      | aux    | ETH/DEC19 | sell | 10     | 2000  | 0                | TYPE_LIMIT | TIF_GTC |
      | aux    | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT | TIF_GTC |
      | aux2   | ETH/DEC19 | sell | 1      | 150   | 0                | TYPE_LIMIT | TIF_GTC |
    Then the opening auction period ends for market "ETH/DEC19"
    And the mark price should be "150" for the market "ETH/DEC19"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/DEC19"

# insurance pool generation - setup orderbook
    When the parties place the following orders:
      | party            | market id | side | volume | price | resulting trades | type       | tif     | reference       |
      | sellSideProvider | ETH/DEC19 | sell | 290    | 150   | 0                | TYPE_LIMIT | TIF_GTC | sell-provider-1 |
      | buySideProvider  | ETH/DEC19 | buy  | 1      | 140   | 0                | TYPE_LIMIT | TIF_GTC | buy-provider-1  |

# insurance pool generation - trade
    When the parties place the following orders:
      | party            | market id | side | volume | price | resulting trades | type       | tif     | reference |
      | designatedLooser | ETH/DEC19 | buy  | 290    | 150   | 1                | TYPE_LIMIT | TIF_GTC | ref-1     |

    Then the parties should have the following account balances:
      | party            | asset | market id | margin  | general |
      | designatedLooser | BTC   | ETH/DEC19 | 11600   | 0       |

#designatedLooser has position of vol 290; price 150; RiskFactor is 0; slippage per unit is 
    Then the parties should have the following margin levels:
      | party            | market id | maintenance | search  | initial  | release |
      | designatedLooser | ETH/DEC19 | 39730       | 127136  | 158920   | 198650   |

# insurance pool generation - modify order book
    Then the parties cancel the following orders:
      | party           | reference      |
      | buySideProvider | buy-provider-1 |
    When the parties place the following orders:
      | party          | market id | side | volume | price | resulting trades | type       | tif     | reference      |
      | buySideProvider | ETH/DEC19 | buy  | 290      | 20    | 0                | TYPE_LIMIT | TIF_GTC | buy-provider-2 |

# insurance pool generation - set new mark price (and trigger closeout)
    When the parties place the following orders:
      | party           | market id | side | volume | price | resulting trades | type       | tif     | reference |
      | sellSideProvider | ETH/DEC19 | sell | 1      | 140   | 0                | TYPE_LIMIT | TIF_GTC | ref-1     |
      | buySideProvider  | ETH/DEC19 | buy  | 1      | 140   | 1                | TYPE_LIMIT | TIF_GTC | ref-2     |

# check positions
    Then the parties should have the following profit and loss:
      | party            | volume | unrealised pnl | realised pnl |
      | designatedLooser | 0      | 0              | -11600           |

# check margin levels
    Then the parties should have the following margin levels:
      | party            | market id | maintenance | search  | initial  | release |
      | designatedLooser | ETH/DEC19 | 0           | 0       | 0        | 0       |
# checking margins
    Then the parties should have the following account balances:
      | party            | asset | market id | margin | general |
      | designatedLooser | BTC   | ETH/DEC19 | 0      | 0       |

# then we make sure the insurance pool collected the funds
    And the insurance pool balance should be "0" for the market "ETH/DEC19"

# now we check what's left in the order book
# we expect 1 order at price of 40 to be left there on the buy side
# we sell a first time 1 to consume the book
# then try to sell 1 again with low price -> result in no trades -> buy side empty
# We expect no orders on the sell side: try to buy 1 for high price -> no trades -> sell side empty

