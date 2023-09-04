import { Module } from '@nestjs/common';

import { PostController } from 'src/modules/command/post/post.controller';

import { PostService } from 'src/modules/command/post/post.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { S3Service } from 'src/modules/command/post/s3.service';

import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';

import { RepositoryModule } from 'src/modules/repositories/repository.module';
import { Post2Service } from 'src/modules/command/post/post2.service';

@Module({
  imports: [RepositoryModule],
  controllers: [PostController],
  providers: [PostService, Post2Service, PrismaService, S3Service],
})
export class PostCommandModule {}
