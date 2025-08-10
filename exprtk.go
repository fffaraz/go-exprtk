package exprtk

// #cgo CXXFLAGS: -flto -fuse-linker-plugin -std=c++11
// #cgo LDFLAGS: -L.
// #include <stdlib.h>
// #include "exprtkwrapper.h"
import "C"

import (
	"errors"
	"unsafe"
)

// GoExprtk ...Exprtk Structure
type GoExprtk struct {
	exprtk C.exprtkWrapper
}

// NewExprtk ... Creates a new object
func NewExprtk() GoExprtk {
	var obj GoExprtk
	obj.exprtk = C.exprtkWrapperInit()
	return obj
}

// SetExpression ... Sets an Expression
func (obj GoExprtk) SetExpression(expr string) {
	cExpr := C.CString(expr)
	defer C.free(unsafe.Pointer(cExpr))
	C.setExpressionString(obj.exprtk, cExpr)
}

// AddDoubleVariable ... Adds variable to the expression
func (obj GoExprtk) AddDoubleVariable(x string) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))
	C.addDoubleVariable(obj.exprtk, cX)
}

// AddStringVariable ... Adds variable to the expression
func (obj GoExprtk) AddStringVariable(x string) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))
	C.addStringVariable(obj.exprtk, cX)
}

// AddVectorVariable ... Adds variable to the expression
func (obj GoExprtk) AddVectorVariable(x string) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))
	C.addVectorVariable(obj.exprtk, cX)
}

// SetDoubleVariableValue ... Sets value to the variable
func (obj GoExprtk) SetDoubleVariableValue(varName string, val float64) {
	cVarName := C.CString(varName)
	defer C.free(unsafe.Pointer(cVarName))
	C.setDoubleVariableValue(obj.exprtk, cVarName, C.double(val))
}

// SetStringVariableValue ... Sets value to the variable
func (obj GoExprtk) SetStringVariableValue(varName, val string) {
	cVarName := C.CString(varName)
	defer C.free(unsafe.Pointer(cVarName))
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	C.setStringVariableValue(obj.exprtk, cVarName, cVal)
}

// SetVectorVariableValue ... Sets value to the variable
func (obj GoExprtk) SetVectorVariableValue(varName string, val []float64) {
	arr := make([]C.double, 0)
	for i := 0; i < len(val); i++ {
		arr = append(arr, C.double(val[i]))
	}
	firstValue := &(arr[0])
	var arrayLength C.int = C.int(len(arr))
	cVarName := C.CString(varName)
	defer C.free(unsafe.Pointer(cVarName))
	C.setVectorVariableValue(obj.exprtk, cVarName, firstValue, arrayLength)
}

// CompileExpression ... Compiles the Expression
func (obj GoExprtk) CompileExpression() error {
	value := C.compileExpression(obj.exprtk)
	if value == 0 {
		return errors.New("failed to compile the expression")
	}
	return nil
}

// GetEvaluatedValue ... Returns the evaluated value
func (obj GoExprtk) GetEvaluatedValue() float64 {
	return float64(C.getEvaluatedValue(obj.exprtk))
}

// Delete ... Destroys the created object and releases the memory
func (obj GoExprtk) Delete() {
	C.deleteExprtk(obj.exprtk)
}
