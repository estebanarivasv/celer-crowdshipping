import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewShippingsComponent } from './view-shippings.component';

describe('ViewRequestsComponent', () => {
  let component: ViewShippingsComponent;
  let fixture: ComponentFixture<ViewShippingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewShippingsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewShippingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
