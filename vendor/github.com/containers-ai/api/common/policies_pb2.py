# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/policies.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='common/policies.proto',
  package='containersai.common',
  syntax='proto3',
  serialized_options=_b('Z#github.com/containers-ai/api/common'),
  serialized_pb=_b('\n\x15\x63ommon/policies.proto\x12\x13\x63ontainersai.common*T\n\x14RecommendationPolicy\x12#\n\x1fRECOMMENDATION_POLICY_UNDEFINED\x10\x00\x12\n\n\x06STABLE\x10\x01\x12\x0b\n\x07\x43OMPACT\x10\x02\x42%Z#github.com/containers-ai/api/commonb\x06proto3')
)

_RECOMMENDATIONPOLICY = _descriptor.EnumDescriptor(
  name='RecommendationPolicy',
  full_name='containersai.common.RecommendationPolicy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='RECOMMENDATION_POLICY_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STABLE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMPACT', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=46,
  serialized_end=130,
)
_sym_db.RegisterEnumDescriptor(_RECOMMENDATIONPOLICY)

RecommendationPolicy = enum_type_wrapper.EnumTypeWrapper(_RECOMMENDATIONPOLICY)
RECOMMENDATION_POLICY_UNDEFINED = 0
STABLE = 1
COMPACT = 2


DESCRIPTOR.enum_types_by_name['RecommendationPolicy'] = _RECOMMENDATIONPOLICY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)