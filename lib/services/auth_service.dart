import 'package:shop_app/models/session_data.dart';

class AuthService {
  static const String relativePath = '/auth';
  final String _baseUrl;
  String _fullPath;

  AuthService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
  }

  Future<SessionData> _authenticate(String email, String password) async {
    return Future.delayed(
      Duration(seconds: 1),
      () => SessionData(
        idToken: 'idTokenMock',
        refreshToken: 'refreshTokenMock',
      ),
    );
  }

  Future<SessionData> signUp(String email, String password) async {
    return _authenticate(email, password);
  }

  Future<SessionData> login(String email, String password) async {
    return _authenticate(email, password);
  }
}
