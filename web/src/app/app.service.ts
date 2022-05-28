import { Injectable } from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {environment} from '../environments/environment';

const AUTH_API = `${environment.apiURL}/api/user/`;

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(private http: HttpClient) { }

  calculateMatrix(rows: string, columns: string) {
      const params = new HttpParams().append('rows', rows).append('cols', columns);
      return this.http.get(AUTH_API + 'spiral', {params});
  }

}
