import { Module } from '@nestjs/common';

import { PostController } from 'src/modules/command/post/post.controller';
import { CommentController } from 'src/modules/command/comment/comment.controller';

import { PostService } from 'src/modules/command/post/post.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentService } from 'src/modules/command/comment/comment.service';
import { S3Service } from 'src/modules/command/post/s3.service';

import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';

import { RepositoryModule } from 'src/modules/repositories/repository.module';
import { Post2Service } from 'src/modules/command/post/post2.service';

@Module({
  imports: [RepositoryModule],
  controllers: [PostController, CommentController],
  providers: [
    PostService,
    Post2Service,
    CommentService,
    PrismaService,
    S3Service,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class CommandModule {}
