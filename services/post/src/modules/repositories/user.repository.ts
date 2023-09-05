import { Injectable } from '@nestjs/common';
import { UserRepository } from './abstract/user.repository';

@Injectable()
export class ConcreteUserRepository extends UserRepository {}
