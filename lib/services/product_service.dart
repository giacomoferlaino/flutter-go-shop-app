import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart';
import 'package:shop_app/models/api_response.dart';

import '../providers/product.dart';
import '../models/request_exception.dart';

class ProductService {
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
    Response response = await post(_fullPath, body: json.encode(product));
    return ApiResponse<Product>.parse(response, _parseProduct);
  }

  Future<ApiResponse<Product>> getAll() async {
    try {
      Response response = await get(_fullPath);
      return ApiResponse<Product>.parse(response, _parseProduct);
    } catch (error) {
      throw handleError(error);
    }
  }

  Future<ApiResponse<Product>> deleteByID(int id) async {
    Response response = await delete(_fullPath + '/$id');
    return ApiResponse<Product>.parse(response, _parseProduct);
  }

  Future<ApiResponse<Product>> updateByID(int id, Product product) async {
    Response response =
        await put(_fullPath + '/$id', body: json.encode(product));
    return ApiResponse<Product>.parse(response, _parseProduct);
  }

  Exception handleError(Error exception) {
    if (exception is SocketException) {
      return RequestException('Internet connection error!');
    }
    return RequestException('An error occured');
  }
}
