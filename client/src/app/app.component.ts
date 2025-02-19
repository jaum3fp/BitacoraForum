import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { HelloWorldService } from './services/hello-world.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {

  title:any = "default"

  constructor(private hwsvc: HelloWorldService) {}
  
  ngOnInit(): void {
    this.hwsvc.getHelloWorld().subscribe(res => {
      this.title = res
    })
  }

}
