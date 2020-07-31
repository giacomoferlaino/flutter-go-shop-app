import 'dart:async';

import 'package:flutter/widgets.dart';
import 'package:get_it/get_it.dart';

import '../models/auth_data.dart';
import '../services/auth_service.dart';
import '../models/api_response.dart';
import '../models/session_data.dart';

class Auth with ChangeNotifier {
  final AuthService authService = GetIt.instance.get<AuthService>();
  Timer _authTimer;

  Future<void> signUp(AuthData authData) async {
    ApiResponse response =
        await authService.signUp(authData.email, authData.password);
    SessionData sessionData = response.data[0];
    final DateTime expirationDate = DateTime.now().add(
      Duration(seconds: sessionData.expiresIn),
    );
    this.authService.registerToken(sessionData.idToken, expirationDate);
    notifyListeners();
  }

  Future<void> login(AuthData authData) async {
    ApiResponse response =
        await authService.login(authData.email, authData.password);
    SessionData sessionData = response.data[0];
    final DateTime expirationDate = DateTime.now().add(
      Duration(seconds: sessionData.expiresIn),
    );
    this.authService.registerToken(sessionData.idToken, expirationDate);
    _autoLogout();
    notifyListeners();
  }

  void logout() {
    this.authService.logout();
    if (_authTimer != null) {
      _authTimer.cancel();
      _authTimer = null;
    }
    notifyListeners();
  }

  void _autoLogout() {
    if (_authTimer != null) _authTimer.cancel();
    final timeOfExpiry =
        this.authService.expirationDate.difference(DateTime.now()).inSeconds;
    _authTimer = Timer(Duration(seconds: timeOfExpiry), logout);
  }
}
