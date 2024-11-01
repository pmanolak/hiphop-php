// Autogenerated by Thrift for thrift/compiler/test/fixtures/basic/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module


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
    premadeCodecTypeSpec_module_MyEnum *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_HackEnum *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i64 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyDataItem *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyDataItemAlias *thrift.TypeSpec = nil
    premadeCodecTypeSpec_bool *thrift.TypeSpec = nil
    premadeCodecTypeSpec_float *thrift.TypeSpec = nil
    premadeCodecTypeSpec_set_float *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyStruct *thrift.TypeSpec = nil
    premadeCodecTypeSpec_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_list_i32 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_set_string *thrift.TypeSpec = nil
    premadeCodecTypeSpec_map_string_i64 *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_Containers *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyEnumAlias *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyUnion *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_ReservedKeyword *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_UnionToBeRenamed *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyException *thrift.TypeSpec = nil
    premadeCodecTypeSpec_module_MyExceptionWithMessage *thrift.TypeSpec = nil
    premadeCodecTypeSpec_void *thrift.TypeSpec = nil
    premadeCodecTypeSpec_binary *thrift.TypeSpec = nil
)

// Premade codec specs initializer
var premadeCodecSpecsInitOnce = sync.OnceFunc(func() {
    premadeCodecTypeSpec_module_MyEnum = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_module_HackEnum = &thrift.TypeSpec{
        CodecEnumSpec: &thrift.CodecEnumSpec{},

    }
    premadeCodecTypeSpec_i64 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I64,
},

    }
    premadeCodecTypeSpec_string = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_STRING,
},

    }
    premadeCodecTypeSpec_module_MyDataItem = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyDataItem() },
},

    }
    premadeCodecTypeSpec_module_MyDataItemAlias = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_module_MyDataItem,
},

    }
    premadeCodecTypeSpec_bool = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_BOOL,
},

    }
    premadeCodecTypeSpec_float = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_FLOAT,
},

    }
    premadeCodecTypeSpec_set_float = &thrift.TypeSpec{
        CodecSetSpec: &thrift.CodecSetSpec{
    ElementWireType: thrift.FLOAT,
	ElementTypeSpec: premadeCodecTypeSpec_float,
},

    }
    premadeCodecTypeSpec_module_MyStruct = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyStruct() },
},

    }
    premadeCodecTypeSpec_i32 = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_I32,
},

    }
    premadeCodecTypeSpec_list_i32 = &thrift.TypeSpec{
        CodecListSpec: &thrift.CodecListSpec{
    ElementWireType: thrift.I32,
	ElementTypeSpec: premadeCodecTypeSpec_i32,
},

    }
    premadeCodecTypeSpec_set_string = &thrift.TypeSpec{
        CodecSetSpec: &thrift.CodecSetSpec{
    ElementWireType: thrift.STRING,
	ElementTypeSpec: premadeCodecTypeSpec_string,
},

    }
    premadeCodecTypeSpec_map_string_i64 = &thrift.TypeSpec{
        CodecMapSpec: &thrift.CodecMapSpec{
	KeyTypeSpec:   premadeCodecTypeSpec_string,
	ValueTypeSpec: premadeCodecTypeSpec_i64,
    KeyWireType:   thrift.STRING,
	ValueWireType: thrift.I64,
},

    }
    premadeCodecTypeSpec_module_Containers = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewContainers() },
},

    }
    premadeCodecTypeSpec_module_MyEnumAlias = &thrift.TypeSpec{
        CodecTypedefSpec: &thrift.CodecTypedefSpec{
	UnderlyingTypeSpec: premadeCodecTypeSpec_module_MyEnum,
},

    }
    premadeCodecTypeSpec_module_MyUnion = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyUnion() },
},

    }
    premadeCodecTypeSpec_module_ReservedKeyword = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewReservedKeyword() },
},

    }
    premadeCodecTypeSpec_module_UnionToBeRenamed = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewUnionToBeRenamed() },
},

    }
    premadeCodecTypeSpec_module_MyException = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyException() },
},

    }
    premadeCodecTypeSpec_module_MyExceptionWithMessage = &thrift.TypeSpec{
        CodecStructSpec: &thrift.CodecStructSpec{
    NewFunc: func() thrift.Struct { return NewMyExceptionWithMessage() },
},

    }
    premadeCodecTypeSpec_void = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_VOID,
},

    }
    premadeCodecTypeSpec_binary = &thrift.TypeSpec{
        CodecPrimitiveSpec: &thrift.CodecPrimitiveSpec{
    PrimitiveType: thrift.CODEC_PRIMITIVE_TYPE_BINARY,
},

    }
})

// Premade struct specs
var (
    premadeStructSpec_MyStruct *thrift.StructSpec = nil
    premadeStructSpec_Containers *thrift.StructSpec = nil
    premadeStructSpec_MyDataItem *thrift.StructSpec = nil
    premadeStructSpec_MyUnion *thrift.StructSpec = nil
    premadeStructSpec_MyException *thrift.StructSpec = nil
    premadeStructSpec_MyExceptionWithMessage *thrift.StructSpec = nil
    premadeStructSpec_ReservedKeyword *thrift.StructSpec = nil
    premadeStructSpec_UnionToBeRenamed *thrift.StructSpec = nil
    premadeStructSpec_reqFooServiceSimpleRPC *thrift.StructSpec = nil
    premadeStructSpec_respFooServiceSimpleRPC *thrift.StructSpec = nil
    premadeStructSpec_reqFB303ServiceSimpleRPC *thrift.StructSpec = nil
    premadeStructSpec_respFB303ServiceSimpleRPC *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePing *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePing *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceGetRandomData *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceGetRandomData *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceSink *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceSink *thrift.StructSpec = nil
    premadeStructSpec_reqMyServicePutDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServicePutDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceHasDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceHasDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceGetDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceGetDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceDeleteDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceDeleteDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceLobDataById *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceLobDataById *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceInvalidReturnForHack *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceInvalidReturnForHack *thrift.StructSpec = nil
    premadeStructSpec_reqMyServiceRpcSkippedCodegen *thrift.StructSpec = nil
    premadeStructSpec_respMyServiceRpcSkippedCodegen *thrift.StructSpec = nil
    premadeStructSpec_reqDbMixedStackArgumentsGetDataByKey0 *thrift.StructSpec = nil
    premadeStructSpec_respDbMixedStackArgumentsGetDataByKey0 *thrift.StructSpec = nil
    premadeStructSpec_reqDbMixedStackArgumentsGetDataByKey1 *thrift.StructSpec = nil
    premadeStructSpec_respDbMixedStackArgumentsGetDataByKey1 *thrift.StructSpec = nil
)

// Premade struct specs initializer
var premadeStructSpecsInitOnce = sync.OnceFunc(func() {
    premadeStructSpec_MyStruct = &thrift.StructSpec{
    Name:               "MyStruct",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "MyIntField",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "MyStringField",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "MyDataField",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyDataItemAlias,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "myEnum",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnum,
            MustBeSetToSerialize: false,
        },        {
            ID:                   5,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "oneway",
            ReflectIndex:         4,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: false,
        },        {
            ID:                   6,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "readonly",
            ReflectIndex:         5,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: false,
        },        {
            ID:                   7,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "idempotent",
            ReflectIndex:         6,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: false,
        },        {
            ID:                   8,
            WireType:             thrift.Type(thrift.SET),
            Name:                 "floatSet",
            ReflectIndex:         7,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_set_float,
            MustBeSetToSerialize: false,
        },        {
            ID:                   9,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "no_hack_codegen_field",
            ReflectIndex:         8,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
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
        8: 7,
        9: 8,
    },
}
    premadeStructSpec_Containers = &thrift.StructSpec{
    Name:               "Containers",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.LIST),
            Name:                 "I32List",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_list_i32,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.SET),
            Name:                 "StringSet",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_set_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.MAP),
            Name:                 "StringToI64Map",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_map_string_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
        3: 2,
    },
}
    premadeStructSpec_MyDataItem = &thrift.StructSpec{
    Name:               "MyDataItem",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_MyUnion = &thrift.StructSpec{
    Name:               "MyUnion",
    IsUnion:            true,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "myEnum",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyEnumAlias,
            MustBeSetToSerialize: true,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myStruct",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyStruct,
            MustBeSetToSerialize: true,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myDataItem",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyDataItem,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.SET),
            Name:                 "floatSet",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_set_float,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
    },
}
    premadeStructSpec_MyException = &thrift.StructSpec{
    Name:               "MyException",
    IsUnion:            false,
    IsException:        true,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "MyIntField",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "MyStringField",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myStruct",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyStruct,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myUnion",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyUnion,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
    },
}
    premadeStructSpec_MyExceptionWithMessage = &thrift.StructSpec{
    Name:               "MyExceptionWithMessage",
    IsUnion:            false,
    IsException:        true,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "MyIntField",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "MyStringField",
            ReflectIndex:         1,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },        {
            ID:                   3,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myStruct",
            ReflectIndex:         2,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyStruct,
            MustBeSetToSerialize: true,
        },        {
            ID:                   4,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "myUnion",
            ReflectIndex:         3,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_module_MyUnion,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
        2: 1,
        3: 2,
        4: 3,
    },
}
    premadeStructSpec_ReservedKeyword = &thrift.StructSpec{
    Name:               "ReservedKeyword",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "reserved_field",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_UnionToBeRenamed = &thrift.StructSpec{
    Name:               "UnionToBeRenamed",
    IsUnion:            true,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "reserved_field",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_reqFooServiceSimpleRPC = &thrift.StructSpec{
    Name:               "reqFooServiceSimpleRPC",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_respFooServiceSimpleRPC = &thrift.StructSpec{
    Name:               "respFooServiceSimpleRPC",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqFB303ServiceSimpleRPC = &thrift.StructSpec{
    Name:               "reqFB303ServiceSimpleRPC",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I32),
            Name:                 "int_parameter",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i32,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respFB303ServiceSimpleRPC = &thrift.StructSpec{
    Name:               "respFB303ServiceSimpleRPC",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRUCT),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_module_ReservedKeyword,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqMyServicePing = &thrift.StructSpec{
    Name:               "reqMyServicePing",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_respMyServicePing = &thrift.StructSpec{
    Name:               "respMyServicePing",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqMyServiceGetRandomData = &thrift.StructSpec{
    Name:               "reqMyServiceGetRandomData",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_respMyServiceGetRandomData = &thrift.StructSpec{
    Name:               "respMyServiceGetRandomData",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqMyServiceSink = &thrift.StructSpec{
    Name:               "reqMyServiceSink",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "sink",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respMyServiceSink = &thrift.StructSpec{
    Name:               "respMyServiceSink",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqMyServicePutDataById = &thrift.StructSpec{
    Name:               "reqMyServicePutDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "data",
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
    premadeStructSpec_respMyServicePutDataById = &thrift.StructSpec{
    Name:               "respMyServicePutDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqMyServiceHasDataById = &thrift.StructSpec{
    Name:               "reqMyServiceHasDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respMyServiceHasDataById = &thrift.StructSpec{
    Name:               "respMyServiceHasDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.BOOL),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_bool,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqMyServiceGetDataById = &thrift.StructSpec{
    Name:               "reqMyServiceGetDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respMyServiceGetDataById = &thrift.StructSpec{
    Name:               "respMyServiceGetDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqMyServiceDeleteDataById = &thrift.StructSpec{
    Name:               "reqMyServiceDeleteDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respMyServiceDeleteDataById = &thrift.StructSpec{
    Name:               "respMyServiceDeleteDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqMyServiceLobDataById = &thrift.StructSpec{
    Name:               "reqMyServiceLobDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.I64),
            Name:                 "id",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_i64,
            MustBeSetToSerialize: false,
        },        {
            ID:                   2,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "data",
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
    premadeStructSpec_respMyServiceLobDataById = &thrift.StructSpec{
    Name:               "respMyServiceLobDataById",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqMyServiceInvalidReturnForHack = &thrift.StructSpec{
    Name:               "reqMyServiceInvalidReturnForHack",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_respMyServiceInvalidReturnForHack = &thrift.StructSpec{
    Name:               "respMyServiceInvalidReturnForHack",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.SET),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_set_float,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqMyServiceRpcSkippedCodegen = &thrift.StructSpec{
    Name:               "reqMyServiceRpcSkippedCodegen",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_respMyServiceRpcSkippedCodegen = &thrift.StructSpec{
    Name:               "respMyServiceRpcSkippedCodegen",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
    },
    FieldSpecIDToIndex: map[int16]int{
    },
}
    premadeStructSpec_reqDbMixedStackArgumentsGetDataByKey0 = &thrift.StructSpec{
    Name:               "reqDbMixedStackArgumentsGetDataByKey0",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "key",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respDbMixedStackArgumentsGetDataByKey0 = &thrift.StructSpec{
    Name:               "respDbMixedStackArgumentsGetDataByKey0",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_binary,
            MustBeSetToSerialize: true,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        0: 0,
    },
}
    premadeStructSpec_reqDbMixedStackArgumentsGetDataByKey1 = &thrift.StructSpec{
    Name:               "reqDbMixedStackArgumentsGetDataByKey1",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   1,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "key",
            ReflectIndex:         0,
            IsOptional:           false,
            ValueTypeSpec:        premadeCodecTypeSpec_string,
            MustBeSetToSerialize: false,
        },    },
    FieldSpecIDToIndex: map[int16]int{
        1: 0,
    },
}
    premadeStructSpec_respDbMixedStackArgumentsGetDataByKey1 = &thrift.StructSpec{
    Name:               "respDbMixedStackArgumentsGetDataByKey1",
    IsUnion:            false,
    IsException:        false,
    FieldSpecs:         []thrift.FieldSpec{
        {
            ID:                   0,
            WireType:             thrift.Type(thrift.STRING),
            Name:                 "success",
            ReflectIndex:         0,
            IsOptional:           true,
            ValueTypeSpec:        premadeCodecTypeSpec_binary,
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
        results = append(results, codecSpecWithFullName{ "module.MyEnum", premadeCodecTypeSpec_module_MyEnum })
        results = append(results, codecSpecWithFullName{ "module.HackEnum", premadeCodecTypeSpec_module_HackEnum })
        results = append(results, codecSpecWithFullName{ "i64", premadeCodecTypeSpec_i64 })
        results = append(results, codecSpecWithFullName{ "string", premadeCodecTypeSpec_string })
        results = append(results, codecSpecWithFullName{ "module.MyDataItem", premadeCodecTypeSpec_module_MyDataItem })
        results = append(results, codecSpecWithFullName{ "module.MyDataItemAlias", premadeCodecTypeSpec_module_MyDataItemAlias })
        results = append(results, codecSpecWithFullName{ "bool", premadeCodecTypeSpec_bool })
        results = append(results, codecSpecWithFullName{ "float", premadeCodecTypeSpec_float })
        results = append(results, codecSpecWithFullName{ "module.MyStruct", premadeCodecTypeSpec_module_MyStruct })
        results = append(results, codecSpecWithFullName{ "i32", premadeCodecTypeSpec_i32 })
        results = append(results, codecSpecWithFullName{ "module.Containers", premadeCodecTypeSpec_module_Containers })
        results = append(results, codecSpecWithFullName{ "module.MyEnumAlias", premadeCodecTypeSpec_module_MyEnumAlias })
        results = append(results, codecSpecWithFullName{ "module.MyUnion", premadeCodecTypeSpec_module_MyUnion })
        results = append(results, codecSpecWithFullName{ "module.ReservedKeyword", premadeCodecTypeSpec_module_ReservedKeyword })
        results = append(results, codecSpecWithFullName{ "module.UnionToBeRenamed", premadeCodecTypeSpec_module_UnionToBeRenamed })
        results = append(results, codecSpecWithFullName{ "module.MyException", premadeCodecTypeSpec_module_MyException })
        results = append(results, codecSpecWithFullName{ "module.MyExceptionWithMessage", premadeCodecTypeSpec_module_MyExceptionWithMessage })
        results = append(results, codecSpecWithFullName{ "void", premadeCodecTypeSpec_void })
        results = append(results, codecSpecWithFullName{ "binary", premadeCodecTypeSpec_binary })
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
