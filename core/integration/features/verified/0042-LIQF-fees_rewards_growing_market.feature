Feature:

  Background:

    Given the simple risk model named "simple-risk-model-1":
      | long | short | max move up | min move down | probability of trading |
      | 0.1  | 0.1   | 500         | 500           | 0.1                    |
    And the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.0004    | 0.001              |
    And the price monitoring named "price-monitoring":
      | horizon | probability | auction extension |
      | 1       | 0.99        | 3                 |
    And the markets:
      | id        | quote name | asset | risk model          | margin calculator         | auction duration | fees          | price monitoring | oracle config          |
      | ETH/MAR22 | USD        | USD   | simple-risk-model-1 | default-margin-calculator | 2                | fees-config-1 | price-monitoring | default-eth-for-future |
    And the following network parameters are set:
      | name                                                | value |
      | market.value.windowLength                           | 1h    |
      | market.stake.target.timeWindow                      | 24h   |
      | market.stake.target.scalingFactor                   | 1     |
      | market.liquidity.targetstake.triggering.ratio       | 0     |
      | market.liquidity.providers.fee.distributionTimeStep | 10m   |

 
  @VirtStake
  Scenario: 001 2 LPs joining at start, unequal commitments. Checking calculation of equity-like-shares and liquidity-fee-distribution in a market with small growth (0042-LIQF-008 0042-LIQF-011)

    # Scenario has 6 market periods:

    # - 0th period (bootstrap period): no LP changes, no trades
    # - 1st period: 1 LPs decrease commitment, some trades occur
    # - 2nd period: 1 LPs increase commitment, some trades occur
    # - 3rd period: 2 LPs decrease commitment, some trades occur
    # - 4th period: 2 LPs increase commitment, some trades occur
    # - 5th period: 1 LPs decrease commitment, 1 LPs increase commitment, some trades occur


    # Scenario moves ahead to next market period by:

    # - moving ahead "1" blocks to trigger the next liquidity distribution
    # - moving ahead "1" blocks to trigger the next market period


    # Following checks occur in each market where trades:

    # - Check transfers from the price taker to the market-liquidity-pool are correct
    # - Check accumulated-liquidity-fees are non-zero and correct
    # - Check equity-like-shares are correct
    # - Check transfers from the market-liquidity-pool to the liquidity-providers are correct
    # - Check accumulated-liquidity-fees are zero

    Given the average block duration is "1801"

    And the parties deposit on asset's general account the following amount:
      | party  | asset | amount    |
      | lp1    | USD   | 100000000 |
      | lp2    | USD   | 100000000 |
      | party1 | USD   | 100000    |
      | party2 | USD   | 100000    |

    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | buy  | MID              | 3          | 1      | amendment  |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment  |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | sell | MID              | 3          | 1      | amendment  |

    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | buy  | BID              | 1          | 2      | submission |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | buy  | MID              | 3          | 1      | amendment  |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | sell | ASK              | 1          | 2      | amendment  |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | sell | MID              | 3          | 1      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/MAR22 | buy  | 50     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/MAR22 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/MAR22 | sell | 50     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |


    # 0th period (bootstrap period): no LP changes, no trades
    Then the opening auction period ends for market "ETH/MAR22"

    And the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1000  | 50   | party2 |

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 500       | 1500      | 5000         | 50000          | 50            |

    And the order book should have the following volumes for market "ETH/MAR22":
      | side | price | volume |
      | buy  | 898   | 280    |
      | buy  | 900   | 1      |
      | buy  | 999   | 77     |
      | sell | 1001  | 75     |
      | sell | 1100  | 1      |
      | sell | 1102  | 228    |
    
    And the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share | average entry valuation |
      | lp1   | 0.08              | 4000                    |
      | lp2   | 0.92              | 50000                   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"


    When the network moves ahead "2" blocks:

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------


    # 1st period: 1 LPs decrease commitment, positive growth:
    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share | average entry valuation |
      | lp1   | 0.08              | 4000                    |
      | lp2   | 0.92              | 50000                   |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0612244897959184 | 4000                    |
      | lp2   | 0.9387755102040816 | 50000                   |

    # -------------------------------------------------------------------------------------------------------------------

    Given the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 51     | 1001  | 1                | TYPE_LIMIT | TIF_GTC |

    Then the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1001  | 51   | lp2    |

    # CALCULATION:
    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(51 * 1001 * 0.002) = ceil(102.102) = 103

    And the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 103    | USD   |

    And the accumulated liquidity fees should be "103" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 500       | 1500      | 10110        | 49000          | 101           |

    And the order book should have the following volumes for market "ETH/MAR22":
      | side | price | volume |
      | buy  | 898   | 274    |
      | buy  | 900   | 1      |
      | buy  | 999   | 75     |
      | sell | 1001  | 74     |
      | sell | 1100  | 1      |
      | sell | 1102  | 223    |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0612244897959184 | 4000                    |
      | lp2   | 0.9387755102040816 | 50000                   |

    # Trigger next liquidity fee distribution without triggering next market period
    And the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0612244897959184 | 4000                    |
      | lp2   | 0.9387755102040816 | 50000                   |

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 6      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 97     | USD   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0612244897959184 | 4000                    |
      | lp2   | 0.9387755102040816 | 50000                   |
    
    # Trigger entry into next market period
    And the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0612244897959184 | 4000                    |
      | lp2   | 0.9387755102040816 | 50000                   |

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------


    # 2nd period: 1 LPs increase commitment, positive growth:
    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0612244897959184 | 4000                    |
      | lp2   | 0.9387755102040816 | 50000                   |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share    | average entry valuation |
      | lp1   | 0.0794314813996635   | 15985.90375384367775    |
      | lp2   | 0.9205685186003365   | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079809          | 15628.748               |
    #  | lp2   | 0.920191          | 50000.000               |

    # -------------------------------------------------------------------------------------------------------------------

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 52      | 1001  | 1                | TYPE_LIMIT | TIF_GTC |

    Then the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1001  | 52   | lp2    |

    # CALCULATION:
    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(52 * 1001 * 0.002) = ceil(104.104) = 105

    And the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 105    | USD   |

    And the accumulated liquidity fees should be "105" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 502       | 1500      | 15315        | 50000          | 153           |

    And the order book should have the following volumes for market "ETH/MAR22":
      | side | price | volume |
      | buy  | 898   | 280    |
      | buy  | 900   | 1      |
      | buy  | 999   | 77     |
      | sell | 1001  | 75     |
      | sell | 1100  | 1      |
      | sell | 1102  | 228    |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share    | average entry valuation           |
      | lp1   | 0.0794314813996635   | 15985.90375384367775              |
      | lp2   | 0.9205685186003365   | 50000                             |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079809          | 15628.748               |
    #  | lp2   | 0.920191          | 50000.000               |

    # Trigger next liquidity fee distribution without triggering next period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share    | average entry valuation           |
      | lp1   | 0.0794314813996635   | 15985.90375384367775              |
      | lp2   | 0.9205685186003365   | 50000                             |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079809          | 15628.748               |
    #  | lp2   | 0.920191          | 50000.000               |

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 8      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 97     | USD   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share    | average entry valuation           |
      | lp1   | 0.0794314813996635   | 15985.90375384367775              |
      | lp2   | 0.9205685186003365   | 50000                             |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079809          | 15628.748               |
    #  | lp2   | 0.920191          | 50000.000               |
    
    # Trigger entry into next market period
    And the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share    | average entry valuation           |
      | lp1   | 0.0794314813996635   | 15985.90375384367775              |
      | lp2   | 0.9205685186003365   | 50000                             |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079809          | 15628.748               |
    #  | lp2   | 0.920191          | 50000.000               |

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------


    # 3rd period: 2 LPs decrease commitment, positive growth:
    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share    | average entry valuation           |
      | lp1   | 0.0794314813996635   | 15985.90375384367775              |
      | lp2   | 0.9205685186003365   | 50000                             |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079809          | 15628.748               |
    #  | lp2   | 0.920191          | 50000.000               |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | sell | MID              | 3          | 1      | amendment |
    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp2 | lp2   | ETH/MAR22 | 45000             | 0.002 | buy  | BID              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 45000             | 0.002 | buy  | MID              | 3          | 1      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 45000             | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 45000             | 0.002 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0620474574136127 | 15985.90375384367775    |
      | lp2   | 0.9379525425863873 | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.062348          | 15628.748               |
    #  | lp2   | 0.937652          | 50000.000               |

    # -------------------------------------------------------------------------------------------------------------------

    Given the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 53     | 1001  | 2                | TYPE_LIMIT | TIF_GTC |

    # CALCULATION:
    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(53 * 1001 * 0.002) = ceil(106.106) = 107
    
    When the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 108    | USD   |
    
    Then the accumulated liquidity fees should be "108" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 502       | 1500      | 20620        | 48000          | 206           |

    And the order book should have the following volumes for market "ETH/MAR22":
      | side | price | volume |
      | buy  | 898   | 268    |
      | buy  | 900   | 1      |
      | buy  | 999   | 73     |
      | sell | 1001  | 73     |
      | sell | 1100  | 1      |
      | sell | 1102  | 219    |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before the network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0620474574136127 | 15985.90375384367775    |
      | lp2   | 0.9379525425863873 | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.062348          | 15628.748               |
    #  | lp2   | 0.937652          | 50000.000               |

    # Trigger next liquidity fee distribution without triggering next period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0620474574136127 | 15985.90375384367775    |
      | lp2   | 0.9379525425863873 | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.062348          | 15628.748               |
    #  | lp2   | 0.937652          | 50000.000               |

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 6      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 102    | USD   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------
    
    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0620474574136127 | 15985.90375384367775    |
      | lp2   | 0.9379525425863873 | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.062348          | 15628.748               |
    #  | lp2   | 0.937652          | 50000.000               |
    
    # Trigger entry into next market period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0620474574136127 | 15985.90375384367775    |
      | lp2   | 0.9379525425863873 | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.062348          | 15628.748               |
    #  | lp2   | 0.937652          | 50000.000               |

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------    


    # 4nd period: 2 LPs increase commitment, positive growth:
    # ------------------------------------------------------------------------------------------------------------------- 

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0620474574136127 | 15985.90375384367775    |
      | lp2   | 0.9379525425863873 | 50000                   |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.062348          | 15628.748               |
    #  | lp2   | 0.937652          | 50000.000               |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 4000              | 0.001 | sell | MID              | 3          | 1      | amendment |
    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | buy  | BID              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | buy  | MID              | 3          | 1      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 46000             | 0.002 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0776542975507553 | 27226.0319392486111616  |
      | lp2   | 0.9223457024492447 | 57008.0995559105993146  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079354          | 24588.910               |
    #  | lp2   | 0.920646          | 50031.943               |

    # ------------------------------------------------------------------------------------------------------------------- 

    Given the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 54     | 1001  | 2                | TYPE_LIMIT | TIF_GTC |

    # CALCULATION:
    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(1001 * 4 * 0.002) = ceil(108.108) = 109

    When the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 110    | USD   |

    Then the accumulated liquidity fees should be "110" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 502       | 1500      | 26026        | 50000          | 260           |

    And the order book should have the following volumes for market "ETH/MAR22":
      | side | price | volume |
      | buy  | 898   | 280    |
      | buy  | 900   | 1      |
      | buy  | 999   | 77     |
      | sell | 1001  | 75     |
      | sell | 1100  | 1      |
      | sell | 1102  | 228    |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0776542975507553 | 27226.0319392486111616  |
      | lp2   | 0.9223457024492447 | 57008.0995559105993146  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079354          | 24588.910               |
    #  | lp2   | 0.920646          | 50031.943               |

    # Trigger next liquidity fee distribution without triggering next period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0776542975507553 | 27226.0319392486111616  |
      | lp2   | 0.9223457024492447 | 57008.0995559105993146  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079354          | 24588.910               |
    #  | lp2   | 0.920646          | 50031.943               |

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 8      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 102    | USD   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------
    
    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0776542975507553 | 27226.0319392486111616  |
      | lp2   | 0.9223457024492447 | 57008.0995559105993146  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079354          | 24588.910               |
    #  | lp2   | 0.920646          | 50031.943               |
    
    # Trigger entry into next market period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0776542975507553 | 27226.0319392486111616  |
      | lp2   | 0.9223457024492447 | 57008.0995559105993146  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079354          | 24588.910               |
    #  | lp2   | 0.920646          | 50031.943               |

    # -------------------------------------------------------------------------------------------------------------------
    # ------------------------------------------------------------------------------------------------------------------- 


    # 5th period: 1 LPs decrease commitment 1 LPs increase commitment, some trades occur
    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0776542975507553 | 27226.0319392486111616  |
      | lp2   | 0.9223457024492447 | 57008.0995559105993146  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.079354          | 24588.910               |
    #  | lp2   | 0.920646          | 50031.943               |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 3000              | 0.001 | sell | MID              | 3          | 1      | amendment |
    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp2 | lp2   | ETH/MAR22 | 47000             | 0.002 | buy  | BID              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 47000             | 0.002 | buy  | MID              | 3          | 1      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 47000             | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/MAR22 | 47000             | 0.002 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0583937269090908 | 27226.0319392486111616  |
      | lp2   | 0.9416062730909092 | 67882.1196053417321887  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.059551          | 24588.910               |
    #  | lp2   | 0.940449          | 50072.555               |

    # -------------------------------------------------------------------------------------------------------------------

    Given the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 55     | 1001  | 2                | TYPE_LIMIT | TIF_GTC |

    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(55 * 1001 * 0.002) = ceil(110.110) = 112

    When the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 112    | USD   |

    # CALCULATION:
    Then the accumulated liquidity fees should be "112" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 502       | 1500      | 31531        | 50000          | 315           |

    And the order book should have the following volumes for market "ETH/MAR22":
      | side | price | volume |
      | buy  | 898   | 279    |
      | buy  | 900   | 1      |
      | buy  | 999   | 76     |
      | sell | 1001  | 76     |
      | sell | 1100  | 1      |
      | sell | 1102  | 228    |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0583937269090908 | 27226.0319392486111616  |
      | lp2   | 0.9416062730909092 | 67882.1196053417321887  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.059551          | 24588.910               |
    #  | lp2   | 0.940449          | 50072.555               |

    # Trigger next liquidity fee distribution without triggering next period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0583937269090908 | 27226.0319392486111616  |
      | lp2   | 0.9416062730909092 | 67882.1196053417321887  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.059551          | 24588.910               |
    #  | lp2   | 0.940449          | 50072.555               |

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 6      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 106    | USD   |

    Then the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------
    
    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0583937269090908 | 27226.0319392486111616  |
      | lp2   | 0.9416062730909092 | 67882.1196053417321887  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.059551          | 24588.910               |
    #  | lp2   | 0.940449          | 50072.555               |
    
    # Trigger entry into next market period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.0583937269090908 | 27226.0319392486111616  |
      | lp2   | 0.9416062730909092 | 67882.1196053417321887  |

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share | average entry valuation |
    #  | lp1   | 0.059551          | 24588.910               |
    #  | lp2   | 0.940449          | 50072.555               |

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------

  @VirtStake
  Scenario: 002 2 LPs joining at start, unequal commitments. Checking calculation of equity-like-shares and liquidity-fee-distribution in a market with large growth (0042-LIQF-008 0042-LIQF-011)

  # Scenario has 6 market periods:

    # - 0th period (bootstrap period): no LP changes, no trades
    # - 1st period: 1 LPs increase commitment, some trades occur
    # - 2nd period: 1 LPs increase commitment, some trades occur


    # Scenario moves ahead to next market period by:

    # - moving ahead "1" blocks to trigger the next liquidity distribution
    # - moving ahead "1" blocks to trigger the next market period


    # Following checks occur in each market where trades:

    # - Check transfers from the price taker to the market-liquidity-pool are correct
    # - Check accumulated-liquidity-fees are non-zero and correct
    # - Check equity-like-shares are correct
    # - Check transfers from the market-liquidity-pool to the liquidity-providers are correct
    # - Check accumulated-liquidity-fees are zero

    Given the average block duration is "1801"

    And the parties deposit on asset's general account the following amount:
      | party  | asset | amount    |
      | lp1    | USD   | 100000000 |
      | lp2    | USD   | 100000000 |
      | party1 | USD   | 100000    |
      | party2 | USD   | 100000    |

    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | lp1   | ETH/MAR22 | 10000             | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | lp1   | ETH/MAR22 | 10000             | 0.001 | buy  | MID              | 3          | 1      | amendment  |
      | lp1 | lp1   | ETH/MAR22 | 10000             | 0.001 | sell | ASK              | 1          | 2      | amendment  |
      | lp1 | lp1   | ETH/MAR22 | 10000             | 0.001 | sell | MID              | 3          | 1      | amendment  |

    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp2 | lp2   | ETH/MAR22 | 20000             | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp2 | lp2   | ETH/MAR22 | 20000             | 0.001 | buy  | MID              | 3          | 1      | amendment  |
      | lp2 | lp2   | ETH/MAR22 | 20000             | 0.001 | sell | ASK              | 1          | 2      | amendment  |
      | lp2 | lp2   | ETH/MAR22 | 20000             | 0.001 | sell | MID              | 3          | 1      | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/MAR22 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/MAR22 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/MAR22 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |


    # 0th period (bootstrap period): no LP changes, no trades
    Then the opening auction period ends for market "ETH/MAR22"

    And the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1000  | 10   | party2 |

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 1       | 500       | 1500      | 1000         | 30000          | 10            |
    
    And the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3333333333333333 | 10000                   |
      | lp2   | 0.6666666666666667 | 20000                   |

    # ERROR:
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Confirm equity-like-shares updated immediately after liquidity amendment
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share  | average entry valuation |
    #  | lp1   | 0.3333333333333333 | 10000                   |
    #  | lp2   | 0.6666666666666667 | 30000                   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"


    When the network moves ahead "2" blocks:

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------


    # 1st period: 1 LPs increase commitment, positive growth:
    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3333333333333333 | 10000                   |
      | lp2   | 0.6666666666666667 | 20000                   |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 11000             | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 11000             | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 11000             | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 11000             | 0.001 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3548387096774194 | 11000                   |
      | lp2   | 0.6451612903225806 | 20000                   |

    # ERROR:
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Confirm equity-like-shares updated immediately after liquidity amendment
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share  | average entry valuation |
    #  | lp1   | 0.3548387096774194 | 11909.091               |
    #  | lp2   | 0.6451612903225806 | 30000                   |

    # -------------------------------------------------------------------------------------------------------------------

    Given the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 16     | 1001  | 1                | TYPE_LIMIT | TIF_GTC |

    Then the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1001  | 16   | lp2    |

    # CALCULATION:
    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(16 * 1001 * 0.001) = ceil(16.016) = 17

    And the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 17     | USD   |

    And the accumulated liquidity fees should be "17" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 500       | 1500      | 2602         | 31000          | 26            |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3548387096774194 | 11000                   |
      | lp2   | 0.6451612903225806 | 20000                   |

    # Trigger next liquidity fee distribution without triggering next market period
    And the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3548387096774194 | 11000                   |
      | lp2   | 0.6451612903225806 | 20000                   |

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 6      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 11     | USD   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3548387096774194 | 11000                   |
      | lp2   | 0.6451612903225806 | 20000                   |
    
    # Trigger entry into next market period
    And the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3548387096774194 | 11000                   |
      | lp2   | 0.6451612903225806 | 20000                   |

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------


    # 2nd period: 1 LPs increase commitment, positive growth:
    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before liquidity amendment
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation |
      | lp1   | 0.3548387096774194 | 11000                   |
      | lp2   | 0.6451612903225806 | 20000                   |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | lp1   | ETH/MAR22 | 12000             | 0.001 | buy  | BID              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 12000             | 0.001 | buy  | MID              | 3          | 1      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 12000             | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/MAR22 | 12000             | 0.001 | sell | MID              | 3          | 1      | amendment |

    # Confirm equity-like-shares updated immediately after liquidity amendment
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation        |
      | lp1   | 0.3641574309414367 | 11653.037790125972454258802688 |
      | lp2   | 0.6358425690585633 | 20000                          |         

    # ERROR:
    # - "equity-like-share" values calculated incorrectly.
    # - "average-entry-valuation" values calculated incorrectly.

    # EXPECTED:
    # Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
    #  | party | equity like share  | average entry valuation |
    #  | lp1   | 0.3704506736874710 | 14360.400               |
    #  | lp2   | 0.6295493263125290 | 30000.000               |

    # -------------------------------------------------------------------------------------------------------------------

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 25     | 1001  | 1                | TYPE_LIMIT | TIF_GTC |

    Then the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 1001  | 25   | lp2    |

    # CALCULATION:
    # liquidity_fee = ceil(volume * price * liquidity_fee_factor) =  ceil(25 * 1001 * 0.001) = ceil(25.025) = 26

    And the following transfers should happen:
      | from   | to     | from account           | to account                  | market id | amount | asset |
      | party1 | market | ACCOUNT_TYPE_GENERAL   | ACCOUNT_TYPE_FEES_LIQUIDITY | ETH/MAR22 | 26    | USD   |

    And the accumulated liquidity fees should be "26" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1001       | TRADING_MODE_CONTINUOUS | 1       | 502       | 1500      | 5105         | 32000          | 51            |

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    Given the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation        |
      | lp1   | 0.3641574309414367 | 11653.037790125972454258802688 |
      | lp2   | 0.6358425690585633 | 20000                          |  

    # Trigger next liquidity fee distribution without triggering next period
    When the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by network moving forwards (as new market period not entered)
    Then the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation        |
      | lp1   | 0.3641574309414367 | 11653.037790125972454258802688 |
      | lp2   | 0.6358425690585633 | 20000                          |  

    And the following transfers should happen:
      | from   | to  | from account                | to account           | market id | amount | asset |
      | market | lp1 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 9      | USD   |
      | market | lp2 | ACCOUNT_TYPE_FEES_LIQUIDITY | ACCOUNT_TYPE_GENERAL | ETH/MAR22 | 17     | USD   |

    And the accumulated liquidity fees should be "0" for the market "ETH/MAR22"

    # -------------------------------------------------------------------------------------------------------------------

    # Check equity-like-shares before network moves forward
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation        |
      | lp1   | 0.3641574309414367 | 11653.037790125972454258802688 |
      | lp2   | 0.6358425690585633 | 20000                          |  
    
    # Trigger entry into next market period
    And the network moves ahead "1" blocks:

    # Confirm equity-like-shares are unchanged by the network moving forwards (as virtual-stakes scaled by same factor, r)
    When the liquidity provider fee shares for the market "ETH/MAR22" should be:
      | party | equity like share  | average entry valuation        |
      | lp1   | 0.3641574309414367 | 11653.037790125972454258802688 |
      | lp2   | 0.6358425690585633 | 20000                          |  

    # -------------------------------------------------------------------------------------------------------------------
    # -------------------------------------------------------------------------------------------------------------------
