import { Module } from '@nestjs/common';
import { DomainModule } from '../domain/domain.module';
import { PostRepository } from '../domain/repositories/post.repository';
import { PrismaPostRepository } from './post.repository';
import { PrismaService } from 'src/prisma/prisma.service';

@Module({
  imports: [DomainModule],
  providers: [
    PrismaService,
    {
      provide: PostRepository,
      useClass: PrismaPostRepository,
    },
  ],
})
export class RepositoryModule {}
