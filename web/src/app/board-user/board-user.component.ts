import { Component, OnInit } from '@angular/core';
import {AppService} from "../app.service";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {Subject} from "rxjs";
import {takeUntil} from "rxjs/operators";

@Component({
  selector: 'app-board-user',
  templateUrl: './board-user.component.html',
  styleUrls: ['./board-user.component.css']
})
export class BoardUserComponent implements OnInit {
  content?: string;

  constructor(private appService: AppService) { }

  ngOnInit(): void {
  }

    title = 'Fibonacci Spiral';

    matrixForm = new FormGroup({
        numRows: new FormControl('', Validators.nullValidator && Validators.required),
        numColumns: new FormControl('', Validators.nullValidator && Validators.required),
    });

    arrayNumbers: any[] = [];
    arrayCount = 0;

    destroy$: Subject<boolean> = new Subject<boolean>();

    onSubmit() {
        this.appService.calculateMatrix(this.matrixForm.value.numRows, this.matrixForm.value.numColumns).pipe(takeUntil(this.destroy$)).subscribe((numbers: any[]) => {
            let arrayFromObject : any;
            arrayFromObject = Object.keys(numbers).map(key => numbers[key])
            arrayFromObject.forEach(function(data){
                numbers = data
            });
            this.arrayCount = numbers.length;
            this.arrayNumbers = numbers;
        });
    }

    ngOnDestroy() {
        this.destroy$.next(true);
        this.destroy$.unsubscribe();
    }

}
