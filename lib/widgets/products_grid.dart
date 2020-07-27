import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:provider/provider.dart';
import 'package:shop_app/providers/shop_filters.dart';
import 'package:shop_app/services/snack_bar_service.dart';

import '../providers/products.dart';
import '../providers/product.dart';
import './product_item.dart';

class ProductsGrid extends StatefulWidget {
  @override
  _ProductsGridState createState() => _ProductsGridState();
}

class _ProductsGridState extends State<ProductsGrid> {
  bool _isInit = true;
  bool _isLoading = false;
  Products _products;
  ShopFilters _filters;
  SnackBarService _snackBar;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    if (_isInit) {
      _products = Provider.of<Products>(context);
      _filters = Provider.of<ShopFilters>(context);
      _snackBar = GetIt.instance.get<SnackBarService>();
      _fetchData();
    }
    _isInit = false;
  }

  Future<void> _fetchData() {
    setState(() {
      _isLoading = true;
    });

    return _products.fetchAll().catchError((excetion) {
      _snackBar.show(
        context: context,
        message: excetion.toString(),
        backgroundColor: Color.fromRGBO(191, 1, 1, 0.7),
      );
    }).then((_) {
      setState(() {
        _isLoading = false;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    final productsData = Provider.of<Products>(context);
    final List<Product> products = _filters.showFavorites
        ? productsData.favoriteItems
        : productsData.items;
    return _isLoading
        ? Center(
            child: CircularProgressIndicator(),
          )
        : RefreshIndicator(
            onRefresh: _fetchData,
            child: GridView.builder(
              padding: const EdgeInsets.all(10),
              itemCount: products.length,
              itemBuilder: (context, index) {
                return ChangeNotifierProvider.value(
                  child: ProductItem(),
                  value: products[index],
                );
              },
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                childAspectRatio: 3 / 2,
                crossAxisSpacing: 10,
                mainAxisSpacing: 10,
              ),
            ),
          );
  }
}
