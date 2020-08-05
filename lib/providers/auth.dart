import 'dart:async';

import 'package:flutter/widgets.dart';
import 'package:get_it/get_it.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../services/auth_service.dart';
import '../models/auth_data.dart';
import '../models/api_response.dart';
import '../models/session_data.dart';
import '../models/active_session.dart';

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
    authService.registerToken(sessionData.idToken, expirationDate);
    notifyListeners();
  }

  Future<void> login(AuthData authData) async {
    ApiResponse response =
        await authService.login(authData.email, authData.password);
    SessionData sessionData = response.data[0];
    final DateTime expirationDate = DateTime.now().add(
      Duration(seconds: sessionData.expiresIn),
    );
    authService.registerToken(sessionData.idToken, expirationDate);
    _autoLogout();
    notifyListeners();
    final prefs = await SharedPreferences.getInstance();
    prefs.setString('activeSession', authService.activeSession.toString());
  }

  Future<bool> autoLogin() async {
    final prefs = await SharedPreferences.getInstance();
    if (!prefs.containsKey('activeSession')) {
      return false;
    }
    final ActiveSession activeSession =
        ActiveSession.parse(prefs.getString('activeSession'));
    if (activeSession.expirationDate.isBefore(DateTime.now())) {
      return false;
    }
    authService.registerToken(
        activeSession.token, activeSession.expirationDate);
    notifyListeners();
    _autoLogout();
    return true;
  }

  Future<void> logout() async {
    authService.logout();
    if (_authTimer != null) {
      _authTimer.cancel();
      _authTimer = null;
    }
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('activeSession');
    notifyListeners();
  }

  void _autoLogout() {
    if (_authTimer != null) _authTimer.cancel();
    final timeOfExpiry =
        authService.expirationDate.difference(DateTime.now()).inSeconds;
    _authTimer = Timer(Duration(seconds: timeOfExpiry), logout);
  }
}
