class RequestException implements Exception {
  final String message;

  RequestException(this.message);

  String toString() {
    return this.message;
  }
}
