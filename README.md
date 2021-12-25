# <img align="right" src="https://user-images.githubusercontent.com/19890545/146812487-90152c62-b2f4-4b3a-b550-6a4edf417817.gif" alt="mizu" title="mizu" /> Mizu

Mizu is Entity Component System framework for Ebiten.

Mizu is based on [ento](https://github.com/wfranczyk/ento), which is made by [wfranczyk](https://github.com/wfranczyk). 

> The name is short for Mizutaki, a japenese dish,
> where chicken pieces and vegetables stewed in a
> simple stock, and eaten with dipping sauce such
> as ponzu.

### Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)

### Features

- Very low boilerplate in exchange for a certain amount of reflection.
- Scene management.
- Compile defined components and entities.

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

// Render every entity
func (r *Render) Draw(screen *ebiten.Image) {
    // But choose the ones you need yourself
    view := w.View(&Pos{}, Rad{})
    for _, e := range view.Filter() {
        var pos *Pos
        var rad *Rad
        entity.Get(&pos, &rad)
        
        ebitenutil.DrawRect(
            screen,
            pos.X-rad.Value/2, pos.Y-rad.Value/2,
            rad.Value, rad.Value,
            colornames.Aliceblue,
        )
    }
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

### Examples

To test the work of the framework, a small particle system is written on it.

Please check out the [examples/particles](https://github.com/sedyh/mizu/tree/master/examples/particles) folder for a demo.
