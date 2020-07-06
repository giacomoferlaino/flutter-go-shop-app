import 'package:flutter/material.dart';

import '../models/cart_item.dart';

class CartElement extends StatelessWidget {
  final CartItem cartItem;

  CartElement(this.cartItem);

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: EdgeInsets.symmetric(
        horizontal: 15,
        vertical: 4,
      ),
      child: Padding(
        padding: EdgeInsets.all(8),
        child: ListTile(
          leading: CircleAvatar(
            child: Padding(
              padding: const EdgeInsets.all(5),
              child: FittedBox(
                child: Text('\$${cartItem.price}'),
              ),
            ),
          ),
          title: Text(cartItem.title),
          subtitle: Text('Total: \$${cartItem.quantity * cartItem.price}'),
          trailing: Text('x ${cartItem.quantity}'),
        ),
      ),
    );
  }
}
