import { Test } from '@nestjs/testing';
import { DatabaseService } from '@way/backend-database';
import { PrismaClient, TenantEntity } from '@prisma/client';
import { lastValueFrom } from 'rxjs';
import { MockProxy, mock, mockDeep, DeepMockProxy } from 'jest-mock-extended';
import { TenantRepository } from './tenant.repository';

describe('TenantRepository', () => {
  let repository: TenantRepository;
  let prisma: DeepMockProxy<PrismaClient>;
  let databaseService: MockProxy<DatabaseService>;

  const mockTenant: TenantEntity = {
    id: '1',
    name: 'Test Tenant',
    description: 'Test Description',
    createdAt: new Date(),
    updatedAt: new Date(),
    deletedAt: null,
  };

  beforeEach(async () => {
    prisma = mockDeep<PrismaClient>({ funcPropSupport: true });
    databaseService = mock<DatabaseService>({
      tenantEntity: prisma.tenantEntity
    });

    const module = await Test.createTestingModule({
      providers: [
        TenantRepository,
        {
          provide: DatabaseService,
          useValue: databaseService,
        },
      ],
    }).compile();

    repository = module.get<TenantRepository>(TenantRepository);
  });

  describe('create$', () => {
    it('should create a new tenant successfully', async () => {
      const newTenant = {
        name: 'New Tenant',
        description: 'New Description',
      };

      prisma.tenantEntity.create.mockResolvedValue({
        ...mockTenant,
        ...newTenant,
      });

      const result = await lastValueFrom(repository.create$(newTenant));

      expect(result).toEqual(expect.objectContaining(newTenant));
      expect(prisma.tenantEntity.create).toHaveBeenCalledWith({
        data: newTenant,
      });
    });

    it('should propagate error if creation fails', async () => {
      const newTenant = {
        name: 'New Tenant',
        description: 'New Description',
      };

      const error = new Error('Database error');
      prisma.tenantEntity.create.mockRejectedValue(error);

      await expect(lastValueFrom(repository.create$(newTenant))).rejects.toThrow(error);
    });
  });

  describe('findById$', () => {
    it('should find an existing tenant', async () => {
      prisma.tenantEntity.findFirst.mockResolvedValue(mockTenant);

      const result = await lastValueFrom(repository.findById$('1'));

      expect(result).toEqual(mockTenant);
      expect(prisma.tenantEntity.findFirst).toHaveBeenCalledWith({
        where: { id: '1', deletedAt: null },
      });
    });

    it('should return null if tenant does not exist', async () => {
      prisma.tenantEntity.findFirst.mockResolvedValue(null);

      const result = await lastValueFrom(repository.findById$('999'));

      expect(result).toBeNull();
    });
  });

  describe('update$', () => {
    it('should update a tenant successfully', async () => {
      const updateData = {
        name: 'Updated Name',
      };

      const updatedTenant = {
        ...mockTenant,
        ...updateData,
      };

      prisma.tenantEntity.update.mockResolvedValue(updatedTenant);

      const result = await lastValueFrom(repository.update$('1', updateData));

      expect(result).toEqual(updatedTenant);
      expect(prisma.tenantEntity.update).toHaveBeenCalledWith({
        where: { id: '1' },
        data: updateData,
      });
    });

    it('should propagate error if update fails', async () => {
      const error = new Error('Database error');
      prisma.tenantEntity.update.mockRejectedValue(error);

      await expect(
        lastValueFrom(repository.update$('1', { name: 'Updated Name' }))
      ).rejects.toThrow(error);
    });
  });

  describe('softDelete$', () => {
    it('should perform soft delete successfully', async () => {
      const deletedTenant = {
        ...mockTenant,
        deletedAt: new Date(),
      };

      prisma.tenantEntity.update.mockResolvedValue(deletedTenant);

      const result = await lastValueFrom(repository.softDelete$('1'));

      expect(result.deletedAt).not.toBeNull();
      expect(prisma.tenantEntity.update).toHaveBeenCalledWith({
        where: { id: '1' },
        data: { deletedAt: expect.any(Date) },
      });
    });
  });

  describe('findAll$', () => {
    it('should find all non-deleted tenants', async () => {
      const tenants = [mockTenant, { ...mockTenant, id: '2' }];
      prisma.tenantEntity.findMany.mockResolvedValue(tenants);

      const result = await lastValueFrom(repository.findAll$());

      expect(result).toEqual(tenants);
      expect(prisma.tenantEntity.findMany).toHaveBeenCalledWith({
        where: { deletedAt: null },
      });
    });

    it('should return empty array if no tenants exist', async () => {
      prisma.tenantEntity.findMany.mockResolvedValue([]);

      const result = await lastValueFrom(repository.findAll$());

      expect(result).toEqual([]);
    });
  });
});
