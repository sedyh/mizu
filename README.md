# <img align="right" width="150" src="https://user-images.githubusercontent.com/19890545/147574976-97120058-08c8-4969-ab72-b5db18571635.png" alt="mizu" title="mizu" /> Mizu

Mizu is Entity Component System framework for Ebitengine.

Mizu is based on [ento](https://github.com/wfranczyk/ento), which is made by [wfranczyk](https://github.com/wfranczyk). 

> The name is short for Mizutaki, a japenese dish,
> where chicken pieces and vegetables stewed in a
> simple stock, and eaten with dipping sauce such
> as ponzu.

### Contents

- [Features](#features)
- [Examples](#examples)
- [Installation](#installation)
- [Usage](#usage)
- [Wiki](#wiki)

### Features

- Very low boilerplate in exchange for a certain amount of reflection.
- Scene management.
- Compile defined components and entities.

### Examples

To check all examples, visit [this](https://github.com/sedyh/mizu/tree/master/examples) page.

<a href="https://github.com/sedyh/mizu/tree/master/examples/particles"> <img width="150" src="https://user-images.githubusercontent.com/19890545/149218102-290ebacd-6cb6-472d-836f-462d4977f1c0.gif"></a> <a href="https://github.com/sedyh/mizu/tree/master/examples/tilemap"><img width="150" src="https://user-images.githubusercontent.com/19890545/149233216-652ad973-4942-494f-a790-5ff059f10559.gif"></a> <a href="https://github.com/sedyh/mizu/tree/master/examples/bunnymark"><img width="150" src="https://user-images.githubusercontent.com/19890545/149235154-52da3044-363e-491a-a25e-80915c5b8df4.gif"></a> <a href="https://github.com/sedyh/mizu/tree/master/examples/platformer"> <img width="150" src="https://user-images.githubusercontent.com/19890545/153062691-573d8647-2793-4b84-a04d-99803fe0f8c0.gif"></a>

### Installation

```
go get github.com/sedyh/mizu
```

### Usage

Import this package

```go
import "github.com/sedyh/mizu/pkg/engine"
```

Define components, attributes of your game objects.

```go
// Position for any entity, if it needs
type Pos struct {
    X, Y float64 // Just a 2D point
}

// Velocity for any entity, if it needs
type Vel struct {
    L, M float64 // Also, 2D point
}

// Radius for any entity, if it needs
type Rad struct {
    Value float64 // Width value
}
```

Define entities, your game objects.

```go
// Your game object
type Ball struct {
    Pos // Ball position
    Vel // Ball velocity
    Rad // Ball radius
}
```

Define systems, your game mechanics that will work for a 
specific set of components.

```go
// You can go through all entities that have a certain set of 
// components specifying the requirements in the fields of the system
type Velocity struct {
    *Pos // Current entity position
    *Vel // Current entity velocity
}

// Apply velocity for each entity that has Pos and Vel
func (v *Velocity) Update(w engine.World) {
    // If they are registered components, they will not be nil
    v.Pos.X += v.Vel.L
    v.Pos.Y += v.Vel.M
}

// When you need many sets of components
// in one system, you can use the views
type Render struct {}

// Render one frame
func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
    // But choose the right entities yourself
    view := w.View(Pos{}, Rad{})
    view.Each(func(entity engine.Entity) {
        var pos *Pos
        var rad *Rad
        entity.Get(&pos, &rad)
        
        ebitenutil.DrawRect(
            screen,
            pos.X-rad.Value/2, pos.Y-rad.Value/2,
            rad.Value, rad.Value,
            colornames.Aliceblue,
        )
    })
}
```

Define scenes, where will all this live.

```go
type Game struct{}

// Main scene, you can use that for settings or main menu
func (g *Game) Setup(w engine.World) {
    w.AddComponents(Pos{}, Vel{}, Rad{})
    w.AddEntities(&Ball{Pos{0, 0}, Vel{4, 4}, Rad{10}})
    w.AddSystems(&Velocity{}, &Render{})
}
```

Run the game.

```go
// Provides its own ebiten.Game implementation
g := engine.NewGame(&Game{})
if err := ebiten.RunGame(g); err != nil {
	log.Fatal(err)
}
```

### Wiki

Please visit our [wiki](../../wiki) for a helpful articles and best practices about Mizu.
