package astcopy_test

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/strparse"
	"github.com/vvakame/astcopy"
)

func Example() {
	x := strparse.Expr(`1 + 2`).(*ast.BinaryExpr)
	y := astcopy.BinaryExpr(x, nil)
	fmt.Println(astequal.Expr(x, y)) // => true

	// Now modify x and make sure y is not modified.
	z := astcopy.BinaryExpr(y, nil)
	x.Op = token.SUB
	fmt.Println(astequal.Expr(y, z)) // => true
	fmt.Println(astequal.Expr(x, y)) // => false

	// Output:
	// true
	// true
	// false
}
