import { Test } from '@nestjs/testing';
import { Tenant } from './tenant.model';
import { mock, MockProxy } from 'jest-mock-extended';
import { lastValueFrom, of, throwError } from 'rxjs';
import { v4 as uuid } from 'uuid';
import { TenantResolver } from './tenant.resolver';
import { TenantService } from './tenant.service';
import { CreateTenantDto } from './dto/create-tenant.dto';
import { UpdateTenantDto } from './dto/update-tenant.dto';

describe('TenantResolver', () => {
  let resolver: TenantResolver;
  let service: MockProxy<TenantService>;

  const mockTenant: Tenant = {
    id: uuid(),
    name: 'Test Tenant',
    description: 'Test Description',
  };

  beforeEach(async () => {
    service = mock<TenantService>();

    const module = await Test.createTestingModule({
      providers: [
        TenantResolver,
        {
          provide: TenantService,
          useValue: service,
        },
      ],
    }).compile();

    resolver = module.get<TenantResolver>(TenantResolver);
  });

  describe('tenant', () => {
    it('should return an existing tenant', async () => {
      service.findById$.mockReturnValue(of(mockTenant));

      const result = await lastValueFrom(resolver.tenant(mockTenant.id));

      expect(result).toEqual(expect.objectContaining(mockTenant));
      expect(service.findById$).toHaveBeenCalledWith(mockTenant.id);
    });

    it('should throw an error if tenant is not found', async () => {
      const error = new Error('Tenant not found');
      service.findById$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(resolver.tenant('non-existent-id'))).rejects.toThrow(error);
    });
  });

  describe('createTenant', () => {
    it('should create a new tenant successfully', async () => {
      const createTenantDto: CreateTenantDto = {
        name: 'New Tenant',
        description: 'New Description',
      };

      const createdTenant: Tenant = {
        id: uuid(),
        ...createTenantDto,
      };

      service.create$.mockReturnValue(of(createdTenant));

      const result = await lastValueFrom(resolver.createTenant(createTenantDto));

      expect(result).toEqual(expect.objectContaining({
        name: createTenantDto.name,
        description: createTenantDto.description,
      }));
      
      expect(service.create$).toHaveBeenCalledWith(expect.objectContaining(createTenantDto));
    });

    it('should throw an error if tenant creation fails', async () => {
      const createTenantDto: CreateTenantDto = {
        name: 'New Tenant',
        description: 'New Description',
      };

      const error = new Error('Unable to create tenant');
      service.create$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(resolver.createTenant(createTenantDto))).rejects.toThrow(error);
    });
  });

  describe('updateTenant', () => {
    it('should update a tenant successfully', async () => {
      const updateTenantDto: UpdateTenantDto = {
        name: 'Updated Tenant',
      };

      const updatedTenant = {
        ...mockTenant,
        ...updateTenantDto,
      };

      service.update$.mockReturnValue(of(updatedTenant));

      const result = await lastValueFrom(resolver.updateTenant(mockTenant.id, updateTenantDto));

      expect(result).toEqual(expect.objectContaining(updateTenantDto));
      expect(service.update$).toHaveBeenCalledWith(mockTenant.id, expect.objectContaining(updateTenantDto));
    });

    it('should throw an error if tenant update fails', async () => {
      const updateTenantDto: UpdateTenantDto = {
        name: 'Updated Tenant',
      };

      const error = new Error('Unable to update tenant');
      service.update$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(resolver.updateTenant(mockTenant.id, updateTenantDto))).rejects.toThrow(error);
    });
  });

  describe('deleteTenant', () => {
    it('should delete a tenant successfully', async () => {
      service.delete$.mockReturnValue(of(undefined));

      const result = await lastValueFrom(resolver.deleteTenant(mockTenant.id));

      expect(result).toBe(true);
      expect(service.delete$).toHaveBeenCalledWith(mockTenant.id);
    });

    it('should return false if tenant deletion fails', async () => {
      const error = new Error('Unable to delete tenant');
      service.delete$.mockReturnValue(throwError(() => error));

      const result = await lastValueFrom(resolver.deleteTenant('non-existent-id'));

      expect(result).toBe(false);
    });
  });
}); 