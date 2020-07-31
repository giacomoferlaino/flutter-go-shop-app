import 'dart:async';
import 'dart:convert';

import 'package:flutter/widgets.dart';
import 'package:get_it/get_it.dart';
import 'package:shared_preferences/shared_preferences.dart';

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
    final prefs = await SharedPreferences.getInstance();
    final userData = json.encode({
      'token': this.authService.token,
      'expirationDate': this.authService.expirationDate.toString()
    });
    prefs.setString('userData', userData);
  }

  Future<bool> authLogin() async {
    final prefs = await SharedPreferences.getInstance();
    if (!prefs.containsKey('userData')) {
      return false;
    }
    final userData =
        json.decode(prefs.getString('userData')) as Map<String, Object>;
    final expiryDate = DateTime.parse(userData['expirationDate']);
    if (expiryDate.isBefore(DateTime.now())) {
      return false;
    }
    this.authService.registerToken(userData['token'], expiryDate);
    notifyListeners();
    _autoLogout();
    return true;
  }

  Future<void> logout() async {
    this.authService.logout();
    if (_authTimer != null) {
      _authTimer.cancel();
      _authTimer = null;
    }
    final prefs = await SharedPreferences.getInstance();
    prefs.remove('userData');
    notifyListeners();
  }

  void _autoLogout() {
    if (_authTimer != null) _authTimer.cancel();
    final timeOfExpiry =
        this.authService.expirationDate.difference(DateTime.now()).inSeconds;
    _authTimer = Timer(Duration(seconds: timeOfExpiry), logout);
  }
}
