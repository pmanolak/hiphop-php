
#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from cpython.ref cimport PyObject
from libcpp.utility cimport move as __move
from folly.iobuf cimport (
    from_unique_ptr as __IOBuf_from_unique_ptr,
    IOBuf as __IOBuf,
)
from thrift.python.serializer import (
    deserialize as __deserialize,
    Protocol as __Protocol,
    serialize_iobuf as __serialize_iobuf,
)
import test.namespace_from_package.module.thrift_types as __thrift_types

cdef api int can_extract__test__namespace_from_package__module__Foo(object __obj) except -1:
    return 1 if isinstance(__obj, __thrift_types.Foo) else 0

cdef api __cIOBuf* extract__test__namespace_from_package__module__Foo(object __obj) except NULL:
    cdef __IOBuf __buf = __serialize_iobuf(__obj, protocol=__Protocol.BINARY)
    return __buf._ours.release()

cdef api object construct__test__namespace_from_package__module__Foo(__unique_ptr[__cIOBuf] __s):
    return __deserialize(
        __thrift_types.Foo,
        __IOBuf_from_unique_ptr(__move(__s)),
        protocol=__Protocol.BINARY
    )

cdef api object init__test__namespace_from_package__module__Foo(object data):
    return __thrift_types.Foo._fbthrift_create(data)
