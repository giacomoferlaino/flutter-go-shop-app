import 'dart:convert';

import 'package:http/http.dart';

import '../services/http_service.dart';
import '../models/session_data.dart';
import '../models/api_response.dart';
import '../models/auth_data.dart';
import '../models/url.dart';
import '../models/active_session.dart';

class AuthService {
  final HttpService httpService;
  static const String relativePath = '/auth';
  final Url _baseUrl;
  ActiveSession activeSession;

  AuthService(this.httpService, this._baseUrl) {
    this.activeSession = ActiveSession(null, null);
  }

  bool get isAuth {
    return token != null;
  }

  String get token {
    if (activeSession.expirationDate != null &&
        activeSession.expirationDate.isAfter(DateTime.now()) &&
        activeSession.token != null) {
      return activeSession.token;
    }
    return null;
  }

  DateTime get expirationDate {
    return this.activeSession.expirationDate;
  }

  void registerToken(String token, DateTime expirationDate) {
    this.activeSession = ActiveSession(token, expirationDate);
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
    this.activeSession = ActiveSession(null, null);
  }
}
