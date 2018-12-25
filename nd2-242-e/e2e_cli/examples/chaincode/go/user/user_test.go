package main

import (
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1",args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", "failed", string(res.Message))
		t.FailNow()
	}
}


func TestUser_Invoke(t *testing.T) {
	scc := new(UserChainCode)
	stub := shim.NewMockStub("ex02", scc)

	checkInvoke(t,stub,[][]byte{[]byte("addUserInfo"), []byte("1"), []byte("1"), []byte("1"), []byte("1"), []byte("1")})

	checkInvoke(t,stub,[][]byte{[]byte("getUserInfo"), []byte("1")})

}
