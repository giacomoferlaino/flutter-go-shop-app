import 'dart:io';

import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;

import '../../models/api_response.dart';
import '../../models/request_exception.dart';

class HttpService {
  Future<ApiResponse<T>> request<T>({
    @required Future<http.Response> Function() request,
    @required T Function(dynamic item) dataParsing,
  }) async {
    try {
      http.Response response = await request();
      return ApiResponse<T>.parse(response, dataParsing);
    } catch (error) {
      throw _handleError(error);
    }
  }

  Exception _handleError(Exception exception) {
    if (exception is SocketException) {
      return RequestException('Unable to contact the server.');
    }
    return RequestException('An error occured');
  }
}
