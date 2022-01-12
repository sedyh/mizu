# <img align="right" src="https://user-images.githubusercontent.com/19890545/149225549-c9263f62-54ec-40d5-80bb-64a872465358.png" alt="mizu" title="Particles" /> Particles

A little particle system made on Mizu.

- To create any effect, you need to create a particle.
- All particles in this system are divided into main and secondary, they are called emitters and particles.
- Emitters contain the enabled root component, while the particles do not.
- This example contains a scene with a two emitters that creates particles each frame and transmits its settings to them.
- The emitter parameters indicate various physical quantities that you can play with to achieve various effects.

### Contents

- [Preview](#preview)
- [Running](#running)

### Preview

![particles-preview](https://user-images.githubusercontent.com/19890545/149218102-290ebacd-6cb6-472d-836f-462d4977f1c0.gif)

### Running

To run the example from sources do the following command:

```
go run github.com/sedyh/mizu/examples/particles@master
```
<sub>Please remember that @master only works since Go 17.</sub>
