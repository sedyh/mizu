# <img align="right" src="https://user-images.githubusercontent.com/19890545/147268423-d643c63a-96d2-40d1-9791-6cd842dc5647.png" alt="bunnymark" title="bunnymark" /> Bunnymark

A popular benchmark made on Mizu.

### Contents

- [Preview](#preview)
- [Running](#running)
- [Performance](#performance)


### Preview

<img src="https://user-images.githubusercontent.com/19890545/147268942-4c939aee-1c30-42d8-b792-39021fd62568.gif">

### Running

To run the benchmark from sources do the following command:

```
go run github.com/sedyh/mizu/examples/bunnymark@master
```
<sub>Please remember that @master only works since Go 17.</sub>

### Performance

Maximum objects at stable 60 FPS

| Software                                     | Hardware                 |  Mizu via systems objects | Ebiten objects |
|----------------------------------------------|--------------------------|---------------------------|----------------|
| Native, MacOS Big Sur 11.1                   | M1 2020                  | 31000                     | 65000          |
| Native, Linux Mint 20.2 Cinnamon             | Ryzen 5 3600, RX 5700 XT | 16000                     | 36000          |
