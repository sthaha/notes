import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';
import 'rxjs/add/observable/from';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/mergeMap';
import 'rxjs/add/operator/merge';
// import fetch from 'node-fetch';


const rs = Observable.of('https://api.github.com/users');


interface FakeResponse {
  json(): string
}

const delayed = (timeout: number, val: any): Promise<FakeResponse> => new Promise((accept) => {
  setTimeout(() => accept(val), timeout)
})

// delayed(1000, 'works').then(console.log)


const fetch = (url: any) => delayed(800, {
  json: () => [
    {login: '1'},
    {login: '2'},
    {login: '3'},
  ]
})

// fetch('foobar').then(x => console.log(x.json()))

rs.flatMap(url => Observable.from(fetch(url)))
  .flatMap(res => res.json())
  .flatMap(users => Observable.of(users))
  // .flatMap(x => {
    // // console.log("........................................" );
    // // console.log(x);
    // // console.log("-----------------------------------------");
    // return x
  // })
  .subscribe((x: any) => console.log(x.login))
