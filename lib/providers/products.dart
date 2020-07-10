import 'package:flutter/material.dart';

import './product.dart';
import '../services/product_service.dart';

class Products with ChangeNotifier {
  final ProductService productService;
  List<Product> _items = [];

  Products(this.productService);

  List<Product> get items {
    return [..._items];
  }

  List<Product> get favoriteItems {
    return _items.where((item) => item.isFavorite).toList();
  }

  Product findByID(String id) {
    return _items.firstWhere((product) => product.id == id);
  }

  Future<void> fetchAll() async {
    List<Product> products = await productService.getAll();
    _items = [...products];
    notifyListeners();
  }

  Future<void> add(Product product) async {
    final String id = await productService.add(product);
    final newProduct = product.clone(id: id);
    _items.insert(0, newProduct);
    notifyListeners();
  }

  void updateByID(String id, Product newProduct) {
    final int prodIndex = _items.indexWhere((product) => product.id == id);
    if (prodIndex < 0) return;
    _items[prodIndex] = newProduct;
    notifyListeners();
  }

  Future<void> deleteByID(String id) async {
    int deletedRows = await productService.deleteByID(id);
    if (deletedRows == 1) {
      _items.removeWhere((product) => product.id == id);
    }
    if (deletedRows > 1) {
      await fetchAll();
    }
    notifyListeners();
  }
}
