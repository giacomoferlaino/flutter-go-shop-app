import 'package:flutter/foundation.dart';

import '../models/cart_item.dart';
import '../providers/product.dart';

class Cart with ChangeNotifier {
  Map<String, CartItem> _items = {};

  Map<String, CartItem> get items {
    return {..._items};
  }

  int get totalQuantity {
    int count = 0;
    _items.forEach((_, cartItem) => count += cartItem.quantity);
    return count;
  }

  int get itemCount {
    return _items.length;
  }

  double get totalAmount {
    double amount = 0;
    _items.forEach(
        (_, cartItem) => amount += (cartItem.price * cartItem.quantity));
    return amount;
  }

  void addItem(Product product) {
    if (_items.containsKey(product.id)) {
      _items.update(
          product.id.toString(),
          (existingCartItem) => CartItem(
                id: existingCartItem.id,
                title: existingCartItem.title,
                price: existingCartItem.price,
                quantity: existingCartItem.quantity + 1,
              ));
    } else {
      _items.putIfAbsent(
          product.id.toString(),
          () => CartItem(
                id: DateTime.now().toString(),
                title: product.title,
                price: product.price,
              ));
    }
    notifyListeners();
  }

  void removeItem(String productId) {
    _items.remove(productId);
    notifyListeners();
  }

  void removeSingleItem(String productId) {
    if (!_items.containsKey(productId)) return;
    if (_items[productId].quantity > 1) {
      _items.update(
          productId,
          (existingCartItem) => CartItem(
              id: existingCartItem.id,
              price: existingCartItem.price,
              title: existingCartItem.title,
              quantity: existingCartItem.quantity - 1));
      notifyListeners();
      return;
    }
    removeItem(productId);
  }

  void clear() {
    _items = {};
    notifyListeners();
  }
}
