// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/markets.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ContinuousTrading) Validate() error {
	return nil
}
func (this *DiscreteTrading) Validate() error {
	return nil
}
func (this *Future) Validate() error {
	if oneOfNester, ok := this.GetOracle().(*Future_EthereumEvent); ok {
		if oneOfNester.EthereumEvent != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.EthereumEvent); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EthereumEvent", err)
			}
		}
	}
	return nil
}
func (this *EthereumEvent) Validate() error {
	return nil
}
func (this *InstrumentMetadata) Validate() error {
	return nil
}
func (this *Instrument) Validate() error {
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	if oneOfNester, ok := this.GetProduct().(*Instrument_Future); ok {
		if oneOfNester.Future != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Future); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Future", err)
			}
		}
	}
	return nil
}
func (this *ModelParamsBS) Validate() error {
	return nil
}
func (this *Forward) Validate() error {
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *ExternalRiskModel) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SimpleRiskModel) Validate() error {
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *SimpleModelParams) Validate() error {
	return nil
}
func (this *ScalingFactors) Validate() error {
	return nil
}
func (this *MarginCalculator) Validate() error {
	if this.ScalingFactors != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScalingFactors); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScalingFactors", err)
		}
	}
	return nil
}
func (this *TradableInstrument) Validate() error {
	if this.Instrument != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Instrument); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Instrument", err)
		}
	}
	if this.MarginCalculator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginCalculator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginCalculator", err)
		}
	}
	if oneOfNester, ok := this.GetRiskModel().(*TradableInstrument_Forward); ok {
		if oneOfNester.Forward != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Forward); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Forward", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRiskModel().(*TradableInstrument_ExternalRiskModel); ok {
		if oneOfNester.ExternalRiskModel != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ExternalRiskModel); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ExternalRiskModel", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRiskModel().(*TradableInstrument_SimpleRiskModel); ok {
		if oneOfNester.SimpleRiskModel != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SimpleRiskModel); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SimpleRiskModel", err)
			}
		}
	}
	return nil
}
func (this *Market) Validate() error {
	if this.TradableInstrument != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TradableInstrument); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TradableInstrument", err)
		}
	}
	if oneOfNester, ok := this.GetTradingMode().(*Market_Continuous); ok {
		if oneOfNester.Continuous != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Continuous); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Continuous", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTradingMode().(*Market_Discrete); ok {
		if oneOfNester.Discrete != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Discrete); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Discrete", err)
			}
		}
	}
	return nil
}
