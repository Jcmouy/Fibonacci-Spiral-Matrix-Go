import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-rows',
  templateUrl: './rows.component.html',
  styleUrls: ['./rows.component.css']
})
export class RowsComponent implements OnInit {

  constructor() { }

  @Input() rows: any[];

  ngOnInit(): void {
  }

}
