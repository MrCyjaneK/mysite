+++
Area = "projects"
CrumbParent = "projects"
Layout = "page"
Type = "project"
Tags = ["OldPC", "Servers"]
Title = "OldPC"
Date = "2021.01.24"
+++

I've deployed OldPC around a year ago, then for a short period of time it was available only as a local server, so that period boring, and then I was like: *Why... Why is is still not available online? That would be cool to have*

Well there was (and it still exist) one tiny problem... I'm behind NAT.

```python
import itertools

def iter_primes():
     # an iterator of all numbers between 2 and +infinity
     numbers = itertools.count(2)

     # generate primes forever
     while True:
         # get the first number from the iterator (always a prime)
         prime = numbers.next()
         yield prime

         # this code iteratively builds up a chain of
         # filters...slightly tricky, but ponder it a bit
         numbers = itertools.ifilter(prime.__rmod__, numbers)

for p in iter_primes():
    if p > 1000:
        break
    print p
```

Maecenas mollis urna at velit dapibus rhoncus. Etiam risus mauris, facilisis vitae erat laoreet, ornare sagittis diam.
Sed sem tellus, egestas eu eleifend et, suscipit sit amet quam. Maecenas eu massa nec urna aliquam maximus finibus sit
amet lectus. Quisque ultricies euismod libero, ac hendrerit enim interdum vitae. Sed quis mauris pharetra, hendrerit
nulla posuere, euismod velit. Vivamus urna tortor, congue non mi in, vehicula commodo leo. Duis volutpat fringilla
finibus. Duis ullamcorper lectus id turpis porta, vitae suscipit orci dictum.

[![Screenshot 1](img/screenshot1-thumb.png)](img/screenshot1.png)
[![Screenshot 2](img/screenshot2-thumb.png)](img/screenshot2.png)

## Instructions ##

Sed rhoncus suscipit blandit. Duis augue nulla, porttitor eu lorem maximus, dictum tempor felis. Nulla gravida cursus
arcu. Nullam lobortis lacus sem, non facilisis massa malesuada nec. Nunc efficitur ligula eu pulvinar pulvinar. Aenean
massa ligula, mattis et semper id, efficitur vel nisi. Sed pretium, quam at mollis sagittis, lorem mi auctor felis, eget
feugiat ipsum nunc a dolor. Proin eu arcu sit amet quam ultrices vehicula. Proin aliquam sem odio.

1.  Suspendisse in sem bibendum urna rutrum ornare.
2.  Nam at dui non nibh elementum facilisis.
3.  Vivamus eleifend lacus quis ex sagittis, venenatis interdum mauris cursus.
4.  Cras imperdiet lorem non libero sollicitudin lacinia.
    *   Etiam luctus sem non ante condimentum, vitae rhoncus risus mattis.
    *   Phasellus fringilla felis vitae accumsan aliquet.

## License ##

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
