import { Module } from '@nestjs/common';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/shared/filters/http-exception.filter';
import { PostQueryModule } from 'src/modules/query/post/post.query.module';
import { CommentQueryModule } from 'src/modules/query/comment/comment.query.module';

@Module({
  imports: [PostQueryModule, CommentQueryModule],
  providers: [{ provide: APP_FILTER, useClass: HttpExceptionFilter }],
})
export class QueryModule {}
