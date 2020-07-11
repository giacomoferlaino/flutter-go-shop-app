import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart';

import '../providers/product.dart';
import '../models/request_exception.dart';

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

  Future<List<Product>> getAll() async {
    try {
      Response response = await get(_fullPath);
      List<dynamic> parsedResponseBody = json.decode(response.body);
      List<Product> products = [];
      parsedResponseBody.forEach((item) => products.add(
            Product(
              id: item['id'],
              description: item['description'],
              imageUrl: item['imageUrl'],
              price: item['price'],
              title: item['title'],
              isFavorite: item['isFavorite'],
            ),
          ));
      return products;
    } catch (error) {
      throw getException(error);
    }
  }

  Future<int> deleteByID(String id) async {
    Response response = await delete(_fullPath + '/$id');
    List<dynamic> parsedBody = json.decode(response.body);
    return parsedBody[0]['deletedRows'];
  }

  Future<int> updateByID(String id, Product product) async {
    Response response =
        await put(_fullPath + '/$id', body: json.encode(product.toMap()));
    List<dynamic> parsedResponseBody = json.decode(response.body);
    return parsedResponseBody[0]['updatedRows'];
  }

  Exception getException(Exception exception) {
    if (exception is SocketException) {
      return RequestException('Internet connection error!');
    }
    return RequestException('An error occured');
  }
}
