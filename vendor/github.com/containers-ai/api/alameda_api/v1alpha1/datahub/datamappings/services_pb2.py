# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/datamappings/services.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import metrics_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2
from alameda_api.v1alpha1.datahub.datamappings import datamapping_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_datamappings_dot_datamapping__pb2
from alameda_api.v1alpha1.datahub.schemas import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/datamappings/services.proto',
  package='containersai.alameda.v1alpha1.datahub.datamappings',
  syntax='proto3',
  serialized_options=b'ZFgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings',
  serialized_pb=b'\n8alameda_api/v1alpha1/datahub/datamappings/services.proto\x12\x32\x63ontainersai.alameda.v1alpha1.datahub.datamappings\x1a\x31\x61lameda_api/v1alpha1/datahub/common/metrics.proto\x1a;alameda_api/v1alpha1/datahub/datamappings/datamapping.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/schemas/types.proto\x1a\x17google/rpc/status.proto\"s\n\x19\x43reateDataMappingsRequest\x12V\n\rdata_mappings\x18\x01 \x03(\x0b\x32?.containersai.alameda.v1alpha1.datahub.datamappings.DataMapping\"\xb8\x01\n\x17ReadDataMappingsRequest\x12N\n\x0bschema_meta\x18\x01 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta\x12M\n\x0bmetric_type\x18\x02 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.MetricType\"\x96\x01\n\x18ReadDataMappingsResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12V\n\rdata_mappings\x18\x02 \x03(\x0b\x32?.containersai.alameda.v1alpha1.datahub.datamappings.DataMapping\"\xba\x01\n\x19\x44\x65leteDataMappingsRequest\x12N\n\x0bschema_meta\x18\x01 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta\x12M\n\x0bmetric_type\x18\x02 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.MetricTypeBHZFgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappingsb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_datamappings_dot_datamapping__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_CREATEDATAMAPPINGSREQUEST = _descriptor.Descriptor(
  name='CreateDataMappingsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.datamappings.CreateDataMappingsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data_mappings', full_name='containersai.alameda.v1alpha1.datahub.datamappings.CreateDataMappingsRequest.data_mappings', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=299,
  serialized_end=414,
)


_READDATAMAPPINGSREQUEST = _descriptor.Descriptor(
  name='ReadDataMappingsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsRequest.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsRequest.metric_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=417,
  serialized_end=601,
)


_READDATAMAPPINGSRESPONSE = _descriptor.Descriptor(
  name='ReadDataMappingsResponse',
  full_name='containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_mappings', full_name='containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsResponse.data_mappings', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=604,
  serialized_end=754,
)


_DELETEDATAMAPPINGSREQUEST = _descriptor.Descriptor(
  name='DeleteDataMappingsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.datamappings.DeleteDataMappingsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='containersai.alameda.v1alpha1.datahub.datamappings.DeleteDataMappingsRequest.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='containersai.alameda.v1alpha1.datahub.datamappings.DeleteDataMappingsRequest.metric_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=757,
  serialized_end=943,
)

_CREATEDATAMAPPINGSREQUEST.fields_by_name['data_mappings'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_datamappings_dot_datamapping__pb2._DATAMAPPING
_READDATAMAPPINGSREQUEST.fields_by_name['schema_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_READDATAMAPPINGSREQUEST.fields_by_name['metric_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2._METRICTYPE
_READDATAMAPPINGSRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_READDATAMAPPINGSRESPONSE.fields_by_name['data_mappings'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_datamappings_dot_datamapping__pb2._DATAMAPPING
_DELETEDATAMAPPINGSREQUEST.fields_by_name['schema_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_DELETEDATAMAPPINGSREQUEST.fields_by_name['metric_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2._METRICTYPE
DESCRIPTOR.message_types_by_name['CreateDataMappingsRequest'] = _CREATEDATAMAPPINGSREQUEST
DESCRIPTOR.message_types_by_name['ReadDataMappingsRequest'] = _READDATAMAPPINGSREQUEST
DESCRIPTOR.message_types_by_name['ReadDataMappingsResponse'] = _READDATAMAPPINGSRESPONSE
DESCRIPTOR.message_types_by_name['DeleteDataMappingsRequest'] = _DELETEDATAMAPPINGSREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateDataMappingsRequest = _reflection.GeneratedProtocolMessageType('CreateDataMappingsRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDATAMAPPINGSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.datamappings.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.datamappings.CreateDataMappingsRequest)
  })
_sym_db.RegisterMessage(CreateDataMappingsRequest)

ReadDataMappingsRequest = _reflection.GeneratedProtocolMessageType('ReadDataMappingsRequest', (_message.Message,), {
  'DESCRIPTOR' : _READDATAMAPPINGSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.datamappings.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsRequest)
  })
_sym_db.RegisterMessage(ReadDataMappingsRequest)

ReadDataMappingsResponse = _reflection.GeneratedProtocolMessageType('ReadDataMappingsResponse', (_message.Message,), {
  'DESCRIPTOR' : _READDATAMAPPINGSRESPONSE,
  '__module__' : 'alameda_api.v1alpha1.datahub.datamappings.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.datamappings.ReadDataMappingsResponse)
  })
_sym_db.RegisterMessage(ReadDataMappingsResponse)

DeleteDataMappingsRequest = _reflection.GeneratedProtocolMessageType('DeleteDataMappingsRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDATAMAPPINGSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.datamappings.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.datamappings.DeleteDataMappingsRequest)
  })
_sym_db.RegisterMessage(DeleteDataMappingsRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
