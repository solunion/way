import { Test } from '@nestjs/testing';
import { PrismaClient, Tenant } from '@prisma/client';
import { DatabaseService } from '@way/backend-database';
import { DeepMockProxy, mock, mockDeep, MockProxy } from 'jest-mock-extended';
import { lastValueFrom } from 'rxjs';
import { v4 as uuid } from 'uuid';
import { TenantRepository } from './tenant.repository';

describe('TenantRepository', () => {
  let repository: TenantRepository;
  let prisma: DeepMockProxy<PrismaClient>;
  let databaseService: MockProxy<DatabaseService>;

  const mockTenant: Tenant = {
    id: uuid(),
    name: 'Test Tenant',
    description: 'Test Description',
    createdAt: new Date(),
    updatedAt: new Date(),
    deletedAt: null,
  };

  beforeEach(async () => {
    prisma = mockDeep<PrismaClient>({ funcPropSupport: true });
    databaseService = mock<DatabaseService>({
      tenant: prisma.tenant,
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

      prisma.tenant.create.mockResolvedValue({
        ...mockTenant,
        ...newTenant,
      });

      const result = await lastValueFrom(repository.create$(newTenant));

      expect(result).toEqual(expect.objectContaining(newTenant));
      expect(prisma.tenant.create).toHaveBeenCalledWith({
        data: newTenant,
      });
    });

    it('should propagate error if creation fails', async () => {
      const newTenant = {
        name: 'New Tenant',
        description: 'New Description',
      };

      const error = new Error('Database error');
      prisma.tenant.create.mockRejectedValue(error);

      await expect(lastValueFrom(repository.create$(newTenant))).rejects.toThrow(error);
    });
  });

  describe('findById$', () => {
    it('should find an existing tenant', async () => {
      prisma.tenant.findFirst.mockResolvedValue(mockTenant);

      const result = await lastValueFrom(repository.findById$(mockTenant.id));

      expect(result).toEqual(mockTenant);
      expect(prisma.tenant.findFirst).toHaveBeenCalledWith({
        where: { id: mockTenant.id, deletedAt: null },
      });
    });

    it('should return null if tenant does not exist', async () => {
      prisma.tenant.findFirst.mockResolvedValue(null);

      const result = await lastValueFrom(repository.findById$(uuid()));

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

      prisma.tenant.update.mockResolvedValue(updatedTenant);

      const result = await lastValueFrom(repository.update$(mockTenant.id, updateData));

      expect(result).toEqual(updatedTenant);
      expect(prisma.tenant.update).toHaveBeenCalledWith({
        where: { id: mockTenant.id },
        data: updateData,
      });
    });

    it('should propagate error if update fails', async () => {
      const error = new Error('Database error');
      prisma.tenant.update.mockRejectedValue(error);

      await expect(lastValueFrom(repository.update$(mockTenant.id, { name: 'Updated Name' }))).rejects.toThrow(error);
    });
  });

  describe('softDelete$', () => {
    it('should perform soft delete successfully', async () => {
      const deletedTenant = {
        ...mockTenant,
        deletedAt: new Date(),
      };

      prisma.tenant.update.mockResolvedValue(deletedTenant);

      const result = await lastValueFrom(repository.softDelete$(mockTenant.id));

      expect(result.deletedAt).not.toBeNull();
      expect(prisma.tenant.update).toHaveBeenCalledWith({
        where: { id: mockTenant.id },
        data: { deletedAt: expect.any(Date) },
      });
    });
  });

  describe('findAll$', () => {
    it('should find all non-deleted tenants', async () => {
      const tenants = [mockTenant, { ...mockTenant, id: uuid() }];
      prisma.tenant.findMany.mockResolvedValue(tenants);

      const result = await lastValueFrom(repository.findAll$());

      expect(result).toEqual(tenants);
      expect(prisma.tenant.findMany).toHaveBeenCalledWith({
        where: { deletedAt: null },
      });
    });

    it('should return empty array if no tenants exist', async () => {
      prisma.tenant.findMany.mockResolvedValue([]);

      const result = await lastValueFrom(repository.findAll$());

      expect(result).toEqual([]);
    });
  });
});
