// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/governance.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NetworkConfiguration) Validate() error {
	if !(this.MinParticipationStake > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MinParticipationStake", fmt.Errorf(`value '%v' must be greater than '0'`, this.MinParticipationStake))
	}
	return nil
}
func (this *Vote) Validate() error {
	if this.Voter == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Voter", fmt.Errorf(`value '%v' must not be an empty string`, this.Voter))
	}
	if _, ok := Vote_Value_name[int32(this.Value)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Value", fmt.Errorf(`value '%v' must be a valid Vote_Value field`, this.Value))
	}
	if !(this.Stake > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Stake", fmt.Errorf(`value '%v' must be greater than '0'`, this.Stake))
	}
	return nil
}
func (this *Proposal) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if _, ok := Proposal_State_name[int32(this.State)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("State", fmt.Errorf(`value '%v' must be a valid Proposal_State field`, this.State))
	}
	if this.Author == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Author", fmt.Errorf(`value '%v' must not be an empty string`, this.Author))
	}
	if nil == this.Proposal {
		return github_com_mwitkow_go_proto_validators.FieldError("Proposal", fmt.Errorf("message must exist"))
	}
	if this.Proposal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proposal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
		}
	}
	for _, item := range this.Votes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Votes", err)
			}
		}
	}
	return nil
}
func (this *Proposal_Terms) Validate() error {
	if nil == this.Parameters {
		return github_com_mwitkow_go_proto_validators.FieldError("Parameters", fmt.Errorf("message must exist"))
	}
	if this.Parameters != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Parameters); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Parameters", err)
		}
	}
	if oneOfNester, ok := this.GetChange().(*Proposal_Terms_UpdateMarket_); ok {
		if oneOfNester.UpdateMarket != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UpdateMarket); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateMarket", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChange().(*Proposal_Terms_NewMarket_); ok {
		if oneOfNester.NewMarket != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NewMarket); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NewMarket", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChange().(*Proposal_Terms_UpdateNetwork_); ok {
		if oneOfNester.UpdateNetwork != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UpdateNetwork); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateNetwork", err)
			}
		}
	}
	return nil
}
func (this *Proposal_Terms_Parameters) Validate() error {
	if !(this.MinParticipationStake > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MinParticipationStake", fmt.Errorf(`value '%v' must be greater than '0'`, this.MinParticipationStake))
	}
	if !(this.MinParticipationStake < 100) {
		return github_com_mwitkow_go_proto_validators.FieldError("MinParticipationStake", fmt.Errorf(`value '%v' must be less than '100'`, this.MinParticipationStake))
	}
	return nil
}
func (this *Proposal_Terms_UpdateMarket) Validate() error {
	return nil
}
func (this *Proposal_Terms_NewMarket) Validate() error {
	return nil
}
func (this *Proposal_Terms_UpdateNetwork) Validate() error {
	if nil == this.Changes {
		return github_com_mwitkow_go_proto_validators.FieldError("Changes", fmt.Errorf("message must exist"))
	}
	if this.Changes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Changes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Changes", err)
		}
	}
	return nil
}
