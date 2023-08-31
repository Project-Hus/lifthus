import { Injectable } from '@nestjs/common';
import { User } from '../domain/aggregates/user/user.model';
import { UserRepository as AbstractUserRepository } from '../domain/repositories/user.repository';

@Injectable()
export class UserRepository extends AbstractUserRepository {
  getUser(id: bigint): User {
    return new User(id);
  }
}
