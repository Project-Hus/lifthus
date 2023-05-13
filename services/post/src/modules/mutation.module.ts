import { Module } from '@nestjs/common';
import { PostController } from './post/post.controller';
import { CommentController } from './comment/comment.controller';
import { PostService } from './post/post.service';
import { PrismaService } from 'src/prisma/prisma.service';

@Module({
  controllers: [PostController, CommentController],
  providers: [PostService, PrismaService],
})
export class MutationModule {}
