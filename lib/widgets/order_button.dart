import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../providers/cart.dart';
import '../providers/orders.dart';

class OrderButton extends StatefulWidget {
  const OrderButton({
    Key key,
    @required this.cart,
  }) : super(key: key);

  final Cart cart;

  @override
  _OrderButtonState createState() => _OrderButtonState();
}

class _OrderButtonState extends State<OrderButton> {
  bool _isLoading = false;

  Future<void> _sendOrder() {
    setState(() {
      _isLoading = true;
    });
    return Provider.of<Orders>(context, listen: false)
        .addOrder(widget.cart.items.values.toList(), widget.cart.totalAmount)
        .then((_) {
      setState(() {
        widget.cart.clear();
        _isLoading = false;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return _isLoading
        ? Center(
            widthFactor: 3,
            child: CircularProgressIndicator(),
          )
        : FlatButton(
            child: const Text('ORDER NOW'),
            onPressed: widget.cart.totalAmount <= 0 ? null : _sendOrder,
            textColor: Theme.of(context).primaryColor,
          );
  }
}
