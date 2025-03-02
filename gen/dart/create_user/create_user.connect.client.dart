//
//  Generated code. Do not modify.
//  source: create_user/create_user.proto
//

import "package:connectrpc/connect.dart" as connect;
import "create_user.pb.dart" as create_usercreate_user;
import "create_user.connect.spec.dart" as specs;

extension type UserServiceClient (connect.Transport _transport) {
  Future<create_usercreate_user.CreateUserResponse> createUser(
    create_usercreate_user.CreateUserRequest input, {
    connect.Headers? headers,
    connect.AbortSignal? signal,
    Function(connect.Headers)? onHeader,
    Function(connect.Headers)? onTrailer,
  }) {
    return connect.Client(_transport).unary(
      specs.UserService.createUser,
      input,
      signal: signal,
      headers: headers,
      onHeader: onHeader,
      onTrailer: onTrailer,
    );
  }
}
