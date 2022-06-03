import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewShippingRequestsComponent } from './view-shipping-requests.component';

describe('ViewShippingRequestsComponent', () => {
  let component: ViewShippingRequestsComponent;
  let fixture: ComponentFixture<ViewShippingRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewShippingRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewShippingRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
