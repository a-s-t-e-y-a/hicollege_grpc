//
//  Generated code. Do not modify.
//  source: on_boarding/on_boarding.proto
//

import "package:connectrpc/connect.dart" as connect;
import "on_boarding.pb.dart" as on_boardingon_boarding;

abstract final class OtpService {
  /// Fully-qualified name of the OtpService service.
  static const name = 'onboardingpb.OtpService';

  static const sendOtp = connect.Spec(
    '/$name/SendOtp',
    connect.StreamType.unary,
    on_boardingon_boarding.SendOtpRequest.new,
    on_boardingon_boarding.SendOtpResponse.new,
  );

  static const verifyOtp = connect.Spec(
    '/$name/VerifyOtp',
    connect.StreamType.unary,
    on_boardingon_boarding.VerifyOtpRequest.new,
    on_boardingon_boarding.VerifyOtpResponse.new,
  );
}
