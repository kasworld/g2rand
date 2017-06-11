// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package g2rand

import (
	"testing"
)

func TestNormIntRange(t *testing.T) {
	rnd := New()
	min, max := 0, 0
	mean, dev := 100, 50
	for i := 0; i < 100; i++ {
		v := rnd.NormIntRange(mean, dev)
		println(v)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	println(min, max)
	println(mean-min, max-mean)
}
