import { Module } from '@nestjs/common';
import { User } from './aggregates/user/user.model';
import { Post } from './aggregates/post/post.model';
import { Comment } from './aggregates/comment/comment.model';
import { Like } from './aggregates/like/like.model';

@Module({
  providers: [User, Post, Comment, Like],
  exports: [User, Post, Comment, Like],
})
export class DomainModule {}
