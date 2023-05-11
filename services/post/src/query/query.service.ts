import { Injectable } from '@nestjs/common';
import { Response } from 'express';

@Injectable()
export class QueryService {
  getHello(): string {
    return 'Hello World!';
  }
  getCookie(res: Response): string {
    res.cookie('abc', 'nanadfsaf');
    return 'abcdef';
  }
}
