import { Test } from '@nestjs/testing';
import { Tenant } from './tenant.model';
import { mock, MockProxy } from 'jest-mock-extended';
import { lastValueFrom, of, throwError } from 'rxjs';
import { v4 as uuid } from 'uuid';
import { TenantController } from './tenant.controller';
import { TenantService } from './tenant.service';
import { CreateTenantDto } from './dto/create-tenant.dto';
import { UpdateTenantDto } from './dto/update-tenant.dto';

describe('TenantController', () => {
  let controller: TenantController;
  let service: MockProxy<TenantService>;

  const mockTenant: Tenant = {
    id: uuid(),
    name: 'Test Tenant',
    description: 'Test Description',
  };

  beforeEach(async () => {
    service = mock<TenantService>();

    const module = await Test.createTestingModule({
      controllers: [TenantController],
      providers: [
        {
          provide: TenantService,
          useValue: service,
        },
      ],
    }).compile();

    controller = module.get<TenantController>(TenantController);
  });

  describe('create$', () => {
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

      const result = await lastValueFrom(controller.create$(createTenantDto));

      expect(result).toEqual(expect.objectContaining({
        name: createTenantDto.name,
        description: createTenantDto.description,
      }));
      
      expect(service.create$).toHaveBeenCalledWith(createTenantDto);
    });

    it('should throw an error if tenant creation fails', async () => {
      const createTenantDto: CreateTenantDto = {
        name: 'New Tenant',
        description: 'New Description',
      };

      const error = new Error('Unable to create tenant');
      service.create$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(controller.create$(createTenantDto))).rejects.toThrow(error);
    });
  });

  describe('getOne$', () => {
    it('should return an existing tenant', async () => {
      service.findById$.mockReturnValue(of(mockTenant));

      const result = await lastValueFrom(controller.getOne$(mockTenant.id));

      expect(result).toEqual(expect.objectContaining(mockTenant));
      expect(service.findById$).toHaveBeenCalledWith(mockTenant.id);
    });

    it('should throw an error if tenant is not found', async () => {
      const error = new Error('Tenant not found');
      service.findById$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(controller.getOne$('non-existent-id'))).rejects.toThrow(error);
    });
  });

  describe('update$', () => {
    it('should update a tenant successfully', async () => {
      const updateTenantDto: UpdateTenantDto = {
        name: 'Updated Tenant',
      };

      const updatedTenant = {
        ...mockTenant,
        ...updateTenantDto,
      };

      service.update$.mockReturnValue(of(updatedTenant));

      const result = await lastValueFrom(controller.update$(mockTenant.id, updateTenantDto));

      expect(result).toEqual(expect.objectContaining(updateTenantDto));
      expect(service.update$).toHaveBeenCalledWith(mockTenant.id, expect.objectContaining(updateTenantDto));
    });

    it('should throw an error if tenant update fails', async () => {
      const updateTenantDto: UpdateTenantDto = {
        name: 'Updated Tenant',
      };

      const error = new Error('Unable to update tenant');
      service.update$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(controller.update$(mockTenant.id, updateTenantDto))).rejects.toThrow(error);
    });
  });

  describe('delete$', () => {
    it('should delete a tenant successfully', async () => {
      service.delete$.mockReturnValue(of(undefined));

      await expect(lastValueFrom(controller.delete$(mockTenant.id))).resolves.toBeUndefined();
      expect(service.delete$).toHaveBeenCalledWith(mockTenant.id);
    });

    it('should throw an error if tenant deletion fails', async () => {
      const error = new Error('Unable to delete tenant');
      service.delete$.mockReturnValue(throwError(() => error));

      await expect(lastValueFrom(controller.delete$('non-existent-id'))).rejects.toThrow(error);
    });
  });
});
