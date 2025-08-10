package exprtk

import (
	"math"
	"reflect"
	"testing"
)

func TestNewExprtk(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

	if reflect.TypeOf(exprtkObj).String() != "exprtk.GoExprtk" {
		t.Error("NewExprtk returned incorrect type")
	}
}

func TestCompileExpression(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression("x.1 + y.1")

	exprtkObj.AddDoubleVariable("[x.1]")
	exprtkObj.AddDoubleVariable("[y.1]")

	err := exprtkObj.CompileExpression()
	if err.Error() != "failed to compile the expression" {
		t.Error("negative case failed")
	}

	exprtkObj.SetExpression("x + y")

	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	err = exprtkObj.CompileExpression()
	if err != nil {
		t.Error("failed to compile the expression")
	}
}

func TestDoubleVariables(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression("(x + 2)*(y-2)")

	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	err := exprtkObj.CompileExpression()
	if err != nil {
		t.Error(err)
		return
	}

	exprtkObj.SetDoubleVariableValue("x", 18)
	exprtkObj.SetDoubleVariableValue("y", 32)

	if exprtkObj.GetEvaluatedValue() != 600 {
		t.Error("Incorrect Value")
	}
}

func TestMixedVariables(t *testing.T) {
	exprtkObj := NewExprtk()
	defer exprtkObj.Delete()

	eqn := "if (eqn == 'avg') avg(x); "
	eqn += "else if (eqn == 'max') max(x); "
	eqn += "else if (eqn == 'min') min(x); "
	eqn += "else if (eqn == 'sum') sum(x); "
	eqn += "else 0; "

	exprtkObj.SetExpression(eqn)
	exprtkObj.AddStringVariable("eqn")
	exprtkObj.AddVectorVariable("x")

	err := exprtkObj.CompileExpression()
	if err != nil {
		t.Error(err)
		return
	}

	array := []float64{1, 2, 3, -4.3, 10, -6.5, 7, 8, -1.3}
	exprtkObj.SetVectorVariableValue("x", array)

	var result float64

	eqnStr := "avg"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	result = exprtkObj.GetEvaluatedValue()
	if math.Round(result*10)/10 != 2.1 {
		t.Error("Incorrect Value for avg", result)
	}

	eqnStr = "max"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	result = exprtkObj.GetEvaluatedValue()
	if result != 10 {
		t.Error("Incorrect Value for max", result)
	}

	eqnStr = "min"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	result = exprtkObj.GetEvaluatedValue()
	if result != -6.5 {
		t.Error("Incorrect Value for min", result)
	}

	eqnStr = "sum"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	result = exprtkObj.GetEvaluatedValue()
	if result != 18.9 {
		t.Error("Incorrect Value for sum", result)
	}

	eqnStr = "xyz"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	result = exprtkObj.GetEvaluatedValue()
	if result != 0.0 {
		t.Error("Incorrect Value for xyz", result)
	}
}
