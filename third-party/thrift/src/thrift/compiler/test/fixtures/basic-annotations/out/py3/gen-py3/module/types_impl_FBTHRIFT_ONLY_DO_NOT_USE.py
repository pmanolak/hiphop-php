#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/basic-annotations/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum
import thrift.py3.types
import module.thrift_enums as _fbthrift_python_enums

_fbthrift__module_name__ = "module.types"



class MyEnum(thrift.py3.types.CompiledEnum, int):
    MyValue1 = 0
    MyValue2 = 1
    DOMAIN = 2

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_MyEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.MyEnum"

    def _to_python(self):
        return _fbthrift_python_enums.MyEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_


    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_

