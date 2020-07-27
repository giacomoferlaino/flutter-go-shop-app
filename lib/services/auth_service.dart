import '../models/session_data.dart';
import '../models/api_response.dart';

class AuthService {
  static const String relativePath = '/auth';
  final String _baseUrl;
  String _fullPath;

  AuthService(this._baseUrl) {
    _fullPath = _baseUrl + relativePath;
  }

  Future<ApiResponse<SessionData>> _authenticate(
      String email, String password) async {
    return Future.delayed(
      Duration(seconds: 1),
      () => ApiResponse<SessionData>(
        meta: MetaData(0, null),
        data: [
          SessionData(
            idToken: 'idToken',
            expiresIn: 60,
          ),
        ],
      ),
    );
  }

  Future<ApiResponse<SessionData>> signUp(String email, String password) async {
    return _authenticate(email, password);
  }

  Future<ApiResponse<SessionData>> login(String email, String password) async {
    return _authenticate(email, password);
  }
}
