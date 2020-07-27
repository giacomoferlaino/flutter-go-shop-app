import 'dart:convert';

import 'package:http/http.dart';

class MetaData {
  final int rows;
  final String error;

  MetaData(this.rows, this.error);
}

class ApiResponse<T> {
  MetaData meta;
  List<T> data;

  ApiResponse({this.meta, this.data});

  ApiResponse.parse(
    Response response,
    T Function(dynamic item) parseObject,
  ) {
    final decodedResponse = json.decode(response.body);
    meta = MetaData(
      decodedResponse['meta']['rows'],
      decodedResponse['meta']['error'],
    );
    if (decodedResponse['data'] == null) {
      data = null;
    } else {
      data = [];
      decodedResponse['data'].forEach((item) => {data.add(parseObject(item))});
    }
  }
}
