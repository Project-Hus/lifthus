import { Injectable } from '@nestjs/common';
import { UserRepository as AbstractUserRepository } from '../domain/repositories/user.repository';

@Injectable()
export class UserRepository extends AbstractUserRepository {}
