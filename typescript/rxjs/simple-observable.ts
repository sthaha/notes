import { Observable } from 'rxjs/Observable';
import { Observer } from 'rxjs/Observer';
import { TeardownLogic } from 'rxjs/Subscription';
import { Subscriber } from 'rxjs/Subscriber';

import 'rxjs/add/observable/of';
import 'rxjs/add/observable/from';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/mergeMap';
import 'rxjs/add/operator/merge';

type SubscribeFn<T> = (x: Subscriber<T>) => TeardownLogic;

function counter(name: string,  end: number, timeout: number): SubscribeFn<any> {
  return function(o: Observer<any>): TeardownLogic {
    let count = 1
    o.next(`${name}: ${count}`)
    const interval = setInterval(() => {
      count++
      if (count <= end) {
        o.next(`${name}: ${count}`)
        return
      }

      o.complete()
      clearInterval(interval)
    }, timeout)

    return () => {clearInterval(interval)}
  }
}
const first =  Observable.create(counter("first ",  5, 1000))
const second = Observable.create(counter("second", 10, 500))
const third =  Observable.create(counter("third ", 15, 250))
first
  .merge(second)
  .merge(third)
  .subscribe((x: any) => console.log(' ... ', x))
