Feature: Setting up 5 traders so that at once all the orders are places they end up with the following margin account balances: tt_5_0: 23 = searchLevel + 1, tt_5_1: 22=searchLevel, tt_5_2: 21=maintenanceLevel+1=searchLevel-1, tt_5_3=maintenanceLevel, tt_5_4=maintenanceLevel-1


  Background:
    Given the insurance pool initial balance for the markets is "0":
    And the execution engine have these markets:
      | name      | quote name | asset | risk model | lamd/long | tau/short | mu/max move up | r/min move down | sigma | release factor | initial factor | search factor | settlement price | auction duration |  maker fee | infrastructure fee | liquidity fee | p. m. update freq. | p. m. horizons | p. m. probs | p. m. durations | prob. of trading | oracle spec pub. keys | oracle spec property | oracle spec property type | oracle spec binding |
      | ETH/DEC19 |  BTC        | BTC   |  simple     | 0.1       | 0.1       | -1             | -1              | -1    | 1.4            | 1.2            | 1.1           | 100              | 1                |  0         | 0                  | 0             | 0                  |                |             |                 | 0.1              | 0xDEADBEEF,0xCAFEDOOD | prices.ETH.value     | TYPE_INTEGER              | prices.ETH.value    |
    And oracles broadcast data signed with "0xDEADBEEF":
      | name             | value |
      | prices.ETH.value | 100   |

  Scenario: https://drive.google.com/file/d/1bYWbNJvG7E-tcqsK26JMu2uGwaqXqm0L/view
    # setup accounts
    Given the following traders:
      | name    | amount    |
      | tt_4    | 500000    |
      | tt_5_0  | 123       |
      | tt_5_1  | 122       |
      | tt_5_2  | 121       |
      | tt_5_3  | 120       |
      | tt_5_4  | 119       |
      | tt_6    | 100000000 |
      | tt_10   | 10000000  |
      | tt_11   | 10000000  |
      | trader1 | 100000000 |
      | trader2 | 100000000 |
    Then I Expect the traders to have new general account:
      | name    | asset |
      | tt_4    | BTC   |
      | tt_5_0  | BTC   |
      | tt_5_1  | BTC   |
      | tt_5_2  | BTC   |
      | tt_5_3  | BTC   |
      | tt_5_4  | BTC   |
      | tt_6    | BTC   |
      | tt_10   | BTC   |
      | tt_11   | BTC   |
      | trader1 | BTC   |
      | trader2 | BTC   |

    Then traders place following orders with references:
      | trader  | market id | side | volume | price | resulting trades | type       | tif     | reference |
      | trader1 | ETH/DEC19 | sell | 1      | 200   | 0                | TYPE_LIMIT | TIF_GTC | t1-s-1    |
      | trader2 | ETH/DEC19 | buy  | 1      | 95    | 0                | TYPE_LIMIT | TIF_GTC | t2-b-1    |
      | trader1 | ETH/DEC19 | buy  | 1      | 100   | 0                | TYPE_LIMIT | TIF_GFA | t1-b-1    |
      | trader2 | ETH/DEC19 | sell | 1      | 100   | 0                | TYPE_LIMIT | TIF_GFA | t2-s-1    |

    Then the opening auction period for market "ETH/DEC19" ends
    And the mark price for the market "ETH/DEC19" is "100"

    # place orders and generate trades
    Then traders place following orders with references:
      | trader | market id | side | volume | price | resulting trades | type        | tif     | reference |
      | tt_10  | ETH/DEC19 | buy  | 10     | 100   | 0                | TYPE_LIMIT  | TIF_GTT | tt_10-1   |
      | tt_11  | ETH/DEC19 | sell | 10     | 100   | 1                | TYPE_LIMIT  | TIF_GTT | tt_11-1   |
      | tt_4   | ETH/DEC19 | buy  | 5      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_4-1    |
      | tt_4   | ETH/DEC19 | buy  | 5      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_4-2    |
      | tt_5_0 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_0-1  |
      | tt_5_1 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_1-1  |
      | tt_5_2 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_2-1  |
      | tt_5_3 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_3-1  |
      | tt_5_4 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_4-1  |
      | tt_6   | ETH/DEC19 | sell | 5      | 150   | 1                | TYPE_LIMIT  | TIF_GTC | tt_6-1    |
      | tt_5_0 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_0-2  |
      | tt_5_1 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_1-2  |
      | tt_5_2 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_2-2  |
      | tt_5_3 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_3-2  |
      | tt_5_4 | ETH/DEC19 | buy  | 1      | 150   | 0                | TYPE_LIMIT  | TIF_GTC | tt_5_4-2  |
      | tt_6   | ETH/DEC19 | sell | 5      | 150   | 1                | TYPE_LIMIT  | TIF_GTC | tt_6-2    |
      | tt_10  | ETH/DEC19 | buy  | 25     | 100   | 0                | TYPE_LIMIT  | TIF_GTC | tt_10-2   |
      | tt_11  | ETH/DEC19 | sell | 25     | 0     | 11               | TYPE_MARKET | TIF_FOK | tt_11-2   |


    And the mark price for the market "ETH/DEC19" is "100"

    # checking margins
    And the margins levels for the traders are:
      | trader | market id | maintenance | search | initial | release |
      | tt_5_0 | ETH/DEC19 | 20          | 22     | 24      | 28      |
      | tt_5_1 | ETH/DEC19 | 20          | 22     | 24      | 28      |
      | tt_5_2 | ETH/DEC19 | 20          | 22     | 24      | 28      |
      | tt_5_3 | ETH/DEC19 | 20          | 22     | 24      | 28      |
      | tt_5_4 | ETH/DEC19 | 0           | 0      | 0       | 0       |

    # checking balances
    Then I expect the trader to have a margin:
      | trader | asset | id        | margin | general |
      | tt_5_0 | BTC   | ETH/DEC19 | 23     | 0       |
      | tt_5_1 | BTC   | ETH/DEC19 | 22     | 0       |
      | tt_5_2 | BTC   | ETH/DEC19 | 21     | 0       |
      | tt_5_3 | BTC   | ETH/DEC19 | 20     | 0       |
      | tt_5_4 | BTC   | ETH/DEC19 | 0      | 0       |
