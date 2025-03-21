
#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from thrift.python.capi.cpp_converter cimport cpp_to_python, python_to_cpp
from libcpp.utility cimport move as cmove

cdef extern from "thrift/compiler/test/fixtures/complex-struct/gen-python-capi/module/thrift_types_capi.h":
    pass

cdef cMyStructFloatFieldThrowExp MyStructFloatFieldThrowExp_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyStructFloatFieldThrowExp](inst))

cdef object MyStructFloatFieldThrowExp_from_cpp(const cMyStructFloatFieldThrowExp& c_struct):
    return cpp_to_python[cMyStructFloatFieldThrowExp](c_struct)

cdef cMyStructMapFloatThrowExp MyStructMapFloatThrowExp_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyStructMapFloatThrowExp](inst))

cdef object MyStructMapFloatThrowExp_from_cpp(const cMyStructMapFloatThrowExp& c_struct):
    return cpp_to_python[cMyStructMapFloatThrowExp](c_struct)

cdef cMyStruct MyStruct_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyStruct](inst))

cdef object MyStruct_from_cpp(const cMyStruct& c_struct):
    return cpp_to_python[cMyStruct](c_struct)

cdef cSimpleStruct SimpleStruct_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cSimpleStruct](inst))

cdef object SimpleStruct_from_cpp(const cSimpleStruct& c_struct):
    return cpp_to_python[cSimpleStruct](c_struct)

cdef cdefaultStruct defaultStruct_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cdefaultStruct](inst))

cdef object defaultStruct_from_cpp(const cdefaultStruct& c_struct):
    return cpp_to_python[cdefaultStruct](c_struct)

cdef cMyStructTypeDef MyStructTypeDef_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyStructTypeDef](inst))

cdef object MyStructTypeDef_from_cpp(const cMyStructTypeDef& c_struct):
    return cpp_to_python[cMyStructTypeDef](c_struct)

cdef cMyDataItem MyDataItem_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyDataItem](inst))

cdef object MyDataItem_from_cpp(const cMyDataItem& c_struct):
    return cpp_to_python[cMyDataItem](c_struct)

cdef cMyUnion MyUnion_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyUnion](inst))

cdef object MyUnion_from_cpp(const cMyUnion& c_struct):
    return cpp_to_python[cMyUnion](c_struct)

cdef cMyUnionFloatFieldThrowExp MyUnionFloatFieldThrowExp_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cMyUnionFloatFieldThrowExp](inst))

cdef object MyUnionFloatFieldThrowExp_from_cpp(const cMyUnionFloatFieldThrowExp& c_struct):
    return cpp_to_python[cMyUnionFloatFieldThrowExp](c_struct)

cdef cComplexNestedStruct ComplexNestedStruct_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cComplexNestedStruct](inst))

cdef object ComplexNestedStruct_from_cpp(const cComplexNestedStruct& c_struct):
    return cpp_to_python[cComplexNestedStruct](c_struct)

cdef cTypeRemapped TypeRemapped_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cTypeRemapped](inst))

cdef object TypeRemapped_from_cpp(const cTypeRemapped& c_struct):
    return cpp_to_python[cTypeRemapped](c_struct)

cdef cemptyXcep emptyXcep_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cemptyXcep](inst))

cdef object emptyXcep_from_cpp(const cemptyXcep& c_struct):
    return cpp_to_python[cemptyXcep](c_struct)

cdef creqXcep reqXcep_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[creqXcep](inst))

cdef object reqXcep_from_cpp(const creqXcep& c_struct):
    return cpp_to_python[creqXcep](c_struct)

cdef coptXcep optXcep_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[coptXcep](inst))

cdef object optXcep_from_cpp(const coptXcep& c_struct):
    return cpp_to_python[coptXcep](c_struct)

cdef ccomplexException complexException_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[ccomplexException](inst))

cdef object complexException_from_cpp(const ccomplexException& c_struct):
    return cpp_to_python[ccomplexException](c_struct)

cdef cContainers Containers_convert_to_cpp(object inst) except *:
    return cmove(python_to_cpp[cContainers](inst))

cdef object Containers_from_cpp(const cContainers& c_struct):
    return cpp_to_python[cContainers](c_struct)

