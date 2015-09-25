// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __testPackageInFile_struct_with_union_field_CachedRef = __testPackageInFile_struct_with_union_field_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_struct_with_union_field_Ref() types.Ref {
	p := types.PackageDef{
		Types: types.MapOfStringToTypeRefDef{

			"StructWithUnionField": __typeRefOfStructWithUnionField(),
		},
	}.New()
	return types.Ref{R: types.RegisterPackage(&p)}
}

// StructWithUnionField

type StructWithUnionField struct {
	m types.Map
}

func NewStructWithUnionField() StructWithUnionField {
	return StructWithUnionField{types.NewMap(
		types.NewString("$name"), types.NewString("StructWithUnionField"),
		types.NewString("$type"), types.MakeTypeRef("StructWithUnionField", __testPackageInFile_struct_with_union_field_CachedRef),
		types.NewString("a"), types.Float32(0),
		types.NewString("$unionIndex"), types.UInt32(0),
		types.NewString("$unionValue"), types.Float64(0),
	)}
}

type StructWithUnionFieldDef struct {
	A            float32
	__unionIndex uint32
	__unionValue interface{}
}

func (def StructWithUnionFieldDef) New() StructWithUnionField {
	return StructWithUnionField{
		types.NewMap(
			types.NewString("$name"), types.NewString("StructWithUnionField"),
			types.NewString("$type"), types.MakeTypeRef("StructWithUnionField", __testPackageInFile_struct_with_union_field_CachedRef),
			types.NewString("a"), types.Float32(def.A),
			types.NewString("$unionIndex"), types.UInt32(def.__unionIndex),
			types.NewString("$unionValue"), def.__unionDefToValue(),
		)}
}

func (s StructWithUnionField) Def() (d StructWithUnionFieldDef) {
	d.A = float32(s.m.Get(types.NewString("a")).(types.Float32))
	d.__unionIndex = uint32(s.m.Get(types.NewString("$unionIndex")).(types.UInt32))
	d.__unionValue = s.__unionValueToDef()
	return
}

func (def StructWithUnionFieldDef) __unionDefToValue() types.Value {
	switch def.__unionIndex {
	case 0:
		return types.Float64(def.__unionValue.(float64))
	case 1:
		return types.NewString(def.__unionValue.(string))
	case 2:
		return def.__unionValue.(types.Blob)
	case 3:
		return def.__unionValue.(types.Value)
	case 4:
		return def.__unionValue.(SetOfUInt8Def).New().NomsValue()
	}
	panic("unreachable")
}

func (s StructWithUnionField) __unionValueToDef() interface{} {
	switch uint32(s.m.Get(types.NewString("$unionIndex")).(types.UInt32)) {
	case 0:
		return float64(s.m.Get(types.NewString("$unionValue")).(types.Float64))
	case 1:
		return s.m.Get(types.NewString("$unionValue")).(types.String).String()
	case 2:
		return s.m.Get(types.NewString("$unionValue")).(types.Blob)
	case 3:
		return s.m.Get(types.NewString("$unionValue"))
	case 4:
		return SetOfUInt8FromVal(s.m.Get(types.NewString("$unionValue"))).Def()
	}
	panic("unreachable")
}

// Creates and returns a Noms Value that describes StructWithUnionField.
func __typeRefOfStructWithUnionField() types.TypeRef {
	return types.MakeStructTypeRef("StructWithUnionField",
		[]types.Field{
			types.Field{"a", types.MakePrimitiveTypeRef(types.Float32Kind), false},
		},
		[]types.Field{
			types.Field{"b", types.MakePrimitiveTypeRef(types.Float64Kind), false},
			types.Field{"c", types.MakePrimitiveTypeRef(types.StringKind), false},
			types.Field{"d", types.MakePrimitiveTypeRef(types.BlobKind), false},
			types.Field{"e", types.MakePrimitiveTypeRef(types.ValueKind), false},
			types.Field{"f", types.MakeCompoundTypeRef("", types.SetKind, types.MakePrimitiveTypeRef(types.UInt8Kind)), false},
		})

}

func StructWithUnionFieldFromVal(val types.Value) StructWithUnionField {
	// TODO: Validate here
	return StructWithUnionField{val.(types.Map)}
}

func (s StructWithUnionField) NomsValue() types.Value {
	return s.m
}

func (s StructWithUnionField) Equals(other StructWithUnionField) bool {
	return s.m.Equals(other.m)
}

func (s StructWithUnionField) Ref() ref.Ref {
	return s.m.Ref()
}

func (s StructWithUnionField) Type() types.TypeRef {
	return s.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (s StructWithUnionField) A() float32 {
	return float32(s.m.Get(types.NewString("a")).(types.Float32))
}

func (s StructWithUnionField) SetA(val float32) StructWithUnionField {
	return StructWithUnionField{s.m.Set(types.NewString("a"), types.Float32(val))}
}

func (s StructWithUnionField) B() (val float64, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 0 {
		return
	}
	return float64(s.m.Get(types.NewString("$unionValue")).(types.Float64)), true
}

func (s StructWithUnionField) SetB(val float64) StructWithUnionField {
	return StructWithUnionField{s.m.Set(types.NewString("$unionIndex"), types.UInt32(0)).Set(types.NewString("$unionValue"), types.Float64(val))}
}

func (def StructWithUnionFieldDef) B() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(float64), true
}

func (def StructWithUnionFieldDef) SetB(val float64) StructWithUnionFieldDef {
	def.__unionIndex = 0
	def.__unionValue = val
	return def
}

func (s StructWithUnionField) C() (val string, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 1 {
		return
	}
	return s.m.Get(types.NewString("$unionValue")).(types.String).String(), true
}

func (s StructWithUnionField) SetC(val string) StructWithUnionField {
	return StructWithUnionField{s.m.Set(types.NewString("$unionIndex"), types.UInt32(1)).Set(types.NewString("$unionValue"), types.NewString(val))}
}

func (def StructWithUnionFieldDef) C() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(string), true
}

func (def StructWithUnionFieldDef) SetC(val string) StructWithUnionFieldDef {
	def.__unionIndex = 1
	def.__unionValue = val
	return def
}

func (s StructWithUnionField) D() (val types.Blob, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 2 {
		return
	}
	return s.m.Get(types.NewString("$unionValue")).(types.Blob), true
}

func (s StructWithUnionField) SetD(val types.Blob) StructWithUnionField {
	return StructWithUnionField{s.m.Set(types.NewString("$unionIndex"), types.UInt32(2)).Set(types.NewString("$unionValue"), val)}
}

func (def StructWithUnionFieldDef) D() (val types.Blob, ok bool) {
	if def.__unionIndex != 2 {
		return
	}
	return def.__unionValue.(types.Blob), true
}

func (def StructWithUnionFieldDef) SetD(val types.Blob) StructWithUnionFieldDef {
	def.__unionIndex = 2
	def.__unionValue = val
	return def
}

func (s StructWithUnionField) E() (val types.Value, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 3 {
		return
	}
	return s.m.Get(types.NewString("$unionValue")), true
}

func (s StructWithUnionField) SetE(val types.Value) StructWithUnionField {
	return StructWithUnionField{s.m.Set(types.NewString("$unionIndex"), types.UInt32(3)).Set(types.NewString("$unionValue"), val)}
}

func (def StructWithUnionFieldDef) E() (val types.Value, ok bool) {
	if def.__unionIndex != 3 {
		return
	}
	return def.__unionValue.(types.Value), true
}

func (def StructWithUnionFieldDef) SetE(val types.Value) StructWithUnionFieldDef {
	def.__unionIndex = 3
	def.__unionValue = val
	return def
}

func (s StructWithUnionField) F() (val SetOfUInt8, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 4 {
		return
	}
	return SetOfUInt8FromVal(s.m.Get(types.NewString("$unionValue"))), true
}

func (s StructWithUnionField) SetF(val SetOfUInt8) StructWithUnionField {
	return StructWithUnionField{s.m.Set(types.NewString("$unionIndex"), types.UInt32(4)).Set(types.NewString("$unionValue"), val.NomsValue())}
}

func (def StructWithUnionFieldDef) F() (val SetOfUInt8Def, ok bool) {
	if def.__unionIndex != 4 {
		return
	}
	return def.__unionValue.(SetOfUInt8Def), true
}

func (def StructWithUnionFieldDef) SetF(val SetOfUInt8Def) StructWithUnionFieldDef {
	def.__unionIndex = 4
	def.__unionValue = val
	return def
}

// SetOfUInt8

type SetOfUInt8 struct {
	s types.Set
}

func NewSetOfUInt8() SetOfUInt8 {
	return SetOfUInt8{types.NewSet()}
}

type SetOfUInt8Def map[uint8]bool

func (def SetOfUInt8Def) New() SetOfUInt8 {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = types.UInt8(d)
		i++
	}
	return SetOfUInt8{types.NewSet(l...)}
}

func (s SetOfUInt8) Def() SetOfUInt8Def {
	def := make(map[uint8]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[uint8(v.(types.UInt8))] = true
		return false
	})
	return def
}

func SetOfUInt8FromVal(p types.Value) SetOfUInt8 {
	return SetOfUInt8{p.(types.Set)}
}

func (s SetOfUInt8) NomsValue() types.Value {
	return s.s
}

func (s SetOfUInt8) Equals(p SetOfUInt8) bool {
	return s.s.Equals(p.s)
}

func (s SetOfUInt8) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfUInt8) Empty() bool {
	return s.s.Empty()
}

func (s SetOfUInt8) Len() uint64 {
	return s.s.Len()
}

func (s SetOfUInt8) Has(p uint8) bool {
	return s.s.Has(types.UInt8(p))
}

type SetOfUInt8IterCallback func(p uint8) (stop bool)

func (s SetOfUInt8) Iter(cb SetOfUInt8IterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(uint8(v.(types.UInt8)))
	})
}

type SetOfUInt8IterAllCallback func(p uint8)

func (s SetOfUInt8) IterAll(cb SetOfUInt8IterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(uint8(v.(types.UInt8)))
	})
}

type SetOfUInt8FilterCallback func(p uint8) (keep bool)

func (s SetOfUInt8) Filter(cb SetOfUInt8FilterCallback) SetOfUInt8 {
	ns := NewSetOfUInt8()
	s.IterAll(func(v uint8) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfUInt8) Insert(p ...uint8) SetOfUInt8 {
	return SetOfUInt8{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfUInt8) Remove(p ...uint8) SetOfUInt8 {
	return SetOfUInt8{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfUInt8) Union(others ...SetOfUInt8) SetOfUInt8 {
	return SetOfUInt8{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfUInt8) Subtract(others ...SetOfUInt8) SetOfUInt8 {
	return SetOfUInt8{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfUInt8) Any() uint8 {
	return uint8(s.s.Any().(types.UInt8))
}

func (s SetOfUInt8) fromStructSlice(p []SetOfUInt8) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfUInt8) fromElemSlice(p []uint8) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.UInt8(v)
	}
	return r
}
