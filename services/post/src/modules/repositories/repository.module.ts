import { Module } from '@nestjs/common';

import { Post } from 'src/modules/domain/aggregates/post/post.model';
import { Comment } from 'src/modules/domain/aggregates/comment/comment.model';

import { UserRepository } from 'src/modules/domain/repositories/user.repository';
import { PostRepository } from 'src/modules/domain/repositories/post.repository';
import { CommentRepository } from 'src/modules/domain/repositories/comment.repository';
import { LikeRepository } from 'src/modules/domain/repositories/like.repository';

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
    { provide: LikeRepository<Post>, useClass: PrismaPostLikeRepository },
    { provide: LikeRepository<Comment>, useClass: PrismaCommentLikeRepository },
  ],
  exports: [
    PrismaService,
    UserRepository,
    PostRepository,
    CommentRepository,
    LikeRepository<Post>,
    LikeRepository<Comment>,
  ],
})
export class RepositoryModule {}
