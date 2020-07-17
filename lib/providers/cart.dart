import 'package:flutter/foundation.dart';

import '../models/cart_item.dart';
import '../providers/product.dart';

class Cart with ChangeNotifier {
  Map<int, CartItem> _items = {};

  Map<int, CartItem> get items {
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
    _items.forEach((_, cartItem) =>
        amount += (cartItem.product.price * cartItem.quantity));
    return amount;
  }

  void addItem(Product product) {
    if (_items.containsKey(product.id)) {
      _items.update(
          product.id,
          (existingCartItem) => CartItem(
                product: existingCartItem.product,
                quantity: existingCartItem.quantity + 1,
              ));
    } else {
      _items.putIfAbsent(
        product.id,
        () => CartItem(
          product: product,
        ),
      );
    }
    notifyListeners();
  }

  void removeItem(int productId) {
    _items.remove(productId);
    notifyListeners();
  }

  void removeSingleItem(int productId) {
    if (!_items.containsKey(productId)) return;
    if (_items[productId].quantity > 1) {
      _items.update(
          productId,
          (existingCartItem) => CartItem(
              product: existingCartItem.product,
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
