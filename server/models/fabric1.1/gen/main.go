// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"go/scanner"
	//"go/token"
	"io/ioutil"

	"github.com/dvyukov/go-fuzz/gen"
)

func main() {
	str, err := ioutil.ReadFile("case") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	data := []byte(str)
	gen.Emit(data, nil, true)
}
