# <img align="right" width="100" src="https://user-images.githubusercontent.com/19890545/147268423-d643c63a-96d2-40d1-9791-6cd842dc5647.png" alt="bunnymark" title="bunnymark" /> Bunnymark

A popular benchmark made on Mizu.

- Close all other programs for more accurate results.
- Press the left mouse button to add some amount of gophers.
- Adjust the number of gophers that appear at a time with the mouse wheel.
- Increase the number of gophers until the FPS starts dropping below 60 to find out your result.
- To understand that the drop in performance is not a one-off - use the graphs on the left, they show TPS, FPS and the number of objects over a certain time.
- Press the right mouse button to disable batching, this will greatly increase the load, but keep in mind that all measurements were taken without coloring.

### Contents

- [Preview](#preview)
- [Running](#running)
- [Performance](#performance)


### Preview

![bunnymark-preview](https://user-images.githubusercontent.com/19890545/149235154-52da3044-363e-491a-a25e-80915c5b8df4.gif)

### Running

To run the example from sources do the following command:

```
go run github.com/sedyh/mizu/examples/bunnymark@master
```
<sub>Please remember that @master only works since Go 17.</sub>

### Performance

Maximum objects at stable 60 FPS

| Software                                     | Hardware                 |  Mizu objects | Ebiten objects |
|----------------------------------------------|--------------------------|---------------|----------------|
| Native, MacOS Big Sur 11.1                   | M1 2020                  | 31000         | 65000          |
| Native, Linux Mint 20.2 Cinnamon             | Ryzen 5 3600, RX 5700 XT | 16000         | 36000          |
