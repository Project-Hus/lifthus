import { Module } from '@nestjs/common';
import { CommentQueryController } from 'src/modules/query/comment/comment.query.controller';
import { CommentQueryService } from 'src/modules/query/comment/comment.query.service';
import { RepositoryModule } from 'src/modules/repositories/repository.module';

@Module({
  imports: [RepositoryModule],
  controllers: [CommentQueryController],
  providers: [CommentQueryService],
})
export class CommentQueryModule {}
