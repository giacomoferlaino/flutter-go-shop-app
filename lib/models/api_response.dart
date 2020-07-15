import 'dart:convert';

import 'package:http/http.dart';

class _MetaData {
  final int rows;

  _MetaData(this.rows);
}

class ApiResponse<T> {
  _MetaData meta;
  List<T> data;

  ApiResponse({this.meta, this.data});

  ApiResponse.parse(
    Response response,
    T Function(dynamic item) parseObject,
  ) {
    final decodedResponse = json.decode(response.body);
    meta = _MetaData(decodedResponse['meta']['rows']);
    if (decodedResponse['data'] == null) {
      data = null;
    } else {
      data = [];
      decodedResponse['data'].forEach((item) => {data.add(parseObject(item))});
    }
  }
}
