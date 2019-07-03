# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from datapipe.recommendations import services_pb2 as datapipe_dot_recommendations_dot_services__pb2


class RecommendationsServiceStub(object):
  """*
  Service for providing data stored in the backend
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListPodRecommendations = channel.unary_unary(
        '/containersai.datapipe.recommendations.RecommendationsService/ListPodRecommendations',
        request_serializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsRequest.SerializeToString,
        response_deserializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsResponse.FromString,
        )
    self.ListAvailablePodRecommendations = channel.unary_unary(
        '/containersai.datapipe.recommendations.RecommendationsService/ListAvailablePodRecommendations',
        request_serializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsRequest.SerializeToString,
        response_deserializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsResponse.FromString,
        )
    self.ListControllerRecommendations = channel.unary_unary(
        '/containersai.datapipe.recommendations.RecommendationsService/ListControllerRecommendations',
        request_serializer=datapipe_dot_recommendations_dot_services__pb2.ListControllerRecommendationsRequest.SerializeToString,
        response_deserializer=datapipe_dot_recommendations_dot_services__pb2.ListControllerRecommendationsResponse.FromString,
        )


class RecommendationsServiceServicer(object):
  """*
  Service for providing data stored in the backend
  """

  def ListPodRecommendations(self, request, context):
    """Used to list pod recommendations
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListAvailablePodRecommendations(self, request, context):
    """Used to list available pod recommendations
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListControllerRecommendations(self, request, context):
    """Used to list controller recommendations
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RecommendationsServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListPodRecommendations': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodRecommendations,
          request_deserializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsRequest.FromString,
          response_serializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsResponse.SerializeToString,
      ),
      'ListAvailablePodRecommendations': grpc.unary_unary_rpc_method_handler(
          servicer.ListAvailablePodRecommendations,
          request_deserializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsRequest.FromString,
          response_serializer=datapipe_dot_recommendations_dot_services__pb2.ListPodRecommendationsResponse.SerializeToString,
      ),
      'ListControllerRecommendations': grpc.unary_unary_rpc_method_handler(
          servicer.ListControllerRecommendations,
          request_deserializer=datapipe_dot_recommendations_dot_services__pb2.ListControllerRecommendationsRequest.FromString,
          response_serializer=datapipe_dot_recommendations_dot_services__pb2.ListControllerRecommendationsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'containersai.datapipe.recommendations.RecommendationsService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))