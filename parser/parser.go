// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

// header of parser.y

import (
	"fmt"
	//"encoding/json"
	//"log"
	"strconv"
	"time"

	"github.com/udhos/gobule/conv"
)

type yySymType struct {
	yys        int
	typeBool   bool
	typeString string // holds: variable, number, or text
	typeScalar scalar
	typeList   []scalar
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault            = 57370
	yyEofCode            = 57344
	TkDot                = 57369
	TkEQ                 = 57363
	TkGE                 = 57367
	TkGT                 = 57365
	TkIdent              = 57358
	TkKeywordAnd         = 57349
	TkKeywordContains    = 57351
	TkKeywordCurrentTime = 57352
	TkKeywordFalse       = 57347
	TkKeywordList        = 57354
	TkKeywordNot         = 57350
	TkKeywordNumber      = 57353
	TkKeywordOr          = 57348
	TkKeywordTrue        = 57346
	TkKeywordVersion     = 57355
	TkLE                 = 57368
	TkLT                 = 57364
	TkNE                 = 57366
	TkNumber             = 57356
	TkParL               = 57359
	TkParR               = 57360
	TkSBktL              = 57361
	TkSBktR              = 57362
	TkText               = 57357
	yyErrCode            = 57345

	yyMaxDepth = 200
	yyTabOfs   = -29
)

var (
	yyPrec = map[int]int{
		TkKeywordOr:  0,
		TkKeywordAnd: 1,
	}

	yyXLAT = map[int]int{
		57360: 0,  // TkParR (28x)
		57356: 1,  // TkNumber (27x)
		57358: 2,  // TkIdent (26x)
		57357: 3,  // TkText (26x)
		57352: 4,  // TkKeywordCurrentTime (24x)
		57353: 5,  // TkKeywordNumber (24x)
		57355: 6,  // TkKeywordVersion (24x)
		57344: 7,  // $end (23x)
		57349: 8,  // TkKeywordAnd (23x)
		57348: 9,  // TkKeywordOr (23x)
		57375: 10, // scalar_exp (15x)
		57362: 11, // TkSBktR (11x)
		57350: 12, // TkKeywordNot (10x)
		57359: 13, // TkParL (9x)
		57363: 14, // TkEQ (8x)
		57367: 15, // TkGE (8x)
		57365: 16, // TkGT (8x)
		57368: 17, // TkLE (8x)
		57364: 18, // TkLT (8x)
		57366: 19, // TkNE (8x)
		57351: 20, // TkKeywordContains (6x)
		57371: 21, // bool_exp (5x)
		57373: 22, // list_exp (5x)
		57347: 23, // TkKeywordFalse (5x)
		57354: 24, // TkKeywordList (5x)
		57346: 25, // TkKeywordTrue (5x)
		57361: 26, // TkSBktL (5x)
		57369: 27, // TkDot (2x)
		57372: 28, // list (1x)
		57374: 29, // prog (1x)
		57370: 30, // $default (0x)
		57345: 31, // error (0x)
	}

	yySymNames = []string{
		"TkParR",
		"TkNumber",
		"TkIdent",
		"TkText",
		"TkKeywordCurrentTime",
		"TkKeywordNumber",
		"TkKeywordVersion",
		"$end",
		"TkKeywordAnd",
		"TkKeywordOr",
		"scalar_exp",
		"TkSBktR",
		"TkKeywordNot",
		"TkParL",
		"TkEQ",
		"TkGE",
		"TkGT",
		"TkLE",
		"TkLT",
		"TkNE",
		"TkKeywordContains",
		"bool_exp",
		"list_exp",
		"TkKeywordFalse",
		"TkKeywordList",
		"TkKeywordTrue",
		"TkSBktL",
		"TkDot",
		"list",
		"prog",
		"$default",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {29, 1},
		2:  {21, 3},
		3:  {21, 3},
		4:  {21, 3},
		5:  {21, 2},
		6:  {21, 1},
		7:  {21, 1},
		8:  {21, 3},
		9:  {21, 4},
		10: {21, 3},
		11: {21, 3},
		12: {21, 3},
		13: {21, 3},
		14: {21, 3},
		15: {21, 3},
		16: {22, 2},
		17: {22, 3},
		18: {22, 4},
		19: {22, 4},
		20: {28, 1},
		21: {28, 2},
		22: {10, 1},
		23: {10, 1},
		24: {10, 1},
		25: {10, 8},
		26: {10, 4},
		27: {10, 4},
		28: {10, 3},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [65][]uint8{
		// 0
		{1: 41, 42, 40, 45, 44, 43, 10: 37, 12: 33, 32, 21: 31, 36, 35, 39, 34, 38, 29: 30},
		{7: 29},
		{7: 28, 88, 89},
		{1: 41, 42, 40, 45, 44, 43, 10: 37, 12: 33, 32, 21: 92, 36, 35, 39, 34, 38},
		{1: 41, 42, 40, 45, 44, 43, 10: 37, 12: 33, 32, 21: 87, 36, 35, 39, 34, 38},
		// 5
		{23, 7: 23, 23, 23},
		{22, 7: 22, 22, 22},
		{12: 83, 20: 82},
		{14: 70, 73, 72, 75, 74, 71},
		{1: 41, 42, 40, 45, 44, 43, 10: 67, 65, 28: 66},
		// 10
		{13: 60},
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 11: 7, 14: 7, 7, 7, 7, 7, 7},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 11: 6, 14: 6, 6, 6, 6, 6, 6},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 11: 5, 14: 5, 5, 5, 5, 5, 5},
		{13: 53},
		// 15
		{13: 48},
		{13: 46},
		{47},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11: 1, 14: 1, 1, 1, 1, 1, 1},
		{2: 50, 49},
		// 20
		{52},
		{51},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 11: 2, 14: 2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 11: 3, 14: 3, 3, 3, 3, 3, 3},
		{1: 54},
		// 25
		{27: 55},
		{1: 56},
		{27: 57},
		{1: 58},
		{59},
		// 30
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 11: 4, 14: 4, 4, 4, 4, 4, 4},
		{2: 62, 61},
		{64},
		{63},
		{12: 10, 20: 10},
		// 35
		{12: 11, 20: 11},
		{12: 13, 20: 13},
		{1: 41, 42, 40, 45, 44, 43, 10: 69, 68},
		{1: 9, 9, 9, 9, 9, 9, 11: 9},
		{12: 12, 20: 12},
		// 40
		{1: 8, 8, 8, 8, 8, 8, 11: 8},
		{1: 41, 42, 40, 45, 44, 43, 10: 81},
		{1: 41, 42, 40, 45, 44, 43, 10: 80},
		{1: 41, 42, 40, 45, 44, 43, 10: 79},
		{1: 41, 42, 40, 45, 44, 43, 10: 78},
		// 45
		{1: 41, 42, 40, 45, 44, 43, 10: 77},
		{1: 41, 42, 40, 45, 44, 43, 10: 76},
		{14, 7: 14, 14, 14},
		{15, 7: 15, 15, 15},
		{16, 7: 16, 16, 16},
		// 50
		{17, 7: 17, 17, 17},
		{18, 7: 18, 18, 18},
		{19, 7: 19, 19, 19},
		{1: 41, 42, 40, 45, 44, 43, 10: 86},
		{20: 84},
		// 55
		{1: 41, 42, 40, 45, 44, 43, 10: 85},
		{20, 7: 20, 20, 20},
		{21, 7: 21, 21, 21},
		{24, 7: 24, 88, 89},
		{1: 41, 42, 40, 45, 44, 43, 10: 37, 12: 33, 32, 21: 91, 36, 35, 39, 34, 38},
		// 60
		{1: 41, 42, 40, 45, 44, 43, 10: 37, 12: 33, 32, 21: 90, 36, 35, 39, 34, 38},
		{25, 7: 25, 88, 89},
		{26, 7: 26, 88, 26},
		{93, 8: 88, 89},
		{27, 7: 27, 27, 27},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 31

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yylex.(*Lex).result.Eval = yyS[yypt-0].typeBool
		}
	case 2:
		{
			yyVAL.typeBool = yyS[yypt-1].typeBool
		}
	case 3:
		{
			yyVAL.typeBool = yyS[yypt-2].typeBool && yyS[yypt-0].typeBool
		}
	case 4:
		{
			yyVAL.typeBool = yyS[yypt-2].typeBool || yyS[yypt-0].typeBool
		}
	case 5:
		{
			yyVAL.typeBool = !yyS[yypt-0].typeBool
		}
	case 6:
		{
			yyVAL.typeBool = true
		}
	case 7:
		{
			yyVAL.typeBool = false
		}
	case 8:
		{
			yyVAL.typeBool = contains(yyS[yypt-2].typeList, yyS[yypt-0].typeScalar)
		}
	case 9:
		{
			yyVAL.typeBool = !contains(yyS[yypt-3].typeList, yyS[yypt-0].typeScalar)
		}
	case 10:
		{
			yyVAL.typeBool = yyS[yypt-2].typeScalar.Equals(yyS[yypt-0].typeScalar)
		}
	case 11:
		{
			yyVAL.typeBool = !yyS[yypt-2].typeScalar.Equals(yyS[yypt-0].typeScalar)
		}
	case 12:
		{
			var eval bool
			if yyS[yypt-2].typeScalar.scalarType != yyS[yypt-0].typeScalar.scalarType {
				yylex.Error("greater-than operator for different types")
			} else {
				eval = yyS[yypt-2].typeScalar.GreaterThan(yyS[yypt-0].typeScalar)
			}
			yyVAL.typeBool = eval
		}
	case 13:
		{
			var eval bool
			if yyS[yypt-2].typeScalar.scalarType != yyS[yypt-0].typeScalar.scalarType {
				yylex.Error("greater-than-or-equal operator for different types")
			} else {
				eval = !yyS[yypt-0].typeScalar.GreaterThan(yyS[yypt-2].typeScalar)
			}
			yyVAL.typeBool = eval
		}
	case 14:
		{
			var eval bool
			if yyS[yypt-2].typeScalar.scalarType != yyS[yypt-0].typeScalar.scalarType {
				yylex.Error("less-than operator for different types")
			} else {
				eval = yyS[yypt-0].typeScalar.GreaterThan(yyS[yypt-2].typeScalar)
			}
			yyVAL.typeBool = eval
		}
	case 15:
		{
			var eval bool
			if yyS[yypt-2].typeScalar.scalarType != yyS[yypt-0].typeScalar.scalarType {
				yylex.Error("less-than-or-equal operator for different types")
			} else {
				eval = !yyS[yypt-2].typeScalar.GreaterThan(yyS[yypt-0].typeScalar)
			}
			yyVAL.typeBool = eval
		}
	case 16:
		{
			yyVAL.typeList = []scalar{}
		}
	case 17:
		{
			yyVAL.typeList = yyS[yypt-1].typeList
		}
	case 18:
		{
			var list []scalar
			{
				s := yyS[yypt-1].typeString
				strList, errParse := parseList(s)
				if errParse != nil {
					yylex.Error(fmt.Sprintf("List(%s): bad list: %v", s, errParse))
				}
				for i, elem := range strList {
					switch v := elem.(type) {
					//case int:
					//    list = append(list, scalar{scalarType: scalarNumber, number: int64(v)})
					case int64:
						list = append(list, scalar{scalarType: scalarNumber, number: v})
					case string:
						list = append(list, scalar{scalarType: scalarText, text: v})
					default:
						yylex.Error(fmt.Sprintf("List(%s): invalid type for element %d: %v", s, i, elem))
					}
				}
			}
			yyVAL.typeList = list
		}
	case 19:
		{
			var list []scalar
			v := yyS[yypt-1].typeString
			l := yylex.(*Lex)
			if varValue, found := l.vars[v]; found {
				// found variable

				if vv, ok := varValue.([]interface{}); ok {

					for i, elem := range vv {
						switch val := elem.(type) {
						case float64:
							list = append(list, scalar{scalarType: scalarNumber, number: int64(val)})
						case int:
							list = append(list, scalar{scalarType: scalarNumber, number: int64(val)})
						case string:
							list = append(list, scalar{scalarType: scalarText, text: val})
						default:
							yylex.Error(fmt.Sprintf("List(%s): invalid type for element %d: %v", v, i, elem))
						}
					}

				} else {
					yylex.Error(fmt.Sprintf("List(%s): unexpected list type (%T): %v", v, varValue, varValue))
				}

			} else {
				yylex.Error(fmt.Sprintf("List(%s): variable undefined", v))
			}
			yyVAL.typeList = list
		}
	case 20:
		{
			l := yylex.(*Lex)
			l.scalarList = []scalar{yyS[yypt-0].typeScalar}
			yyVAL.typeList = l.scalarList
		}
	case 21:
		{
			l := yylex.(*Lex)
			l.scalarList = append(l.scalarList, yyS[yypt-0].typeScalar)
			yyVAL.typeList = l.scalarList
		}
	case 22:
		{
			yyVAL.typeScalar = scalar{scalarType: scalarText, text: yyS[yypt-0].typeString}
		}
	case 23:
		{
			s := yyS[yypt-0].typeString
			n, errConv := parseInt(s)
			if errConv != nil {
				yylex.Error(fmt.Sprintf("bad number conversion: '%s': %v", s, errConv))
			}
			yyVAL.typeScalar = scalar{scalarType: scalarNumber, number: n}
		}
	case 24:
		{
			v := yyS[yypt-0].typeString
			l := yylex.(*Lex)
			value := scalar{scalarType: scalarText}
			if varValue, found := l.vars[v]; found {
				switch val := varValue.(type) {
				case string:
					value.text = val
				case int:
					value.text = strconv.Itoa(val)
				default:
					yylex.Error(fmt.Sprintf("unexpected variable type: '%s'", v))
				}
			} else {
				value.text = fmt.Sprintf("variable undefined:'%s'", v)
				yylex.Error(value.text)
			}
			yyVAL.typeScalar = value
		}
	case 25:
		{
			s1 := yyS[yypt-5].typeString
			s2 := yyS[yypt-3].typeString
			s3 := yyS[yypt-1].typeString

			v1, errConv1 := parseInt(s1)
			if errConv1 != nil {
				yylex.Error(fmt.Sprintf("bad Version(version) number conversion 1: '%s': %v", s1, errConv1))
			}
			v2, errConv2 := parseInt(s2)
			if errConv2 != nil {
				yylex.Error(fmt.Sprintf("bad Version(version) number conversion 2: '%s': %v", s2, errConv2))
			}
			v3, errConv3 := parseInt(s3)
			if errConv3 != nil {
				yylex.Error(fmt.Sprintf("bad Version(version) number conversion 3: '%s': %v", s3, errConv3))
			}

			yyVAL.typeScalar = scalar{scalarType: scalarNumber, number: conv.VersionToNumber(v1, v2, v3)}
		}
	case 26:
		{
			s := yyS[yypt-1].typeString
			n, errConv := parseInt(s)
			if errConv != nil {
				yylex.Error(fmt.Sprintf("bad Number(text) conversion: '%s': %v", s, errConv))
			}
			yyVAL.typeScalar = scalar{scalarType: scalarNumber, number: n}
		}
	case 27:
		{
			v := yyS[yypt-1].typeString
			l := yylex.(*Lex)
			value := scalar{scalarType: scalarNumber}
			if varValue, found := l.vars[v]; found {
				// found variable
				switch val := varValue.(type) {
				case string:
					n, errConv := parseInt(val)
					if errConv != nil {
						yylex.Error(fmt.Sprintf("bad Number(variable) conversion: %s='%s': %v", v, val, errConv))
					}
					value.number = n
				case int:
					value.number = int64(val)
				case float64:
					value.number = int64(val)
				default:
					yylex.Error(fmt.Sprintf("unexpected Number(variable) var type: '%s': %q", v, varValue))
				}
			} else {
				value.text = fmt.Sprintf("Number() variable undefined:'%s'", v)
				yylex.Error(value.text)
			}
			yyVAL.typeScalar = value
		}
	case 28:
		{
			now := time.Now()
			n := now.Hour()*10000 + now.Minute()*100 + now.Second()
			//fmt.Printf("currenttime: %d\n", n)
			yyVAL.typeScalar = scalar{scalarType: scalarNumber, number: int64(n)}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
