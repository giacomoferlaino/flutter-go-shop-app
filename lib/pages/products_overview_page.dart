import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../providers/cart.dart';
import '../providers/products.dart';
import '../widgets/products_grid.dart';
import '../widgets/app_drawer.dart';
import '../widgets/badge.dart';
import './cart_page.dart';

enum FilterOptions {
  Favorites,
  All,
}

class ProductsOverviewPage extends StatefulWidget {
  static const routeName = '/product-overview';

  @override
  _ProductsOverviewPageState createState() => _ProductsOverviewPageState();
}

class _ProductsOverviewPageState extends State<ProductsOverviewPage> {
  bool _showOnlyFavorites = false;
  bool _isLoading = false;
  Products products;
  NavigatorState navigator;

  @override
  void initState() {
    super.initState();
    products = Provider.of<Products>(context, listen: false);
    navigator = Navigator.of(context);
    _isLoading = true;
    products.fetchAll().then((_) => setState(() {
          _isLoading = false;
        }));
  }

  void _toggleFavorites() {
    setState(() {
      _showOnlyFavorites = !_showOnlyFavorites;
    });
  }

  void _goToShoppingCart() {
    navigator.pushNamed(CartPage.routeName);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('MyShop'),
        actions: <Widget>[
          IconButton(
            icon: Icon(
                _showOnlyFavorites ? Icons.favorite : Icons.favorite_border),
            onPressed: _toggleFavorites,
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
              onPressed: _goToShoppingCart,
            ),
          ),
        ],
      ),
      drawer: AppDrawer(),
      body: _isLoading
          ? Center(
              child: CircularProgressIndicator(),
            )
          : RefreshIndicator(
              onRefresh: () => products.fetchAll(),
              child: ProductsGrid(_showOnlyFavorites),
            ),
    );
  }
}
