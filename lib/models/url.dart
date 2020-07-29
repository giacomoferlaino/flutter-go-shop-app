import 'package:flutter/foundation.dart';

class Url {
  final String scheme;
  final String host;
  final int port;

  Url({
    @required this.scheme,
    @required this.host,
    @required this.port,
  });
}
