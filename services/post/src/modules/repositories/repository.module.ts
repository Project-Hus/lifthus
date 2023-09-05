import { Module } from '@nestjs/common';

import { Post } from 'src/domain/aggregates/post/post.model';
import { Comment } from 'src/domain/aggregates/comment/comment.model';

import { UserRepository } from 'src/domain/repositories/user.repository';
import { PostRepository } from 'src/domain/repositories/post.repository';
import { CommentRepository } from 'src/domain/repositories/comment.repository';
import {
  CommentLikeRepository,
  LikeRepository,
  PostLikeRepository,
} from 'src/domain/repositories/like.repository';

import { PrismaService } from 'src/prisma/prisma.service';
import { ConcreteUserRepository } from 'src/modules/repositories/user.repository';
import { PrismaPostRepository } from 'src/modules/repositories/post.repository';
import { PrismaCommentRepository } from 'src/modules/repositories/comment.repository';
import {
  PrismaCommentLikeRepository,
  PrismaPostLikeRepository,
} from 'src/modules/repositories/like.repository';

@Module({
  providers: [
    PrismaService,
    { provide: UserRepository, useClass: ConcreteUserRepository },
    { provide: PostRepository, useClass: PrismaPostRepository },
    { provide: CommentRepository, useClass: PrismaCommentRepository },
    { provide: PostLikeRepository, useClass: PrismaPostLikeRepository },
    {
      provide: CommentLikeRepository,
      useClass: PrismaCommentLikeRepository,
    },
  ],
  exports: [
    PrismaService,
    UserRepository,
    PostRepository,
    CommentRepository,
    PostLikeRepository,
    CommentLikeRepository,
  ],
})
export class RepositoryModule {}
