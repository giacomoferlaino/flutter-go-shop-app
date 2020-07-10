import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import './providers/cart.dart';
import './providers/products.dart';
import './providers/orders.dart';
import './pages/orders_page.dart';
import './pages/cart_page.dart';
import './pages/products_overview_page.dart';
import './pages/product_detailt_page.dart';
import './pages/user_products_page.dart';
import './pages/edit_product_page.dart';
import './services/product_service.dart';

const String baseUrl = 'http://10.0.2.2:8080';
final ProductService productService = ProductService(baseUrl);

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(
          create: (context) => Products(productService),
        ),
        ChangeNotifierProvider(
          create: (context) => Cart(),
        ),
        ChangeNotifierProvider(
          create: (context) => Orders(),
        )
      ],
      child: MaterialApp(
        title: 'MyShop',
        theme: ThemeData(
          primarySwatch: Colors.purple,
          accentColor: Colors.deepOrange,
          fontFamily: 'Lato',
        ),
        initialRoute: ProductsOverviewPage.routeName,
        routes: {
          ProductsOverviewPage.routeName: (context) => ProductsOverviewPage(),
          ProductDetailPage.routeName: (context) => ProductDetailPage(),
          CartPage.routeName: (context) => CartPage(),
          OrdersPage.routeName: (context) => OrdersPage(),
          UserProductsPage.routeName: (context) => UserProductsPage(),
          EditProductPage.routeName: (context) => EditProductPage(),
        },
      ),
    );
  }
}
