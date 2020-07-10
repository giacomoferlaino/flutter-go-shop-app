import 'dart:convert';

import 'package:http/http.dart';

import '../providers/product.dart';

class ProductService {
  static const String relativePath = '/product';
  final String _baseUrl;
  String _fullPath;

  ProductService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
  }

  Future<String> add(Product product) async {
    Response response =
        await post(_fullPath, body: json.encode(product.toMap()));
    final String id = json.decode(response.body)['id'];
    return id;
  }
}
