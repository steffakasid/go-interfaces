package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoFunctionReturnsStruct1(t *testing.T) {
	someIface := doSomething(true)
	assert.IsType(t, &struct1{}, someIface)
	assert.Implements(t, (*realIface)(nil), someIface)
}

func TestDoFunctionReturnsStruct2(t *testing.T) {
	someIface := doSomething(false)
	assert.IsType(t, &struct2{}, someIface)
	assert.Implements(t, (*realIface)(nil), someIface)
}

func TestIface1(t *testing.T) {
	iFace = &struct1{"Hans"} // must use the reference to the object here!
	assert.IsType(t, &struct1{}, iFace)
	assert.Implements(t, (*realIface)(nil), iFace)
}

func TestRealIface(t *testing.T) {
	var iFace realIface
	iFace = &struct1{"Hans"} // must use the reference to the object here!
	assert.IsType(t, &struct1{}, iFace)
	assert.Implements(t, (*realIface)(nil), iFace)
}
