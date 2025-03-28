// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package gql

import (
	"context"

	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"
	proto "code.vegaprotocol.io/vega/protos/vega"
)

type nodeResolver VegaResolverRoot

func (r *nodeResolver) RankingScore(ctx context.Context, obj *proto.Node) (proto.RankingScore, error) {
	return *obj.RankingScore, nil
}

func (r *nodeResolver) RewardScore(ctx context.Context, obj *proto.Node) (proto.RewardScore, error) {
	return *obj.RewardScore, nil
}

func (r *nodeResolver) DelegationsConnection(ctx context.Context, node *proto.Node, partyID *string, pagination *v2.Pagination) (*v2.DelegationsConnection, error) {
	var nodeID *string
	if node != nil {
		nodeID = &node.Id
	}
	return handleDelegationConnectionRequest(ctx, r.tradingDataClientV2, partyID, nodeID, nil, pagination)
}
