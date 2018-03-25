import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';
import 'rxjs/add/observable/from';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/mergeMap';
import 'rxjs/add/operator/merge';
import fetch from 'node-fetch';


// const rs = Observable.of('https://api.github.com/users');

// rs.flatMap(url => Observable.from(fetch(url)))
  // .flatMap(res => res.json())
  // .subscribe(res => console.log(res))
//
const stream = Observable.from([1, 2, 3])
// stream.subscribe(x => console.log('...', x))


// create an observable that emits at regular interval
//

function counter(name: string,  end: number, timeout: number) {
  return function(o: any) {
    let count = 0
    const interval = setInterval(() => {
      count++
      if (count < end) {
        o.next(`${name}: ${count}`)
        return
      }

      // terminate
      o.complete(count)
      clearInterval(interval)
    }, timeout)
  }
}
const first = Observable.create(counter("first",  5, 1000))
const second = Observable.create(counter("second", 10, 500))
first.merge(second).subscribe((x: any) => console.log(' ...e: ', x))
