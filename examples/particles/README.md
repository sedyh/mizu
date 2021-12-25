# <img align="right" src="https://user-images.githubusercontent.com/19890545/146812487-90152c62-b2f4-4b3a-b550-6a4edf417817.gif" alt="mizu" title="Particles" /> Particles

A little particle system made on Mizu.

- To create any effect, you need to create a particle.
- All particles in this system are divided into main and secondary, they are called emitters and particles.
- Emitters contain the enabled root component, while the particles do not.
- This example contains a scene with a single emitter that creates particles each frame and transmits its settings to them.
- The emitter parameters indicate various physical quantities that you can play with to achieve various effects.

### Contents

- [Preview](#preview)
- [Running](#running)

### Preview

![particles-preview](https://user-images.githubusercontent.com/19890545/147395575-8d56c480-12fe-4b35-822a-ae4fd4ae9915.gif)

### Running

To run the particle system from sources do the following command:

```
go run github.com/sedyh/mizu/examples/particles@master
```
<sub>Please remember that @master only works since Go 17.</sub>
