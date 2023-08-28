import { Module } from '@nestjs/common';
import { Post } from './aggregates/post/post.model';
import { Comment } from './aggregates/post/comment.model';

@Module({
  imports: [],
  providers: [Post, Comment],
})
export class DomainModule {}
