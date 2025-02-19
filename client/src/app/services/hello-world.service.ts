import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class HelloWorldService {

  constructor(private client: HttpClient) {}

  getHelloWorld() {
    return this.client.get("http://127.21.0.12:8080", {responseType: "text"})
  }

}