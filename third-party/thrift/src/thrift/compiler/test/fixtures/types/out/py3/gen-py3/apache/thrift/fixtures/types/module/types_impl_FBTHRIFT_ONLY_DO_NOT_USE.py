#
# Autogenerated by Thrift for thrift/compiler/test/fixtures/types/src/module.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import enum
import thrift.py3.types
import apache.thrift.fixtures.types.module.thrift_enums as _fbthrift_python_enums

_fbthrift__module_name__ = "apache.thrift.fixtures.types.module.types"



class has_bitwise_ops(thrift.py3.types.CompiledEnum):
    none = 0
    zero = 1
    one = 2
    two = 4
    three = 8

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_has_bitwise_ops()

    @staticmethod
    def __get_thrift_name__():
        return "module.has_bitwise_ops"

    def _to_python(self):
        return _fbthrift_python_enums.has_bitwise_ops(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __lt__(self, other):
        if isinstance(other, has_bitwise_ops):
            return self._fbthrift_value_ < other._fbthrift_value_

        raise NotImplementedError(
            "'<' only implemented for comparisons with has_bitwise_ops"
        )

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_

class is_unscoped(thrift.py3.types.CompiledEnum):
    hello = 0
    world = 1

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_is_unscoped()

    @staticmethod
    def __get_thrift_name__():
        return "module.is_unscoped"

    def _to_python(self):
        return _fbthrift_python_enums.is_unscoped(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __lt__(self, other):
        if isinstance(other, is_unscoped):
            return self._fbthrift_value_ < other._fbthrift_value_

        raise NotImplementedError(
            "'<' only implemented for comparisons with is_unscoped"
        )

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_

class MyForwardRefEnum(thrift.py3.types.CompiledEnum):
    ZERO = 0
    NONZERO = 12

    __module__ = _fbthrift__module_name__
    __slots__ = ()

    @staticmethod
    def __get_metadata__():
        return _fbthrift_python_enums.gen_metadata_enum_MyForwardRefEnum()

    @staticmethod
    def __get_thrift_name__():
        return "module.MyForwardRefEnum"

    def _to_python(self):
        return _fbthrift_python_enums.MyForwardRefEnum(self._fbthrift_value_)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self._fbthrift_value_

    def __lt__(self, other):
        if isinstance(other, MyForwardRefEnum):
            return self._fbthrift_value_ < other._fbthrift_value_

        raise NotImplementedError(
            "'<' only implemented for comparisons with MyForwardRefEnum"
        )

    def __int__(self):
        return self._fbthrift_value_

    def __index__(self):
        return self._fbthrift_value_

