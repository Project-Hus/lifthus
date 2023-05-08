import { Injectable } from '@nestjs/common';

@Injectable()
export class GetService {
  getHello(): string {
    return 'Hello World!';
  }
}
