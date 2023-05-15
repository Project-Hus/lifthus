import { Module } from '@nestjs/common';
import { QueryController } from './query/query.controller';
import { PostQueryService } from './query/post.query.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentQueryService } from './query/comment.query.service';

@Module({
  controllers: [QueryController],
  providers: [PrismaService, PostQueryService, CommentQueryService],
})
export class QueryModule {}
