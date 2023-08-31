import { Module } from '@nestjs/common';
import { CommentQueryController } from './comment/query.controller';
import { PostQueryService } from './post/post.query.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentQueryService } from './comment/comment.query.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';
import { RepositoryModule } from '../repositories/repository.module';
import { DomainModule } from '../domain/domain.module';
import { PostQueryController } from './post/post.query.controller';

@Module({
  controllers: [CommentQueryController, PostQueryController],
  imports: [DomainModule, RepositoryModule],
  providers: [
    PrismaService,
    PostQueryService,
    CommentQueryService,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class QueryModule {}
