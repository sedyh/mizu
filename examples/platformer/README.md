# <img align="right" width="80" src="https://user-images.githubusercontent.com/19890545/153058733-49e120de-9067-4cb8-8906-fb66222f1971.png" alt="platformer" title="platformer" /> Platformer

An example of a platformer with collision detection and resolution based on SweptAABB.

- Use WASD to control the player, space to jump, - and + to zoom in and out.
- The player and the crate will collide with tiles and each other.

Limitations:

- Only linear movement.
- No resizing while movement.
- Dynamic vs static detection and resolution.

### Contents

- [Preview](#preview)
- [Running](#running)
- [Performance](#performance)


### Preview

![platformer-preview](https://user-images.githubusercontent.com/19890545/153053601-a282aa37-be45-46b9-b258-d8765d89ffa5.gif)

### Running

To run the example from sources do the following command:

```
go run github.com/sedyh/mizu/examples/platformer@master
```
<sub>Please remember that @master only works since Go 17.</sub>
