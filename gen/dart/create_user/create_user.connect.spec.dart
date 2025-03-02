//
//  Generated code. Do not modify.
//  source: create_user/create_user.proto
//

import "package:connectrpc/connect.dart" as connect;
import "create_user.pb.dart" as create_usercreate_user;

abstract final class UserService {
  /// Fully-qualified name of the UserService service.
  static const name = 'userpb.UserService';

  static const createUser = connect.Spec(
    '/$name/CreateUser',
    connect.StreamType.unary,
    create_usercreate_user.CreateUserRequest.new,
    create_usercreate_user.CreateUserResponse.new,
  );
}
