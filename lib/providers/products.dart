import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';

import './product.dart';
import '../services/product_service.dart';
import '../models/api_response.dart';

class Products with ChangeNotifier {
  final ProductService productService = GetIt.instance.get<ProductService>();
  List<Product> _items = [];

  Products();

  List<Product> get items {
    return [..._items];
  }

  List<Product> get favoriteItems {
    return _items.where((item) => item.isFavorite).toList();
  }

  Product findByID(int id) {
    return _items.firstWhere((product) => product.id == id);
  }

  Future<void> fetchAll() async {
    ApiResponse response = await productService.getAll();
    _items = [...response.data];
    notifyListeners();
  }

  Future<void> add(Product product) async {
    ApiResponse response = await productService.add(product);
    Product createdProduct = response.data[0];
    final newProduct = product.clone(id: createdProduct.id);
    _items.insert(0, newProduct);
    notifyListeners();
  }

  Future<void> updateByID(int id, Product product) async {
    ApiResponse response = await productService.updateByID(id, product);
    if (response.meta.rows == 1) {
      final int prodIndex = _items.indexWhere((product) => product.id == id);
      if (prodIndex < 0) return;
      _items[prodIndex] = product;
      notifyListeners();
    }
    if (response.meta.rows > 1) {
      await fetchAll();
    }
  }

  Future<void> deleteByID(int id) async {
    ApiResponse response = await productService.deleteByID(id);
    if (response.meta.rows == 1) {
      _items.removeWhere((product) => product.id == id);
      notifyListeners();
    }
    if (response.meta.rows > 1) {
      await fetchAll();
    }
  }
}
