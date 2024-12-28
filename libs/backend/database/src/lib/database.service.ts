import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import { PrismaClient } from '@prisma/client';
import { from, Observable } from 'rxjs';

@Injectable()
export class DatabaseService extends PrismaClient implements OnModuleInit, OnModuleDestroy {
  onModuleInit = (): Observable<void> => {
    return from(this.$connect());
  };

  onModuleDestroy(): Observable<void> {
    return from(this.$disconnect());
  }
}
