Feature: Position resolution case 3

  Background:
    Given the insurance pool initial balance for the markets is "0":
    And the execution engine have these markets:
      | name      | quote name | asset | risk model | lamd/long | tau/short | mu/max move up | r/min move down | sigma | release factor | initial factor | search factor | settlement price | auction duration |  maker fee | infrastructure fee | liquidity fee | p. m. update freq. | p. m. horizons | p. m. probs | p. m. durations | prob. of trading | oracle spec pub. keys | oracle spec property | oracle spec property type | oracle spec binding |
      | ETH/DEC19 |  BTC        | BTC   |  simple     | 0         | 0         | 0              | 0.016           | 2.0   | 5              | 4              | 3.2           | 42               | 0                |  0         | 0                  | 0             | 0                  |                |             |                 | 0.1              | 0xDEADBEEF,0xCAFEDOOD | prices.ETH.value     | TYPE_INTEGER              | prices.ETH.value    |
    And oracles broadcast data signed with "0xDEADBEEF":
      | name             | value |
      | prices.ETH.value | 42    |

  Scenario: https://docs.google.com/spreadsheets/d/1D433fpt7FUCk04dZ9FHDVy-4hA6Bw_a2
# setup accounts
    Given the following traders:
      | name             | amount        |
      | sellSideProvider | 1000000000000 |
      | buySideProvider  | 1000000000000 |
      | designatedLooser | 12000         |
    Then I Expect the traders to have new general account:
      | name             | asset |
      | designatedLooser | BTC   |
      | sellSideProvider | BTC   |
      | buySideProvider  | BTC   |

# insurance pool generation - setup orderbook
    Then traders place following orders with references:
      | trader           | market id | side | volume | price | resulting trades | type       | tif     | reference       |
      | sellSideProvider | ETH/DEC19 | sell | 290    | 150   | 0                | TYPE_LIMIT | TIF_GTC | sell-provider-1 |
      | buySideProvider  | ETH/DEC19 | buy  | 1      | 140   | 0                | TYPE_LIMIT | TIF_GTC | buy-provider-1  |

# insurance pool generation - trade
    Then traders place following orders:
      | trader           | market id | side | volume | price | resulting trades | type       | tif     |
      | designatedLooser | ETH/DEC19 | buy  | 290    | 150   | 1                | TYPE_LIMIT | TIF_GTC |

    Then the margins levels for the traders are:
      | trader           | id        | maintenance | search | initial | release |
      | designatedLooser | ETH/DEC19 | 2900        | 9280   | 11600   | 14500   |

# insurance pool generation - modify order book
    Then traders cancels the following orders reference:
      | trader          | reference      |
      | buySideProvider | buy-provider-1 |
    Then traders place following orders with references:
      | trader          | market id | side | volume | price | resulting trades | type       | tif     | reference      |
      | buySideProvider | ETH/DEC19 | buy  | 300    | 40    | 0                | TYPE_LIMIT | TIF_GTC | buy-provider-2 |

# check the trader accounts
    Then I expect the trader to have a margin:
      | trader           | asset | id        | margin | general |
      | designatedLooser | BTC   | ETH/DEC19 | 11600  | 400     |

# insurance pool generation - set new mark price (and trigger closeout)
    Then traders place following orders:
      | trader           | market id | side | volume | price | resulting trades | type       | tif     |
      | sellSideProvider | ETH/DEC19 | sell | 1      | 120   | 0                | TYPE_LIMIT | TIF_GTC |
      | buySideProvider  | ETH/DEC19 | buy  | 1      | 120   | 1                | TYPE_LIMIT | TIF_GTC |

# check positions
    Then position API produce the following:
      | trader           | volume | unrealisedPNL | realisedPNL |
      | designatedLooser | 0      | 0             | -12000      |

# checking margins
    Then I expect the trader to have a margin:
      | trader           | asset | id        | margin | general |
      | designatedLooser | BTC   | ETH/DEC19 | 0      | 0       |

# then we make sure the insurance pool collected the funds
    And the insurance pool balance is "3300" for the market "ETH/DEC19"

# now we check what's left in the orderbook
# we expect 10 orders at price of 40 to be left there on the buy side
# we sell a first time 10 to consume the book
# then try to sell 1 again with low price -> result in no trades -> buy side empty
# We expect no orders on the sell side: try to buy 1 for high price -> no trades -> sell side empty
    Then traders place following orders:
      | trader           | market id | side | volume | price | resulting trades | type       | tif     |
      | sellSideProvider | ETH/DEC19 | sell | 10     | 40    | 1                | TYPE_LIMIT | TIF_FOK |
      | sellSideProvider | ETH/DEC19 | sell | 1      | 1     | 0                | TYPE_LIMIT | TIF_FOK |
      | buySideProvider  | ETH/DEC19 | buy  | 1      | 1000  | 0                | TYPE_LIMIT | TIF_FOK |
