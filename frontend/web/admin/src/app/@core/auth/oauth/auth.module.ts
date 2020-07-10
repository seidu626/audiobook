import { AuthGuard } from './services/auth-guard.service';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';

export const COMPONENTS = [
];

@NgModule({
  imports: [CommonModule, ReactiveFormsModule],
  declarations: COMPONENTS
})
export class AuthModule {
  static forRoot() {
    return {
      ngModule: RootAuthModule,
      providers: [AuthGuard]
    };
  }
}

@NgModule({
  imports: [
    AuthModule
  ],
})

export class RootAuthModule {}
