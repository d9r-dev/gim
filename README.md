Just another text editor. On my journey to learn another programming
language other than javascript I came across the list of 
interesting projects every software developer should try from 
Austin Henley [see](https://austinhenley.com/blog/challengingprojects.html).

## Data Structure
First things first. What data structure should one use for displaying the textfiles? Following resources were usefull:
* [Text Editor: Data Structures](https://austinhenley.com/blog/challengingprojects.html)

Basically, there is no way you could use an array for obvious reasons.
Other than that there are some data structures to consider:

* [Rope](https://en.wikipedia.org/wiki/Rope_(data_structure))
* [Gap Buffer](https://en.wikipedia.org/wiki/Gap_buffer)
* [Piece Table](https://en.wikipedia.org/wiki/Piece_table)

To make a long story short: Ropes are hard to write and maintain. 
Gap Buffers are okay but can be inefficient when editing large files and jumping around.
Piece Tables are pretty much optimal for this purpose, and they are used for example by Word.
Also, you can implement undo/redo pretty easily with a piece table.
