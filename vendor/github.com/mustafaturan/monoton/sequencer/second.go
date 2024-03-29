// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"
	"math"

	"github.com/mustafaturan/monoton/mtimer"
)

// NewSecond returns the preconfigured second sequencer
func NewSecond() *Sequence {
	second := uint(time.Second)
	return &Sequence{
		now:     func() uint { return mtimer.Now() / second },
		max:     math.MaxUint32,
		maxTime: math.MaxUint32,
	}
}
