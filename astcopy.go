// Package astcopy implements Go AST reflection-free deep copy operations.
package astcopy

import (
	"go/ast"
)

// CopyNodeMap hold mapping copied node to original node.
type CopyNodeMap map[ast.Node]ast.Node

// Node returns x node deep copy.
// Copy of nil argument is nil.
func Node(x ast.Node, nMap CopyNodeMap) ast.Node {
	return copyNode(x, nMap)
}

// NodeList returns xs node slice deep copy.
// Copy of nil argument is nil.
func NodeList(xs []ast.Node, nMap CopyNodeMap) []ast.Node {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Node, len(xs))
	for i := range xs {
		cp[i] = copyNode(xs[i], nMap)
	}
	return cp
}

// Expr returns x expression deep copy.
// Copy of nil argument is nil.
func Expr(x ast.Expr, nMap CopyNodeMap) ast.Expr {
	return copyExpr(x, nMap)
}

// ExprList returns xs expression slice deep copy.
// Copy of nil argument is nil.
func ExprList(xs []ast.Expr, nMap CopyNodeMap) []ast.Expr {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Expr, len(xs))
	for i := range xs {
		cp[i] = copyExpr(xs[i], nMap)
	}
	return cp
}

// Stmt returns x statement deep copy.
// Copy of nil argument is nil.
func Stmt(x ast.Stmt, nMap CopyNodeMap) ast.Stmt {
	return copyStmt(x, nMap)
}

// StmtList returns xs statement slice deep copy.
// Copy of nil argument is nil.
func StmtList(xs []ast.Stmt, nMap CopyNodeMap) []ast.Stmt {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Stmt, len(xs))
	for i := range xs {
		cp[i] = copyStmt(xs[i], nMap)
	}
	return cp
}

// Decl returns x declaration deep copy.
// Copy of nil argument is nil.
func Decl(x ast.Decl, nMap CopyNodeMap) ast.Decl {
	return copyDecl(x, nMap)
}

// DeclList returns xs declaration slice deep copy.
// Copy of nil argument is nil.
func DeclList(xs []ast.Decl, nMap CopyNodeMap) []ast.Decl {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Decl, len(xs))
	for i := range xs {
		cp[i] = copyDecl(xs[i], nMap)
	}
	return cp
}

// BadExpr returns x deep copy.
// Copy of nil argument is nil.
func BadExpr(x *ast.BadExpr, nMap CopyNodeMap) *ast.BadExpr {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// Ident returns x deep copy.
// Copy of nil argument is nil.
func Ident(x *ast.Ident, nMap CopyNodeMap) *ast.Ident {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// IdentList returns xs identifier slice deep copy.
// Copy of nil argument is nil.
func IdentList(xs []*ast.Ident, nMap CopyNodeMap) []*ast.Ident {
	if xs == nil {
		return nil
	}
	cp := make([]*ast.Ident, len(xs))
	for i := range xs {
		cp[i] = Ident(xs[i], nMap)
	}
	return cp
}

// Ellipsis returns x deep copy.
// Copy of nil argument is nil.
func Ellipsis(x *ast.Ellipsis, nMap CopyNodeMap) *ast.Ellipsis {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Elt = copyExpr(x.Elt, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// BasicLit returns x deep copy.
// Copy of nil argument is nil.
func BasicLit(x *ast.BasicLit, nMap CopyNodeMap) *ast.BasicLit {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// FuncLit returns x deep copy.
// Copy of nil argument is nil.
func FuncLit(x *ast.FuncLit, nMap CopyNodeMap) *ast.FuncLit {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Type = FuncType(x.Type, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// CompositeLit returns x deep copy.
// Copy of nil argument is nil.
func CompositeLit(x *ast.CompositeLit, nMap CopyNodeMap) *ast.CompositeLit {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Type = copyExpr(x.Type, nMap)
	cp.Elts = ExprList(x.Elts, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ParenExpr returns x deep copy.
// Copy of nil argument is nil.
func ParenExpr(x *ast.ParenExpr, nMap CopyNodeMap) *ast.ParenExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// SelectorExpr returns x deep copy.
// Copy of nil argument is nil.
func SelectorExpr(x *ast.SelectorExpr, nMap CopyNodeMap) *ast.SelectorExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	cp.Sel = Ident(x.Sel, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// IndexExpr returns x deep copy.
// Copy of nil argument is nil.
func IndexExpr(x *ast.IndexExpr, nMap CopyNodeMap) *ast.IndexExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	cp.Index = copyExpr(x.Index, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// SliceExpr returns x deep copy.
// Copy of nil argument is nil.
func SliceExpr(x *ast.SliceExpr, nMap CopyNodeMap) *ast.SliceExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	cp.Low = copyExpr(x.Low, nMap)
	cp.High = copyExpr(x.High, nMap)
	cp.Max = copyExpr(x.Max, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// TypeAssertExpr returns x deep copy.
// Copy of nil argument is nil.
func TypeAssertExpr(x *ast.TypeAssertExpr, nMap CopyNodeMap) *ast.TypeAssertExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	cp.Type = copyExpr(x.Type, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// CallExpr returns x deep copy.
// Copy of nil argument is nil.
func CallExpr(x *ast.CallExpr, nMap CopyNodeMap) *ast.CallExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Fun = copyExpr(x.Fun, nMap)
	cp.Args = ExprList(x.Args, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// StarExpr returns x deep copy.
// Copy of nil argument is nil.
func StarExpr(x *ast.StarExpr, nMap CopyNodeMap) *ast.StarExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// UnaryExpr returns x deep copy.
// Copy of nil argument is nil.
func UnaryExpr(x *ast.UnaryExpr, nMap CopyNodeMap) *ast.UnaryExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// BinaryExpr returns x deep copy.
// Copy of nil argument is nil.
func BinaryExpr(x *ast.BinaryExpr, nMap CopyNodeMap) *ast.BinaryExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	cp.Y = copyExpr(x.Y, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// KeyValueExpr returns x deep copy.
// Copy of nil argument is nil.
func KeyValueExpr(x *ast.KeyValueExpr, nMap CopyNodeMap) *ast.KeyValueExpr {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Key = copyExpr(x.Key, nMap)
	cp.Value = copyExpr(x.Value, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ArrayType returns x deep copy.
// Copy of nil argument is nil.
func ArrayType(x *ast.ArrayType, nMap CopyNodeMap) *ast.ArrayType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Len = copyExpr(x.Len, nMap)
	cp.Elt = copyExpr(x.Elt, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// StructType returns x deep copy.
// Copy of nil argument is nil.
func StructType(x *ast.StructType, nMap CopyNodeMap) *ast.StructType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Fields = FieldList(x.Fields, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// Field returns x deep copy.
// Copy of nil argument is nil.
func Field(x *ast.Field, nMap CopyNodeMap) *ast.Field {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Names = IdentList(x.Names, nMap)
	cp.Type = copyExpr(x.Type, nMap)
	cp.Tag = BasicLit(x.Tag, nMap)
	cp.Doc = CommentGroup(x.Doc, nMap)
	cp.Comment = CommentGroup(x.Comment, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// FieldList returns x deep copy.
// Copy of nil argument is nil.
func FieldList(x *ast.FieldList, nMap CopyNodeMap) *ast.FieldList {
	if x == nil {
		return nil
	}
	cp := *x
	if x.List != nil {
		cp.List = make([]*ast.Field, len(x.List))
		for i := range x.List {
			cp.List[i] = Field(x.List[i], nMap)
		}
	}
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// FuncType returns x deep copy.
// Copy of nil argument is nil.
func FuncType(x *ast.FuncType, nMap CopyNodeMap) *ast.FuncType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Params = FieldList(x.Params, nMap)
	cp.Results = FieldList(x.Results, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// InterfaceType returns x deep copy.
// Copy of nil argument is nil.
func InterfaceType(x *ast.InterfaceType, nMap CopyNodeMap) *ast.InterfaceType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Methods = FieldList(x.Methods, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// MapType returns x deep copy.
// Copy of nil argument is nil.
func MapType(x *ast.MapType, nMap CopyNodeMap) *ast.MapType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Key = copyExpr(x.Key, nMap)
	cp.Value = copyExpr(x.Value, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ChanType returns x deep copy.
// Copy of nil argument is nil.
func ChanType(x *ast.ChanType, nMap CopyNodeMap) *ast.ChanType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Value = copyExpr(x.Value, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// BlockStmt returns x deep copy.
// Copy of nil argument is nil.
func BlockStmt(x *ast.BlockStmt, nMap CopyNodeMap) *ast.BlockStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.List = StmtList(x.List, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ImportSpec returns x deep copy.
// Copy of nil argument is nil.
func ImportSpec(x *ast.ImportSpec, nMap CopyNodeMap) *ast.ImportSpec {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Name = Ident(x.Name, nMap)
	cp.Path = BasicLit(x.Path, nMap)
	cp.Doc = CommentGroup(x.Doc, nMap)
	cp.Comment = CommentGroup(x.Comment, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ValueSpec returns x deep copy.
// Copy of nil argument is nil.
func ValueSpec(x *ast.ValueSpec, nMap CopyNodeMap) *ast.ValueSpec {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Names = IdentList(x.Names, nMap)
	cp.Values = ExprList(x.Values, nMap)
	cp.Type = copyExpr(x.Type, nMap)
	cp.Doc = CommentGroup(x.Doc, nMap)
	cp.Comment = CommentGroup(x.Comment, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// TypeSpec returns x deep copy.
// Copy of nil argument is nil.
func TypeSpec(x *ast.TypeSpec, nMap CopyNodeMap) *ast.TypeSpec {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Name = Ident(x.Name, nMap)
	cp.Type = copyExpr(x.Type, nMap)
	cp.Doc = CommentGroup(x.Doc, nMap)
	cp.Comment = CommentGroup(x.Comment, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// Spec returns x deep copy.
// Copy of nil argument is nil.
func Spec(x ast.Spec, nMap CopyNodeMap) ast.Spec {
	if x == nil {
		return nil
	}

	switch x := x.(type) {
	case *ast.ImportSpec:
		return ImportSpec(x, nMap)
	case *ast.ValueSpec:
		return ValueSpec(x, nMap)
	case *ast.TypeSpec:
		return TypeSpec(x, nMap)
	default:
		panic("unhandled spec")
	}
}

// SpecList returns xs spec slice deep copy.
// Copy of nil argument is nil.
func SpecList(xs []ast.Spec, nMap CopyNodeMap) []ast.Spec {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Spec, len(xs))
	for i := range xs {
		cp[i] = Spec(xs[i], nMap)
	}
	return cp
}

// BadStmt returns x deep copy.
// Copy of nil argument is nil.
func BadStmt(x *ast.BadStmt, nMap CopyNodeMap) *ast.BadStmt {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// DeclStmt returns x deep copy.
// Copy of nil argument is nil.
func DeclStmt(x *ast.DeclStmt, nMap CopyNodeMap) *ast.DeclStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Decl = copyDecl(x.Decl, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// EmptyStmt returns x deep copy.
// Copy of nil argument is nil.
func EmptyStmt(x *ast.EmptyStmt, nMap CopyNodeMap) *ast.EmptyStmt {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// LabeledStmt returns x deep copy.
// Copy of nil argument is nil.
func LabeledStmt(x *ast.LabeledStmt, nMap CopyNodeMap) *ast.LabeledStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Label = Ident(x.Label, nMap)
	cp.Stmt = copyStmt(x.Stmt, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ExprStmt returns x deep copy.
// Copy of nil argument is nil.
func ExprStmt(x *ast.ExprStmt, nMap CopyNodeMap) *ast.ExprStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// SendStmt returns x deep copy.
// Copy of nil argument is nil.
func SendStmt(x *ast.SendStmt, nMap CopyNodeMap) *ast.SendStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Chan = copyExpr(x.Chan, nMap)
	cp.Value = copyExpr(x.Value, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// IncDecStmt returns x deep copy.
// Copy of nil argument is nil.
func IncDecStmt(x *ast.IncDecStmt, nMap CopyNodeMap) *ast.IncDecStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.X = copyExpr(x.X, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// AssignStmt returns x deep copy.
// Copy of nil argument is nil.
func AssignStmt(x *ast.AssignStmt, nMap CopyNodeMap) *ast.AssignStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Lhs = ExprList(x.Lhs, nMap)
	cp.Rhs = ExprList(x.Rhs, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// GoStmt returns x deep copy.
// Copy of nil argument is nil.
func GoStmt(x *ast.GoStmt, nMap CopyNodeMap) *ast.GoStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Call = CallExpr(x.Call, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// DeferStmt returns x deep copy.
// Copy of nil argument is nil.
func DeferStmt(x *ast.DeferStmt, nMap CopyNodeMap) *ast.DeferStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Call = CallExpr(x.Call, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ReturnStmt returns x deep copy.
// Copy of nil argument is nil.
func ReturnStmt(x *ast.ReturnStmt, nMap CopyNodeMap) *ast.ReturnStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Results = ExprList(x.Results, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// BranchStmt returns x deep copy.
// Copy of nil argument is nil.
func BranchStmt(x *ast.BranchStmt, nMap CopyNodeMap) *ast.BranchStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Label = Ident(x.Label, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// IfStmt returns x deep copy.
// Copy of nil argument is nil.
func IfStmt(x *ast.IfStmt, nMap CopyNodeMap) *ast.IfStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Init = copyStmt(x.Init, nMap)
	cp.Cond = copyExpr(x.Cond, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	cp.Else = copyStmt(x.Else, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// CaseClause returns x deep copy.
// Copy of nil argument is nil.
func CaseClause(x *ast.CaseClause, nMap CopyNodeMap) *ast.CaseClause {
	if x == nil {
		return nil
	}
	cp := *x
	cp.List = ExprList(x.List, nMap)
	cp.Body = StmtList(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// SwitchStmt returns x deep copy.
// Copy of nil argument is nil.
func SwitchStmt(x *ast.SwitchStmt, nMap CopyNodeMap) *ast.SwitchStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Init = copyStmt(x.Init, nMap)
	cp.Tag = copyExpr(x.Tag, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// TypeSwitchStmt returns x deep copy.
// Copy of nil argument is nil.
func TypeSwitchStmt(x *ast.TypeSwitchStmt, nMap CopyNodeMap) *ast.TypeSwitchStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Init = copyStmt(x.Init, nMap)
	cp.Assign = copyStmt(x.Assign, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// CommClause returns x deep copy.
// Copy of nil argument is nil.
func CommClause(x *ast.CommClause, nMap CopyNodeMap) *ast.CommClause {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Comm = copyStmt(x.Comm, nMap)
	cp.Body = StmtList(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// SelectStmt returns x deep copy.
// Copy of nil argument is nil.
func SelectStmt(x *ast.SelectStmt, nMap CopyNodeMap) *ast.SelectStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Body = BlockStmt(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// ForStmt returns x deep copy.
// Copy of nil argument is nil.
func ForStmt(x *ast.ForStmt, nMap CopyNodeMap) *ast.ForStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Init = copyStmt(x.Init, nMap)
	cp.Cond = copyExpr(x.Cond, nMap)
	cp.Post = copyStmt(x.Post, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// RangeStmt returns x deep copy.
// Copy of nil argument is nil.
func RangeStmt(x *ast.RangeStmt, nMap CopyNodeMap) *ast.RangeStmt {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Key = copyExpr(x.Key, nMap)
	cp.Value = copyExpr(x.Value, nMap)
	cp.X = copyExpr(x.X, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// Comment returns x deep copy.
// Copy of nil argument is nil.
func Comment(x *ast.Comment, nMap CopyNodeMap) *ast.Comment {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// CommentGroup returns x deep copy.
// Copy of nil argument is nil.
func CommentGroup(x *ast.CommentGroup, nMap CopyNodeMap) *ast.CommentGroup {
	if x == nil {
		return nil
	}
	cp := *x
	if x.List != nil {
		cp.List = make([]*ast.Comment, len(x.List))
		for i := range x.List {
			cp.List[i] = Comment(x.List[i], nMap)
		}
	}
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// File returns x deep copy.
// Copy of nil argument is nil.
func File(x *ast.File, nMap CopyNodeMap) *ast.File {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Doc = CommentGroup(x.Doc, nMap)
	cp.Name = Ident(x.Name, nMap)
	cp.Decls = DeclList(x.Decls, nMap)
	cp.Imports = make([]*ast.ImportSpec, len(x.Imports))
	for i := range x.Imports {
		cp.Imports[i] = ImportSpec(x.Imports[i], nMap)
	}
	cp.Unresolved = IdentList(x.Unresolved, nMap)
	cp.Comments = make([]*ast.CommentGroup, len(x.Comments))
	for i := range x.Comments {
		cp.Comments[i] = CommentGroup(x.Comments[i], nMap)
	}
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// Package returns x deep copy.
// Copy of nil argument is nil.
func Package(x *ast.Package, nMap CopyNodeMap) *ast.Package {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Files = make(map[string]*ast.File)
	for filename, f := range x.Files {
		cp.Files[filename] = f
	}
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// BadDecl returns x deep copy.
// Copy of nil argument is nil.
func BadDecl(x *ast.BadDecl, nMap CopyNodeMap) *ast.BadDecl {
	if x == nil {
		return nil
	}
	cp := *x
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// GenDecl returns x deep copy.
// Copy of nil argument is nil.
func GenDecl(x *ast.GenDecl, nMap CopyNodeMap) *ast.GenDecl {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Specs = SpecList(x.Specs, nMap)
	cp.Doc = CommentGroup(x.Doc, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

// FuncDecl returns x deep copy.
// Copy of nil argument is nil.
func FuncDecl(x *ast.FuncDecl, nMap CopyNodeMap) *ast.FuncDecl {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Recv = FieldList(x.Recv, nMap)
	cp.Name = Ident(x.Name, nMap)
	cp.Type = FuncType(x.Type, nMap)
	cp.Body = BlockStmt(x.Body, nMap)
	cp.Doc = CommentGroup(x.Doc, nMap)
	if nMap != nil {
		nMap[&cp] = x
	}
	return &cp
}

func copyNode(x ast.Node, nMap CopyNodeMap) ast.Node {
	switch x := x.(type) {
	case ast.Expr:
		return copyExpr(x, nMap)
	case ast.Stmt:
		return copyStmt(x, nMap)
	case ast.Decl:
		return copyDecl(x, nMap)

	case ast.Spec:
		return Spec(x, nMap)
	case *ast.FieldList:
		return FieldList(x, nMap)
	case *ast.Comment:
		return Comment(x, nMap)
	case *ast.CommentGroup:
		return CommentGroup(x, nMap)
	case *ast.File:
		return File(x, nMap)
	case *ast.Package:
		return Package(x, nMap)

	default:
		panic("unhandled node")
	}
}

func copyExpr(x ast.Expr, nMap CopyNodeMap) ast.Expr {
	if x == nil {
		return nil
	}

	switch x := x.(type) {
	case *ast.BadExpr:
		return BadExpr(x, nMap)
	case *ast.Ident:
		return Ident(x, nMap)
	case *ast.Ellipsis:
		return Ellipsis(x, nMap)
	case *ast.BasicLit:
		return BasicLit(x, nMap)
	case *ast.FuncLit:
		return FuncLit(x, nMap)
	case *ast.CompositeLit:
		return CompositeLit(x, nMap)
	case *ast.ParenExpr:
		return ParenExpr(x, nMap)
	case *ast.SelectorExpr:
		return SelectorExpr(x, nMap)
	case *ast.IndexExpr:
		return IndexExpr(x, nMap)
	case *ast.SliceExpr:
		return SliceExpr(x, nMap)
	case *ast.TypeAssertExpr:
		return TypeAssertExpr(x, nMap)
	case *ast.CallExpr:
		return CallExpr(x, nMap)
	case *ast.StarExpr:
		return StarExpr(x, nMap)
	case *ast.UnaryExpr:
		return UnaryExpr(x, nMap)
	case *ast.BinaryExpr:
		return BinaryExpr(x, nMap)
	case *ast.KeyValueExpr:
		return KeyValueExpr(x, nMap)
	case *ast.ArrayType:
		return ArrayType(x, nMap)
	case *ast.StructType:
		return StructType(x, nMap)
	case *ast.FuncType:
		return FuncType(x, nMap)
	case *ast.InterfaceType:
		return InterfaceType(x, nMap)
	case *ast.MapType:
		return MapType(x, nMap)
	case *ast.ChanType:
		return ChanType(x, nMap)

	default:
		panic("unhandled expr")
	}
}

func copyStmt(x ast.Stmt, nMap CopyNodeMap) ast.Stmt {
	if x == nil {
		return nil
	}

	switch x := x.(type) {
	case *ast.BadStmt:
		return BadStmt(x, nMap)
	case *ast.DeclStmt:
		return DeclStmt(x, nMap)
	case *ast.EmptyStmt:
		return EmptyStmt(x, nMap)
	case *ast.LabeledStmt:
		return LabeledStmt(x, nMap)
	case *ast.ExprStmt:
		return ExprStmt(x, nMap)
	case *ast.SendStmt:
		return SendStmt(x, nMap)
	case *ast.IncDecStmt:
		return IncDecStmt(x, nMap)
	case *ast.AssignStmt:
		return AssignStmt(x, nMap)
	case *ast.GoStmt:
		return GoStmt(x, nMap)
	case *ast.DeferStmt:
		return DeferStmt(x, nMap)
	case *ast.ReturnStmt:
		return ReturnStmt(x, nMap)
	case *ast.BranchStmt:
		return BranchStmt(x, nMap)
	case *ast.BlockStmt:
		return BlockStmt(x, nMap)
	case *ast.IfStmt:
		return IfStmt(x, nMap)
	case *ast.CaseClause:
		return CaseClause(x, nMap)
	case *ast.SwitchStmt:
		return SwitchStmt(x, nMap)
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt(x, nMap)
	case *ast.CommClause:
		return CommClause(x, nMap)
	case *ast.SelectStmt:
		return SelectStmt(x, nMap)
	case *ast.ForStmt:
		return ForStmt(x, nMap)
	case *ast.RangeStmt:
		return RangeStmt(x, nMap)

	default:
		panic("unhandled stmt")
	}
}

func copyDecl(x ast.Decl, nMap CopyNodeMap) ast.Decl {
	if x == nil {
		return nil
	}

	switch x := x.(type) {
	case *ast.BadDecl:
		return BadDecl(x, nMap)
	case *ast.GenDecl:
		return GenDecl(x, nMap)
	case *ast.FuncDecl:
		return FuncDecl(x, nMap)

	default:
		panic("unhandled decl")
	}
}
