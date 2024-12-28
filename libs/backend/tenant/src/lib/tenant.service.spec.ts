import { Test } from '@nestjs/testing';
import { Tenant } from '@prisma/client';
import { mock, MockProxy } from 'jest-mock-extended';
import { lastValueFrom, of, throwError } from 'rxjs';
import { v4 as uuid } from 'uuid';
import { Tenant as TenantModel } from './tenant.model';
import { TenantRepository } from './tenant.repository';
import { TenantService } from './tenant.service';

describe('TenantService', () => {
  let service: TenantService;
  let repository: MockProxy<TenantRepository>;

  const mockTenant: Tenant = {
    id: uuid(),
    name: 'Test Tenant',
    description: 'Test Description',
    createdAt: new Date(),
    updatedAt: new Date(),
    deletedAt: null,
  };

  const mockTenantId = mockTenant.id;

  beforeEach(async () => {
    repository = mock<TenantRepository>();

    const module = await Test.createTestingModule({
      providers: [
        TenantService,
        {
          provide: TenantRepository,
          useValue: repository,
        },
      ],
    }).compile();

    service = module.get<TenantService>(TenantService);
  });

  describe('create$', () => {
    it('should create a new tenant successfully', async () => {
      const newTenant = {
        name: 'New Tenant',
        description: 'New Description',
      };

      repository.create$.mockReturnValue(of({ ...mockTenant, ...newTenant }));

      const result = await lastValueFrom(service.create$(newTenant));

      expect(result).toEqual(expect.objectContaining(newTenant));
      expect(repository.create$).toHaveBeenCalledWith(newTenant);
    });

    it('should throw an error if tenant creation fails', async () => {
      const newTenant = {
        name: 'New Tenant',
        description: 'New Description',
      };

      const error = new Error('Unable to create tenant');
      repository.create$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(service.create$(newTenant))).rejects.toThrow(error);
    });
  });

  describe('findById$', () => {
    it('should find an existing tenant', async () => {
      repository.findById$.mockReturnValue(of(mockTenant));

      const result = await lastValueFrom(service.findById$('1'));

      expect(result).toEqual(expect.objectContaining(mockTenant));
      expect(repository.findById$).toHaveBeenCalledWith('1');
    });

    it('should throw an error if tenant is not found', async () => {
      repository.findById$.mockReturnValue(of(null));

      await expect(lastValueFrom(service.findById$('999'))).rejects.toThrow();
    });
  });

  describe('update$', () => {
    it('should update a tenant successfully', async () => {
      const updateData: Partial<TenantModel> = {
        name: 'Updated Name',
      };

      const updatedTenant = {
        ...mockTenant,
        ...updateData,
      };

      repository.update$.mockReturnValue(of(updatedTenant));

      const result = await lastValueFrom(service.update$('1', updateData));

      expect(result).toEqual(expect.objectContaining(updatedTenant));
      expect(repository.update$).toHaveBeenCalledWith('1', updateData);
    });

    it('should throw an error if tenant update fails', async () => {
      const updateData: Partial<TenantModel> = {
        name: 'Updated Name',
      };

      const error = new Error('Unable to update tenant');
      repository.update$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(service.update$('1', updateData))).rejects.toThrow(error);
    });
  });

  describe('delete$', () => {
    it('should delete a tenant successfully', async () => {
      const deletedTenant = {
        ...mockTenant,
        deletedAt: new Date(),
      };
      repository.softDelete$.mockReturnValue(of(deletedTenant));

      await expect(lastValueFrom(service.delete$(mockTenantId))).resolves.toBeUndefined();
      expect(repository.softDelete$).toHaveBeenCalledWith(mockTenantId);
    });

    it('should throw an error if tenant deletion fails', async () => {
      const error = new Error('Unable to delete tenant');
      repository.softDelete$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(service.delete$('non-existent-id'))).rejects.toThrow(error);
    });
  });
});
