import 'dart:convert';

import 'package:http/http.dart';

import '../models/session_data.dart';
import '../models/api_response.dart';
import '../models/auth_data.dart';
import '../models/url.dart';
import '../services/http_service.dart';

class AuthService {
  final HttpService httpService;
  static const String relativePath = '/auth';
  final Url _baseUrl;
  String _token;
  DateTime _expirationDate;

  AuthService(this.httpService, this._baseUrl);

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

  void registerToken(String token, DateTime expirationDate) {
    this._token = token;
    this._expirationDate = expirationDate;
  }

  SessionData _parseSessionData(dynamic item) {
    return SessionData(
      idToken: item['idToken'],
      expiresIn: item['expiresIn'],
    );
  }

  Future<ApiResponse<SessionData>> signUp(String email, String password) async {
    AuthData authData = AuthData(
      email: email,
      password: password,
    );
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      path: relativePath + '/signup',
    );
    await httpService.request<SessionData>(
      request: () => post(uri.toString(), body: json.encode(authData)),
      dataParsing: (item) => null,
    );
    return login(authData.email, authData.password);
  }

  Future<ApiResponse<SessionData>> login(String email, String password) async {
    AuthData authData = AuthData(
      email: email,
      password: password,
    );
    final Uri uri = Uri(
      scheme: _baseUrl.scheme,
      host: _baseUrl.host,
      port: _baseUrl.port,
      path: relativePath + '/login',
    );
    return httpService.request<SessionData>(
      request: () => post(uri, body: json.encode(authData)),
      dataParsing: _parseSessionData,
    );
  }

  void logout() {
    this._token = null;
    this._expirationDate = null;
  }
}
