import 'dart:math';

import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

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
            subtitle: Text(
                DateFormat('dd/MM/yyyy hh:mm').format(widget.order.dateTime)),
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
              height: min(widget.order.products.length * 20.0 + 10, 100),
              child: ListView(
                children: widget.order.products
                    .map((product) => Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: <Widget>[
                            Text(
                              product.title,
                              style: TextStyle(
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                            Text('${product.quantity}x \$${product.price}',
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
