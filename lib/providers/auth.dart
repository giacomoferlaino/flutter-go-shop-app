import 'package:flutter/widgets.dart';
import 'package:get_it/get_it.dart';

import '../models/auth_data.dart';
import '../services/auth_service.dart';
import '../models/api_response.dart';
import '../models/session_data.dart';

class Auth with ChangeNotifier {
  final AuthService authService = GetIt.instance.get<AuthService>();
  String _token;
  DateTime _expirationDate;
  String _userId;

  bool get isAuth {
    return token != null;
  }

  String get token {
    if (_expirationDate != null &&
        _expirationDate.isAfter(DateTime.now()) &&
        _token != null) {
      return _token;
    }
    return null;
  }

  Future<void> signUp(AuthData authData) async {
    ApiResponse response =
        await authService.signUp(authData.email, authData.password);
    SessionData sessionData = response.data[0];
    _token = sessionData.idToken;
    _expirationDate =
        DateTime.now().add(Duration(seconds: sessionData.expiresIn));
    notifyListeners();
  }

  Future<void> login(AuthData authData) async {
    ApiResponse response =
        await authService.login(authData.email, authData.password);
    SessionData sessionData = response.data[0];
    _token = sessionData.idToken;
    _expirationDate =
        DateTime.now().add(Duration(seconds: sessionData.expiresIn));
    notifyListeners();
  }
}
