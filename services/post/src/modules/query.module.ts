import { Module } from '@nestjs/common';
import { QueryController } from './query/query.controller';
import { PostQueryService } from './query/post.query.service';

@Module({
  controllers: [QueryController],
  providers: [PostQueryService],
})
export class QueryModule {}
