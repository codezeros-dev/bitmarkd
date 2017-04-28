// Copyright (c) 2014-2017 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package constants

import (
	"time"
)

// the time for a record to expire
const (
	ReservoirTimeout = 24 * time.Hour
)

// the maximum time before unverified asset is expired
const (
	AssetTimeout = ReservoirTimeout + time.Hour
)

// the time between rebroadcasts of unconfirment transactions
const (
	RebroadcastInterval = 15 * time.Minute
)