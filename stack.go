// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import "loov.dev/lensm/internal/go/buildcfg"

const (
	STACKSYSTEM = 0
	StackSystem = STACKSYSTEM
	StackBig    = 4096
	StackSmall  = 128
)

// Initialize StackGuard and StackLimit according to target system.
var StackGuard = 928*stackGuardMultiplier() + StackSystem
var StackLimit = StackGuard - StackSystem - StackSmall

// stackGuardMultiplier returns a multiplier to apply to the default
// stack guard size. Larger multipliers are used for non-optimized
// builds that have larger stack frames or for specific targets.
func stackGuardMultiplier() int {
	if buildcfg.GOOS == "aix" {
		return 2
	}
	return stackGuardMultiplierDefault
}
