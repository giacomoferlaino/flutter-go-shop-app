import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart';
import 'package:shop_app/models/api_response.dart';

import '../models/order_item.dart';
import '../models/request_exception.dart';

class OrderService {
  static const String relativePath = '/order';
  final String _baseUrl;
  String _fullPath;

  OrderService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
  }

  OrderItem _parseOrder(dynamic item) {
    return OrderItem(
      id: item['id'],
      amount: item['amount'],
      cartItems: item['products'],
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
