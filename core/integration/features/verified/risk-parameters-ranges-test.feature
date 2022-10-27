Feature: test risk model parameter ranges
  Background:

    Given the log normal risk model named "log-normal-risk-model-0":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 0.1 | 0  | 0 | 1.0   |
    #risk factor short:3.5569036
    #risk factor long:0.801225765

    # test risk aversion
    Given the log normal risk model named "log-normal-risk-model-11":
      | risk aversion | tau | mu | r | sigma |
      | 0.00000001    | 0.1 | 0  | 0 | 1.0   |
    #risk factor short:4.9256840
    #risk factor long:0.847272675
    Given the log normal risk model named "log-normal-risk-model-12":
      | risk aversion | tau | mu | r | sigma |
      | 0.9           | 0.1 | 0  | 0 | 1.0   |
    #risk factor short:0.0499488
    #risk factor long:0.076967499

    # test tau
    Given the log normal risk model named "log-normal-risk-model-21":
      | risk aversion | tau        | mu | r | sigma |
      | 0.000001      | 0.00000001 | 0  | 0 | 1.0   |
    #risk factor short:0.0004950
    #risk factor long:0.000494716
    Given the log normal risk model named "log-normal-risk-model-22":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 0.1 | 0  | 0 | 1.0   |
    #risk factor short:3.5569036
    #risk factor long:0.801225765
    Given the log normal risk model named "log-normal-risk-model-23":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 1   | 0  | 0 | 1.0   |
    #risk factor short:86.2176101 (it can not end the auction so this is tested separately in scenario 002)
    #risk factor long:0.996594553

    # test mu
    Given the log normal risk model named "log-normal-risk-model-31":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 0.1 | -2 | 0 | 1.0   |
    #risk factor short:2.7308771
    #risk factor long:0.803203602
    # actual mu = -2*0.1=-0.2
    Given the log normal risk model named "log-normal-risk-model-32":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 0.1 | 1  | 0 | 1.0   |
    #risk factor short:4.0361573
    #risk factor long:0.800229405
    # actual mu = 1*0.1=0.1
    Given the log normal risk model named "log-normal-risk-model-33":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 0.1 | 2  | 0 | 1.0   |
    #risk factor short:4.5658146
    #risk factor long:0.799228051
    # actual mu = 2*0.1=0.2
    Given the log normal risk model named "log-normal-risk-model-34":
      | risk aversion | tau | mu  | r | sigma |
      | 0.000001      | 0.1 | -20 | 0 | 1.0   |
    #risk factor short:-0.3832902
    #risk factor long:0.820141635
    # actual mu = -20*0.1=-2
    # mu=-20 will make the prob_of_trading for any price over mid price (1000) very small (kept at 1.00E-08), hence LP vol extremely large, it can not end the auction so this is tested separately in scenario 002
    Given the log normal risk model named "log-normal-risk-model-35":
      | risk aversion | tau | mu  | r | sigma |
      | 0.000001      | 0.1 | -20 | 0 | 1.0   |
    #risk factor short:32.6712163
    #risk factor long:0.780320497
    # actual mu = 20*0.1=2
    # mu=-20 will make the prob_of_trading for any price over mid price (1000) very small (kept at 1.00E-08), hence LP vol extremely large, it can not end the auction so this is tested separately in scenario 002

    # test r
    Given the log normal risk model named "log-normal-risk-model-41":
      | risk aversion | tau | mu | r   | sigma |
      | 0.000001      | 0.1 | 0  | -20 | 1.0   |
    #risk factor short:3.5569036
    #risk factor long:0.801225765
    Given the log normal risk model named "log-normal-risk-model-42":
      | risk aversion | tau | mu | r   | sigma |
      | 0.000001      | 0.1 | 0  | 0.5 | 1.0   |
    #risk factor short:3.5569036
    #risk factor long:0.801225765
    Given the log normal risk model named "log-normal-risk-model-43":
      | risk aversion | tau | mu | r  | sigma |
      | 0.000001      | 0.1 | 0  | 20 | 1.0   |
    #risk factor short:3.5569036
    #risk factor long:0.801225765

    # test sigma
    Given the log normal risk model named "log-normal-risk-model-51":
      | risk aversion | tau | mu | r | sigma |
      | 0.000001      | 0.1 | 0  | 0 | 1     |
    #risk factor short:3.5569036
    #risk factor long:0.801225765
    # tested in scenario 002
    Given the log normal risk model named "log-normal-risk-model-52":
      | risk aversion | tau | mu | r  | sigma |
      | 0.000001      | 0.1 | 0  | 20 | 100   |
    #risk factor short:3.5569036
    #risk factor long:0.801225765
    # tested in scenario 002

    And the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.004     | 0.001              |
    And the price monitoring named "price-monitoring-1":
      | horizon | probability | auction extension |
      | 43200   | 0.99        | 300               |

    And the markets:
      | id        | quote name | asset | risk model               | margin calculator         | auction duration | fees          | price monitoring   | oracle config          |
      | ETH/MAR0  | ETH        | USD   | log-normal-risk-model-0  | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR11 | ETH        | USD   | log-normal-risk-model-11 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR12 | ETH        | USD   | log-normal-risk-model-12 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR21 | ETH        | USD   | log-normal-risk-model-21 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR22 | ETH        | USD   | log-normal-risk-model-22 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR23 | ETH        | USD   | log-normal-risk-model-23 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR31 | ETH        | USD   | log-normal-risk-model-31 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR32 | ETH        | USD   | log-normal-risk-model-32 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR33 | ETH        | USD   | log-normal-risk-model-33 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR34 | ETH        | USD   | log-normal-risk-model-34 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR35 | ETH        | USD   | log-normal-risk-model-35 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR41 | ETH        | USD   | log-normal-risk-model-41 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR42 | ETH        | USD   | log-normal-risk-model-42 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR43 | ETH        | USD   | log-normal-risk-model-43 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR51 | ETH        | USD   | log-normal-risk-model-52 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
      | ETH/MAR52 | ETH        | USD   | log-normal-risk-model-52 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |

    And the parties deposit on asset's general account the following amount:
      | party  | asset | amount         |
      | party0 | USD   | 50000000000000 |
      | party1 | USD   | 50000000000000 |
      | party2 | USD   | 50000000000000 |
      | party3 | USD   | 50000000000000 |

  Scenario: 001, test different value of risk parameters within defined ranges

    And the following network parameters are set:
      | name                                          | value |
      | market.stake.target.timeWindow                | 24h   |
      | market.stake.target.scalingFactor             | 1     |
      | market.liquidity.bondPenaltyParameter         | 0.2   |
      | market.liquidity.targetstake.triggering.ratio | 0.1   |

    And the average block duration is "1"

    And the parties submit the following liquidity provision:
      | id   | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1  | party0 | ETH/MAR0  | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
      | lp1  | party0 | ETH/MAR0  | 50000             | 0.001 | buy  | BID              | 500        | 20     | amendment  |
      | lp2  | party0 | ETH/MAR11 | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
      | lp2  | party0 | ETH/MAR11 | 50000             | 0.001 | buy  | BID              | 500        | 20     | amendment  |
      | lp3  | party0 | ETH/MAR12 | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
      | lp3  | party0 | ETH/MAR12 | 50000             | 0.001 | buy  | BID              | 500        | 20     | amendment  |
      | lp4  | party0 | ETH/MAR21 | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
      | lp4  | party0 | ETH/MAR21 | 50000             | 0.001 | buy  | BID              | 500        | 20     | amendment  |
      | lp5  | party0 | ETH/MAR22 | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
      | lp5  | party0 | ETH/MAR22 | 50000             | 0.001 | buy  | BID              | 500        | 20     | amendment  |
      | lp6  | party0 | ETH/MAR31 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp6  | party0 | ETH/MAR31 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp7  | party0 | ETH/MAR32 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp7  | party0 | ETH/MAR32 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp8  | party0 | ETH/MAR33 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp8  | party0 | ETH/MAR33 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp9  | party0 | ETH/MAR41 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp9  | party0 | ETH/MAR41 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp10 | party0 | ETH/MAR42 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp10 | party0 | ETH/MAR42 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp11 | party0 | ETH/MAR43 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp11 | party0 | ETH/MAR43 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp12 | party0 | ETH/MAR51 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp12 | party0 | ETH/MAR51 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |
      | lp13 | party0 | ETH/MAR52 | 50000             | 0.001 | buy  | BID              | 500        | 20     | submission |
      | lp13 | party0 | ETH/MAR52 | 50000             | 0.001 | sell | ASK              | 500        | 20     | amendment  |


    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/MAR0  | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-11  |
      | party1 | ETH/MAR0  | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-12  |
      | party2 | ETH/MAR0  | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-13 |
      | party2 | ETH/MAR0  | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-14 |
      | party1 | ETH/MAR11 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-21  |
      | party1 | ETH/MAR11 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-22  |
      | party2 | ETH/MAR11 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-23 |
      | party2 | ETH/MAR11 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-24 |
      | party1 | ETH/MAR12 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-31  |
      | party1 | ETH/MAR12 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-32  |
      | party2 | ETH/MAR12 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-33 |
      | party2 | ETH/MAR12 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-34 |
      | party1 | ETH/MAR21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-41  |
      | party1 | ETH/MAR21 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-42  |
      | party2 | ETH/MAR21 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-43 |
      | party2 | ETH/MAR21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-44 |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/MAR22 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-11  |
      | party1 | ETH/MAR22 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-12  |
      | party2 | ETH/MAR22 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-13 |
      | party2 | ETH/MAR22 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-14 |
      | party1 | ETH/MAR31 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-21  |
      | party1 | ETH/MAR31 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-22  |
      | party2 | ETH/MAR31 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-23 |
      | party2 | ETH/MAR31 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-24 |
      | party1 | ETH/MAR32 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-31  |
      | party1 | ETH/MAR32 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-32  |
      | party2 | ETH/MAR32 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-33 |
      | party2 | ETH/MAR32 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-34 |
      | party1 | ETH/MAR33 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-41  |
      | party1 | ETH/MAR33 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-42  |
      | party2 | ETH/MAR33 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-43 |
      | party2 | ETH/MAR33 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-44 |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/MAR41 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-11  |
      | party1 | ETH/MAR41 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-12  |
      | party2 | ETH/MAR41 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-13 |
      | party2 | ETH/MAR41 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-14 |
      | party1 | ETH/MAR42 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-21  |
      | party1 | ETH/MAR42 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-22  |
      | party2 | ETH/MAR42 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-23 |
      | party2 | ETH/MAR42 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-24 |
      | party1 | ETH/MAR43 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-31  |
      | party1 | ETH/MAR43 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-32  |
      | party2 | ETH/MAR43 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-33 |
      | party2 | ETH/MAR43 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-34 |
    # | party1 | ETH/MAR51 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-41  |
    # | party1 | ETH/MAR51 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-42  |
    # | party2 | ETH/MAR51 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-43 |
    # | party2 | ETH/MAR51 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-44 |

    # And the parties place the following orders:
    #   | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
    #   | party1 | ETH/MAR52 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-11  |
    #   | party1 | ETH/MAR52 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-12  |
    #   | party2 | ETH/MAR52 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-13 |
    #   | party2 | ETH/MAR52 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-14 |

    When the opening auction period ends for market "ETH/MAR0"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR0"
    When the opening auction period ends for market "ETH/MAR11"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR11"
    When the opening auction period ends for market "ETH/MAR12"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR12"
    When the opening auction period ends for market "ETH/MAR21"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR21"
    When the opening auction period ends for market "ETH/MAR22"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR22"
    When the opening auction period ends for market "ETH/MAR31"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR31"
    When the opening auction period ends for market "ETH/MAR32"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR32"
    When the opening auction period ends for market "ETH/MAR33"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR33"

    When the opening auction period ends for market "ETH/MAR41"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR41"
    When the opening auction period ends for market "ETH/MAR42"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR42"
    When the opening auction period ends for market "ETH/MAR43"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR43"

    # When the opening auction period ends for market "ETH/MAR51"
    # And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR51"
    # When the opening auction period ends for market "ETH/MAR52"
    # And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR52"

    And the market data for the market "ETH/MAR0" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 35569        | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 3.5569036 =35569
    And the market data for the market "ETH/MAR11" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 49256        | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 4.9256840 = 49256
    And the market data for the market "ETH/MAR12" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 746          | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.076967499 = 746
    And the market data for the market "ETH/MAR21" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 4            | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.0004950 = 4
    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 35569        | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 3.5569036 =35569
    And the market data for the market "ETH/MAR31" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 906       | 1096      | 27308        | 50000          | 10            |
    # target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 2.7308771 = 27308
    And the market data for the market "ETH/MAR32" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 910       | 1100      | 40361        | 50000          | 10            |
    # target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 4.0361573 = 40361
    And the market data for the market "ETH/MAR33" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 911       | 1102      | 45658        | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 4.5658146 = 45658

    And the market data for the market "ETH/MAR41" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 35569        | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.16882368861315200 =1689
    And the market data for the market "ETH/MAR42" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 35569        | 50000          | 10            |
    # target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.36483236867768200 = 3648
    And the market data for the market "ETH/MAR43" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 35569        | 50000          | 10            |
    # target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.13281340025639400 = 1328
    # And the market data for the market "ETH/MAR51" should be:
    #   | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
    #   | 1000       | TRADING_MODE_CONTINUOUS | 43200   | 909       | 1099      | 35569        | 50000          | 10            |
    #target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.13281340025639400 = 1328

    # And the market data for the market "ETH/MAR52" should be:
    #   | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
    #   | 1000       | TRADING_MODE_CONTINUOUS | 43200    | 909       | 1099      | 1328         | 50000          | 10            |
    # # target_stake = mark_price x max_oi x target_stake_scaling_factor x rf_short = 1000 x 10 x 1 x 0.13281340025639400 = 1328

    Then the order book should have the following volumes for market "ETH/MAR0":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR11":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR12":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR21":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR31":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR32":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR33":
      | side | price | volume |
      | sell | 1100  | 1      |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR41":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR42":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |
    Then the order book should have the following volumes for market "ETH/MAR43":
      | side | price | volume |
      | sell | 1100  | 92     |
      | sell | 1120  | 0      |
      | buy  | 900   | 113    |
      | buy  | 880   | 0      |

    And the parties should have the following account balances:
      | party  | asset | market id | margin | general        | bond  |
      | party0 | USD   | ETH/MAR21 | 67     | 49999995616917 | 50000 |
      | party1 | USD   | ETH/MAR21 | 1207   | 49999999890451 | 0     |
      | party2 | USD   | ETH/MAR21 | 1207   | 49999999536766 | 0     |

  Scenario: 002, test mu=-20, and few other odd situations

    And the following network parameters are set:
      | name                                          | value |
      | market.stake.target.timeWindow                | 24h   |
      | market.stake.target.scalingFactor             | 1     |
      | market.liquidity.bondPenaltyParameter         | 0.2   |
      | market.liquidity.targetstake.triggering.ratio | 0.1   |

    And the average block duration is "1"

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/MAR34 | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
      | lp1 | party0 | ETH/MAR34 | 50000             | 0.001 | buy  | BID              | 500        | 20     | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
      | party1 | ETH/MAR34 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-11  |
      | party1 | ETH/MAR34 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-12  |
      | party2 | ETH/MAR34 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-13 |
      | party2 | ETH/MAR34 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-14 |

    When the opening auction period ends for market "ETH/MAR34"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR34"

    Then the parties should have the following margin levels:
      | party  | market id | maintenance | search | initial | release |
      | party0 | ETH/MAR34 | 16270       | 17897  | 19524   | 22778   |
      | party1 | ETH/MAR34 | 2598        | 2857   | 3117    | 3637    |
      | party2 | ETH/MAR34 | 2858        | 3143   | 3429    | 4001    |