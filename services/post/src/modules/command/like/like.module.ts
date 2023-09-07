import { Module } from '@nestjs/common';
import { LikeController } from './like.controller';
import { LikeService } from './like.service';
import { RepositoryModule } from 'src/modules/repositories/repository.module';

@Module({
  imports: [RepositoryModule],
  controllers: [LikeController],
  providers: [LikeService],
})
export class LikeModule {}
