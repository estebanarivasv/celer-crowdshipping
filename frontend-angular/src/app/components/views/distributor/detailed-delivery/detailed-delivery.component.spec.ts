import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailedDeliveryComponent } from './detailed-delivery.component';

describe('DetailedDeliveryComponent', () => {
  let component: DetailedDeliveryComponent;
  let fixture: ComponentFixture<DetailedDeliveryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailedDeliveryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailedDeliveryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
