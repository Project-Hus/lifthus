import { Injectable } from '@nestjs/common';
import { Response } from 'express';

@Injectable()
export class CommentQueryService {
  getHello(): string {
    return 'Hello World!';
  }
}
