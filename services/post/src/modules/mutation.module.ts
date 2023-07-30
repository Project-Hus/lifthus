import { Module } from '@nestjs/common';
import { PostController } from './post/post.controller';
import { CommentController } from './comment/comment.controller';
import { PostService } from './post/post.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentService } from './comment/comment.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';

@Module({
  controllers: [PostController, CommentController],
  providers: [
    PostService,
    CommentService,
    PrismaService,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class MutationModule {}
