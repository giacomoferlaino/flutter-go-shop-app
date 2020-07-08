import 'package:flutter/foundation.dart';

class Product with ChangeNotifier {
  final String id;
  final String title;
  final String description;
  final double price;
  final String imageUrl;
  bool isFavorite;

  Product({
    @required this.id,
    @required this.title,
    @required this.description,
    @required this.price,
    @required this.imageUrl,
    this.isFavorite = false,
  });

  Product clone({
    String id,
    String title,
    String description,
    double price,
    String imageUrl,
    bool isFavorite,
  }) {
    return Product(
      id: id == null ? this.id : id,
      title: title == null ? this.title : title,
      description: description == null ? this.description : description,
      price: price == null ? this.price : price,
      imageUrl: imageUrl == null ? this.imageUrl : imageUrl,
      isFavorite: isFavorite == null ? this.isFavorite : isFavorite,
    );
  }

  void toggleFavorite() {
    isFavorite = !isFavorite;
    notifyListeners();
  }
}
