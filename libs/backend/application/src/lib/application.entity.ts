import { Application } from '@prisma/client';

export class ApplicationEntity implements Application {
  id!: string;
  name!: string;
  description!: string | null;
  tenantId!: string | null;
  createdAt!: Date;
  updatedAt!: Date;
  deletedAt!: Date | null;
}
