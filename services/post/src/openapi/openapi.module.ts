import { Module } from '@nestjs/common';
import { AppModule } from 'src/modules/app.module';

@Module({
  imports: [AppModule],
})
export class OpenapiModule {}
