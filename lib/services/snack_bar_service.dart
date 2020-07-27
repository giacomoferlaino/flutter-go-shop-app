import 'package:flutter/material.dart';

class SnackBarService {
  final Duration duration;

  SnackBarService({@required this.duration});

  void show({
    @required BuildContext context,
    @required String message,
    SnackBarAction action,
    Color backgroundColor,
  }) {
    TextAlign textAlignment =
        action != null ? TextAlign.left : TextAlign.center;
    ScaffoldState scaffold = Scaffold.of(context);
    scaffold.hideCurrentSnackBar();
    scaffold.showSnackBar(SnackBar(
      backgroundColor: backgroundColor,
      content: Text(
        message,
        textAlign: textAlignment,
      ),
      duration: duration,
      action: action,
    ));
  }
}
