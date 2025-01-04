import { IsNotEmpty, IsUUID, MaxLength, MinLength } from 'class-validator';

export class NewRuleSet {
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  tenantId?: string;
  parent?: RuleSet;
  children?: RuleSet[];

  constructor(name: string, tenantId?: string, parent?: RuleSet, children?: RuleSet[]) {
    this.name = name;
    this.tenantId = tenantId;
    this.parent = parent;
    this.children = children;
  }
}

export class RuleSet extends NewRuleSet {
  @IsUUID()
  id: string;

  constructor(id: string, name: string, tenantId?: string, parent?: RuleSet, children?: RuleSet[]) {
    super(name, tenantId, parent, children);
    this.id = id;
  }
} 