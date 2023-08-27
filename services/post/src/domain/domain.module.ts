import { Module } from '@nestjs/common';
import { Post } from './post.model';
import { Comment } from './comment.model';

@Module({
  imports: [],
  providers: [Post, Comment],
})
export class DomainModule {}
