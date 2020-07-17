import 'package:flutter/foundation.dart';
import 'package:get_it/get_it.dart';

import '../models/order_item.dart';
import '../models/cart_item.dart';
import '../models/api_response.dart';
import '../services/order_service.dart';

class Orders with ChangeNotifier {
  final OrderService orderService = GetIt.instance.get<OrderService>();
  List<OrderItem> _orders = [];

  List<OrderItem> get orders {
    return [..._orders];
  }

  void addOrder(List<CartItem> cartProducts, double total) async {
    OrderItem order = OrderItem(
      id: null,
      amount: total,
      dateTime: DateTime.now(),
      cartItems: cartProducts,
    );
    ApiResponse response = await orderService.add(order);
    OrderItem createdOrder = response.data[0];
    OrderItem newOrder = order.clone(id: createdOrder.id);
    _orders.insert(0, newOrder);
    notifyListeners();
  }
}
