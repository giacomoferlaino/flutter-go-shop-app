import 'package:flutter/widgets.dart';
import 'package:get_it/get_it.dart';

import '../models/auth_data.dart';
import '../services/auth_service.dart';
import '../models/api_response.dart';
import '../models/session_data.dart';

class Auth with ChangeNotifier {
  final AuthService authService = GetIt.instance.get<AuthService>();

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
    notifyListeners();
  }
}
