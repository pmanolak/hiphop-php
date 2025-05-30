
#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/refs/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libcpp.memory cimport make_shared, unique_ptr
from cython.operator cimport dereference as deref, address
from libcpp.utility cimport move as cmove
cimport module.types as _fbthrift_ctypes
from thrift.py3.serializer cimport (
    cserialize as __cserialize,
    cdeserialize as __cdeserialize,
)
from thrift.python.protocol cimport Protocol
cimport folly.iobuf as _folly__iobuf


cdef shared_ptr[_fbthrift_cbindings.cMyUnion] MyUnion_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cMyUnion](deref(
        (<_fbthrift_ctypes.MyUnion?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object MyUnion_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyUnion]& c_struct):
    return _fbthrift_ctypes.MyUnion._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cNonTriviallyDestructibleUnion] NonTriviallyDestructibleUnion_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cNonTriviallyDestructibleUnion](deref(
        (<_fbthrift_ctypes.NonTriviallyDestructibleUnion?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object NonTriviallyDestructibleUnion_from_cpp(const shared_ptr[_fbthrift_cbindings.cNonTriviallyDestructibleUnion]& c_struct):
    return _fbthrift_ctypes.NonTriviallyDestructibleUnion._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMyField] MyField_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cMyField](deref(
        (<_fbthrift_ctypes.MyField?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object MyField_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyField]& c_struct):
    return _fbthrift_ctypes.MyField._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cMyStruct] MyStruct_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cMyStruct](deref(
        (<_fbthrift_ctypes.MyStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object MyStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cMyStruct]& c_struct):
    return _fbthrift_ctypes.MyStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithUnion] StructWithUnion_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithUnion](deref(
        (<_fbthrift_ctypes.StructWithUnion?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithUnion_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithUnion]& c_struct):
    return _fbthrift_ctypes.StructWithUnion._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cRecursiveStruct] RecursiveStruct_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cRecursiveStruct](deref(
        (<_fbthrift_ctypes.RecursiveStruct?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object RecursiveStruct_from_cpp(const shared_ptr[_fbthrift_cbindings.cRecursiveStruct]& c_struct):
    return _fbthrift_ctypes.RecursiveStruct._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithContainers] StructWithContainers_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithContainers](deref(
        (<_fbthrift_ctypes.StructWithContainers?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithContainers_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithContainers]& c_struct):
    return _fbthrift_ctypes.StructWithContainers._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithSharedConst] StructWithSharedConst_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithSharedConst](deref(
        (<_fbthrift_ctypes.StructWithSharedConst?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithSharedConst_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithSharedConst]& c_struct):
    return _fbthrift_ctypes.StructWithSharedConst._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cEmpty] Empty_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cEmpty](deref(
        (<_fbthrift_ctypes.Empty?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object Empty_from_cpp(const shared_ptr[_fbthrift_cbindings.cEmpty]& c_struct):
    return _fbthrift_ctypes.Empty._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithRef] StructWithRef_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithRef](deref(
        (<_fbthrift_ctypes.StructWithRef?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithRef_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithRef]& c_struct):
    return _fbthrift_ctypes.StructWithRef._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithBox] StructWithBox_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithBox](deref(
        (<_fbthrift_ctypes.StructWithBox?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithBox_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithBox]& c_struct):
    return _fbthrift_ctypes.StructWithBox._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithInternBox] StructWithInternBox_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithInternBox](deref(
        (<_fbthrift_ctypes.StructWithInternBox?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithInternBox_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithInternBox]& c_struct):
    return _fbthrift_ctypes.StructWithInternBox._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithTerseInternBox] StructWithTerseInternBox_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithTerseInternBox](deref(
        (<_fbthrift_ctypes.StructWithTerseInternBox?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithTerseInternBox_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithTerseInternBox]& c_struct):
    return _fbthrift_ctypes.StructWithTerseInternBox._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cAdaptedStructWithInternBox] AdaptedStructWithInternBox_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cAdaptedStructWithInternBox](deref(
        (<_fbthrift_ctypes.AdaptedStructWithInternBox?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object AdaptedStructWithInternBox_from_cpp(const shared_ptr[_fbthrift_cbindings.cAdaptedStructWithInternBox]& c_struct):
    return _fbthrift_ctypes.AdaptedStructWithInternBox._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cAdaptedStructWithTerseInternBox] AdaptedStructWithTerseInternBox_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cAdaptedStructWithTerseInternBox](deref(
        (<_fbthrift_ctypes.AdaptedStructWithTerseInternBox?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object AdaptedStructWithTerseInternBox_from_cpp(const shared_ptr[_fbthrift_cbindings.cAdaptedStructWithTerseInternBox]& c_struct):
    return _fbthrift_ctypes.AdaptedStructWithTerseInternBox._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithRefTypeUnique] StructWithRefTypeUnique_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithRefTypeUnique](deref(
        (<_fbthrift_ctypes.StructWithRefTypeUnique?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithRefTypeUnique_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithRefTypeUnique]& c_struct):
    return _fbthrift_ctypes.StructWithRefTypeUnique._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithRefTypeShared] StructWithRefTypeShared_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithRefTypeShared](deref(
        (<_fbthrift_ctypes.StructWithRefTypeShared?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithRefTypeShared_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithRefTypeShared]& c_struct):
    return _fbthrift_ctypes.StructWithRefTypeShared._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithRefTypeSharedConst] StructWithRefTypeSharedConst_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithRefTypeSharedConst](deref(
        (<_fbthrift_ctypes.StructWithRefTypeSharedConst?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithRefTypeSharedConst_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithRefTypeSharedConst]& c_struct):
    return _fbthrift_ctypes.StructWithRefTypeSharedConst._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithRefAndAnnotCppNoexceptMoveCtor] StructWithRefAndAnnotCppNoexceptMoveCtor_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithRefAndAnnotCppNoexceptMoveCtor](deref(
        (<_fbthrift_ctypes.StructWithRefAndAnnotCppNoexceptMoveCtor?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithRefAndAnnotCppNoexceptMoveCtor_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithRefAndAnnotCppNoexceptMoveCtor]& c_struct):
    return _fbthrift_ctypes.StructWithRefAndAnnotCppNoexceptMoveCtor._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)

cdef shared_ptr[_fbthrift_cbindings.cStructWithString] StructWithString_convert_to_cpp(object inst) except*:
    return make_shared[_fbthrift_cbindings.cStructWithString](deref(
        (<_fbthrift_ctypes.StructWithString?>inst)._cpp_obj_FBTHRIFT_ONLY_DO_NOT_USE
    ))

cdef object StructWithString_from_cpp(const shared_ptr[_fbthrift_cbindings.cStructWithString]& c_struct):
    return _fbthrift_ctypes.StructWithString._create_FBTHRIFT_ONLY_DO_NOT_USE(c_struct)


