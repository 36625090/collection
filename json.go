/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"encoding/json"
)

func Dump(in any) string {
	data, err := json.Marshal(in)
	if nil != err {
		return "ERROR: " + err.Error()
	}
	return string(data)
}
