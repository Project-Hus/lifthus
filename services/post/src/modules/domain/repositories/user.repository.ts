import { Injectable } from '@nestjs/common';
import { User } from '../aggregates/user/user.model';

@Injectable()
export class UserRepository {
  constructor() {}
  getUser(id: bigint): User {
    return User.create(id);
  }
}
