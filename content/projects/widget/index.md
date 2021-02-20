+++
Area = "projects"
CrumbParent = "projects"
GitHub = "foo"
Layout = "page"
Tags = ["widget", "cpp", "mit license"]
Title = "Widget"
+++

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean in facilisis libero. Cras tincidunt augue non commodo
dictum. Ut finibus dignissim leo ut fringilla. Maecenas justo urna, vulputate id tellus at, fermentum pellentesque
lorem. Nulla a dolor a orci ornare hendrerit. Aliquam eget lorem ullamcorper, laoreet tellus quis, rhoncus leo. In quis
dolor varius, pretium lacus at, interdum est. Pellentesque sed metus condimentum felis imperdiet gravida.

Ut viverra, augue vel ultrices mollis, leo lacus dictum erat, vel lacinia arcu nisi id lorem. Mauris mollis elementum
viverra. Fusce convallis velit mollis viverra cursus. Curabitur bibendum magna tortor, mattis pretium quam rutrum eu. Ut
volutpat turpis ut metus sagittis lobortis. Praesent et urna sed est blandit posuere non ut enim. Vestibulum venenatis
mauris quis aliquet congue. Donec faucibus fermentum tortor non congue.

```cpp
#include <iostream>
#include <vector>
#include <stdexcept>

int main() {
    try {
        std::vector<int> vec{3,4,3,1};
        int i{vec.at(4)}; // Throws an exception, std::out_of_range (indexing for vec is from 0-3 not 1-4)
    }

    // An exception handler, catches std::out_of_range, which is thrown by vec.at(4)
    catch (std::out_of_range& e) {
        std::cerr << "Accessing a non-existent element: " << e.what() << '\n';
    }

    // To catch any other standard library exceptions (they derive from std::exception)
    catch (std::exception& e) {
        std::cerr << "Exception thrown: " << e.what() << '\n';
    }

    // Catch any unrecognised exceptions (i.e. those which don't derive from std::exception)
    catch (...) {
        std::cerr << "Some fatal error\n";
    }
}
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
