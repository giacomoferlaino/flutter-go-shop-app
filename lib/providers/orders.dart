import 'package:flutter/foundation.dart';
import 'package:get_it/get_it.dart';

import '../models/order_item.dart';
import '../models/cart_item.dart';
import '../models/api_response.dart';
import '../services/order_service.dart';
import '../services/user_service.dart';

class Orders with ChangeNotifier {
  final OrderService orderService = GetIt.instance.get<OrderService>();
  final UserService userService = GetIt.instance.get<UserService>();
  List<OrderItem> _items = [];

  List<OrderItem> get items {
    return [..._items];
  }

  Future<void> fetchAll() async {
    ApiResponse response = await userService.getOrders();
    _items = response.data.reversed.toList();
    notifyListeners();
  }

  Future<void> addOrder(List<CartItem> cartProducts, double total) async {
    OrderItem order = OrderItem(
      id: null,
      amount: total,
      dateTime: DateTime.now(),
      cartItems: cartProducts,
    );
    ApiResponse response = await userService.addOrder(order);
    OrderItem createdOrder = response.data[0];
    OrderItem newOrder = order.clone(id: createdOrder.id);
    _items.insert(0, newOrder);
    notifyListeners();
  }
}
