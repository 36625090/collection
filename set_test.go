/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import "testing"

func TestSet_Append(t *testing.T) {
	set := NewSet[int](6, 712, 898, 13, 12)
	set.Append(120)
	set.Sorted()
	t.Log(set)
	set.Reverse()
	t.Log(set)
}
