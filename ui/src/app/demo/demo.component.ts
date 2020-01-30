import { Component, OnInit } from '@angular/core';
import {Apollo} from 'apollo-angular';
import gql from 'graphql-tag';
@Component({
  selector: 'app-demo',
  templateUrl: './demo.component.html',
  styleUrls: ['./demo.component.css']
})
export class DemoComponent implements OnInit {

  count:number = 0;
  users:any[];
  loading :boolean;
  error :any = null;
  constructor(private apollo: Apollo) { }

  ngOnInit() {
    this.apollo
    .watchQuery({
      query: gql`
        {
          users {
            count
            list{
              uuid
              firstname
              lastname
              email
            }
            
          }
        }
      `,
    })
    .valueChanges.subscribe(result => {
      this.count= result.data && result.data['users'].count;
      this.users= result.data && result.data['users'].list;
      this.loading = result.loading;
      this.error = result.errors;
    });
  }

}
