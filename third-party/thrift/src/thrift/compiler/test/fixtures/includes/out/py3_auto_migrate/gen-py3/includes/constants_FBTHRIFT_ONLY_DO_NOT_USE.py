#
# Autogenerated by Thrift for includes.thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
import thrift.py3.types

import includes.types
import transitive.types


ExampleIncluded = includes.types.Included(MyIntField=2, MyTransitiveField=transitive.types.Foo(a=2))
IncludedConstant = 42
