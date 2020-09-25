// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"fmt"

	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var arg [][]byte

func Fuzz(data []byte) int {
	cc := new(Chaincode)               // 创建Chaincode对象
	stub := shim.NewMockStub("source", cc) // 创建MockStub对象

	var i interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	return Assert(stub, i)
}

var K string

func Assert(stub *shim.MockStub, i interface{}) int {
	switch t := i.(type) {
	case map[string]interface{}:
		for k, v := range t {
			switch t1 := v.(type) {
			case map[string]interface{}:
				//fmt.Println(k, " : ")
				arg = append(arg, []byte(k))
				K = k
				Assert(stub, t1)
			case []interface{}:
				//fmt.Println(k, " : ")
				for k1, v1 := range t1 {
					switch t2 := v1.(type) {
					case map[string]interface{}:
						fmt.Println(k1, " : ")
						Assert(stub, t2)
					default:
						fmt.Println(k1, " : ", v1)
					}
				}
			default:
				//fmt.Println(k, " : ", v)
				p, ok := v.(string)
				if ok != true {
					//fmt.Println("err")
				}
				arg = append(arg, []byte(p))
			}
		}
		//fmt.Println(arg)
		if K == "Init" {
			res := stub.MockInit("1", [][]byte{arg[1]})
			if res.Status != 200 {
				return 0
			}
		} else {
			res := stub.MockInvoke("1", arg)
			if res.Status != 200 {
				return 0
			}
		}
		arg = [][]byte{}
	}
	return 1
}
