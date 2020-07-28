class AuthData {
  String email;
  String password;

  AuthData({this.email, this.password});

  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'password': password,
    };
  }
}
