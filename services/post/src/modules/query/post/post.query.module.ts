import { Module } from '@nestjs/common';
import { PostQueryService } from 'src/modules/query/post/post.query.service';
import { PostQueryController } from 'src/modules/query/post/post.query.controller';
import { RepositoryModule } from 'src/modules/repositories/repository.module';

@Module({
  imports: [RepositoryModule],
  controllers: [PostQueryController],
  providers: [PostQueryService],
})
export class PostQueryModule {}
