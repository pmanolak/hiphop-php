// Autogenerated by Thrift for thrift/compiler/test/fixtures/go-typedef/src/module1.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module1


import (
    "reflect"
    "sync"

    module0 "module0"
    module2 "module2"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

var _ = module0.GoUnusedProtection__
var _ = module2.GoUnusedProtection__
// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = reflect.Ptr

// Premade codec specs
var (
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Plate *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Year *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Drivers *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Accessory *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_module1_Accessory *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_PartName *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_i32_module1_PartName *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Automobile *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i64 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_MapKey *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_module1_MapKey_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_MapContainer *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Car *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Pair *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_module1_Automobile *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_module1_Car *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Collection *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_State *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module1_Enum *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_module1_Plate = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_module1_Year = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_list_string = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.STRING,
	ElementTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_module1_Drivers = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_list_string,
},

    }
    premadeCodecTypeSpec_module1_Accessory = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: module0.GetCodecTypeSpec("module0.Accessory"),
},

    }
    premadeCodecTypeSpec_list_module1_Accessory = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.STRUCT,
	ElementTypeSpec: premadeCodecTypeSpec_module1_Accessory,
},

    }
    premadeCodecTypeSpec_module1_PartName = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: module0.GetCodecTypeSpec("module0.PartName"),
},

    }
    premadeCodecTypeSpec_map_i32_module1_PartName = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_i32,
	ValueTypeSpec: premadeCodecTypeSpec_module1_PartName,
    KeyWireType:   thrift.I32,
	ValueWireType: thrift.STRUCT,
},

    }
    premadeCodecTypeSpec_module1_Automobile = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewAutomobile() },
},

    }
    premadeCodecTypeSpec_i64 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I64,
},

    }
    premadeCodecTypeSpec_module1_MapKey = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMapKey() },
},

    }
    premadeCodecTypeSpec_map_module1_MapKey_string = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_module1_MapKey,
	ValueTypeSpec: premadeCodecTypeSpec_string,
    KeyWireType:   thrift.STRUCT,
	ValueWireType: thrift.STRING,
},

    }
    premadeCodecTypeSpec_module1_MapContainer = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMapContainer() },
},

    }
    premadeCodecTypeSpec_module1_Car = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_module1_Automobile,
},

    }
    premadeCodecTypeSpec_module1_Pair = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewPair() },
},

    }
    premadeCodecTypeSpec_list_module1_Automobile = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.STRUCT,
	ElementTypeSpec: premadeCodecTypeSpec_module1_Automobile,
},

    }
    premadeCodecTypeSpec_list_module1_Car = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.STRUCT,
	ElementTypeSpec: premadeCodecTypeSpec_module1_Car,
},

    }
    premadeCodecTypeSpec_module1_Collection = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewCollection() },
},

    }
    premadeCodecTypeSpec_module1_State = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_module1_Enum = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: module2.GetCodecTypeSpec("module2.Enum"),
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_Automobile *thrift.StructSpec = nil
    premadeStructSpec_MapKey *thrift.StructSpec = nil
    premadeStructSpec_MapContainer *thrift.StructSpec = nil
    premadeStructSpec_Pair *thrift.StructSpec = nil
    premadeStructSpec_Collection *thrift.StructSpec = nil
    premadeStructSpec_reqFinderByPlate *thrift.StructSpec = nil
    premadeStructSpec_respFinderByPlate *thrift.StructSpec = nil
    premadeStructSpec_reqFinderAliasByPlate *thrift.StructSpec = nil
    premadeStructSpec_respFinderAliasByPlate *thrift.StructSpec = nil
    premadeStructSpec_reqFinderPreviousPlate *thrift.StructSpec = nil
    premadeStructSpec_respFinderPreviousPlate *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_Automobile = &thrift.StructSpec{
    Name:               "Automobile",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "plate",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "previous_plate",
            ReflectIndex:         1,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: true,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "first_plate",
            ReflectIndex:         2,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "year",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Year,
            MustBeSetToSerialize: false,
        },        {
            ID:                   5,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "drivers",
            ReflectIndex:         4,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Drivers,
            MustBeSetToSerialize: false,
        },        {
            ID:                   6,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "Accessories",
            ReflectIndex:         5,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_module1_Accessory,
            MustBeSetToSerialize: false,
        },        {
            ID:                   7,
            WireType:             thrift.Type(thrift.MAP),
            Name:                 "PartNames",
            ReflectIndex:         6,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_i32_module1_PartName,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
        5: 4,
        6: 5,
        7: 6,
    },
}
    premadeStructSpec_MapKey = &thrift.StructSpec{
    Name:               "MapKey",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "num",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "strval",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
    },
}
    premadeStructSpec_MapContainer = &thrift.StructSpec{
    Name:               "MapContainer",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.MAP),
            Name:                 "mapval",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_module1_MapKey_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_Pair = &thrift.StructSpec{
    Name:               "Pair",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "automobile",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Automobile,
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "car",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Car,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
    },
}
    premadeStructSpec_Collection = &thrift.StructSpec{
    Name:               "Collection",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "automobiles",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_module1_Automobile,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "cars",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_module1_Car,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
    },
}
    premadeStructSpec_reqFinderByPlate = &thrift.StructSpec{
    Name:               "reqFinderByPlate",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "plate",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respFinderByPlate = &thrift.StructSpec{
    Name:               "respFinderByPlate",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Automobile,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqFinderAliasByPlate = &thrift.StructSpec{
    Name:               "reqFinderAliasByPlate",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "plate",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respFinderAliasByPlate = &thrift.StructSpec{
    Name:               "respFinderAliasByPlate",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Car,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqFinderPreviousPlate = &thrift.StructSpec{
    Name:               "reqFinderPreviousPlate",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "plate",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respFinderPreviousPlate = &thrift.StructSpec{
    Name:               "respFinderPreviousPlate",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module1_Plate,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
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
        results = append(results, codecSpecWithFullName{ "module1.Plate", premadeCodecTypeSpec_module1_Plate })
        results = append(results, codecSpecWithFullName{ "i32", premadeCodecTypeSpec_i32 })
        results = append(results, codecSpecWithFullName{ "module1.Year", premadeCodecTypeSpec_module1_Year })
        results = append(results, codecSpecWithFullName{ "module1.Drivers", premadeCodecTypeSpec_module1_Drivers })
        results = append(results, codecSpecWithFullName{ "module1.Accessory", premadeCodecTypeSpec_module1_Accessory })
        results = append(results, codecSpecWithFullName{ "module1.PartName", premadeCodecTypeSpec_module1_PartName })
        results = append(results, codecSpecWithFullName{ "module1.Automobile", premadeCodecTypeSpec_module1_Automobile })
        results = append(results, codecSpecWithFullName{ "i64", premadeCodecTypeSpec_i64 })
        results = append(results, codecSpecWithFullName{ "module1.MapKey", premadeCodecTypeSpec_module1_MapKey })
        results = append(results, codecSpecWithFullName{ "module1.MapContainer", premadeCodecTypeSpec_module1_MapContainer })
        results = append(results, codecSpecWithFullName{ "module1.Car", premadeCodecTypeSpec_module1_Car })
        results = append(results, codecSpecWithFullName{ "module1.Pair", premadeCodecTypeSpec_module1_Pair })
        results = append(results, codecSpecWithFullName{ "module1.Collection", premadeCodecTypeSpec_module1_Collection })
        results = append(results, codecSpecWithFullName{ "module1.State", premadeCodecTypeSpec_module1_State })
        results = append(results, codecSpecWithFullName{ "module1.Enum", premadeCodecTypeSpec_module1_Enum })
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
