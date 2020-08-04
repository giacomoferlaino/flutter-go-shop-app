import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';

import './product.dart';
import '../services/product_service.dart';
import '../services/user_service.dart';
import '../models/api_response.dart';

class Products with ChangeNotifier {
  final ProductService productService = GetIt.instance.get<ProductService>();
  final UserService userService = GetIt.instance.get<UserService>();
  List<Product> _items = [];
  List<Product> _ownedItems = [];
  List<Product> _favoriteItems = [];

  Products();

  List<Product> get items {
    return [..._items];
  }

  List<Product> get favoriteItems {
    return [..._favoriteItems];
  }

  List<Product> get ownedItems {
    return [..._ownedItems];
  }

  bool isFavorite(int productId) {
    for (Product product in _favoriteItems) {
      if (product.id == productId) return true;
    }
    return false;
  }

  Product findByID(int id) {
    return _items.firstWhere((product) => product.id == id);
  }

  Future<void> fetchAll() async {
    ApiResponse response = await productService.getAll();
    _items = [...response.data];
    response = await userService.getFavoriteProducts();
    _favoriteItems = [...response.data];
    response = await userService.getProducts();
    _ownedItems = [...response.data];
    notifyListeners();
  }

  Future<void> add(Product product) async {
    await userService.addProduct(product);
    fetchAll();
  }

  Future<void> updateByID(int id, Product product) async {
    await productService.updateByID(id, product);
    await fetchAll();
  }

  Future<void> deleteByID(int id) async {
    await productService.deleteByID(id);
    await fetchAll();
  }

  Future<void> addFavorite(Product product) async {
    await userService.addFavoriteProduct(product.id);
    fetchAll();
  }

  Future<void> removeFavorite(Product product) async {
    await userService.removeFavoriteProduct(product.id);
    fetchAll();
  }
}
