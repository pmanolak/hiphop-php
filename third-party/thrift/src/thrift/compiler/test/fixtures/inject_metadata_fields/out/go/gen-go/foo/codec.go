// Autogenerated by Thrift for foo.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package foo


import (
    "reflect"
    "sync"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = reflect.Ptr

// Premade codec specs
var (
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_foo_Fields *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_foo_Fields = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewFields() },
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Fields *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Fields = &thrift.StructSpec{
    Name:               "Fields",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   100,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "injected_field",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   101,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "injected_structured_annotation_field",
            ReflectIndex:         1,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },        {
            ID:                   102,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "injected_unstructured_annotation_field",
            ReflectIndex:         2,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        100: 0,
        101: 1,
        102: 2,
    },
}
})

// Helper type to allow us to store codec specs in a slice at compile time,
// and put them in a map at runtime. See comment at the top of template
// about a compilation limitation that affects map literals.
type codecSpecWithFullName struct {
    fullName string
    typeSpec *thrift.TypeSpec
}

var premadeCodecSpecsSliceOnce = sync.OnceValue(
    func() []codecSpecWithFullName {
        // Relies on premade codec specs initialization
        premadeCodecSpecsInitOnce()
        results := make([]codecSpecWithFullName, 0)
        results = append(results, codecSpecWithFullName{ "string", premadeCodecTypeSpec_string })
        results = append(results, codecSpecWithFullName{ "foo.Fields", premadeCodecTypeSpec_foo_Fields })
        return results
    },
)

var premadeCodecSpecsMapOnce = sync.OnceValue(
    func() map[string]*thrift.TypeSpec {
        codecSpecsWithFullName := premadeCodecSpecsSliceOnce()
        results := make(map[string]*thrift.TypeSpec, len(codecSpecsWithFullName))
        for _, value := range codecSpecsWithFullName {
            results[value.fullName] = value.typeSpec
        }
        return results
    },
)

func init() {
    premadeCodecSpecsInitOnce()
    premadeStructSpecsInitOnce()
}

// GetMetadataThriftType (INTERNAL USE ONLY).
// Returns metadata TypeSpec for a given full type name.
func GetCodecTypeSpec(fullName string) *thrift.TypeSpec {
    return premadeCodecSpecsMapOnce()[fullName]
}
