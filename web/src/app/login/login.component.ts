import { Component, OnInit } from '@angular/core';
import { AuthService } from '../_services/auth.service';
import { TokenStorageService } from '../_services/token-storage.service';
import {AbstractControl, FormBuilder, FormGroup, Validators} from "@angular/forms";
import { Router } from '@angular/router';
import {LoginResponse} from "./login-response";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {
  submitted = false;
  form: FormGroup;

  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  roles: string[] = [];

  constructor(
      private formBuilder: FormBuilder,
      private authService: AuthService,
      private tokenStorage: TokenStorageService,
      private router: Router) { }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
      this.roles = this.tokenStorage.getUser().roles;
    }
      this.form = this.formBuilder.group(
          {
              username: [
                  '',
                  [
                      Validators.required,
                      Validators.minLength(6),
                      Validators.maxLength(20)
                  ]
              ],
              password: [
                  '',
                  [
                      Validators.required,
                      Validators.minLength(6),
                      Validators.maxLength(40)
                  ]
              ]
          }
      );
  }

  get f(): { [key: string]: AbstractControl } {
      return this.form.controls;
  }

  onSubmit(): void {
    this.submitted = true;
    this.authService.login(this.form.value.username, this.form.value.password).subscribe(
        response => {
            const loginResponse: LoginResponse = JSON.parse(JSON.stringify(response));
            console.log('accessToken Bearer ' + loginResponse.jwt);
            this.tokenStorage.saveToken(loginResponse.jwt);
            this.tokenStorage.saveUser(response);
            this.isLoginFailed = false;
            this.isLoggedIn = true;
            this.roles = this.tokenStorage.getUser().roles;
            this.reloadPage();
        },
        err => {
          this.errorMessage = err.error.message;
          this.isLoginFailed = true;
        }
    );
  }

  reloadPage(): void {
    window.location.reload();
  }

}
