import 'package:flutter/foundation.dart';
import 'package:get_it/get_it.dart';

import '../services/product_service.dart';

class Product with ChangeNotifier {
  final ProductService productService = GetIt.instance.get<ProductService>();
  final int id;
  final String title;
  final String description;
  final double price;
  final String imageUrl;

  Product({
    @required this.id,
    @required this.title,
    @required this.description,
    @required this.price,
    @required this.imageUrl,
  });

  Product clone({
    int id,
    String title,
    String description,
    double price,
    String imageUrl,
  }) {
    return Product(
      id: id == null ? this.id : id,
      title: title == null ? this.title : title,
      description: description == null ? this.description : description,
      price: price == null ? this.price : price,
      imageUrl: imageUrl == null ? this.imageUrl : imageUrl,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'title': title,
      'description': description,
      'price': price,
      'imageUrl': imageUrl,
    };
  }
}
