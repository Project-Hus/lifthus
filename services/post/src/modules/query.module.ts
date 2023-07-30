import { Module } from '@nestjs/common';
import { QueryController } from './query/query.controller';
import { PostQueryService } from './query/post.query.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentQueryService } from './query/comment.query.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';

@Module({
  controllers: [QueryController],
  providers: [
    PrismaService,
    PostQueryService,
    CommentQueryService,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class QueryModule {}
