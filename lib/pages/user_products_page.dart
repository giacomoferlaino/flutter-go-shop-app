import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../providers/products.dart';
import '../widgets/user_product_item.dart';
import '../widgets/app_drawer.dart';
import '../pages/edit_product_page.dart';

class UserProductsPage extends StatelessWidget {
  static const String routeName = '/user-products';

  @override
  Widget build(BuildContext context) {
    final Products products = Provider.of<Products>(context);
    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'Your Products!',
        ),
        actions: <Widget>[
          IconButton(
            icon: const Icon(Icons.add),
            onPressed: () {
              Navigator.of(context).pushNamed(EditProductPage.routeName);
            },
          ),
        ],
      ),
      drawer: AppDrawer(),
      body: RefreshIndicator(
        onRefresh: () => products.fetchAll(),
        child: Padding(
          padding: EdgeInsets.all(8),
          child: ListView.builder(
            itemCount: products.items.length,
            itemBuilder: (_, index) => UserProductItem(products.items[index]),
          ),
        ),
      ),
    );
  }
}
