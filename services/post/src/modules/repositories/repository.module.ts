import { Module } from '@nestjs/common';

import { UserRepository } from 'src/modules/repositories/abstract/user.repository';
import { PostRepository } from 'src/modules/repositories/abstract/post.repository';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';
import {
  CommentLikeRepository,
  PostLikeRepository,
} from 'src/modules/repositories/abstract/like.repository';

import { PrismaService } from 'src/modules/repositories/prisma/prisma.service';
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
