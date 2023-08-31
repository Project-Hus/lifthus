import { Module } from '@nestjs/common';
import { DomainModule } from '../domain/domain.module';
import { PostRepository } from '../domain/repositories/post.repository';
import { PrismaPostRepository } from './post.repository';
import { PrismaService } from 'src/prisma/prisma.service';
import { UserRepository } from './user.repository';
import { UserRepository as AbstractUserRepository } from '../domain/repositories/user.repository';

@Module({
  imports: [DomainModule],
  providers: [
    PrismaService,
    {
      provide: UserRepository,
      useClass: AbstractUserRepository,
    },
    {
      provide: PostRepository,
      useClass: PrismaPostRepository,
    },
  ],
})
export class RepositoryModule {}
