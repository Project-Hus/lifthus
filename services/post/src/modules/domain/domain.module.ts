import { Module } from '@nestjs/common';
import { Like } from './aggregates/like/like.model';
import { Post } from './aggregates/post/post.model';
import { User } from './aggregates/user/user.model';
import { Comment } from './aggregates/comment/comment.model';
import { PostSummary } from './aggregates/post/postSummary.model';

@Module({
  imports: [],
  providers: [User, PostSummary, Post, Comment, Like],
})
export class DomainModule {}
