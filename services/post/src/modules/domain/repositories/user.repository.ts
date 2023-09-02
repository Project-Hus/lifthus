import { Injectable, Logger } from '@nestjs/common';
import { User } from '../aggregates/user/user.model';

@Injectable()
export abstract class UserRepository {
  getUser(id: bigint): User {
    return User.create(id);
  }
}
