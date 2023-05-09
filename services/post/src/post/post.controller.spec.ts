import { Test, TestingModule } from '@nestjs/testing';
import { GetController } from '../get/get.controller';
import { GetService } from 'src/get/get.service';

describe('AppController', () => {
  let getController: GetController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [GetController],
      providers: [GetService],
    }).compile();

    getController = app.get<GetController>(GetController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(getController.getHello()).toBe('Hello World!');
    });
  });
});
