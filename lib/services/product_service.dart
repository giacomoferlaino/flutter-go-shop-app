import 'dart:convert';

import 'package:http/http.dart';

import '../models/api_response.dart';
import '../models/url.dart';
import '../providers/product.dart';
import './auth_service.dart';
import './http_service.dart';

class ProductService {
  final HttpService httpService;
  final AuthService authService;
  static const List<String> relativePath = ['product'];
  final Url _baseUrl;

  ProductService(this.httpService, this.authService, this._baseUrl);

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

  Future<ApiResponse<Product>> add(
    Product product,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: relativePath,
    );
    return httpService.request<Product>(
      request: () => post(
        uri,
        body: json.encode(product),
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> getAll() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: relativePath,
    );
    return httpService.request<Product>(
      request: () => get(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> deleteByID(
    int id,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: relativePath,
      queryParameters: {'id': id},
    );
    return httpService.request<Product>(
      request: () => delete(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> updateByID(
    int id,
    Product product,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: relativePath,
      queryParameters: {'id': id.toString()},
    );
    return httpService.request<Product>(
      request: () => put(
        uri,
        body: json.encode(product),
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }
}
