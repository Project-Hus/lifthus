import { Module } from '@nestjs/common';
import { CommentQueryController } from './comment/comment.query.controller';
import { PostQueryService } from 'src/modules/query/post/post.query.service';
import { CommentQueryService } from 'src/modules/query/comment/comment.query.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';
import { PostQueryController } from 'src/modules/query/post/post.query.controller';
import { RepositoryModule } from 'src/modules/repositories/repository.module';

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
