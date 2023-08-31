import { Module } from '@nestjs/common';
import { DomainModule } from '../domain/domain.module';
import { PostRepository } from '../domain/repositories/post.repository';
import { PrismaPostRepository } from './post.repository';

@Module({
  imports: [DomainModule],
  providers: [
    {
      provide: PostRepository,
      useClass: PrismaPostRepository,
    },
  ],
})
export class RepositoryModule {}
