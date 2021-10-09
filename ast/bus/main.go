package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

type Driver struct {
	Orders       int
	DrivingYears int `json:"driving_years"`
}

func getDriverRemote() []byte {
	return []byte(`{"orders":100000,"driving_years":18}`)
}

func isOldDriver(d *Driver) bool {
	if d.Orders > 10000 && d.DrivingYears > 5 {
		return true
	}
	return false
}

func judge(bop ast.Node, m map[string]int64) bool {
	if isLeaf(bop) {
		expr := bop.(*ast.BinaryExpr)
		x := expr.X.(*ast.Ident)
		y := expr.Y.(*ast.BasicLit)

		if expr.Op == token.GTR {
			left := m[x.Name]
			right, _ := strconv.ParseInt(y.Value, 10, 64)
			return left > right
		}
		return false
	}

	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		println("this cannot be true")
		return false
	}

	switch expr.Op {
	case token.LAND:
		return judge(expr.X, m) && judge(expr.Y, m)
	case token.LOR:
		return judge(expr.X, m) || judge(expr.Y, m)
	}
	println("unsupported operator")
	return false
}

func isLeaf(bop ast.Node) bool {
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		return false
	}

	_, okL := expr.X.(*ast.Ident)
	_, okR := expr.Y.(*ast.BasicLit)

	if okL && okR {
		return true
	}

	return false
}

func Eval(m map[string]int64, expr string) (bool, error) {
	exprAst, err := parser.ParseExpr(expr)
	if nil != err {
		return false, err
	}

	fset := token.NewFileSet()
	ast.Print(fset, exprAst)

	return judge(exprAst, m), nil
}

func main() {
	bs := getDriverRemote()
	var d Driver
	json.Unmarshal(bs, &d)
	fmt.Println(d)
	fmt.Println(isOldDriver(&d))

	m := map[string]int64{
		"orders":        100000,
		"driving_years": 100,
	}
	rule := "orders > 10000 && driving_years > 50 && orders > 100"
	fmt.Println(Eval(m, rule))
}
