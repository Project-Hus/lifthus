import { Module } from '@nestjs/common';
import { CommentQueryController } from './comment/comment.query.controller';
import { PostQueryService } from './post/post.query.service';
import { CommentQueryService } from './comment/comment.query.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';
import { PostQueryController } from './post/post.query.controller';
import { RepositoryModule } from '../repositories/repository.module';
import { DomainModule } from '../domain/domain.module';

@Module({
  imports: [RepositoryModule],
  controllers: [CommentQueryController, PostQueryController],
  providers: [
    PostQueryService,
    CommentQueryService,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class QueryModule {}
