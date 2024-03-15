
# Magex

> **Warning**
>
> This is the **dev** branch. It means, that
> product of this version in the development
> status and have unstable API.

A game engine designed for **Survio**. It can be
used to write your own games. But, this product
**isn't provides stability**. Originally posted
for mod development.

Project **is open for any contributions!**

## Progress

**Progress in the moment of the first push:**
- Most of engine components are documented, but
not all. Documentation will be written at the
refactoring stage. Sorry.
- Physics: Rigidbody in the progress. Does not
work in this moment.
- SmartText: Is the smart system, that can
colorize the text by tags like: `[#fffff]`,
`[bg#ff0000]`, `[red]`, `[reset]`, etc.

- **These componenents are ready to use**:
- - `BgColor` -- *Provides background color*
- - `Camera` -- *Provides camera functional*
- - `Rigidbody` -- *Wraps the game object transform*
*over the 2D physics engine.*
- - - **NOTE ABOUT RIGIDBODY**. *there is also a*
*pseudocomponent* `Shapes`. *It defines functions*
*that helps to initiate the shapes*.
- - `Sprite` - *Draw some sprite on the camera*
- - `Tilemap` - *Easily places a tiles with the*
*own coordinate system*
- - `Transform` - *Base of almost any game object*
- - **UI**:
- - - `Button` - *Simple button that can handle*
*mouse clicks or hovering*
- - - `UISprite` - *Simplified version of the*
*Sprite. Used only for drawing sprites on the*
**UI** *level*
- - - `Text` - *Simple text*
- - - `UI Colliders` (*Box*, *Circle*) - *Simple*
*colliders based on the* [collision2d](https://github.com/Tarliton/collision2d)
*library.*

## Known issues

- **Rigidbody** is unstable, doesn't for in this
moment.
- **Camera** scaling is work, but these is no some
algorithms to scale the camera for the center
- **Tilemap** displays tiles incorrectly (one-two
pixels of the tile textures overlaps with other)
- **GOMap** calls event-functions with incorrect
priority. Incorrect texture mapping occurs
(blinking).

if you know a solution to the problem, please
create an [issue](https://github.com/ftdot/magex/issues/new)
or create a PR.
