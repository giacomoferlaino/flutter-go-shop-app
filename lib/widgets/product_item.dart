import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:provider/provider.dart';

import '../providers/cart.dart';
import '../providers/product.dart';
import '../providers/products.dart';
import '../pages/product_detailt_page.dart';
import '../services/snack_bar_service.dart';

class ProductItem extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final SnackBarService snackBarService =
        GetIt.instance.get<SnackBarService>();
    final Product product = Provider.of<Product>(context, listen: false);
    final Products products = Provider.of<Products>(context, listen: false);
    final bool isFavorite = products.isFavorite(product.id);
    return Container(
      decoration: BoxDecoration(
          border: Border.all(width: 0.5),
          borderRadius: BorderRadius.circular(10)),
      child: ClipRRect(
        borderRadius: BorderRadius.circular(10),
        child: GridTile(
          footer: GridTileBar(
            backgroundColor: Colors.black87,
            leading: Consumer<Product>(
              builder: (context, product, _) => IconButton(
                icon: Icon(isFavorite ? Icons.favorite : Icons.favorite_border),
                onPressed: () {
                  isFavorite
                      ? products.removeFavorite(product)
                      : products.addFavorite(product);
                },
                color: Theme.of(context).accentColor,
              ),
            ),
            trailing: Consumer<Cart>(
              builder: (context, cart, _) => IconButton(
                icon: Icon(Icons.shopping_cart),
                onPressed: () {
                  cart.addItem(product);
                  snackBarService.show(
                    context: context,
                    message: 'Item added to cart!',
                    action: SnackBarAction(
                      label: 'UNDO',
                      onPressed: () {
                        cart.removeSingleItem(product.id);
                      },
                    ),
                  );
                },
                color: Theme.of(context).accentColor,
              ),
            ),
            title: Text(
              product.title,
              textAlign: TextAlign.center,
            ),
          ),
          child: GestureDetector(
            onTap: () {
              Navigator.of(context).pushNamed(ProductDetailPage.routeName,
                  arguments: product.id);
            },
            child: FadeInImage(
              placeholder: AssetImage('assets/images/product-placeholder.png'),
              image: NetworkImage(product.imageUrl),
              fit: BoxFit.cover,
            ),
          ),
        ),
      ),
    );
  }
}
