import { Injectable } from '@nestjs/common';
import { User } from '../domain/aggregates/user/user.model';

@Injectable()
export class UserRepository {
  constructor() {}
  getUser(id: bigint): User {
    return new User(id);
  }
}
