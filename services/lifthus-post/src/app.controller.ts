import { Controller, Get, Header } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('/post')
  getHello(): string {
    console.log('DSFSDGDSGSDGSDFGDSFGFI');
    return this.appService.getHello();
  }
}
