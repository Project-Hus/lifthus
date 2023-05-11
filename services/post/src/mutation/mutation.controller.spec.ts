import { Test, TestingModule } from '@nestjs/testing';
import { query } from 'express';
import { QueryController } from 'src/query/query.controller';
import { QueryService } from 'src/query/query.service';

describe('AppController', () => {
  let queryController: QueryController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [QueryController],
      providers: [QueryService],
    }).compile();

    queryController = app.get<QueryController>(QueryController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(queryController.getHello()).toBe('Hello World!');
    });
  });
});
