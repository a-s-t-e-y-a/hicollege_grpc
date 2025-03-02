//
//  Generated code. Do not modify.
//  source: on_boarding/on_boarding.proto
//

import "package:connectrpc/connect.dart" as connect;
import "on_boarding.pb.dart" as on_boardingon_boarding;
import "on_boarding.connect.spec.dart" as specs;

extension type OtpServiceClient (connect.Transport _transport) {
  Future<on_boardingon_boarding.SendOtpResponse> sendOtp(
    on_boardingon_boarding.SendOtpRequest input, {
    connect.Headers? headers,
    connect.AbortSignal? signal,
    Function(connect.Headers)? onHeader,
    Function(connect.Headers)? onTrailer,
  }) {
    return connect.Client(_transport).unary(
      specs.OtpService.sendOtp,
      input,
      signal: signal,
      headers: headers,
      onHeader: onHeader,
      onTrailer: onTrailer,
    );
  }

  Future<on_boardingon_boarding.VerifyOtpResponse> verifyOtp(
    on_boardingon_boarding.VerifyOtpRequest input, {
    connect.Headers? headers,
    connect.AbortSignal? signal,
    Function(connect.Headers)? onHeader,
    Function(connect.Headers)? onTrailer,
  }) {
    return connect.Client(_transport).unary(
      specs.OtpService.verifyOtp,
      input,
      signal: signal,
      headers: headers,
      onHeader: onHeader,
      onTrailer: onTrailer,
    );
  }
}
