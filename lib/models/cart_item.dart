import 'package:flutter/foundation.dart';

import '../providers/product.dart';

class CartItem {
  final Product product;
  final int quantity;

  CartItem({
    @required this.product,
    this.quantity = 1,
  });

  Map<String, dynamic> toJson() {
    return {
      'product': product,
      'quantity': quantity,
    };
  }
}
