import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../providers/product.dart';
import '../providers/products.dart';
import '../pages/edit_product_page.dart';

class UserProductItem extends StatefulWidget {
  final Product product;

  UserProductItem(this.product);

  @override
  _UserProductItemState createState() => _UserProductItemState();
}

class _UserProductItemState extends State<UserProductItem> {
  bool _isLoading = false;

  _deleteProduct(Products products) {
    setState(() {
      _isLoading = true;
    });
    products.deleteByID(widget.product.id).then((_) => setState(() {
          _isLoading = false;
        }));
  }

  @override
  Widget build(BuildContext context) {
    Products products = Provider.of<Products>(context, listen: false);
    return _isLoading
        ? Center(
            child: CircularProgressIndicator(),
          )
        : Card(
            child: ListTile(
              title: Text(widget.product.title),
              leading: CircleAvatar(
                backgroundImage: NetworkImage(widget.product.imageUrl),
              ),
              trailing: Container(
                width: 100,
                child: Row(
                  children: <Widget>[
                    IconButton(
                        icon: Icon(Icons.edit),
                        onPressed: () {
                          Navigator.of(context).pushNamed(
                              EditProductPage.routeName,
                              arguments: widget.product.id);
                        },
                        color: Theme.of(context).primaryColor),
                    IconButton(
                        icon: Icon(Icons.delete),
                        onPressed: () => _deleteProduct(products),
                        color: Theme.of(context).errorColor),
                  ],
                ),
              ),
            ),
          );
  }
}
