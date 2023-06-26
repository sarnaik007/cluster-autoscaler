/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package estimator

import "time"

type thresholdBinpackingLimit struct {
	maxNodes    int
	maxDuration time.Duration
}

func (l *thresholdBinpackingLimit) GetNodeLimit() int {
	return l.maxNodes
}

func (l *thresholdBinpackingLimit) GetDurationLimit() time.Duration {
	return l.maxDuration
}

// NewThresholdBinpackingLimit returns a BinpackingLimit that caps maximum node
// count and maximum duration of binpacking by given static values
func NewThresholdBinpackingLimit(maxNodes int, maxDuration time.Duration) BinpackingLimit {
	return &thresholdBinpackingLimit{
		maxNodes:    maxNodes,
		maxDuration: maxDuration,
	}
}
