package bytecode

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"gosh/compiler/token"
	"gosh/core/object"
	"reflect"
	"strconv"
)

func (v *GoshVisitor) VisitEXP(ctx *parser.EXPContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *GoshVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	slog.Trace("enter exp stmt:", ctx.GetText())
	kids := ctx.GetChildren()
	for i := 0; i < len(kids); i++ {
		kid := kids[i]
		//slog.Info(reflect.TypeOf(kid))
		switch t := kid.(type) {
		case *parser.ExpressionContext:
			t.Accept(v)
		case *parser.BinOPContext:
			if t.GetChildCount() != 1 {
				slog.Error("wrong Op number")
			}
			kid := t.GetChild(0)
			if kidt, ok := kid.(*antlr.TerminalNodeImpl); ok {
				switch kidt.GetSymbol().GetTokenType() {
				case parser.GoshParserADD:
					tmp := kids[i+1]
					if tmp2, ok := tmp.(*parser.ExpressionContext); ok {
						slog.Trace("递归exp")
						tmp2.Accept(v)
						i++
						v.emit(token.OpAdd)
					} else {
						slog.Error("转型失败")
					}
				case parser.GoshParserSUB:
					tmp := kids[i+1]
					if tmp2, ok := tmp.(*parser.ExpressionContext); ok {
						slog.Trace("递归exp")
						tmp2.Accept(v)
						i++
						v.emit(token.OpSub)
					} else {
						slog.Error("转型失败")
					}

				default:
					slog.Warn("尚未支持 ", kidt.GetText())
				}
			} else {
				slog.Error("wrong Op ", reflect.TypeOf(kid))
			}

		case *antlr.TerminalNodeImpl:
			switch t.GetSymbol().GetTokenType() {
			case parser.GoshParserNumber:
				input := t.GetText()
				slog.Trace("解析:Number ", input)
				var res int
				if intValue, err := strconv.ParseInt(input, 10, 32); err == nil {
					res = v.addConstant(&object.Int{Value: intValue})
				} else {
					if floatValue, err := strconv.ParseFloat(input, 32); err == nil {
						res = v.addConstant(&object.Float{Value: floatValue})
					} else {
						res = v.addConstant(&object.String{Value: input})
					}
				}
				// 常量操作符、所在常量表的下标
				v.emit(token.OpConstant, res)
			case parser.GoshParserID:
				slog.Trace("解析:ID")
				// TODO: 解析该id
			case parser.GoshParserL_PAREN:
				continue
			case parser.GoshParserR_PAREN:
				continue
			default:
				slog.Error("未解析的符号", t.GetSymbol().GetTokenType())
			}
		case *parser.MulDivOPContext:
			tmp := kids[i+1]
			if tmp2, ok := tmp.(*parser.ExpressionContext); ok {
				slog.Trace("递归exp")
				tmp2.Accept(v)
				i++
			} else {
				slog.Error("转型失败")
			}
			//slog.Debug(t.GetText())
			if t.GetText() == "*" {
				v.emit(token.OpMul)
			} else {
				v.emit(token.OpDiv)
			}
		case *parser.UnOPContext:
			tmp := kids[i+1]
			if tmp2, ok := tmp.(*parser.ExpressionContext); ok {
				tmp2.Accept(v)
				i++
			} else {
				slog.Error("解析错误")
			}
			if t.GetText() == "-" {
				v.emit(token.OpMinus)
			}

		}

	}
	return nil
}

func (v *GoshVisitor) VisitConstvalue(ctx *parser.ConstvalueContext) interface{} {
	slog.Trace("enter constValue stmt:", ctx.GetText())
	input := ctx.GetText()
	var res int
	if intValue, err := strconv.ParseInt(input, 10, 32); err == nil {
		res = v.addConstant(&object.Int{Value: intValue})
	} else {
		if floatValue, err := strconv.ParseFloat(input, 32); err == nil {
			res = v.addConstant(&object.Float{Value: floatValue})
		} else {
			res = v.addConstant(&object.String{Value: input})
		}
	}
	// 常量操作符、所在常量表的下标
	v.emit(token.OpConstant, res)
	return nil
}
