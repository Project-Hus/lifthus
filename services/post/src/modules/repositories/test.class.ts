import { Injectable } from '@nestjs/common';

@Injectable()
export class ABC {
  constructor() {}
  hi(): string {
    return 'hi';
  }
}

@Injectable()
export class CON extends ABC {
  _hi(): string {
    return 'tmp';
  }
}
