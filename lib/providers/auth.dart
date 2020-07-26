import 'package:flutter/widgets.dart';
import 'package:get_it/get_it.dart';

import '../models/auth_data.dart';
import '../services/auth_service.dart';

class Auth with ChangeNotifier {
  final AuthService authService = GetIt.instance.get<AuthService>();
  String _token;
  DateTime _expirationDate;
  String _userId;

  Future<void> signUp(AuthData authData) async {
    return authService.signUp(authData.email, authData.password);
  }

  Future<void> login(AuthData authData) async {
    return authService.login(authData.email, authData.password);
  }
}
