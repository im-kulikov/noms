// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_sha1_b3ecb0f_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("A",
			[]types.Field{
				types.Field{"A", types.MakeCompoundType(types.ListKind, types.MakeCompoundType(types.ListKind, types.MakePrimitiveType(types.BlobKind))), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_sha1_b3ecb0f_CachedRef = types.RegisterPackage(&p)
}

// A

type A struct {
	_A ListOfListOfBlob

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewA(cs chunks.ChunkStore) A {
	return A{
		_A: NewListOfListOfBlob(cs),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type ADef struct {
	A ListOfListOfBlobDef
}

func (def ADef) New(cs chunks.ChunkStore) A {
	return A{
		_A:  def.A.New(cs),
		cs:  cs,
		ref: &ref.Ref{},
	}
}

func (s A) Def() (d ADef) {
	d.A = s._A.Def()
	return
}

var __typeForA types.Type

func (m A) Type() types.Type {
	return __typeForA
}

func init() {
	__typeForA = types.MakeType(__genPackageInFile_sha1_b3ecb0f_CachedRef, 0)
	types.RegisterStruct(__typeForA, builderForA, readerForA)
}

func builderForA(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := A{ref: &ref.Ref{}, cs: cs}
	s._A = values[i].(ListOfListOfBlob)
	i++
	return s
}

func readerForA(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(A)
	values = append(values, s._A)
	return values
}

func (s A) Equals(other types.Value) bool {
	return other != nil && __typeForA.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s A) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s A) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForA.Chunks()...)
	chunks = append(chunks, s._A.Chunks()...)
	return
}

func (s A) ChildValues() (ret []types.Value) {
	ret = append(ret, s._A)
	return
}

func (s A) A() ListOfListOfBlob {
	return s._A
}

func (s A) SetA(val ListOfListOfBlob) A {
	s._A = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfListOfBlob

type ListOfListOfBlob struct {
	l   types.List
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewListOfListOfBlob(cs chunks.ChunkStore) ListOfListOfBlob {
	return ListOfListOfBlob{types.NewTypedList(cs, __typeForListOfListOfBlob), cs, &ref.Ref{}}
}

type ListOfListOfBlobDef []ListOfBlobDef

func (def ListOfListOfBlobDef) New(cs chunks.ChunkStore) ListOfListOfBlob {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New(cs)
	}
	return ListOfListOfBlob{types.NewTypedList(cs, __typeForListOfListOfBlob, l...), cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) Def() ListOfListOfBlobDef {
	d := make([]ListOfBlobDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(ListOfBlob).Def()
	}
	return d
}

func (l ListOfListOfBlob) Equals(other types.Value) bool {
	return other != nil && __typeForListOfListOfBlob.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfListOfBlob) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfListOfBlob) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfListOfBlob) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfListOfBlob.
var __typeForListOfListOfBlob types.Type

func (m ListOfListOfBlob) Type() types.Type {
	return __typeForListOfListOfBlob
}

func init() {
	__typeForListOfListOfBlob = types.MakeCompoundType(types.ListKind, types.MakeCompoundType(types.ListKind, types.MakePrimitiveType(types.BlobKind)))
	types.RegisterValue(__typeForListOfListOfBlob, builderForListOfListOfBlob, readerForListOfListOfBlob)
}

func builderForListOfListOfBlob(cs chunks.ChunkStore, v types.Value) types.Value {
	return ListOfListOfBlob{v.(types.List), cs, &ref.Ref{}}
}

func readerForListOfListOfBlob(v types.Value) types.Value {
	return v.(ListOfListOfBlob).l
}

func (l ListOfListOfBlob) Len() uint64 {
	return l.l.Len()
}

func (l ListOfListOfBlob) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfListOfBlob) Get(i uint64) ListOfBlob {
	return l.l.Get(i).(ListOfBlob)
}

func (l ListOfListOfBlob) Slice(idx uint64, end uint64) ListOfListOfBlob {
	return ListOfListOfBlob{l.l.Slice(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) Set(i uint64, val ListOfBlob) ListOfListOfBlob {
	return ListOfListOfBlob{l.l.Set(i, val), l.cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) Append(v ...ListOfBlob) ListOfListOfBlob {
	return ListOfListOfBlob{l.l.Append(l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) Insert(idx uint64, v ...ListOfBlob) ListOfListOfBlob {
	return ListOfListOfBlob{l.l.Insert(idx, l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) Remove(idx uint64, end uint64) ListOfListOfBlob {
	return ListOfListOfBlob{l.l.Remove(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) RemoveAt(idx uint64) ListOfListOfBlob {
	return ListOfListOfBlob{(l.l.RemoveAt(idx)), l.cs, &ref.Ref{}}
}

func (l ListOfListOfBlob) fromElemSlice(p []ListOfBlob) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfListOfBlobIterCallback func(v ListOfBlob, i uint64) (stop bool)

func (l ListOfListOfBlob) Iter(cb ListOfListOfBlobIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(ListOfBlob), i)
	})
}

type ListOfListOfBlobIterAllCallback func(v ListOfBlob, i uint64)

func (l ListOfListOfBlob) IterAll(cb ListOfListOfBlobIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(ListOfBlob), i)
	})
}

func (l ListOfListOfBlob) IterAllP(concurrency int, cb ListOfListOfBlobIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(ListOfBlob), i)
	})
}

type ListOfListOfBlobFilterCallback func(v ListOfBlob, i uint64) (keep bool)

func (l ListOfListOfBlob) Filter(cb ListOfListOfBlobFilterCallback) ListOfListOfBlob {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(ListOfBlob), i)
	})
	return ListOfListOfBlob{out, l.cs, &ref.Ref{}}
}

// ListOfBlob

type ListOfBlob struct {
	l   types.List
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewListOfBlob(cs chunks.ChunkStore) ListOfBlob {
	return ListOfBlob{types.NewTypedList(cs, __typeForListOfBlob), cs, &ref.Ref{}}
}

type ListOfBlobDef []types.Blob

func (def ListOfBlobDef) New(cs chunks.ChunkStore) ListOfBlob {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d
	}
	return ListOfBlob{types.NewTypedList(cs, __typeForListOfBlob, l...), cs, &ref.Ref{}}
}

func (l ListOfBlob) Def() ListOfBlobDef {
	d := make([]types.Blob, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(types.Blob)
	}
	return d
}

func (l ListOfBlob) Equals(other types.Value) bool {
	return other != nil && __typeForListOfBlob.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfBlob) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfBlob) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfBlob) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfBlob.
var __typeForListOfBlob types.Type

func (m ListOfBlob) Type() types.Type {
	return __typeForListOfBlob
}

func init() {
	__typeForListOfBlob = types.MakeCompoundType(types.ListKind, types.MakePrimitiveType(types.BlobKind))
	types.RegisterValue(__typeForListOfBlob, builderForListOfBlob, readerForListOfBlob)
}

func builderForListOfBlob(cs chunks.ChunkStore, v types.Value) types.Value {
	return ListOfBlob{v.(types.List), cs, &ref.Ref{}}
}

func readerForListOfBlob(v types.Value) types.Value {
	return v.(ListOfBlob).l
}

func (l ListOfBlob) Len() uint64 {
	return l.l.Len()
}

func (l ListOfBlob) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfBlob) Get(i uint64) types.Blob {
	return l.l.Get(i).(types.Blob)
}

func (l ListOfBlob) Slice(idx uint64, end uint64) ListOfBlob {
	return ListOfBlob{l.l.Slice(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfBlob) Set(i uint64, val types.Blob) ListOfBlob {
	return ListOfBlob{l.l.Set(i, val), l.cs, &ref.Ref{}}
}

func (l ListOfBlob) Append(v ...types.Blob) ListOfBlob {
	return ListOfBlob{l.l.Append(l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfBlob) Insert(idx uint64, v ...types.Blob) ListOfBlob {
	return ListOfBlob{l.l.Insert(idx, l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfBlob) Remove(idx uint64, end uint64) ListOfBlob {
	return ListOfBlob{l.l.Remove(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfBlob) RemoveAt(idx uint64) ListOfBlob {
	return ListOfBlob{(l.l.RemoveAt(idx)), l.cs, &ref.Ref{}}
}

func (l ListOfBlob) fromElemSlice(p []types.Blob) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfBlobIterCallback func(v types.Blob, i uint64) (stop bool)

func (l ListOfBlob) Iter(cb ListOfBlobIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(types.Blob), i)
	})
}

type ListOfBlobIterAllCallback func(v types.Blob, i uint64)

func (l ListOfBlob) IterAll(cb ListOfBlobIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(types.Blob), i)
	})
}

func (l ListOfBlob) IterAllP(concurrency int, cb ListOfBlobIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(types.Blob), i)
	})
}

type ListOfBlobFilterCallback func(v types.Blob, i uint64) (keep bool)

func (l ListOfBlob) Filter(cb ListOfBlobFilterCallback) ListOfBlob {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(types.Blob), i)
	})
	return ListOfBlob{out, l.cs, &ref.Ref{}}
}
