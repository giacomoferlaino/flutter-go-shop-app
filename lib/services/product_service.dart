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
  static const List<String> userProductsPath = ['user', 'product'];
  static const List<String> userFavoritesPath = ['user', 'product', 'favorite'];
  static const List<String> globalProductsPath = ['product'];
  final Url _baseUrl;

  ProductService(this.httpService, this.authService, this._baseUrl);

  Product _parseProduct(dynamic item) {
    return Product(
      id: item['id'],
      description: item['description'],
      imageUrl: item['imageUrl'],
      price: item['price'],
      title: item['title'],
    );
  }

  Future<ApiResponse<Product>> add(
    Product product,
  ) async {
    Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: globalProductsPath,
    );
    ApiResponse apiResponse = await httpService.request<Product>(
      request: () => post(
        uri,
        body: json.encode(product),
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
    Product newProduct = apiResponse.data[0];
    uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: userProductsPath,
      queryParameters: {'id': newProduct.id.toString()},
    );
    await httpService.request<Product>(
      request: () => post(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
    return apiResponse;
  }

  Future<ApiResponse<Product>> getAll() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: userProductsPath,
    );
    return httpService.request<Product>(
      request: () => get(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> getFavorites() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: userFavoritesPath,
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
      pathSegments: userProductsPath,
      queryParameters: {'id': id.toString()},
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
      pathSegments: globalProductsPath,
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

  Future<ApiResponse<Product>> addFavorite(
    int id,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: userFavoritesPath,
      queryParameters: {'id': id.toString()},
    );
    return httpService.request<Product>(
      request: () => post(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }

  Future<ApiResponse<Product>> removeFavorite(
    int id,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: userFavoritesPath,
      queryParameters: {'id': id.toString()},
    );
    return httpService.request<Product>(
      request: () => delete(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: _parseProduct,
    );
  }
}
