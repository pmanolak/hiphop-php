#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/mcpp2-compare/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from cpython.ref cimport PyObject
from libc.stdint cimport (
    int8_t as cint8_t,
    int16_t as cint16_t,
    int32_t as cint32_t,
    int64_t as cint64_t,
)
from libcpp cimport bool as cbool
from libcpp.map cimport map as cmap, pair as cpair
from libcpp.memory cimport shared_ptr, unique_ptr
from libcpp.set cimport set as cset
from libcpp.string cimport string
from libcpp.vector cimport vector

from folly cimport cFollyFuture, cFollyTry, cFollyUnit
cimport folly.iobuf as _fbthrift_iobuf
from thrift.python.common cimport cRpcOptions
from thrift.py3.client cimport cClientWrapper

cimport module.types as _module_types
cimport module.cbindings as _module_cbindings

cimport includes.types as _includes_types
cimport includes.cbindings as _includes_cbindings

cdef extern from "thrift/compiler/test/fixtures/mcpp2-compare/gen-cpp2/module_clients.h" namespace "::some::valid::ns":
  cdef cppclass cEmptyServiceAsyncClient "::some::valid::ns::EmptyServiceAsyncClient":
      pass

cdef extern from "<utility>" namespace "std":
  cdef unique_ptr[cEmptyServiceClientWrapper] move(unique_ptr[cEmptyServiceClientWrapper])

cdef extern from "thrift/compiler/test/fixtures/mcpp2-compare/gen-cpp2/module_clients.h" namespace "::some::valid::ns":
  cdef cppclass cReturnServiceAsyncClient "::some::valid::ns::ReturnServiceAsyncClient":
      pass

cdef extern from "<utility>" namespace "std":
  cdef unique_ptr[cReturnServiceClientWrapper] move(unique_ptr[cReturnServiceClientWrapper])

cdef extern from "thrift/compiler/test/fixtures/mcpp2-compare/gen-cpp2/module_clients.h" namespace "::some::valid::ns":
  cdef cppclass cParamServiceAsyncClient "::some::valid::ns::ParamServiceAsyncClient":
      pass

cdef extern from "<utility>" namespace "std":
  cdef unique_ptr[cParamServiceClientWrapper] move(unique_ptr[cParamServiceClientWrapper])

cdef extern from "thrift/lib/cpp/TProcessorEventHandler.h" namespace "::apache::thrift":
  cdef cppclass cTProcessorEventHandler "apache::thrift::TProcessorEventHandler":
    pass

cdef extern from "thrift/compiler/test/fixtures/mcpp2-compare/gen-py3/module/clients_wrapper.h" namespace "::some::valid::ns":
  cdef cppclass cEmptyServiceClientWrapper "::some::valid::ns::EmptyServiceClientWrapper":
    void setPersistentHeader(const string& key, const string& value)
    void addEventHandler(const shared_ptr[cTProcessorEventHandler]& handler)



  cdef cppclass cReturnServiceClientWrapper "::some::valid::ns::ReturnServiceClientWrapper":
    void setPersistentHeader(const string& key, const string& value)
    void addEventHandler(const shared_ptr[cTProcessorEventHandler]& handler)

    cFollyFuture[cFollyUnit] noReturn(cRpcOptions, )
    cFollyFuture[cbool] boolReturn(cRpcOptions, )
    cFollyFuture[cint16_t] i16Return(cRpcOptions, )
    cFollyFuture[cint32_t] i32Return(cRpcOptions, )
    cFollyFuture[cint64_t] i64Return(cRpcOptions, )
    cFollyFuture[float] floatReturn(cRpcOptions, )
    cFollyFuture[double] doubleReturn(cRpcOptions, )
    cFollyFuture[string] stringReturn(cRpcOptions, )
    cFollyFuture[string] binaryReturn(cRpcOptions, )
    cFollyFuture[cmap[string,cint64_t]] mapReturn(cRpcOptions, )
    cFollyFuture[cint32_t] simpleTypedefReturn(cRpcOptions, )
    cFollyFuture[vector[cmap[_module_cbindings.cEmpty,_module_cbindings.cMyStruct]]] complexTypedefReturn(cRpcOptions, )
    cFollyFuture[vector[vector[vector[cmap[_module_cbindings.cEmpty,_module_cbindings.cMyStruct]]]]] list_mostComplexTypedefReturn(cRpcOptions, )
    cFollyFuture[_module_cbindings.cMyEnumA] enumReturn(cRpcOptions, )
    cFollyFuture[vector[_module_cbindings.cMyEnumA]] list_EnumReturn(cRpcOptions, )
    cFollyFuture[_module_cbindings.cMyStruct] structReturn(cRpcOptions, )
    cFollyFuture[cset[_module_cbindings.cMyStruct]] set_StructReturn(cRpcOptions, )
    cFollyFuture[_module_cbindings.cComplexUnion] unionReturn(cRpcOptions, )
    cFollyFuture[vector[_module_cbindings.cComplexUnion]] list_UnionReturn(cRpcOptions, )
    cFollyFuture[_fbthrift_iobuf.cIOBuf] readDataEb(cRpcOptions, 
      cint64_t arg_size,)
    cFollyFuture[unique_ptr[_fbthrift_iobuf.cIOBuf]] readData(cRpcOptions, 
      cint64_t arg_size,)


  cdef cppclass cParamServiceClientWrapper "::some::valid::ns::ParamServiceClientWrapper":
    void setPersistentHeader(const string& key, const string& value)
    void addEventHandler(const shared_ptr[cTProcessorEventHandler]& handler)

    cFollyFuture[cFollyUnit] void_ret_i16_param(cRpcOptions, 
      cint16_t arg_param1,)
    cFollyFuture[cFollyUnit] void_ret_byte_i16_param(cRpcOptions, 
      cint8_t arg_param1,
      cint16_t arg_param2,)
    cFollyFuture[cFollyUnit] void_ret_map_param(cRpcOptions, 
      cmap[string,cint64_t] arg_param1,)
    cFollyFuture[cFollyUnit] void_ret_map_setlist_param(cRpcOptions, 
      cmap[string,cint64_t] arg_param1,
      cset[vector[string]] arg_param2,)
    cFollyFuture[cFollyUnit] void_ret_map_typedef_param(cRpcOptions, 
      cint32_t arg_param1,)
    cFollyFuture[cFollyUnit] void_ret_enum_param(cRpcOptions, 
      _module_cbindings.cMyEnumA arg_param1,)
    cFollyFuture[cFollyUnit] void_ret_struct_param(cRpcOptions, 
      _module_cbindings.cMyStruct arg_param1,)
    cFollyFuture[cFollyUnit] void_ret_listunion_param(cRpcOptions, 
      vector[_module_cbindings.cComplexUnion] arg_param1,)
    cFollyFuture[cbool] bool_ret_i32_i64_param(cRpcOptions, 
      cint32_t arg_param1,
      cint64_t arg_param2,)
    cFollyFuture[cbool] bool_ret_map_param(cRpcOptions, 
      cmap[string,cint64_t] arg_param1,)
    cFollyFuture[cbool] bool_ret_union_param(cRpcOptions, 
      _module_cbindings.cComplexUnion arg_param1,)
    cFollyFuture[cint64_t] i64_ret_float_double_param(cRpcOptions, 
      float arg_param1,
      double arg_param2,)
    cFollyFuture[cint64_t] i64_ret_string_typedef_param(cRpcOptions, 
      string arg_param1,
      cset[vector[vector[cmap[_module_cbindings.cEmpty,_module_cbindings.cMyStruct]]]] arg_param2,)
    cFollyFuture[cint64_t] i64_ret_i32_i32_i32_i32_i32_param(cRpcOptions, 
      cint32_t arg_param1,
      cint32_t arg_param2,
      cint32_t arg_param3,
      cint32_t arg_param4,
      cint32_t arg_param5,)
    cFollyFuture[double] double_ret_setstruct_param(cRpcOptions, 
      cset[_module_cbindings.cMyStruct] arg_param1,)
    cFollyFuture[string] string_ret_string_param(cRpcOptions, 
      string arg_param1,)
    cFollyFuture[string] binary_ret_binary_param(cRpcOptions, 
      string arg_param1,)
    cFollyFuture[cmap[string,cint64_t]] map_ret_bool_param(cRpcOptions, 
      cbool arg_param1,)
    cFollyFuture[vector[cbool]] list_ret_map_setlist_param(cRpcOptions, 
      cmap[cint32_t,vector[string]] arg_param1,
      vector[string] arg_param2,)
    cFollyFuture[cmap[cset[vector[cint32_t]],cmap[vector[cset[string]],string]]] mapsetlistmapliststring_ret_listlistlist_param(cRpcOptions, 
      vector[vector[vector[vector[cint32_t]]]] arg_param1,)
    cFollyFuture[cint32_t] typedef_ret_i32_param(cRpcOptions, 
      cint32_t arg_param1,)
    cFollyFuture[vector[cint32_t]] listtypedef_ret_typedef_param(cRpcOptions, 
      vector[cmap[_module_cbindings.cEmpty,_module_cbindings.cMyStruct]] arg_param1,)
    cFollyFuture[_module_cbindings.cMyEnumA] enum_ret_double_param(cRpcOptions, 
      double arg_param1,)
    cFollyFuture[_module_cbindings.cMyEnumA] enum_ret_double_enum_param(cRpcOptions, 
      double arg_param1,
      _module_cbindings.cMyEnumA arg_param2,)
    cFollyFuture[vector[_module_cbindings.cMyEnumA]] listenum_ret_map_param(cRpcOptions, 
      cmap[string,cint64_t] arg_param1,)
    cFollyFuture[_module_cbindings.cMyStruct] struct_ret_i16_param(cRpcOptions, 
      cint16_t arg_param1,)
    cFollyFuture[cset[_module_cbindings.cMyStruct]] setstruct_ret_set_param(cRpcOptions, 
      cset[string] arg_param1,)
    cFollyFuture[_module_cbindings.cComplexUnion] union_ret_i32_i32_param(cRpcOptions, 
      cint32_t arg_param1,
      cint32_t arg_param2,)
    cFollyFuture[vector[_module_cbindings.cComplexUnion]] listunion_string_param(cRpcOptions, 
      string arg_param1,)

