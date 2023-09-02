import { Module } from '@nestjs/common';

import { PrismaService } from 'src/prisma/prisma.service';

import { ConcreteUserRepository } from './user.repository';
import { PrismaPostRepository } from './post.repository';
import { PrismaCommentRepository } from './comment.repository';
import {
  PrismaCommentLikeRepository,
  PrismaPostLikeRepository,
} from './like.repository';
import { LikeRepository } from '../domain/repositories/like.repository';
import { PostRepository } from '../domain/repositories/post.repository';
import { CommentRepository } from '../domain/repositories/comment.repository';
import { UserRepository } from '../domain/repositories/user.repository';
import { Post } from '../domain/aggregates/post/post.model';
import { Comment } from '../domain/aggregates/comment/comment.model';
import { DomainModule } from '../domain/domain.module';

//
export abstract class ABC {
  hi(): string {
    return 'hi';
  }
}

export class CON extends ABC {}
//

const postRepositoryProvider = {
  provide: PostRepository,
  useClass: PrismaPostRepository,
};

const userRepositoryProvider = {
  provide: UserRepository,
  useClass: ConcreteUserRepository,
};

@Module({
  providers: [
    PrismaService,
    userRepositoryProvider,
    postRepositoryProvider,
    // { provide: CommentRepository, useClass: PrismaCommentRepository },
    // { provide: LikeRepository<Post>, useClass: PrismaPostLikeRepository },
    // { provide: LikeRepository<Comment>, useClass: PrismaCommentLikeRepository },
    { provide: ABC, useClass: CON },
  ],
  exports: [
    PrismaService,
    userRepositoryProvider,
    postRepositoryProvider,
    // CommentRepository,
    // LikeRepository<Post>,
    // LikeRepository<Comment>,
    ABC,
  ],
})
export class RepositoryModule {}
