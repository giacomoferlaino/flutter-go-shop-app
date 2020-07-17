import 'package:flutter/foundation.dart';
import 'package:intl/intl.dart';

import './cart_item.dart';

class OrderItem {
  static const String dateTimeFormat = 'dd/MM/yyyy hh:mm';
  final int id;
  final double amount;
  final List<CartItem> cartItems;
  final DateTime dateTime;

  OrderItem({
    @required this.id,
    @required this.amount,
    @required this.cartItems,
    @required this.dateTime,
  });

  OrderItem clone({
    int id,
    double amount,
    List<CartItem> products,
    DateTime dateTime,
  }) {
    return OrderItem(
      id: id == null ? this.id : id,
      amount: amount == null ? this.amount : amount,
      cartItems: products == null ? this.cartItems : products,
      dateTime: dateTime == null ? this.dateTime : dateTime,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'amount': amount,
      'cartItems': cartItems,
      'dateTime': dateTime.toString(),
    };
  }

  String formatDateTime() {
    return DateFormat(OrderItem.dateTimeFormat).format(dateTime);
  }
}
