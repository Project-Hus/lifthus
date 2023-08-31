import { Module } from '@nestjs/common';
import { DomainModule } from '../domain/domain.module';
import { PostRepository } from '../domain/repositories/post.repository';
import { PrismaPostRepository } from './post.repository';
import { PrismaService } from 'src/prisma/prisma.service';
import { UserRepository } from './user.repository';
import { UserRepository as AbstractUserRepository } from '../domain/repositories/user.repository';
import { LikeRepository } from '../domain/repositories/like.repository';
import { Post } from '../domain/aggregates/post/post.model';
import {
  PrismaCommentLikeRepository,
  PrismaPostLikeRepository,
} from './like.repository';
import { CommentRepository } from '../domain/repositories/comment.repository';
import { PrismaCommentRepository } from './comment.repository';

@Module({
  imports: [DomainModule],
  providers: [
    PrismaService,
    {
      provide: UserRepository,
      useClass: AbstractUserRepository,
    },
    {
      provide: PostRepository,
      useClass: PrismaPostRepository,
    },
    { provide: CommentRepository, useClass: PrismaCommentRepository },
    PrismaPostLikeRepository,
    PrismaCommentLikeRepository,
  ],
  exports: [
    UserRepository,
    PostRepository,
    CommentRepository,
    PrismaPostLikeRepository,
    PrismaCommentLikeRepository,
  ],
})
export class RepositoryModule {}
