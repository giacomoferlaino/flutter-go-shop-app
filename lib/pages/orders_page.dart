import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:provider/provider.dart';

import '../providers/orders.dart';
import '../widgets/order_element.dart';
import '../widgets/app_drawer.dart';
import '../services/snack_bar_service.dart';

class OrdersPage extends StatefulWidget {
  static const String routeName = '/orders';

  @override
  _OrdersPageState createState() => _OrdersPageState();
}

class _OrdersPageState extends State<OrdersPage> {
  bool _isLoading = false;
  bool _isInit = true;
  Orders _orders;
  SnackBarService _snackBar;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    if (_isInit) {
      _orders = Provider.of<Orders>(context);
      _snackBar = GetIt.instance.get<SnackBarService>();
      _fetchData();
      _isInit = false;
    }
  }

  Future<void> _fetchData() {
    setState(() {
      _isLoading = true;
    });

    return _orders.fetchAll().catchError((excetion) {
      _snackBar.show(context: context, message: excetion.toString());
    }).then((_) {
      setState(() {
        _isLoading = false;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Your Orders'),
      ),
      drawer: AppDrawer(),
      body: _isLoading
          ? Center(
              child: CircularProgressIndicator(),
            )
          : RefreshIndicator(
              onRefresh: _fetchData,
              child: ListView.builder(
                itemCount: _orders.items.length,
                itemBuilder: (context, index) =>
                    OrderElement(_orders.items[index]),
              ),
            ),
    );
  }
}
