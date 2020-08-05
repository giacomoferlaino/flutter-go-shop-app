import 'dart:convert';

class ActiveSession {
  String _token;
  DateTime _expirationDate;

  ActiveSession(this._token, this._expirationDate);

  ActiveSession.parse(String jsonText) {
    Map<String, String> parsedJson = json.decode(jsonText);
    this._token = parsedJson['token'];
    this._expirationDate = parsedJson['expirationDate'] != null
        ? DateTime.parse(parsedJson['expirationDate'])
        : null;
  }

  String get token {
    return _token;
  }

  DateTime get expirationDate {
    return _expirationDate;
  }

  String toString() {
    return json.encode({
      token: token,
      expirationDate: expirationDate.toString(),
    });
  }
}
