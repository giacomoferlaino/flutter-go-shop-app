import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shop_app/providers/shop_filters.dart';

import '../providers/cart.dart';
import '../widgets/products_grid.dart';
import '../widgets/app_drawer.dart';
import '../widgets/badge.dart';
import './cart_page.dart';

enum FilterOptions {
  Favorites,
  All,
}

class ProductsOverviewPage extends StatelessWidget {
  static const routeName = '/product-overview';

  @override
  Widget build(BuildContext context) {
    NavigatorState navigator = Navigator.of(context);
    ShopFilters shopFilters = Provider.of<ShopFilters>(context);
    return Scaffold(
      appBar: AppBar(
        title: const Text('MyShop'),
        actions: <Widget>[
          IconButton(
            icon: Icon(shopFilters.showFavorites
                ? Icons.favorite
                : Icons.favorite_border),
            onPressed: shopFilters.toggleFavorites,
            tooltip: 'Show only favorites',
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 8),
            child: VerticalDivider(
              color: Colors.white54,
              thickness: 1,
            ),
          ),
          Consumer<Cart>(
            builder: (_, cart, child) => Badge(
              child: child,
              value: cart.totalQuantity.toString(),
            ),
            child: IconButton(
              icon: Icon(
                Icons.shopping_cart,
              ),
              onPressed: () => navigator.pushNamed(CartPage.routeName),
            ),
          ),
        ],
      ),
      drawer: AppDrawer(),
      body: ProductsGrid(),
    );
  }
}
