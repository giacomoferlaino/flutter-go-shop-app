import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart';

import '../models/order_item.dart';
import '../models/request_exception.dart';
import '../models/api_response.dart';
import '../models/cart_item.dart';
import '../providers/product.dart';

class OrderService {
  static const String relativePath = '/order';
  final String _baseUrl;
  String _fullPath;

  OrderService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
  }

  OrderItem _parseOrder(dynamic item) {
    final List<CartItem> cartItems = [];
    item['cartItems'].forEach((item) {
      Product product = Product(
        id: item['product']['id'],
        description: item['product']['description'],
        imageUrl: item['product']['imageUrl'],
        price: item['product']['price'],
        isFavorite: item['product']['isFavorite'],
        title: item['product']['title'],
      );
      return cartItems.add(CartItem(
        product: product,
        quantity: item['quantity'],
      ));
    });
    return OrderItem(
      id: item['id'],
      amount: item['amount'],
      cartItems: cartItems,
      dateTime: DateTime.parse(item['dateTime']),
    );
  }

  Future<ApiResponse<OrderItem>> add(OrderItem order) async {
    Response response = await post(_fullPath, body: json.encode(order));
    return ApiResponse<OrderItem>.parse(response, _parseOrder);
  }

  Future<ApiResponse<OrderItem>> getAll() async {
    try {
      Response response = await get(_fullPath);
      return ApiResponse<OrderItem>.parse(response, _parseOrder);
    } catch (error) {
      throw getException(error);
    }
  }

  Future<ApiResponse<OrderItem>> deleteByID(int id) async {
    Response response = await delete(_fullPath + '/$id');
    return ApiResponse<OrderItem>.parse(response, _parseOrder);
  }

  Future<ApiResponse<OrderItem>> updateByID(int id, OrderItem product) async {
    Response response =
        await put(_fullPath + '/$id', body: json.encode(product));
    return ApiResponse<OrderItem>.parse(response, _parseOrder);
  }

  Exception getException(Exception exception) {
    if (exception is SocketException) {
      return RequestException('Internet connection error!');
    }
    return RequestException('An error occured');
  }
}
