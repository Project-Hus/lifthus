import { MiddlewareConsumer, Module } from '@nestjs/common';
import { PostController } from './post/post.controller';
import { CommentController } from './comment/comment.controller';
import { PostService } from './post/post.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentService } from './comment/comment.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';

import upload from 'src/common/utils/multer';
import { S3Service } from './post/s3.service';

@Module({
  controllers: [PostController, CommentController],
  providers: [
    PostService,
    CommentService,
    PrismaService,
    S3Service,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class MutationModule {}