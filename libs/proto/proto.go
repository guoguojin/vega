// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package proto

import "github.com/golang/protobuf/proto"

func Marshal(m proto.Message) ([]byte, error) {
	buf := proto.NewBuffer(nil)
	buf.SetDeterministic(true)
	if err := buf.Marshal(m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}
