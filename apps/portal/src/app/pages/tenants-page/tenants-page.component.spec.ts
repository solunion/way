import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TenantsPageComponent } from './tenants-page.component';

describe('TenantsPageComponent', () => {
  let component: TenantsPageComponent;
  let fixture: ComponentFixture<TenantsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TenantsPageComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TenantsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
