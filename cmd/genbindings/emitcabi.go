package main

import (
	"fmt"
	"sort"
	"strings"
)

// cppComment renders a string safely in a C++ block comment.
// It strips interior nested comments.
func cppComment(s string) string {
	// Remove nested comments
	uncomment := strings.NewReplacer("/*", "", "*/", "")
	return "/* " + uncomment.Replace(s) + " */ "
}

func (p CppParameter) RenderTypeCabi() string {

	if p.ParameterType == "QString" {
		return "struct miqt_string"

	} else if p.ParameterType == "QByteArray" {
		return "struct miqt_string"

	} else if inner, ok := p.QListOf(); ok {
		return "struct miqt_array " + cppComment("of "+inner.RenderTypeCabi())

	} else if inner, ok := p.QSetOf(); ok {
		return "struct miqt_array " + cppComment("set of "+inner.RenderTypeCabi())

	} else if inner1, inner2, ok := p.QMapOf(); ok {
		return "struct miqt_map " + cppComment("of "+inner1.RenderTypeCabi()+" to "+inner2.RenderTypeCabi())

	} else if inner1, inner2, ok := p.QPairOf(); ok {
		return "struct miqt_map " + cppComment("tuple of "+inner1.RenderTypeCabi()+" and "+inner2.RenderTypeCabi())

	} else if (p.Pointer || p.ByRef) && p.QtClassType() {
		if p.PointerCount > 1 {
			return cabiClassName(p.ParameterType) + strings.Repeat("*", p.PointerCount)
		}
		return cabiClassName(p.ParameterType) + "*"

	} else if p.QtClassType() && !p.Pointer {
		// Even if C++ returns by value, CABI is returning a heap copy (new'd, not malloc'd)
		return cabiClassName(p.ParameterType) + "*"
	}

	ret := p.ParameterType
	switch p.ParameterType {
	case "uchar":
		ret = "unsigned char"
	case "uint":
		ret = "unsigned int"
	case "ulong":
		ret = "unsigned long"
	case "qint8":
		ret = "int8_t"
	case "quint8":
		ret = "uint8_t"
	case "qint16", "short":
		ret = "int16_t"
	case "quint16", "ushort", "unsigned short":
		ret = "uint16_t"
	case "qint32":
		ret = "int32_t"
	case "quint32":
		ret = "uint32_t"
	case "qlonglong", "qint64":
		ret = "int64_t"
	case "qulonglong", "quint64":
		ret = "uint64_t"
	case "qfloat16":
		ret = "_Float16" // No idea where this typedef comes from, but it exists
	case "qreal":
		ret = "double"
	case "qintptr", "QIntegerForSizeof<void *>::Signed": // long long int
		ret = "intptr_t" // long int
	case "quintptr", "uintptr", "QIntegerForSizeof<void *>::Unsigned":
		ret = "uintptr_t"
	case "qsizetype", "qptrdiff", "QIntegerForSizeof<std::size_t>::Signed":
		ret = "ptrdiff_t"
	}

	if p.Const {
		// This is needed for const-correctness for calling some overloads
		// e.g. QShortcut ctor taking (QWidget* parent, const char* member) signal -
		// the signal/slot requires that member is const, not just plain char*
		ret = "const " + ret
	}

	if ft, ok := p.QFlagsOf(); ok {
		if e, ok := KnownEnums[ft.ParameterType]; ok {
			ret = e.Enum.UnderlyingType.RenderTypeCabi()
		} else {
			ret = "int"
		}

	} else if e, ok := KnownEnums[p.ParameterType]; ok {
		ret = e.Enum.UnderlyingType.RenderTypeCabi()

	}

	if p.Pointer {
		ret += strings.Repeat("*", p.PointerCount)
	} else if p.ByRef {
		ret += "*"
	}

	return ret
}

// RenderTypeQtCpp renders the Qt C++ type in the original form from the function
// definition, before any typedefs or transformations were applied.
func (p CppParameter) RenderTypeQtCpp() string {
	return p.GetQtCppType().RenderTypeIntermediateCpp()
}

// RenderTypeIntermediateCpp renders the Qt C++ type WITHOUT resolving the
// interior QtCppOriginalType. This is used for intermediate const_cast<>s.
func (p CppParameter) RenderTypeIntermediateCpp() string {
	cppType := p.ParameterType

	if p.Const {
		cppType = "const " + cppType
	}
	if p.Pointer {
		cppType += strings.Repeat("*", p.PointerCount)
	}
	if p.ByRef {
		cppType += "&"
	}

	return cppType
}

// emitParametersCpp emits the parameter definitions exactly how Qt C++ defines them.
func emitParametersCpp(m CppMethod) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		tmp = append(tmp, p.RenderTypeQtCpp()+" "+p.ParameterName)
	}

	return strings.Join(tmp, `, `)
}

func emitParameterNames(m CppMethod) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		tmp = append(tmp, p.ParameterName)
	}

	return strings.Join(tmp, `, `)
}

func emitParameterTypesCpp(m CppMethod, includeHidden bool) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		tmp = append(tmp, p.RenderTypeQtCpp())
	}
	if includeHidden {
		for _, p := range m.HiddenParams {
			tmp = append(tmp, p.RenderTypeQtCpp())
		}
	}

	return strings.Join(tmp, `, `)
}

func emitParametersCabi(m CppMethod, selfType string) string {
	tmp := make([]string, 0, len(m.Parameters)+1)

	if !m.IsStatic && selfType != "" {
		tmp = append(tmp, selfType+" self")
	}

	for _, p := range m.Parameters {
		tmp = append(tmp, p.RenderTypeCabi()+" "+p.ParameterName)
	}

	return strings.Join(tmp, ", ")
}

func emitParametersCABI2CppForwarding(params []CppParameter, indent string) (preamble string, forwarding string) {
	tmp := make([]string, 0, len(params)+1)

	for _, p := range params {
		addPre, addFwd := emitCABI2CppForwarding(p, indent)
		preamble += addPre
		tmp = append(tmp, addFwd)
	}

	return preamble, strings.Join(tmp, ", ")
}

func makeNamePrefix(in string) string {
	replacer := strings.NewReplacer(`[`, `_`, `]`, "", `.`, `_`)
	return replacer.Replace(in)
}

func emitCABI2CppForwarding(p CppParameter, indent string) (preamble string, forwarding string) {

	nameprefix := makeNamePrefix(p.ParameterName)

	if p.ParameterType == "QString" {
		// The CABI received parameter is a struct miqt_string, passed by value
		// C++ needs it as a QString. Create one on the stack for automatic cleanup
		// The caller will free the miqt_string
		preamble += indent + "QString " + nameprefix + "_QString = QString::fromUtf8(" + p.ParameterName + ".data, " + p.ParameterName + ".len);\n"
		return preamble, nameprefix + "_QString"

	} else if p.ParameterType == "QByteArray" {
		// The caller will free the miqt_string data
		// This ctor makes a deep copy, on the stack which will be dtor'd by RAII
		preamble += indent + "QByteArray " + nameprefix + "_QByteArray(" + p.ParameterName + ".data, " + p.ParameterName + ".len);\n"
		return preamble, nameprefix + "_QByteArray"

	} else if listType, ok := p.QListOf(); ok {

		preamble += indent + p.GetQtCppType().ParameterType + " " + nameprefix + "_QList;\n"
		preamble += indent + nameprefix + "_QList.reserve(" + p.ParameterName + ".len);\n"

		preamble += indent + listType.RenderTypeCabi() + "* " + nameprefix + "_arr = static_cast<" + listType.RenderTypeCabi() + "*>(" + p.ParameterName + ".data);\n"
		preamble += indent + "for(size_t i = 0; i < " + p.ParameterName + ".len; ++i) {\n"

		listType.ParameterName = nameprefix + "_arr[i]"
		addPre, addFwd := emitCABI2CppForwarding(listType, indent+"\t")
		preamble += addPre
		preamble += indent + "\t" + nameprefix + "_QList.push_back(" + addFwd + ");\n"

		preamble += indent + "}\n"

		// Support passing QList<>* (very rare, but used in qnetwork)
		if p.Pointer {
			return preamble, "&" + nameprefix + "_QList"
		} else {
			return preamble, nameprefix + "_QList"
		}

	} else if kType, vType, ok := p.QMapOf(); ok {
		preamble += indent + p.GetQtCppType().ParameterType + " " + nameprefix + "_QMap;\n"

		// This container may be a QMap or a QHash
		// QHash supports .reserve(), but QMap doesn't
		if strings.HasPrefix(p.ParameterType, "QHash<") {
			preamble += indent + nameprefix + "_QMap.reserve(" + p.ParameterName + ".len);\n"
		}

		preamble += indent + kType.RenderTypeCabi() + "* " + nameprefix + "_karr = static_cast<" + kType.RenderTypeCabi() + "*>(" + p.ParameterName + ".keys);\n"
		preamble += indent + vType.RenderTypeCabi() + "* " + nameprefix + "_varr = static_cast<" + vType.RenderTypeCabi() + "*>(" + p.ParameterName + ".values);\n"
		preamble += indent + "for(size_t i = 0; i < " + p.ParameterName + ".len; ++i) {\n"

		kType.ParameterName = nameprefix + "_karr[i]"
		addPreK, addFwdK := emitCABI2CppForwarding(kType, indent+"\t")
		preamble += addPreK

		vType.ParameterName = nameprefix + "_varr[i]"
		addPreV, addFwdV := emitCABI2CppForwarding(vType, indent+"\t")
		preamble += addPreV

		preamble += indent + "\t" + nameprefix + "_QMap[" + addFwdK + "] = " + addFwdV + ";\n"

		preamble += indent + "}\n"
		return preamble, nameprefix + "_QMap"

	} else if kType, vType, ok := p.QPairOf(); ok {
		preamble += indent + p.GetQtCppType().ParameterType + " " + nameprefix + "_QPair;\n"

		preamble += indent + kType.RenderTypeCabi() + "* " + nameprefix + "_first_arr = static_cast<" + kType.RenderTypeCabi() + "*>(" + p.ParameterName + ".keys);\n"
		preamble += indent + vType.RenderTypeCabi() + "* " + nameprefix + "_second_arr = static_cast<" + vType.RenderTypeCabi() + "*>(" + p.ParameterName + ".values);\n"

		kType.ParameterName = nameprefix + "_first_arr[0]"
		addPreK, addFwdK := emitCABI2CppForwarding(kType, indent+"\t")
		preamble += addPreK

		vType.ParameterName = nameprefix + "_second_arr[0]"
		addPreV, addFwdV := emitCABI2CppForwarding(vType, indent+"\t")
		preamble += addPreV

		preamble += indent + nameprefix + "_QPair.first = " + addFwdK + ";\n"
		preamble += indent + nameprefix + "_QPair.second = " + addFwdV + ";\n"

		return preamble, nameprefix + "_QPair"

	} else if p.IsFlagType() || p.IntType() || p.IsKnownEnum() {
		castSrc := p.ParameterName
		castType := p.RenderTypeQtCpp()

		if p.ByRef { // e.g. QDataStream::operator>>() overloads
			castSrc = "*" + castSrc
		}

		if p.QtCppOriginalType != nil && p.QtCppOriginalType.Const != p.Const {
			return preamble, "static_cast<" + p.RenderTypeQtCpp() + ">(const_cast<" + p.RenderTypeIntermediateCpp() + ">(" + p.ParameterName + "))"
		}

		if p.ParameterType == "qint64" ||
			p.ParameterType == "quint64" ||
			p.ParameterType == "qlonglong" ||
			p.ParameterType == "qulonglong" ||
			p.GetQtCppType().ParameterType == "qintptr" ||
			p.GetQtCppType().ParameterType == "qsizetype" || // Qt 6 qversionnumber.h: invalid ‘static_cast’ from type ‘ptrdiff_t*’ {aka ‘long int*’} to type ‘qsizetype*’ {aka ‘long long int*’}
			p.ParameterType == "qint8" ||
			(p.IsFlagType() && p.ByRef) {
			// QDataStream::operator>>() by reference (qint64)
			// QLockFile::getLockInfo() by pointer
			// QTextStream::operator>>() by reference (qlonglong + qulonglong)
			// QDataStream::operator>>() qint8
			// CABI has these as int64_t* (long int) which fails a static_cast to qint64& (long long int&)
			// Hack a hard C-style cast
			return preamble, "(" + castType + ")(" + castSrc + ")"
		} else {
			// Use static_cast<> safely
			return preamble, "static_cast<" + castType + ">(" + castSrc + ")"
		}

	} else if _, ok := p.QSetOf(); ok {
		panic("QSet<> arguments are not yet implemented") // n.b. doesn't seem to exist in QtCore/QtGui/QtWidgets at all

	} else if p.ByRef {
		if p.Pointer {
			// By ref and by pointer
			// This happens for QDataStream &QDataStream::operator>>(char *&s)
			// We are only using one level of indirection
			return preamble, p.ParameterName
		} else {
			// By ref and not by pointer
			// We changed RenderTypeCabi() to render this as a pointer
			// Need to dereference so we can pass as reference to the actual Qt C++ function
			//tmp = append(tmp, "*"+p.ParameterName)
			return preamble, "*" + p.ParameterName
		}

	} else if p.QtClassType() && !p.Pointer {
		// CABI takes all Qt types by pointer, even if C++ wants them by value
		// Dereference the passed-in pointer
		if strings.Contains(p.ParameterName, `[`) {
			return preamble, "*(" + p.ParameterName + ")" // Extra brackets aren't necessary, just nice
		}
		return preamble, "*" + p.ParameterName

	} else {
		return preamble, p.ParameterName
	}

}

// emitAssignCppToCabi transforms and assigns rvalue to the assignExpression.
// Sample assignExpression: `return `, `auto foo = `
// Sample rvalue: `foo`, `foo(xyz)`
// The return is a complete statement including trailing newline.
func emitAssignCppToCabi(assignExpression string, p CppParameter, rvalue string) string {

	shouldReturn := assignExpression // n.b. already has indent
	afterCall := ""
	assignExpression = strings.TrimLeft(assignExpression, " \t")
	indent := shouldReturn[0 : len(shouldReturn)-len(assignExpression)]

	shouldReturn = shouldReturn[len(indent):]

	namePrefix := makeNamePrefix(p.ParameterName)

	if p.Void() {
		shouldReturn = ""
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.ParameterType == "QString" {

		if p.Pointer {
			// e.g. QTextStream::String()
			// These are rare, and probably expected to be lightweight references
			// But, a copy is the best we can project it as
			// Un-pointer-ify
			shouldReturn = ifv(p.Const, "const ", "") + "QString* " + namePrefix + "_ret = "
			afterCall = indent + "// Convert QString pointer from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory\n"
			afterCall += indent + "QByteArray " + namePrefix + "_b = " + namePrefix + "_ret->toUtf8();\n"

		} else {
			shouldReturn = ifv(p.Const, "const ", "") + "QString " + p.ParameterName + "_ret = "
			afterCall = indent + "// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory\n"
			afterCall += indent + "QByteArray " + namePrefix + "_b = " + namePrefix + "_ret.toUtf8();\n"
		}

		afterCall += indent + "struct miqt_string " + namePrefix + "_ms;\n"
		afterCall += indent + namePrefix + "_ms.len = " + namePrefix + "_b.length();\n"
		afterCall += indent + namePrefix + "_ms.data = static_cast<char*>(malloc(" + namePrefix + "_ms.len));\n"
		afterCall += indent + "memcpy(" + namePrefix + "_ms.data, " + namePrefix + "_b.data(), " + namePrefix + "_ms.len);\n"
		afterCall += indent + assignExpression + namePrefix + "_ms;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.ParameterType == "QByteArray" {
		// C++ has given us a QByteArray. CABI needs this as a struct miqt_string
		// Do not free the data, the caller will free it

		shouldReturn = ifv(p.Const, "const ", "") + "QByteArray " + p.ParameterName + "_qb = "

		afterCall += indent + "struct miqt_string " + namePrefix + "_ms;\n"
		afterCall += indent + namePrefix + "_ms.len = " + namePrefix + "_qb.length();\n"
		afterCall += indent + namePrefix + "_ms.data = static_cast<char*>(malloc(" + namePrefix + "_ms.len));\n"
		afterCall += indent + "memcpy(" + namePrefix + "_ms.data, " + namePrefix + "_qb.data(), " + namePrefix + "_ms.len);\n"
		afterCall += indent + assignExpression + namePrefix + "_ms;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if t, ok := p.QListOf(); ok {

		// In some cases rvalue is a function call and the temporary
		// is necessary; in some cases it's a literal and the temporary is
		// elided; but in some cases it's a Qt class and the temporary goes
		// through a copy constructor
		// TODO Detect safe cases where this can be optimized

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "

		afterCall += indent + "// Convert QList<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + t.RenderTypeCabi() + "* " + namePrefix + "_arr = static_cast<" + t.RenderTypeCabi() + "*>(malloc(sizeof(" + t.RenderTypeCabi() + ") * " + namePrefix + "_ret.length()));\n"
		afterCall += indent + "for (size_t i = 0, e = " + namePrefix + "_ret.length(); i < e; ++i) {\n"
		afterCall += emitAssignCppToCabi(indent+"\t"+namePrefix+"_arr[i] = ", t, namePrefix+"_ret[i]")
		afterCall += indent + "}\n"

		afterCall += indent + "struct miqt_array " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = " + namePrefix + "_ret.length();\n"
		afterCall += indent + "" + namePrefix + "_out.data = static_cast<void*>(" + namePrefix + "_arr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if t, ok := p.QSetOf(); ok {

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "

		afterCall += indent + "// Convert QSet<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + t.RenderTypeCabi() + "* " + namePrefix + "_arr = static_cast<" + t.RenderTypeCabi() + "*>(malloc(sizeof(" + t.RenderTypeCabi() + ") * " + namePrefix + "_ret.size()));\n"
		afterCall += indent + "int " + namePrefix + "_ctr = 0;\n"
		afterCall += indent + "QSetIterator<" + t.RenderTypeQtCpp() + "> " + namePrefix + "_itr(" + namePrefix + "_ret);\n"
		afterCall += indent + "while (" + namePrefix + "_itr.hasNext()) {\n"
		afterCall += emitAssignCppToCabi(indent+"\t"+namePrefix+"_arr["+namePrefix+"_ctr++] = ", t, namePrefix+"_itr.next()")
		afterCall += indent + "}\n"

		afterCall += indent + "struct miqt_array " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = " + namePrefix + "_ret.size();\n"
		afterCall += indent + "" + namePrefix + "_out.data = static_cast<void*>(" + namePrefix + "_arr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if kType, vType, ok := p.QMapOf(); ok {
		// QMap<K,V>

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "

		afterCall += indent + "// Convert QMap<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + kType.RenderTypeCabi() + "* " + namePrefix + "_karr = static_cast<" + kType.RenderTypeCabi() + "*>(malloc(sizeof(" + kType.RenderTypeCabi() + ") * " + namePrefix + "_ret.size()));\n"
		afterCall += indent + "" + vType.RenderTypeCabi() + "* " + namePrefix + "_varr = static_cast<" + vType.RenderTypeCabi() + "*>(malloc(sizeof(" + vType.RenderTypeCabi() + ") * " + namePrefix + "_ret.size()));\n"

		afterCall += indent + "int " + namePrefix + "_ctr = 0;\n"
		afterCall += indent + "for (auto " + namePrefix + "_itr = " + namePrefix + "_ret.keyValueBegin(); " + namePrefix + "_itr != " + namePrefix + "_ret.keyValueEnd(); ++" + namePrefix + "_itr) {\n"
		afterCall += emitAssignCppToCabi(indent+"\t"+namePrefix+"_karr["+namePrefix+"_ctr] = ", kType, namePrefix+"_itr->first")
		afterCall += emitAssignCppToCabi(indent+"\t"+namePrefix+"_varr["+namePrefix+"_ctr] = ", vType, namePrefix+"_itr->second")
		afterCall += indent + "\t" + namePrefix + "_ctr++;\n"

		afterCall += indent + "}\n"

		afterCall += indent + "struct miqt_map " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = " + namePrefix + "_ret.size();\n"
		afterCall += indent + "" + namePrefix + "_out.keys = static_cast<void*>(" + namePrefix + "_karr);\n"
		afterCall += indent + "" + namePrefix + "_out.values = static_cast<void*>(" + namePrefix + "_varr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if kType, vType, ok := p.QPairOf(); ok {
		// QPair<T1,T2>

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "

		afterCall += indent + "// Convert QPair<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + kType.RenderTypeCabi() + "* " + namePrefix + "_first_arr = static_cast<" + kType.RenderTypeCabi() + "*>(malloc(sizeof(" + kType.RenderTypeCabi() + ")));\n"
		afterCall += indent + "" + vType.RenderTypeCabi() + "* " + namePrefix + "_second_arr = static_cast<" + vType.RenderTypeCabi() + "*>(malloc(sizeof(" + vType.RenderTypeCabi() + ")));\n"

		afterCall += emitAssignCppToCabi(indent+namePrefix+"_first_arr[0] = ", kType, namePrefix+"_ret.first")
		afterCall += emitAssignCppToCabi(indent+namePrefix+"_second_arr[0] = ", vType, namePrefix+"_ret.second")

		afterCall += indent + "struct miqt_map " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = 1;\n"
		afterCall += indent + "" + namePrefix + "_out.keys = static_cast<void*>(" + namePrefix + "_first_arr);\n"
		afterCall += indent + "" + namePrefix + "_out.values = static_cast<void*>(" + namePrefix + "_second_arr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.QtClassType() && p.ByRef {
		// It's a pointer in disguise, just needs one cast
		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "
		afterCall += indent + "// Cast returned reference into pointer\n"
		if p.Const {
			nonConst := p // copy
			nonConst.Const = false
			nonConst.ByRef = false
			nonConst.Pointer = true
			nonConst.PointerCount = 1
			afterCall += indent + "" + assignExpression + "const_cast<" + nonConst.RenderTypeQtCpp() + ">(&" + namePrefix + "_ret);\n"
		} else {
			afterCall += indent + "" + assignExpression + "&" + namePrefix + "_ret;\n"
		}
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.QtClassType() && !p.Pointer {

		// Elide temporary and emit directly from the rvalue
		return indent + assignExpression + "new " + p.ParameterType + "(" + rvalue + ");\n"

	} else if p.IsFlagType() || p.IsKnownEnum() || p.QtCppOriginalType != nil {
		// Needs an explicit cast
		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "

		if p.QtCppOriginalType != nil && p.QtCppOriginalType.Const != p.Const {
			afterCall += indent + "" + assignExpression + "const_cast<" + p.RenderTypeCabi() + ">(static_cast<" + p.RenderTypeIntermediateCpp() + ">(" + namePrefix + "_ret));\n"
		} else if p.QtCppOriginalType != nil && p.QtCppOriginalType.ParameterType == "qintptr" {
			// Hard int cast
			afterCall += indent + "" + assignExpression + "(" + p.RenderTypeCabi() + ")(" + namePrefix + "_ret);\n"
		} else {
			afterCall += indent + "" + assignExpression + "static_cast<" + p.RenderTypeCabi() + ">(" + namePrefix + "_ret);\n"
		}
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.Const {
		shouldReturn += "(" + p.RenderTypeCabi() + ") "
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else {
		// Basic type
		if p.ByRef {
			// The C++ type is a reference, the CABI type is a pointer type
			shouldReturn += "&"
		}

		return indent + shouldReturn + rvalue + ";\n" + afterCall
	}

}

func getCppZeroValue(p CppParameter) string {
	if p.Pointer {
		return "nullptr"
	} else if p.IsKnownEnum() {
		return "(" + p.RenderTypeQtCpp() + ")(0)"
	} else if p.IntType() {
		return "0"
	} else if p.ParameterType == "bool" {
		return "false"
	} else {
		return p.RenderTypeQtCpp() + "()"
	}
}

// getReferencedTypes finds all referenced Qt types in this file.
func getReferencedTypes(src *CppParsedHeader) []string {

	foundTypes := map[string]struct{}{}

	var maybeAddType func(p CppParameter)
	maybeAddType = func(p CppParameter) {
		if p.QtClassType() {
			foundTypes[p.ParameterType] = struct{}{}
		}
		if t, ok := p.QListOf(); ok {
			foundTypes["QList"] = struct{}{} // FIXME or QVector?
			maybeAddType(t)
		}
		if kType, vType, ok := p.QMapOf(); ok {
			foundTypes["QMap"] = struct{}{} // FIXME or QHash?
			maybeAddType(kType)
			maybeAddType(vType)
		}
		if kType, vType, ok := p.QPairOf(); ok {
			foundTypes["QPair"] = struct{}{}
			maybeAddType(kType)
			maybeAddType(vType)
		}
		if t, ok := p.QSetOf(); ok {
			foundTypes["QSet"] = struct{}{}
			maybeAddType(t)
		}
	}

	for _, c := range src.Classes {

		foundTypes[c.ClassName] = struct{}{}

		for _, ctor := range c.Ctors {
			for _, p := range ctor.Parameters {
				maybeAddType(p)
			}
		}
		for _, m := range c.Methods {
			for _, p := range m.Parameters {
				maybeAddType(p)
			}
			maybeAddType(m.ReturnType)
		}
		for _, vm := range c.VirtualMethods() {
			for _, p := range vm.Parameters {
				maybeAddType(p)
			}
			maybeAddType(vm.ReturnType)
		}
		for _, cn := range c.AllInherits() {
			maybeAddType(CppParameter{
				ParameterType: cn,
			})
		}
	}

	// Some types (e.g. QRgb) are found but are typedefs, not classes
	for _, td := range src.Typedefs {
		delete(foundTypes, td.Alias)
	}

	// Convert to sorted list
	foundTypesList := make([]string, 0, len(foundTypes))
	for ft := range foundTypes {
		if !AllowClass(ft) {
			continue
		}

		foundTypesList = append(foundTypesList, ft)
	}
	sort.Strings(foundTypesList)

	return foundTypesList
}

// cabiClassName returns the Go / CABI class name for a Qt C++ class.
// Normally this is the same, except for class types that are nested inside another class definition.
func cabiClassName(className string) string {

	// Many types are defined in qnamespace.h under Qt::
	// The Go implementation is always called qt.Foo, and these names don't
	// collide with anything, so strip the redundant prefix
	className = strings.TrimPrefix(className, `Qt::`)

	// Must use __ to avoid subclass/method name collision e.g. QPagedPaintDevice::Margins
	return strings.Replace(className, `::`, `__`, -1)
}

func cabiPreventStructDeclaration(className string) bool {
	switch className {
	case "QList", "QString", "QSet", "QMap", "QHash", "QPair", "QVector", "QByteArray":
		return true // These types are reprojected
	default:
		return false
	}
}

func emitBindingHeader(src *CppParsedHeader, filename string, packageName string) (string, error) {
	ret := strings.Builder{}

	includeGuard := "MIQT_" + strings.ToUpper(strings.Replace(strings.Replace(packageName, `/`, `_`, -1), `-`, `_`, -1)) + "_GEN_" + strings.ToUpper(strings.Replace(strings.Replace(filename, `.`, `_`, -1), `-`, `_`, -1))

	bindingInclude := "../libmiqt/libmiqt.h"

	if strings.Contains(packageName, `/`) {
		bindingInclude = "../" + bindingInclude
	}

	ret.WriteString(`#pragma once
#ifndef ` + includeGuard + `
#define ` + includeGuard + `

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "` + bindingInclude + `"

#ifdef __cplusplus
extern "C" {
#endif

`)

	foundTypesList := getReferencedTypes(src)

	ret.WriteString("#ifdef __cplusplus\n")

	for _, ft := range foundTypesList {
		if cabiPreventStructDeclaration(ft) {
			continue
		}

		if strings.Contains(ft, `::`) {
			// Forward declarations of inner classes are not yet supported in C++
			// @ref https://stackoverflow.com/q/1021793

			ret.WriteString(`#if defined(WORKAROUND_INNER_CLASS_DEFINITION_` + cabiClassName(ft) + ")\n")
			ret.WriteString(`typedef ` + ft + " " + cabiClassName(ft) + ";\n")
			ret.WriteString("#else\n")
			ret.WriteString(`class ` + cabiClassName(ft) + ";\n")
			ret.WriteString("#endif\n")

		} else {
			ret.WriteString(`class ` + ft + ";\n")
		}
	}

	ret.WriteString("#else\n")

	for _, ft := range foundTypesList {
		if cabiPreventStructDeclaration(ft) {
			continue
		}
		ret.WriteString(`typedef struct ` + cabiClassName(ft) + " " + cabiClassName(ft) + ";\n")
	}

	ret.WriteString("#endif\n")

	ret.WriteString("\n")

	for _, c := range src.Classes {

		methodPrefixName := cabiClassName(c.ClassName)

		for i, ctor := range c.Ctors {
			ret.WriteString(fmt.Sprintf("void %s_new%s(%s);\n", methodPrefixName, maybeSuffix(i), emitParametersCabiConstructor(&c, &ctor)))
		}

		for _, m := range c.Methods {
			ret.WriteString(fmt.Sprintf("%s %s_%s(%s);\n", m.ReturnType.RenderTypeCabi(), methodPrefixName, m.SafeMethodName(), emitParametersCabi(m, ifv(m.IsConst, "const ", "")+methodPrefixName+"*")))

			if m.IsSignal {
				ret.WriteString(fmt.Sprintf("%s %s_connect_%s(%s* self, intptr_t slot);\n", m.ReturnType.RenderTypeCabi(), methodPrefixName, m.SafeMethodName(), methodPrefixName))
			}
		}

		for _, m := range c.VirtualMethods() {
			ret.WriteString(fmt.Sprintf("void %s_override_virtual_%s(%s* self, intptr_t slot);\n", methodPrefixName, m.SafeMethodName(), "void" /*methodPrefixName*/))

			ret.WriteString(fmt.Sprintf("%s %s_virtualbase_%s(%s);\n", m.ReturnType.RenderTypeCabi(), methodPrefixName, m.SafeMethodName(), emitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void" /*methodPrefixName*/ +"*")))
		}

		// delete
		if c.CanDelete {
			ret.WriteString(fmt.Sprintf("void %s_Delete(%s* self, bool isSubclass);\n", methodPrefixName, methodPrefixName))
		}

		ret.WriteString("\n")
	}

	ret.WriteString(
		`#ifdef __cplusplus
} /* extern C */
#endif 

#endif
`)
	return ret.String(), nil
}

func fullyQualifiedConstructor(className string) string {
	parts := strings.Split(className, `::`)
	return className + "::" + parts[len(parts)-1]
}

func emitParametersCabiConstructor(c *CppClass, ctor *CppMethod) string {

	plist := slice_copy(ctor.Parameters) // semi-shallow copy

	plist = append(plist, CppParameter{
		ParameterName: cabiClassName("outptr_" + cabiClassName(c.ClassName)),
		ParameterType: c.ClassName,
		Pointer:       true,
		PointerCount:  2,
	})
	for _, baseClass := range c.AllInherits() {
		plist = append(plist, CppParameter{
			ParameterName: cabiClassName("outptr_" + cabiClassName(baseClass)),
			ParameterType: baseClass,
			Pointer:       true,
			PointerCount:  2,
		})
	}

	slist := make([]string, 0, len(plist))
	for _, p := range plist {
		slist = append(slist, p.RenderTypeCabi()+" "+p.ParameterName)
	}

	return strings.Join(slist, `, `)
}

func emitBindingCpp(src *CppParsedHeader, filename string) (string, error) {
	ret := strings.Builder{}

	for _, ref := range getReferencedTypes(src) {

		if ref == "QString" {
			ret.WriteString("#include <QString>\n")
			ret.WriteString("#include <QByteArray>\n")
			ret.WriteString("#include <cstring>\n")
			continue
		}

		if strings.Contains(ref, `::`) {
			ret.WriteString(`#define WORKAROUND_INNER_CLASS_DEFINITION_` + cabiClassName(ref) + "\n")
			continue
		}

		if !ImportHeaderForClass(ref) {
			continue
		}

		ret.WriteString(`#include <` + ref + ">\n")
	}

	ret.WriteString(`#include <` + filename + ">\n")
	ret.WriteString(`#include "gen_` + filename + "\"\n")
	ret.WriteString("#include \"_cgo_export.h\"\n\n")

	for _, c := range src.Classes {

		methodPrefixName := cabiClassName(c.ClassName)
		cppClassName := c.ClassName
		virtualMethods := c.VirtualMethods()

		if len(virtualMethods) > 0 {

			overriddenClassName := "MiqtVirtual" + strings.Replace(cppClassName, `::`, ``, -1)

			ret.WriteString("class " + overriddenClassName + " : public virtual " + cppClassName + " {\n" +
				"public:\n" +
				"\n",
			)

			for _, ctor := range c.Ctors {
				ret.WriteString("\t" + overriddenClassName + "(" + emitParametersCpp(ctor) + "): " + cppClassName + "(" + emitParameterNames(ctor) + ") {};\n")
			}
			ret.WriteString("\n")

			if !c.CanDelete {
				ret.WriteString(
					"private:\n" +
						"\tvirtual ~" + overriddenClassName + "();\n" + //  = delete;\n" +
						"\n" +
						"public:\n" +
						"\n",
				)
			} else {
				ret.WriteString(
					"\tvirtual ~" + overriddenClassName + "() = default;\n" +
						"\n",
				)
			}

			for _, m := range virtualMethods {

				{
					var maybeReturn, maybeReturn2 string
					var returnTransformP, returnTransformF string
					if !m.ReturnType.Void() {
						maybeReturn = "return "

						maybeReturn2 = m.ReturnType.RenderTypeCabi() + " callback_return_value = "
						returnParam := m.ReturnType // copy
						returnParam.ParameterName = "callback_return_value"
						returnTransformP, returnTransformF = emitCABI2CppForwarding(returnParam, "\t\t")
					}

					handleVarname := "handle__" + m.SafeMethodName()

					ret.WriteString(
						"\t// cgo.Handle value for overwritten implementation\n" +
							"\tintptr_t " + handleVarname + " = 0;\n" +
							"\n",
					)

					// In the case of method overloads, we always need to use the
					// original method name (CppCallTarget), not the MethodName

					ret.WriteString(
						"\t// Subclass to allow providing a Go implementation\n" +
							"\tvirtual " + m.ReturnType.RenderTypeQtCpp() + " " + m.CppCallTarget() + "(" + emitParametersCpp(m) + ") " + ifv(m.IsConst, "const ", "") + "override {\n",
					)

					ret.WriteString("\t\tif (" + handleVarname + " == 0) {\n")
					if m.IsPureVirtual {
						if m.ReturnType.Void() {
							ret.WriteString("\t\t\treturn; // Pure virtual, there is no base we can call\n")
						} else {
							ret.WriteString("\t\t\treturn " + getCppZeroValue(m.ReturnType) + "; // Pure virtual, there is no base we can call\n")
						}
					} else {
						ret.WriteString("\t\t\t" + maybeReturn + methodPrefixName + "::" + m.CppCallTarget() + "(" + emitParameterNames(m) + ");\n")

						if m.ReturnType.Void() {
							ret.WriteString("\t\t\treturn;\n")
						}
					}
					ret.WriteString("\t\t}\n")

					paramArgs := []string{}
					if m.IsConst {
						// We're calling a Cgo-exported function, but Cgo can't
						// describe a const pointer to a custom class, unless
						// it's a primitive or wrapped in a typedef.
						// Just strip the const_cast away
						paramArgs = append(paramArgs, "const_cast<"+overriddenClassName+"*>(this)")
					} else {
						paramArgs = append(paramArgs, "this")
					}
					paramArgs = append(paramArgs, handleVarname)

					var signalCode string

					for i, p := range m.Parameters {
						signalCode += emitAssignCppToCabi(fmt.Sprintf("\t\t%s sigval%d = ", p.RenderTypeCabi(), i+1), p, p.ParameterName)
						paramArgs = append(paramArgs, fmt.Sprintf("sigval%d", i+1))
					}

					ret.WriteString(
						"\t\t\n" +
							signalCode + "\n" +
							"\t\t" + maybeReturn2 + "miqt_exec_callback_" + methodPrefixName + "_" + m.SafeMethodName() + "(" + strings.Join(paramArgs, `, `) + ");\n" +
							returnTransformP + "\n" +
							"\t\t" + ifv(maybeReturn == "", "", "return "+returnTransformF+";") + "\n" +
							"\t}\n" +

							"\n",
					)
				}

				// If there is a base version of this method, add a helper to
				// allow calling it

				if !m.IsPureVirtual {

					// The virtualbase wrapper needs to take CABI parameters, not
					// real Qt parameters, in case there are protected enum types
					// (e.g. QAbstractItemView::CursorAction)

					var parametersCabi []string
					for _, p := range m.Parameters {
						parametersCabi = append(parametersCabi, p.RenderTypeCabi()+" "+p.ParameterName)
					}
					vbpreamble, vbforwarding := emitParametersCABI2CppForwarding(m.Parameters, "\t\t")

					vbCallTarget := methodPrefixName + "::" + m.CppCallTarget() + "(" + vbforwarding + ")"

					ret.WriteString(
						"\t// Wrapper to allow calling protected method\n" +
							"\t" + m.ReturnType.RenderTypeCabi() + " virtualbase_" + m.SafeMethodName() + "(" + strings.Join(parametersCabi, ", ") + ") " + ifv(m.IsConst, "const ", "") + "{\n" +
							vbpreamble + "\n" +
							emitAssignCppToCabi("\t\treturn ", m.ReturnType, vbCallTarget) + "\n" +
							"\t}\n" +

							"\n",
					)

				}
			}

			ret.WriteString(
				"};\n" +
					"\n")

			cppClassName = overriddenClassName
		}

		for i, ctor := range c.Ctors {

			// The returned ctor needs to return a C++ pointer for not just the
			// class itself, but also all of the inherited base classes
			// That's because C++ virtual inheritance shifts the pointer; we
			// need all the base pointers to call base methods from CGO
			// Supply them all as out-parameters so we only need one roundtrip

			preamble, forwarding := emitParametersCABI2CppForwarding(ctor.Parameters, "\t")

			ret.WriteString(
				"void " + methodPrefixName + "_new" + maybeSuffix(i) + "(" + emitParametersCabiConstructor(&c, &ctor) + ") {\n",
			)

			if ctor.LinuxOnly {
				ret.WriteString(
					"#ifndef Q_OS_LINUX\n" +
						"\treturn;\n" +
						"#else\n",
				)
			}

			ret.WriteString(
				preamble +
					"\t" + cppClassName + "* ret = new " + cppClassName + "(" + forwarding + ");\n" + // Subclass class name
					"\t*outptr_" + cabiClassName(c.ClassName) + " = ret;\n", // Original class name
			)
			for _, baseClass := range c.AllInherits() {
				ret.WriteString("\t*outptr_" + cabiClassName(baseClass) + " = static_cast<" + baseClass + "*>(ret);\n")
			}

			if ctor.LinuxOnly {
				ret.WriteString(
					"#endif\n",
				)
			}

			ret.WriteString(
				"}\n" +
					"\n",
			)

		}

		for _, m := range c.Methods {

			// Protected virtual methods will be bound separately (the only
			// useful thing is to expose calling the virtual base)
			// Protected non-virtual methods should always be hidden
			if m.IsProtected {
				continue
			}

			// Need to take an extra 'self' parameter

			preamble, forwarding := emitParametersCABI2CppForwarding(m.Parameters, "\t")

			// callTarget is an rvalue representing the full C++ function call.
			callTarget := "self->"
			if m.IsStatic {
				callTarget = c.ClassName + "::"
			}

			callTarget += m.CppCallTarget() + "(" + forwarding + ")"

			// Qt 6.8 moved many operator== implementations from class methods
			// into global operators.
			// By using infix syntax, either can be called
			if m.IsReadonlyOperator() && len(m.Parameters) == 1 {
				operator := m.CppCallTarget()[8:]
				callTarget = "(*self " + operator + " " + forwarding + ")"
			}

			if m.LinuxOnly {
				ret.WriteString(fmt.Sprintf(
					"%s %s_%s(%s) {\n"+
						"#ifdef Q_OS_LINUX\n"+
						"%s"+
						"%s"+
						"#else\n"+
						"\t%s _ret_invalidOS;\n"+
						"\treturn _ret_invalidOS;\n"+
						"#endif\n"+
						"}\n"+
						"\n",
					m.ReturnType.RenderTypeCabi(), methodPrefixName, m.SafeMethodName(), emitParametersCabi(m, ifv(m.IsConst, "const ", "")+methodPrefixName+"*"),
					preamble,
					emitAssignCppToCabi("\treturn ", m.ReturnType, callTarget),
					m.ReturnType.RenderTypeCabi(),
				))

			} else if m.BecomesNonConstInVersion != nil {

				nonConstCallTarget := "const_cast<" + methodPrefixName + "*>(self)->" + m.CppCallTarget() + "(" + forwarding + ")"

				ret.WriteString("" +
					m.ReturnType.RenderTypeCabi() + " " + methodPrefixName + "_" + m.SafeMethodName() + "(" + emitParametersCabi(m, ifv(m.IsConst, "const ", "")+methodPrefixName+"*") + ") {\n" +
					preamble + "\n" +
					"// This method was changed from const to non-const in Qt " + *m.BecomesNonConstInVersion + "\n" +
					"#if QT_VERSION < QT_VERSION_CHECK(" + strings.Replace(*m.BecomesNonConstInVersion, `.`, `,`, -1) + ",0)\n" +
					emitAssignCppToCabi("\treturn ", m.ReturnType, callTarget) +
					"#else\n" +
					emitAssignCppToCabi("\treturn ", m.ReturnType, nonConstCallTarget) +
					"#endif\n" +
					"}\n" +
					"\n",
				)

			} else {

				ret.WriteString(fmt.Sprintf(
					"%s %s_%s(%s) {\n"+
						"%s"+
						"%s"+
						"}\n"+
						"\n",
					m.ReturnType.RenderTypeCabi(), methodPrefixName, m.SafeMethodName(), emitParametersCabi(m, ifv(m.IsConst, "const ", "")+methodPrefixName+"*"),
					preamble,
					emitAssignCppToCabi("\treturn ", m.ReturnType, callTarget),
				))

			}

			if m.IsSignal {

				bindingFunc := "miqt_exec_callback_" + cabiClassName(c.ClassName) + "_" + m.SafeMethodName()

				// If there are hidden parameters, the type of the signal itself
				// needs to include them
				exactSignal := `static_cast<void (` + c.ClassName + `::*)(` + emitParameterTypesCpp(m, true) + `)` + ifv(m.IsConst, ` const`, ``) + `>(&` + c.ClassName + `::` + m.CppCallTarget() + `)`

				paramArgs := []string{"slot"}
				paramArgDefs := []string{"intptr_t cb"}

				var signalCode string

				for i, p := range m.Parameters {
					signalCode += emitAssignCppToCabi(fmt.Sprintf("\t\t%s sigval%d = ", p.RenderTypeCabi(), i+1), p, p.ParameterName)
					paramArgs = append(paramArgs, fmt.Sprintf("sigval%d", i+1))
					paramArgDefs = append(paramArgDefs, p.RenderTypeCabi()+" "+p.ParameterName)
				}

				signalCode += "\t\t" + bindingFunc + "(" + strings.Join(paramArgs, `, `) + ");\n"

				ret.WriteString(
					`void ` + methodPrefixName + `_connect_` + m.SafeMethodName() + `(` + methodPrefixName + `* self, intptr_t slot) {` + "\n" +
						"\t" + cppClassName + `::connect(self, ` + exactSignal + `, self, [=](` + emitParametersCpp(m) + `) {` + "\n" +
						signalCode +
						"\t});\n" +
						"}\n" +
						"\n",
				)
			}

		}

		// Virtual override helpers
		for _, m := range virtualMethods {

			// Virtual methods: Allow overriding
			// (Never use a const self*)
			// The pointer that we are passed is the base type, not the subclassed
			// type. First cast the void* to the base type, and only then,
			// upclass it

			ret.WriteString(
				`void ` + methodPrefixName + `_override_virtual_` + m.SafeMethodName() + `(void* self, intptr_t slot) {` + "\n" +
					"\tdynamic_cast<" + cppClassName + "*>( (" + cabiClassName(c.ClassName) + "*)(self) )->handle__" + m.SafeMethodName() + " = slot;\n" +
					"}\n" +
					"\n",
			)

			// 2. Add CABI function to call the base method

			if !m.IsPureVirtual {
				// This is not generally exposed in the Go binding, but when overriding
				// the method, allows Go code to call super()

				// It uses CABI-CABI, the CABI-QtC++ type conversion will be done
				// inside the class method so as to allow for accessing protected
				// types.
				// Both the parameters and return type are given in CABI format.

				var parameterNames []string
				for _, param := range m.Parameters {
					parameterNames = append(parameterNames, param.ParameterName)
				}

				// callTarget is an rvalue representing the full C++ function call.
				// These are never static

				callTarget := "( (" + ifv(m.IsConst, "const ", "") + cppClassName + "*)(self) )->virtualbase_" + m.SafeMethodName() + "(" + strings.Join(parameterNames, `, `) + ")"

				ret.WriteString(
					m.ReturnType.RenderTypeCabi() + " " + methodPrefixName + "_virtualbase_" + m.SafeMethodName() + "(" + emitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void*") + ") {\n" +
						"\t" + ifv(m.ReturnType.Void(), "", "return ") + callTarget + ";\n" +
						"}\n" +
						"\n",
				)

			}

		}

		// Delete
		if c.CanDelete {
			ret.WriteString(
				"void " + methodPrefixName + "_Delete(" + methodPrefixName + "* self, bool isSubclass) {\n" +
					"\tif (isSubclass) {\n" +
					"\t\tdelete dynamic_cast<" + cppClassName + "*>( self );\n" +
					"\t} else {\n" +
					"\t\tdelete self;\n" +
					"\t}\n" +
					"}\n" +
					"\n",
			)
		}
	}

	return ret.String(), nil
}
