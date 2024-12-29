import { Test } from '@nestjs/testing';
import { PrismaClient } from '@prisma/client';
import { DeepMockProxy, mockDeep, mockReset } from 'jest-mock-extended';
import { lastValueFrom } from 'rxjs';
import { DatabaseService } from './database.service';

describe('DatabaseService', () => {
  let service: DatabaseService;
  let prismaClient: DeepMockProxy<PrismaClient>;

  beforeEach(async () => {
    prismaClient = mockDeep<PrismaClient>();

    const module = await Test.createTestingModule({
      providers: [
        DatabaseService,
        {
          provide: PrismaClient,
          useValue: prismaClient,
        },
      ],
    }).compile();

    service = module.get<DatabaseService>(DatabaseService);
  });

  afterEach(() => {
    mockReset(prismaClient);
  });

  describe('onModuleInit', () => {
    it('should call $connect on PrismaClient', async () => {
      await lastValueFrom(service.onModuleInit());
      jest.spyOn(service, '$connect').mockResolvedValue(prismaClient.$connect());

      expect(prismaClient.$connect).toHaveBeenCalledTimes(1);
    });

    it('should propagate error if $connect fails', async () => {
      const error = new Error('Connection error');
      jest.spyOn(service, '$connect').mockRejectedValue(error);

      await expect(lastValueFrom(service.onModuleInit())).rejects.toThrow(error);
    });
  });

  describe('onModuleDestroy', () => {
    it('should call $disconnect on PrismaClient', async () => {
      await lastValueFrom(service.onModuleDestroy());
      jest.spyOn(service, '$disconnect').mockResolvedValue(prismaClient.$disconnect());
      expect(prismaClient.$disconnect).toHaveBeenCalledTimes(1);
    });

    it('should propagate error if $disconnect fails', async () => {
      const error = new Error('Disconnection error');
      jest.spyOn(service, '$disconnect').mockRejectedValue(error);

      await expect(lastValueFrom(service.onModuleDestroy())).rejects.toThrow(error);
    });
  });
});
