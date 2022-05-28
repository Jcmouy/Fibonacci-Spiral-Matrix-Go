import { Component, OnInit } from '@angular/core';
import Typed from 'typed.js';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
  content?: string;
  typed: any;

  constructor() { }

  ngOnInit(): void {
    this.setTypedText();
  }

    setTypedText() {

        let stringsText = [];

        stringsText = ['Welcome to Fibonacci Spiral Matrix', 'Discover the magic of the mathematics',
            'and coding'];

        const options = {
            strings: stringsText,
            typeSpeed: 80,
            backSpeed: 10,
            showCursor: true,
            cursorChar: '',
            loop: true
        };

        this.typed = new Typed('.typing-element', options);
        this.typed.reset(true);
    }

}