import 'dart:convert';

import 'package:get_it/get_it.dart';
import 'package:http/http.dart';
import 'package:shop_app/models/api_response.dart';

import '../providers/product.dart';
import 'http_service.dart';

class ProductService {
  final HttpService httpService = GetIt.instance.get<HttpService>();
  static const String relativePath = '/product';
  final String _baseUrl;
  String _fullPath;

  ProductService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
  }

  Product _parseProduct(dynamic item) {
    return Product(
      id: item['id'],
      description: item['description'],
      imageUrl: item['imageUrl'],
      price: item['price'],
      title: item['title'],
      isFavorite: item['isFavorite'],
    );
  }

  Future<ApiResponse<Product>> add(Product product) async {
    return httpService.request<Product>(
      request: () => post(_fullPath, body: json.encode(product)),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> getAll() async {
    return httpService.request<Product>(
      request: () => get(_fullPath),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> deleteByID(int id) async {
    return httpService.request<Product>(
      request: () => delete(_fullPath + '/$id'),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> updateByID(int id, Product product) async {
    return httpService.request<Product>(
      request: () => put(_fullPath + '/$id', body: json.encode(product)),
      dataParsing: _parseProduct,
    );
  }
}
