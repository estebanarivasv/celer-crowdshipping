import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewShippingOfferComponent } from './new-shipping-offer.component';

describe('NewShippingOfferComponent', () => {
  let component: NewShippingOfferComponent;
  let fixture: ComponentFixture<NewShippingOfferComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewShippingOfferComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewShippingOfferComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
