import 'dart:io';

import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;

import '../models/api_response.dart';
import '../models/request_exception.dart';

class HttpService {
  Future<ApiResponse<T>> request<T>({
    @required Future<http.Response> Function() request,
    @required T Function(dynamic item) dataParsing,
  }) async {
    try {
      http.Response response = await request();
      ApiResponse<T> apiResponse = ApiResponse<T>.parse(response, dataParsing);
      String error = apiResponse.meta.error;
      if (error != null && error != '') {
        throw RequestException(error);
      }
      return apiResponse;
    } catch (error) {
      throw _handleError(error);
    }
  }

  Exception _handleError(dynamic exception) {
    if (exception is RequestException) {
      return exception;
    }
    if (exception is SocketException) {
      return RequestException('Unable to contact the server.');
    }
    return RequestException('An error occured');
  }
}
