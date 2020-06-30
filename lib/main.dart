import 'package:flutter/material.dart';
import 'package:shop_app/pages/product_detailt_page.dart';

import './pages/products_overview_page.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'MyShop',
      theme: ThemeData(
        primarySwatch: Colors.purple,
        accentColor: Colors.deepOrange,
        fontFamily: 'Lato',
      ),
      home: ProductsOverviewPage(),
      routes: {
        ProductDetailPage.routeName: (context) => ProductDetailPage(),
      },
    );
  }
}
