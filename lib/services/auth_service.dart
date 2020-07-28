import 'dart:convert';

import 'package:get_it/get_it.dart';
import 'package:http/http.dart';

import '../models/session_data.dart';
import '../models/api_response.dart';
import '../models/auth_data.dart';
import '../services/http_service.dart';

class AuthService {
  final HttpService httpService = GetIt.instance.get<HttpService>();
  static const String relativePath = '/auth';
  final String _baseUrl;
  String _fullPath;

  AuthService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
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
    await httpService.request<SessionData>(
      request: () => post(_fullPath + '/signup', body: json.encode(authData)),
      dataParsing: (item) => null,
    );
    return login(authData.email, authData.password);
  }

  Future<ApiResponse<SessionData>> login(String email, String password) async {
    AuthData authData = AuthData(
      email: email,
      password: password,
    );
    return httpService.request<SessionData>(
      request: () => post(_fullPath + '/login', body: json.encode(authData)),
      dataParsing: _parseSessionData,
    );
  }
}
