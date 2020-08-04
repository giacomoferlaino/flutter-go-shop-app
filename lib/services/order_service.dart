import 'dart:convert';

import 'package:http/http.dart';

import '../models/order_item.dart';
import '../models/api_response.dart';
import '../models/cart_item.dart';
import '../models/url.dart';
import '../providers/product.dart';
import './auth_service.dart';
import './http_service.dart';

class OrderService {
  final HttpService _httpService;
  final AuthService _authService;
  static const List<String> _orderPath = ['order'];
  final Url _baseUrl;

  OrderService(this._httpService, this._authService, this._baseUrl);

  static OrderItem parseOrder(dynamic item) {
    final List<CartItem> cartItems = [];
    if (item['cartItems'] != null) {
      item['cartItems'].forEach((item) {
        Product product = Product(
          id: item['product']['id'],
          description: item['product']['description'],
          imageUrl: item['product']['imageUrl'],
          price: item['product']['price'],
          title: item['product']['title'],
        );
        return cartItems.add(CartItem(
          product: product,
          quantity: item['quantity'],
        ));
      });
    }
    return OrderItem(
      id: item['id'],
      amount: item['amount'],
      cartItems: cartItems,
      dateTime: DateTime.parse(item['dateTime']),
    );
  }

  Future<ApiResponse<OrderItem>> add(
    OrderItem order,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _orderPath,
    );
    return _httpService.request(
      request: () => post(
        uri,
        body: json.encode(order),
        headers: {'Authorization': this._authService.token},
      ),
      dataParsing: parseOrder,
    );
  }

  Future<ApiResponse<OrderItem>> getAll() async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _orderPath,
    );
    return _httpService.request(
      request: () => get(
        uri,
        headers: {'Authorization': this._authService.token},
      ),
      dataParsing: parseOrder,
    );
  }

  Future<ApiResponse<OrderItem>> deleteByID(
    int id,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _orderPath,
      queryParameters: {'id': id},
    );
    return _httpService.request(
        request: () => delete(
              uri,
              headers: {'Authorization': this._authService.token},
            ),
        dataParsing: parseOrder);
  }

  Future<ApiResponse<OrderItem>> updateByID(
    int id,
    OrderItem product,
  ) async {
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      pathSegments: _orderPath,
      queryParameters: {'id': id},
    );
    return _httpService.request(
      request: () => put(
        uri,
        body: json.encode(product),
        headers: {'Authorization': this._authService.token},
      ),
      dataParsing: parseOrder,
    );
  }
}
