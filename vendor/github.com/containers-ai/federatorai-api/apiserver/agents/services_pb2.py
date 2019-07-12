# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: apiserver/agents/services.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='apiserver/agents/services.proto',
  package='containersai.apiserver.agents',
  syntax='proto3',
  serialized_options=_b('Z9github.com/containers-ai/federatorai-api/apiserver/agents'),
  serialized_pb=_b('\n\x1f\x61piserver/agents/services.proto\x12\x1d\x63ontainersai.apiserver.agents\x1a\x17google/rpc/status.proto\"\x11\n\x0fRegisterRequest\"\x12\n\x10RegisterResponse\"\x18\n\x16GetAgentVersionRequest\"\x19\n\x17GetAgentVersionResponse\"\x16\n\x14\x44ownloadAgentRequest2\xdf\x02\n\rAgentsService\x12m\n\x08Register\x12..containersai.apiserver.agents.RegisterRequest\x1a/.containersai.apiserver.agents.RegisterResponse\"\x00\x12\x82\x01\n\x0fGetAgentVersion\x12\x35.containersai.apiserver.agents.GetAgentVersionRequest\x1a\x36.containersai.apiserver.agents.GetAgentVersionResponse\"\x00\x12Z\n\rDownloadAgent\x12\x33.containersai.apiserver.agents.DownloadAgentRequest\x1a\x12.google.rpc.Status\"\x00\x42;Z9github.com/containers-ai/federatorai-api/apiserver/agentsb\x06proto3')
  ,
  dependencies=[google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_REGISTERREQUEST = _descriptor.Descriptor(
  name='RegisterRequest',
  full_name='containersai.apiserver.agents.RegisterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=91,
  serialized_end=108,
)


_REGISTERRESPONSE = _descriptor.Descriptor(
  name='RegisterResponse',
  full_name='containersai.apiserver.agents.RegisterResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=110,
  serialized_end=128,
)


_GETAGENTVERSIONREQUEST = _descriptor.Descriptor(
  name='GetAgentVersionRequest',
  full_name='containersai.apiserver.agents.GetAgentVersionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=130,
  serialized_end=154,
)


_GETAGENTVERSIONRESPONSE = _descriptor.Descriptor(
  name='GetAgentVersionResponse',
  full_name='containersai.apiserver.agents.GetAgentVersionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=156,
  serialized_end=181,
)


_DOWNLOADAGENTREQUEST = _descriptor.Descriptor(
  name='DownloadAgentRequest',
  full_name='containersai.apiserver.agents.DownloadAgentRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=183,
  serialized_end=205,
)

DESCRIPTOR.message_types_by_name['RegisterRequest'] = _REGISTERREQUEST
DESCRIPTOR.message_types_by_name['RegisterResponse'] = _REGISTERRESPONSE
DESCRIPTOR.message_types_by_name['GetAgentVersionRequest'] = _GETAGENTVERSIONREQUEST
DESCRIPTOR.message_types_by_name['GetAgentVersionResponse'] = _GETAGENTVERSIONRESPONSE
DESCRIPTOR.message_types_by_name['DownloadAgentRequest'] = _DOWNLOADAGENTREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RegisterRequest = _reflection.GeneratedProtocolMessageType('RegisterRequest', (_message.Message,), dict(
  DESCRIPTOR = _REGISTERREQUEST,
  __module__ = 'apiserver.agents.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.apiserver.agents.RegisterRequest)
  ))
_sym_db.RegisterMessage(RegisterRequest)

RegisterResponse = _reflection.GeneratedProtocolMessageType('RegisterResponse', (_message.Message,), dict(
  DESCRIPTOR = _REGISTERRESPONSE,
  __module__ = 'apiserver.agents.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.apiserver.agents.RegisterResponse)
  ))
_sym_db.RegisterMessage(RegisterResponse)

GetAgentVersionRequest = _reflection.GeneratedProtocolMessageType('GetAgentVersionRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETAGENTVERSIONREQUEST,
  __module__ = 'apiserver.agents.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.apiserver.agents.GetAgentVersionRequest)
  ))
_sym_db.RegisterMessage(GetAgentVersionRequest)

GetAgentVersionResponse = _reflection.GeneratedProtocolMessageType('GetAgentVersionResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETAGENTVERSIONRESPONSE,
  __module__ = 'apiserver.agents.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.apiserver.agents.GetAgentVersionResponse)
  ))
_sym_db.RegisterMessage(GetAgentVersionResponse)

DownloadAgentRequest = _reflection.GeneratedProtocolMessageType('DownloadAgentRequest', (_message.Message,), dict(
  DESCRIPTOR = _DOWNLOADAGENTREQUEST,
  __module__ = 'apiserver.agents.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.apiserver.agents.DownloadAgentRequest)
  ))
_sym_db.RegisterMessage(DownloadAgentRequest)


DESCRIPTOR._options = None

_AGENTSSERVICE = _descriptor.ServiceDescriptor(
  name='AgentsService',
  full_name='containersai.apiserver.agents.AgentsService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=208,
  serialized_end=559,
  methods=[
  _descriptor.MethodDescriptor(
    name='Register',
    full_name='containersai.apiserver.agents.AgentsService.Register',
    index=0,
    containing_service=None,
    input_type=_REGISTERREQUEST,
    output_type=_REGISTERRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetAgentVersion',
    full_name='containersai.apiserver.agents.AgentsService.GetAgentVersion',
    index=1,
    containing_service=None,
    input_type=_GETAGENTVERSIONREQUEST,
    output_type=_GETAGENTVERSIONRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DownloadAgent',
    full_name='containersai.apiserver.agents.AgentsService.DownloadAgent',
    index=2,
    containing_service=None,
    input_type=_DOWNLOADAGENTREQUEST,
    output_type=google_dot_rpc_dot_status__pb2._STATUS,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_AGENTSSERVICE)

DESCRIPTOR.services_by_name['AgentsService'] = _AGENTSSERVICE

# @@protoc_insertion_point(module_scope)