import 'dart:math';

import 'package:flutter/material.dart';

import '../models/order_item.dart';

class OrderElement extends StatefulWidget {
  final OrderItem order;

  OrderElement(this.order);

  @override
  _OrderElementState createState() => _OrderElementState();
}

class _OrderElementState extends State<OrderElement> {
  bool _expanded = false;

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: EdgeInsets.all(10),
      child: Column(
        children: <Widget>[
          ListTile(
            title: Text('\$${widget.order.amount}'),
            subtitle: Text(widget.order.formatDateTime()),
            trailing: IconButton(
              icon: Icon(_expanded ? Icons.expand_less : Icons.expand_more),
              onPressed: () {
                setState(() {
                  _expanded = !_expanded;
                });
              },
            ),
          ),
          if (_expanded)
            Container(
              padding: EdgeInsets.all(15),
              height: min(widget.order.cartItems.length * 20.0 + 10, 100),
              child: ListView(
                children: widget.order.cartItems
                    .map((cartItem) => Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: <Widget>[
                            Text(
                              cartItem.product.title,
                              style: TextStyle(
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                            Text(
                                '${cartItem.quantity}x \$${cartItem.product.price}',
                                style: TextStyle(
                                  fontSize: 18,
                                  color: Colors.grey,
                                )),
                          ],
                        ))
                    .toList(),
              ),
            )
        ],
      ),
    );
  }
}
