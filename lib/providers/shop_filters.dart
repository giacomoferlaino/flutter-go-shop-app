import 'package:flutter/material.dart';

class ShopFilters with ChangeNotifier {
  bool showFavorites = false;

  void toggleFavorites() {
    showFavorites = !showFavorites;
    notifyListeners();
  }
}
