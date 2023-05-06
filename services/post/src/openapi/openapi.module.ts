import { Module } from '@nestjs/common';
import { OpenapiService } from './openapi.service';
import { OpenapiController } from './openapi.controller';
import { AppModule } from 'src/app.module';

@Module({
  imports: [AppModule],
  controllers: [OpenapiController],
  providers: [OpenapiService],
})
export class OpenapiModule {}
