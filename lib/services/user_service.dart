import 'dart:convert';

import 'package:http/http.dart';

import '../models/url.dart';
import '../models/api_response.dart';
import '../models/order_item.dart';
import '../providers/product.dart';
import '../services/auth_service.dart';
import '../services/http_service.dart';
import '../services/product_service.dart';
import '../services/order_service.dart';

class UserService {
  final HttpService httpService;
  final AuthService authService;
  static const List<String> _userProductsPath = ['user', 'product'];
  static const List<String> _userFavoriteProductsPath = [
    'user',
    'product',
    'favorite',
  ];
  static const List<String> _userOrdersPath = ['user', 'order'];
  final Url _baseUrl;

  UserService(
    this.httpService,
    this.authService,
    this._baseUrl,
  );

  Future<ApiResponse<Product>> addProduct(Product product) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userProductsPath,
    );
    return httpService.request(
      request: () => post(
        uri,
        body: json.encode(product),
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: ProductService.parseProduct,
    );
  }

  Future<ApiResponse<Product>> getProducts() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userProductsPath,
    );
    return httpService.request<Product>(
      request: () => get(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: ProductService.parseProduct,
    );
  }

  Future<ApiResponse<Product>> getFavoriteProducts() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userFavoriteProductsPath,
    );
    return httpService.request<Product>(
      request: () => get(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: ProductService.parseProduct,
    );
  }

  Future<ApiResponse<Product>> addFavoriteProduct(
    int id,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userFavoriteProductsPath,
      queryParameters: {'id': id.toString()},
    );
    return httpService.request<Product>(
      request: () => post(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: (dynamic item) => null,
    );
  }

  Future<ApiResponse<Product>> removeFavoriteProduct(
    int id,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userFavoriteProductsPath,
      queryParameters: {'id': id.toString()},
    );
    return httpService.request<Product>(
      request: () => delete(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: (dynamic item) => null,
    );
  }

  Future<ApiResponse<OrderItem>> getOrders() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userOrdersPath,
    );
    return httpService.request<OrderItem>(
      request: () => get(
        uri,
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: OrderService.parseOrder,
    );
  }

  Future<ApiResponse<OrderItem>> addOrder(OrderItem order) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _userOrdersPath,
    );
    return httpService.request(
      request: () => post(
        uri,
        body: json.encode(order),
        headers: {'Authorization': this.authService.token},
      ),
      dataParsing: OrderService.parseOrder,
    );
  }
}
